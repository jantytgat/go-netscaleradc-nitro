package config

var (
	CmdActionAllow = CmdPolicyAction{"ALLOW"}
	CmdActionDeny  = CmdPolicyAction{"DENY"}
)

type CmdPolicyAction struct {
	string
}

type SystemCmdPolicy struct {
	PolicyName string   `json:"policyname,omitempty" nitro:"permission=readwrite"`
	Action     string   `json:"action" nitro:"permission=readwrite"`
	CmdSpec    string   `json:"cmdspec,omitempty" nitro:"permission=readwrite"`
	Builtin    []string `json:"builtin,omitempty" nitro:"permission=readonly"`
	Feature    string   `json:"feature,omitempty" nitro:"permission=readonly"`
	Count      float64  `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r SystemCmdPolicy) GetTypeName() string {
	return "systemcmdpolicy"
}

func NewSystemCmdPolicyAddRequest(name string, action CmdPolicyAction, spec string) SystemCmdPolicy {
	return SystemCmdPolicy{
		PolicyName: name,
		Action:     action.string,
		CmdSpec:    spec,
	}
}
