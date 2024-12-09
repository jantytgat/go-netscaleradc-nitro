package config

type DnsAddRec struct {
	Hostname    string  `json:"hostname,omitempty" nitro:"permission=readwrite"`
	IpAddress   string  `json:"ipaddress,omitempty" nitro:"permission=readwrite"`
	Ttl         float64 `json:"ttl,omitempty" nitro:"permission=readwrite"`
	EcsSubnet   string  `json:"ecssubnet,omitempty" nitro:"permission=readwrite"`
	Type        string  `json:"type,omitempty" nitro:"permission=readwrite"`
	NodeId      float64 `json:"nodeid,omitempty" nitro:"permission=readwrite"`
	Vservername string  `json:"vservername,omitempty" nitro:"permission=readonly"`
	Authtype    string  `json:"authtype,omitempty" nitro:"permission=readonly"`
	Count       float64 `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r DnsAddRec) GetTypeName() string {
	return "dnsaddrec"
}

func NewDnsAddRecAddRequest(hostname string, ipaddress string, ttl float64) DnsAddRec {
	return DnsAddRec{
		Hostname:  hostname,
		IpAddress: ipaddress,
		Ttl:       ttl,
	}
}
