package config

var ServerFieldNames = struct {
	AppflowLog                string
	Autoscale                 string
	BoundTrafficDomain        string
	Cacheable                 string
	ClientIp                  string
	ClientKeepAlive           string
	ClientTimeout             string
	Comment                   string
	Compression               string
	Count                     string
	Delay                     string
	Domain                    string
	DomainResolveNow          string
	DomainResolveRetry        string
	DownstateFlush            string
	DupPort                   string
	DupServiceType            string
	Graceful                  string
	Internal                  string
	IpAddress                 string
	Ipv6Address               string
	MaxBandwidth              string
	MaxRequests               string
	Name                      string
	NewName                   string
	Port                      string
	QueryType                 string
	ServerConfigFlags         string
	ServerState               string
	ServerTimeout             string
	ServiceType               string
	State                     string
	StateChangeTimeSec        string
	SurgeProtection           string
	Svcitmactsvcs             string
	Svcitmboundsvcs           string
	Svcitmpriority            string
	TcpBuffering              string
	TicksSinceLastStateChange string
	TrafficDomain             string
	TranslationIp             string
	TranslationMask           string
	Usip                      string
	Weight                    string
}{
	AppflowLog:                "appflowlog",
	Autoscale:                 "autoscale",
	BoundTrafficDomain:        "boundtd",
	Cacheable:                 "cacheable",
	ClientIp:                  "cip",
	ClientKeepAlive:           "cka",
	ClientTimeout:             "clttimeout",
	Comment:                   "comment",
	Compression:               "cmp",
	Count:                     "__count",
	Delay:                     "delay",
	Domain:                    "domain",
	DomainResolveNow:          "domainresolvenow",
	DomainResolveRetry:        "domainresolveretry",
	DownstateFlush:            "downstateflush",
	DupPort:                   "dup_port",
	DupServiceType:            "dup_svctype",
	Graceful:                  "graceful",
	Internal:                  "internal",
	IpAddress:                 "ipaddress",
	Ipv6Address:               "ipv6address",
	MaxBandwidth:              "maxbandwidth",
	MaxRequests:               "maxreq",
	Name:                      "name",
	NewName:                   "newname",
	Port:                      "port",
	QueryType:                 "querytype",
	ServerConfigFlags:         "svrcfgflags",
	ServerState:               "svrstate",
	ServerTimeout:             "svrtimeout",
	ServiceType:               "svctype",
	State:                     "state",
	StateChangeTimeSec:        "statechangetimesec",
	SurgeProtection:           "sp",
	Svcitmactsvcs:             "svcitmactsvcs",
	Svcitmboundsvcs:           "svcitmboundsvcs",
	Svcitmpriority:            "svcitmpriority",
	TcpBuffering:              "tcpb",
	TicksSinceLastStateChange: "tickssincelaststatechange",
	TrafficDomain:             "td",
	TranslationIp:             "translationip",
	TranslationMask:           "translationmask",
	Usip:                      "usip",
	Weight:                    "weight",
}

