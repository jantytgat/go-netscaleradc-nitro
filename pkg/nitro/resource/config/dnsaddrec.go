package config

var DnsAddressRecFieldNames = struct {
	Authtype    string
	Count       string
	EcsSubnet   string
	Hostname    string
	IpAddress   string
	NodeId      string
	Ttl         string
	Type        string
	Vservername string
}{
	Authtype:    "authtype",
	Count:       "__count",
	EcsSubnet:   "ecssubnet",
	Hostname:    "hostname",
	IpAddress:   "ipaddress",
	NodeId:      "nodeid",
	Ttl:         "ttl",
	Type:        "type",
	Vservername: "vservername",
}

type DnsAddressRecord struct {
	Authtype    string  `json:"authtype,omitempty" nitro:"permission=readonly"`
	Count       float64 `json:"__count,omitempty" nitro:"permission=readonly"`
	EcsSubnet   string  `json:"ecssubnet,omitempty" nitro:"permission=readwrite"`
	Hostname    string  `json:"hostname,omitempty" nitro:"permission=readwrite"`
	IpAddress   string  `json:"ipaddress,omitempty" nitro:"permission=readwrite"`
	NodeId      float64 `json:"nodeid,omitempty" nitro:"permission=readwrite"`
	Ttl         float64 `json:"ttl,omitempty" nitro:"permission=readwrite"`
	Type        string  `json:"type,omitempty" nitro:"permission=readwrite"`
	Vservername string  `json:"vservername,omitempty" nitro:"permission=readonly"`
}

func (r DnsAddressRecord) GetTypeName() string {
	return "dnsaddrec"
}

func NewDnsAddRecAddRequest(hostname string, ipaddress string, ttl float64) DnsAddressRecord {
	return DnsAddressRecord{
		Hostname:  hostname,
		IpAddress: ipaddress,
		Ttl:       ttl,
	}
}
