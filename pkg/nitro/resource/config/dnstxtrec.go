package config

type DnsTxtRec struct {
	Domain     string   `json:"domain,omitempty" nitro:"permission=readwrite"`
	Data       []string `json:"string,omitempty" nitro:"permission=readwrite"`
	Ttl        float64  `json:"ttl,omitempty" nitro:"permission=readwrite"`
	RecordId   string   `json:"recordid,omitempty" nitro:"permission=readwrite"`
	EcsSubnet  string   `json:"ecssubnet,omitempty" nitro:"permission=readwrite"`
	RecordType string   `json:"type,omitempty" nitro:"permission=readwrite"`
	NodeId     string   `json:"nodeid,omitempty" nitro:"permission=readwrite"`
	AuthType   string   `json:"authtype,omitempty" nitro:"permission=readonly"`
	Count      float64  `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r DnsTxtRec) GetTypeName() string {
	return "dnstxtrec"
}

func NewDnsTxtRecAddRequest(domain string, data []string, ttl float64) DnsTxtRec {
	return DnsTxtRec{
		Domain: domain,
		Data:   data,
		Ttl:    ttl,
	}
}