type Server struct {
	AppflowLog                string  `json:"appflowlog,omitempty" nitro:"permission=readwrite"`
	Autoscale                 string  `json:"autoscale,omitempty" nitro:"permission=readonly"`
	BoundTrafficDomain        string  `json:"boundtd,omitempty" nitro:"permission=readonly"`
	Cacheable                 string  `json:"cacheable,omitempty" nitro:"permission=readonly"`
	ClientIp                  string  `json:"cip,omitempty" nitro:"permission=readonly"`
	ClientKeepAlive           string  `json:"cka,omitempty" nitro:"permission=readonly"`
	ClientTimeout             float64 `json:"clttimeout,omitempty" nitro:"permission=readonly"`
	Comment                   string  `json:"comment,omitempty" nitro:"permission=readwrite"`
	Compression               string  `json:"cmp,omitempty" nitro:"permission=readonly"`
	Count                     float64 `json:"__count,omitempty" nitro:"permission=readonly"`
	Delay                     float64 `json:"delay,omitempty" nitro:"permission=readwrite"`
	Domain                    string  `json:"domain,omitempty" nitro:"permission=readwrite"`
	DomainResolveNow          bool    `json:"domainresolvenow,omitempty" nitro:"permission=readwrite"`
	DomainResolveRetry        int     `json:"domainresolveretry,omitempty" nitro:"permission=readwrite"`
	DownstateFlush            string  `json:"downstateflush,omitempty" nitro:"permission=readwrite"`
	DupPort                   float64 `json:"dup_port,omitempty" nitro:"permission=readonly"`
	DupServiceType            string  `json:"dup_svctype,omitempty" nitro:"permission=readonly"`
	Graceful                  string  `json:"graceful,omitempty" nitro:"permission=readwrite"`
	Internal                  bool    `json:"internal,omitempty" nitro:"permission=readwrite"`
	IpAddress                 string  `json:"ipaddress,omitempty" nitro:"permission=readwrite"`
	Ipv6Address               string  `json:"ipv6address,omitempty" nitro:"permission=readwrite"`
	MaxBandwidth              string  `json:"maxbandwidth,omitempty" nitro:"permission=readwrite"`
	MaxRequests               string  `json:"maxreq,omitempty" nitro:"permission=readwrite"`
	Name                      string  `json:"name,omitempty" nitro:"permission=readwrite"`
	NewName                   string  `json:"newname,omitempty" nitro:"permission=readwrite"`
	Port                      int     `json:"port,omitempty" nitro:"permission=readwrite"`
	QueryType                 string  `json:"querytype,omitempty" nitro:"permission=readwrite"`
	ServerConfigFlags         string  `json:"svrcfgflags,omitempty" nitro:"permission=readonly"`
	ServerState               string  `json:"svrstate,omitempty" nitro:"permission=readonly"`
	ServerTimeout             int     `json:"svrtimeout,omitempty" nitro:"permission=readonly"`
	ServiceType               string  `json:"svctype,omitempty" nitro:"permission=readonly"`
	State                     string  `json:"state,omitempty" nitro:"permission=readwrite"`
	StateChangeTimeSec        string  `json:"statechangetimesec,omitempty" nitro:"permission=readonly"`
	SurgeProtection           string  `json:"sp,omitempty" nitro:"permission=readonly"`
	Svcitmactsvcs             string  `json:"svcitmactsvcs,omitempty" nitro:"permission=readonly"`
	Svcitmboundsvcs           string  `json:"svcitmboundsvcs,omitempty" nitro:"permission=readonly"`
	Svcitmpriority            string  `json:"svcitmpriority,omitempty" nitro:"permission=readonly"`
	TcpBuffering              string  `json:"tcpb,omitempty" nitro:"permission=readonly"`
	TicksSinceLastStateChange string  `json:"tickssincelaststatechange,omitempty" nitro:"permission=readonly"`
	TrafficDomain             string  `json:"td,omitempty" nitro:"permission=readwrite"`
	TranslationIp             string  `json:"translationip,omitempty" nitro:"permission=readwrite"`
	TranslationMask           string  `json:"translationmask,omitempty" nitro:"permission=readwrite"`
	Usip                      string  `json:"usip,omitempty" nitro:"permission=readonly"`
	Weight                    string  `json:"weight,omitempty" nitro:"permission=readwrite"`
}

func (r Server) GetTypeName() string {
	return "server"
}

func NewServerCreateByIpv4Request(name string, ip string) Server {
	return Server{
		Name:      name,
		IpAddress: ip,
	}
}

func NewServerCreateByIpv6Request(name string, ip string) Server {
	return Server{
		Name:        name,
		Ipv6Address: ip,
	}
}

func NewServerCreateByDomainRequest(name string, domain string) Server {
	return Server{
		Name:   name,
		Domain: domain,
	}
}

func NewServerDisableRequest(name string, delay float64, graceful bool) Server {
	r := Server{
		Name:  name,
		Delay: delay,
	}

	switch graceful {
	case true:
		r.Graceful = "YES"
	case false:
		r.Graceful = "NO"
	}

	return r
}

func NewServerRenameRequest(oldName string, newName string) Server {
	return Server{
		Name:    oldName,
		NewName: newName,
	}
}
