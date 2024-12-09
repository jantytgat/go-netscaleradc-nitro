package config

type SystemUserSystemCmdPolicyBinding struct {
	Priority   float64 `json:"priority,omitempty" nitro:"permission=readwrite"`
	PolicyName string  `json:"policyname,omitempty" nitro:"permission=readwrite"`
	Username   string  `json:"username,omitempty" nitro:"permission=readwrite"`
	Count      float64 `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r SystemUserSystemCmdPolicyBinding) GetTypeName() string {
	return "systemuser_systemcmdpolicy_binding"
}

func NewSystemUserSystemCmdPolicyBindingAddRequest(username string, policyname string, priority float64) SystemUserSystemCmdPolicyBinding {
	return SystemUserSystemCmdPolicyBinding{
		Priority:   priority,
		PolicyName: policyname,
		Username:   username,
	}
}
