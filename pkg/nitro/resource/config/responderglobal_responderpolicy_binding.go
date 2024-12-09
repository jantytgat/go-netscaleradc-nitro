package config

type ResponderGlobalResponderPolicyBinding struct {
	Priority               string  `json:"priority,omitempty" nitro:"permission=readwrite"`
	GlobalBindType         string  `json:"globalbindtype,omitempty" nitro:"permission=readwrite"`
	GotoPriorityExpression string  `json:"gotopriorityexpression,omitempty" nitro:"permission=readwrite"`
	PolicyName             string  `json:"policyname,omitempty" nitro:"permission=readwrite"`
	Type                   string  `json:"type,omitempty" nitro:"permission=readwrite"`
	LabelType              string  `json:"labeltype,omitempty" nitro:"permission=readwrite"`
	LabelName              string  `json:"labelname,omitempty" nitro:"permission=readwrite"`
	Invoke                 bool    `json:"invoke,omitempty" nitro:"permission=readwrite"`
	BoundPolicyNumber      string  `json:"numpol,omitempty" nitro:"permission=readonly"`
	FlowType               string  `json:"flowtype,omitempty" nitro:"permission=readonly"`
	Count                  float64 `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r ResponderGlobalResponderPolicyBinding) GetTypeName() string {
	return "responderglobal_responderpolicy_binding"
}

func NewResponderGlobalResponderPolicyBindingAddRequest(name string, bindType string, priority string, gotoPriorityExpression string) ResponderGlobalResponderPolicyBinding {
	return ResponderGlobalResponderPolicyBinding{
		PolicyName:             name,
		Type:                   bindType,
		Priority:               priority,
		GotoPriorityExpression: gotoPriorityExpression,
	}
}
