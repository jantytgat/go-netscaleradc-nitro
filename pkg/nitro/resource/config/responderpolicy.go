package config

var ResponderPolicyFieldNames = struct {
	Action          string
	ActivePolicy    string
	AppFlowAction   string
	Builtin         string
	Comment         string
	Count           string
	Feature         string
	Hits            string
	LogAction       string
	Name            string
	NewName         string
	Priority        string
	Rule            string
	UndefinedAction string
	UndefinedHits   string
}{
	Action:          "action",
	ActivePolicy:    "activepolicy",
	AppFlowAction:   "appflowaction",
	Builtin:         "builtin",
	Comment:         "comment",
	Count:           "__count",
	Feature:         "feature",
	Hits:            "hits",
	LogAction:       "logaction",
	Name:            "name",
	NewName:         "newname",
	Priority:        "priority",
	Rule:            "rule",
	UndefinedAction: "undefaction",
	UndefinedHits:   "undefhits",
}

type ResponderPolicy struct {
	Action          string   `json:"action,omitempty" nitro:"permission=readwrite"`
	ActivePolicy    int      `json:"activepolicy,omitempty" nitro:"permission=readonly"`
	AppFlowAction   string   `json:"appflowaction,omitempty" nitro:"permission=readwrite"`
	Builtin         []string `json:"builtin,omitempty" nitro:"permission=readonly"`
	Comment         string   `json:"comment,omitempty" nitro:"permission=readwrite"`
	Count           float64  `json:"__count,omitempty" nitro:"permission=readonly"`
	Feature         string   `json:"feature,omitempty" nitro:"permission=readonly"`
	Hits            string   `json:"hits,omitempty" nitro:"permission=readonly"`
	LogAction       string   `json:"logaction,omitempty" nitro:"permission=readwrite"`
	Name            string   `json:"name,omitempty" nitro:"permission=readwrite"`
	NewName         string   `json:"newname,omitempty" nitro:"permission=readwrite"`
	Priority        string   `json:"priority,omitempty" nitro:"permission=readonly"`
	Rule            string   `json:"rule,omitempty" nitro:"permission=readwrite"`
	UndefinedAction string   `json:"undefaction,omitempty" nitro:"permission=readwrite"`
	UndefinedHits   string   `json:"undefhits,omitempty" nitro:"permission=readonly"`
}

func (r ResponderPolicy) GetTypeName() string {
	return "responderpolicy"
}

func NewResponderPolicyAddRequest(name string, rule string, action string, undefinedAction string) ResponderPolicy {
	return ResponderPolicy{
		Name:            name,
		Rule:            rule,
		Action:          action,
		UndefinedAction: undefinedAction,
	}
}
