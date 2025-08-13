package config

var (
	CmdActionAllow = CmdPolicyAction{"ALLOW"}
	CmdActionDeny  = CmdPolicyAction{"DENY"}
)

type CmdPolicyAction struct {
	string
}

var SystemCmdPolicyFieldNames = struct {
	Action     string
	Builtin    string
	CmdSpec    string
	Count      string
	Feature    string
	PolicyName string
}{
	Action:     "action",
	Builtin:    "builtin",
	CmdSpec:    "cmdspec",
	Count:      "__count",
	Feature:    "feature",
	PolicyName: "policyname",
}

type SystemCmdPolicy struct {
	Action     string   `json:"action" nitro:"permission=readwrite"`
	Builtin    []string `json:"builtin,omitempty" nitro:"permission=readonly"`
	CmdSpec    string   `json:"cmdspec,omitempty" nitro:"permission=readwrite"`
	Count      float64  `json:"__count,omitempty" nitro:"permission=readonly"`
	Feature    string   `json:"feature,omitempty" nitro:"permission=readonly"`
	PolicyName string   `json:"policyname,omitempty" nitro:"permission=readwrite"`
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
