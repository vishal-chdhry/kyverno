package api

import "sync"

type RuleResponseList struct {
	sync.Mutex
	responses []*RuleResponse
}

func NewRuleResponseList() *RuleResponseList {
	return &RuleResponseList{
		responses: make([]*RuleResponse, 0),
	}
}

func (r *RuleResponseList) Add(resp *RuleResponse) {
	r.Lock()
	defer r.Unlock()

	if r.responses == nil {
		r.responses = make([]*RuleResponse, 0)
	}
	r.responses = append(r.responses, resp)
}

func (r *RuleResponseList) GetAll() []*RuleResponse {
	return r.responses
}
