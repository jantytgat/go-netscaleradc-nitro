package config

type ResponderAction struct {
	Name               string   `json:"name,omitempty" nitro:"permission=readwrite"`
	Type               string   `json:"type,omitempty" nitro:"permission=readwrite"`
	Target             string   `json:"target,omitempty" nitro:"permission=readwrite"`
	HtmlPage           string   `json:"htmlpage,omitempty" nitro:"permission=readwrite"`
	BypassSafetyCheck  string   `json:"bypasssafetycheck,omitempty" nitro:"permission=readwrite"`
	Comment            string   `json:"comment,omitempty" nitro:"permission=readwrite"`
	ResponseStatusCode float64  `json:"responsestatuscode,omitempty" nitro:"permission=readwrite"`
	ReasonPhrase       string   `json:"reasonphrase,omitempty" nitro:"permission=readwrite"`
	Headers            []string `json:"headers,omitempty" nitro:"permission=readwrite"`
	NewName            string   `json:"newName,omitempty" nitro:"permission=readwrite"`
	Hits               string   `json:"hits,omitempty" nitro:"permission=readonly"`
	ReferenceCount     string   `json:"referencecount,omitempty" nitro:"permission=readonly"`
	UndefinedHits      float64  `json:"undefinedhits,omitempty" nitro:"permission=readonly"`
	Builtin            []string `json:"builtin,omitempty" nitro:"permission=readonly"`
	Feature            string   `json:"feature,omitempty" nitro:"permission=readonly"`
	Count              float64  `json:"__count,omitempty" nitro:"permission=readonly"`
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
