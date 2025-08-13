package config

var SystemUserSystemCmdPolicyBindingFieldNames = struct {
	Count      string
	PolicyName string
	Priority   string
	Username   string
}{
	Count:      "__count",
	PolicyName: "policyname",
	Priority:   "priority",
	Username:   "username",
}

type SystemUserSystemCmdPolicyBinding struct {
	Count      float64 `json:"__count,omitempty" nitro:"permission=readonly"`
	PolicyName string  `json:"policyname,omitempty" nitro:"permission=readwrite"`
	Priority   float64 `json:"priority,omitempty" nitro:"permission=readwrite"`
	Username   string  `json:"username,omitempty" nitro:"permission=readwrite"`
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
