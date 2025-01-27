package cel

import (
	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	admissionregistrationv1beta1 "k8s.io/api/admissionregistration/v1beta1"
	"k8s.io/apiserver/pkg/admission/plugin/cel"
	"k8s.io/apiserver/pkg/admission/plugin/policy/validating"
	"k8s.io/apiserver/pkg/admission/plugin/webhook/matchconditions"
	"k8s.io/apiserver/pkg/cel/environment"
)

type Compiler struct {
	compositedCompiler cel.CompositedCompiler
	// CEL expressions
	validateExpressions        []admissionregistrationv1beta1.Validation
	auditAnnotationExpressions []admissionregistrationv1beta1.AuditAnnotation
	matchExpressions           []admissionregistrationv1.MatchCondition
	variables                  []admissionregistrationv1beta1.Variable
}

func NewCompiler(
	validations []admissionregistrationv1beta1.Validation,
	auditAnnotations []admissionregistrationv1beta1.AuditAnnotation,
	matchConditions []admissionregistrationv1.MatchCondition,
	variables []admissionregistrationv1beta1.Variable,
) (*Compiler, error) {
	compositedCompiler, err := cel.NewCompositedCompiler(environment.MustBaseEnvSet(environment.DefaultCompatibilityVersion(), false))
	if err != nil {
		return nil, err
	}
	return &Compiler{
		compositedCompiler:         *compositedCompiler,
		validateExpressions:        validations,
		auditAnnotationExpressions: auditAnnotations,
		matchExpressions:           matchConditions,
		variables:                  variables,
	}, nil
}

func (c Compiler) CompileVariables(optionalVars cel.OptionalVariableDeclarations) {
	c.compositedCompiler.CompileAndStoreVariables(
		c.convertVariables(),
		optionalVars,
		environment.StoredExpressions,
	)
}

func (c Compiler) CompileValidateExpressions(optionalVars cel.OptionalVariableDeclarations) cel.ConditionEvaluator {
	return c.compositedCompiler.CompileCondition(
		c.convertValidations(),
		optionalVars,
		environment.StoredExpressions,
	)
}

func (c Compiler) CompileMessageExpressions(optionalVars cel.OptionalVariableDeclarations) cel.ConditionEvaluator {
	return c.compositedCompiler.CompileCondition(
		c.convertMessageExpressions(),
		optionalVars,
		environment.StoredExpressions,
	)
}

func (c Compiler) CompileAuditAnnotationsExpressions(optionalVars cel.OptionalVariableDeclarations) cel.ConditionEvaluator {
	return c.compositedCompiler.CompileCondition(
		c.convertAuditAnnotations(),
		optionalVars,
		environment.StoredExpressions,
	)
}

func (c Compiler) CompileMatchExpressions(optionalVars cel.OptionalVariableDeclarations) cel.ConditionEvaluator {
	return c.compositedCompiler.CompileCondition(
		c.convertMatchExpressions(),
		optionalVars,
		environment.StoredExpressions,
	)
}

func (c Compiler) convertValidations() []cel.ExpressionAccessor {
	celExpressionAccessor := make([]cel.ExpressionAccessor, len(c.validateExpressions))
	for i, validation := range c.validateExpressions {
		validation := validating.ValidationCondition{
			Expression: validation.Expression,
			Message:    validation.Message,
			Reason:     validation.Reason,
		}
		celExpressionAccessor[i] = &validation
	}
	return celExpressionAccessor
}

func (c Compiler) convertMessageExpressions() []cel.ExpressionAccessor {
	celExpressionAccessor := make([]cel.ExpressionAccessor, len(c.validateExpressions))
	for i, validation := range c.validateExpressions {
		if validation.MessageExpression != "" {
			condition := validating.MessageExpressionCondition{
				MessageExpression: validation.MessageExpression,
			}
			celExpressionAccessor[i] = &condition
		}
	}
	return celExpressionAccessor
}

func (c Compiler) convertAuditAnnotations() []cel.ExpressionAccessor {
	celExpressionAccessor := make([]cel.ExpressionAccessor, len(c.auditAnnotationExpressions))
	for i, validation := range c.auditAnnotationExpressions {
		validation := validating.AuditAnnotationCondition{
			Key:             validation.Key,
			ValueExpression: validation.ValueExpression,
		}
		celExpressionAccessor[i] = &validation
	}
	return celExpressionAccessor
}

func (c Compiler) convertMatchExpressions() []cel.ExpressionAccessor {
	celExpressionAccessor := make([]cel.ExpressionAccessor, len(c.matchExpressions))
	for i, condition := range c.matchExpressions {
		condition := matchconditions.MatchCondition{
			Name:       condition.Name,
			Expression: condition.Expression,
		}
		celExpressionAccessor[i] = &condition
	}
	return celExpressionAccessor
}

func (c Compiler) convertVariables() []cel.NamedExpressionAccessor {
	namedExpressions := make([]cel.NamedExpressionAccessor, len(c.variables))
	for i, variable := range c.variables {
		namedExpressions[i] = &validating.Variable{
			Name:       variable.Name,
			Expression: variable.Expression,
		}
	}
	return namedExpressions
}
