package config

var ResponderActionFieldNames = struct {
	Builtin            string
	BypassSafetyCheck  string
	Comment            string
	Count              string
	Feature            string
	Headers            string
	Hits               string
	HtmlPage           string
	Name               string
	NewName            string
	ReasonPhrase       string
	ReferenceCount     string
	ResponseStatusCode string
	Target             string
	Type               string
	UndefinedHits      string
}{
	Builtin:            "builtin",
	BypassSafetyCheck:  "bypasssafetycheck",
	Comment:            "comment",
	Count:              "__count",
	Feature:            "feature",
	Headers:            "headers",
	Hits:               "hits",
	HtmlPage:           "htmlpage",
	Name:               "name",
	NewName:            "newName",
	ReasonPhrase:       "reasonphrase",
	ReferenceCount:     "referencecount",
	ResponseStatusCode: "responsestatuscode",
	Target:             "target",
	Type:               "type",
	UndefinedHits:      "undefhits",
}

type ResponderAction struct {
	Builtin            []string `json:"builtin,omitempty" nitro:"permission=readonly"`
	BypassSafetyCheck  string   `json:"bypasssafetycheck,omitempty" nitro:"permission=readwrite"`
	Comment            string   `json:"comment,omitempty" nitro:"permission=readwrite"`
	Count              float64  `json:"__count,omitempty" nitro:"permission=readonly"`
	Feature            string   `json:"feature,omitempty" nitro:"permission=readonly"`
	Headers            []string `json:"headers,omitempty" nitro:"permission=readwrite"`
	Hits               string   `json:"hits,omitempty" nitro:"permission=readonly"`
	HtmlPage           string   `json:"htmlpage,omitempty" nitro:"permission=readwrite"`
	Name               string   `json:"name,omitempty" nitro:"permission=readwrite"`
	NewName            string   `json:"newName,omitempty" nitro:"permission=readwrite"`
	ReasonPhrase       string   `json:"reasonphrase,omitempty" nitro:"permission=readwrite"`
	ReferenceCount     string   `json:"referencecount,omitempty" nitro:"permission=readonly"`
	ResponseStatusCode float64  `json:"responsestatuscode,omitempty" nitro:"permission=readwrite"`
	Target             string   `json:"target,omitempty" nitro:"permission=readwrite"`
	Type               string   `json:"type,omitempty" nitro:"permission=readwrite"`
	UndefinedHits      string   `json:"undefhits,omitempty" nitro:"permission=readonly"`
}

func (r ResponderAction) GetTypeName() string {
	return "responderaction"
}

func NewResponderActionAddRequest(name string, respondertype string, target string) ResponderAction {
	return ResponderAction{
		Name:   name,
		Type:   respondertype,
		Target: target,
	}
}
