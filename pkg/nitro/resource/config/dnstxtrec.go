package config

var DnsTxtRecordFieldNames = struct {
	AuthType   string
	Count      string
	Data       string
	Domain     string
	EcsSubnet  string
	NodeId     string
	RecordId   string
	RecordType string
	Ttl        string
}{
	AuthType:   "authtype",
	Count:      "__count",
	Data:       "string",
	Domain:     "domain",
	EcsSubnet:  "ecssubnet",
	NodeId:     "nodeid",
	RecordId:   "recordid",
	RecordType: "type",
	Ttl:        "ttl",
}

type DnsTxtRecord struct {
	AuthType   string   `json:"authtype,omitempty" nitro:"permission=readonly"`
	Count      float64  `json:"__count,omitempty" nitro:"permission=readonly"`
	Data       []string `json:"string,omitempty" nitro:"permission=readwrite"`
	Domain     string   `json:"domain,omitempty" nitro:"permission=readwrite"`
	EcsSubnet  string   `json:"ecssubnet,omitempty" nitro:"permission=readwrite"`
	NodeId     string   `json:"nodeid,omitempty" nitro:"permission=readwrite"`
	RecordId   string   `json:"recordid,omitempty" nitro:"permission=readwrite"`
	RecordType string   `json:"type,omitempty" nitro:"permission=readwrite"`
	Ttl        float64  `json:"ttl,omitempty" nitro:"permission=readwrite"`
}

func (r DnsTxtRecord) GetTypeName() string {
	return "dnstxtrec"
}

func NewDnsTxtRecAddRequest(domain string, data []string, ttl float64) DnsTxtRecord {
	return DnsTxtRecord{
		Domain: domain,
		Data:   data,
		Ttl:    ttl,
	}
}
