package factories

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	kyvernov1 "github.com/kyverno/kyverno/api/kyverno/v1"
	"github.com/kyverno/kyverno/pkg/clients/dclient"
	engineapi "github.com/kyverno/kyverno/pkg/engine/api"
	enginecontext "github.com/kyverno/kyverno/pkg/engine/context"
	"github.com/kyverno/kyverno/pkg/engine/context/loaders"
	"github.com/kyverno/kyverno/pkg/engine/jmespath"
	"github.com/kyverno/kyverno/pkg/logging"
	"github.com/kyverno/kyverno/pkg/registryclient"
)

type ContextLoaderFactoryOptions func(*contextLoader)

func DefaultContextLoaderFactory(cmResolver engineapi.ConfigmapResolver, opts ...ContextLoaderFactoryOptions) engineapi.ContextLoaderFactory {
	return func(_ kyvernov1.PolicyInterface, _ kyvernov1.Rule) engineapi.ContextLoader {
		cl := &contextLoader{
			logger:     logging.WithName("DefaultContextLoaderFactory"),
			cmResolver: cmResolver,
		}
		for _, o := range opts {
			o(cl)
		}
		return cl
	}
}

func WithInitializer(initializer engineapi.Initializer) ContextLoaderFactoryOptions {
	return func(cl *contextLoader) {
		cl.initializers = append(cl.initializers, initializer)
	}
}

type contextLoader struct {
	logger       logr.Logger
	cmResolver   engineapi.ConfigmapResolver
	initializers []engineapi.Initializer
}

func (l *contextLoader) Load(
	ctx context.Context,
	jp jmespath.Interface,
	client dclient.Interface,
	rclient registryclient.Client,
	contextEntries []kyvernov1.ContextEntry,
	jsonContext enginecontext.Interface,
) error {
	for _, init := range l.initializers {
		if err := init(jsonContext); err != nil {
			return err
		}
	}
	for _, entry := range contextEntries {
		deferredLoader, err := l.newDeferredLoader(ctx, jp, client, rclient, entry, jsonContext)
		if err != nil {
			return fmt.Errorf("failed to create deferred loader for context entry %s", entry.Name)
		}
		if deferredLoader != nil {
			if err := jsonContext.AddDeferredLoader(entry.Name, deferredLoader); err != nil {
				return err
			}
		}
	}
	return nil
}

func (l *contextLoader) newDeferredLoader(
	ctx context.Context,
	jp jmespath.Interface,
	client dclient.Interface,
	rclient registryclient.Client,
	entry kyvernov1.ContextEntry,
	jsonContext enginecontext.Interface,
) (enginecontext.Loader, error) {
	if entry.ConfigMap != nil {
		if l.cmResolver != nil {
			l := loaders.NewConfigMapLoader(ctx, l.logger, entry, l.cmResolver, jsonContext)
			return l, nil
		} else {
			l.logger.Info("disabled loading of ConfigMap context entry %s", entry.Name)
			return nil, nil
		}
	} else if entry.APICall != nil {
		if client != nil {
			l := loaders.NewAPILoader(ctx, l.logger, entry, jsonContext, jp, client)
			return l, nil
		} else {
			l.logger.Info("disabled loading of APICall context entry %s", entry.Name)
			return nil, nil
		}
	} else if entry.ImageRegistry != nil {
		if rclient != nil {
			l := loaders.NewImageDataLoader(ctx, l.logger, entry, jsonContext, jp, rclient)
			return l, nil
		} else {
			l.logger.Info("disabled loading of ImageRegistry context entry %s", entry.Name)
			return nil, nil
		}
	} else if entry.Variable != nil {
		l := loaders.NewVariableLoader(l.logger, entry, jsonContext, jp)
		return l, nil
	}

	return nil, fmt.Errorf("missing ConfigMap|APICall|ImageRegistry|Variable in context entry %s", entry.Name)
}
