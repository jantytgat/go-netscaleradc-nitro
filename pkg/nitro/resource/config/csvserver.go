package config

import "log/slog"

var CsVserverFieldNames = struct {
	AppflowLog                    string
	Authentication                string
	Authentication401             string
	AuthenticationHost            string
	AuthenticationProfile         string
	AuthenticationVserverName     string
	BackupIp                      string
	BackupPersistenceTimeout      string
	BackupVserver                 string
	Bindpoint                     string
	Cacheable                     string
	CacheType                     string
	CacheVserver                  string
	CaseSensitive                 string
	ClientTimeout                 string
	Comment                       string
	CookieDomain                  string
	CookieName                    string
	CookieTimeout                 string
	Count                         string
	CurrentState                  string
	DatabaseProfileName           string
	DisablePrimaryOnDown          string
	DnsOverHttps                  string
	DnsProfileName                string
	DnsRecordType                 string
	DnsVserverName                string
	Domain                        string
	DomainName                    string
	DownstateFlush                string
	Dtls                          string
	GreaterThan2GBTransactions    string
	Hits                          string
	Homepage                      string
	HttpProfileName               string
	HttpRedirectUrl               string
	IcmpVserverResponse           string
	InsertVserverIpPort           string
	Invoke                        string
	Ip                            string
	IpMask                        string
	IpPattern                     string
	IpSet                         string
	Ipv46                         string
	Ipv6PersistenceMaskLength     string
	L2ConnectionParameters        string
	LbVserver                     string
	ListenPolicy                  string
	ListenPriority                string
	MssqlServerVersion            string
	MysqlCharacterSet             string
	MysqlProtocolVersion          string
	MysqlServerCapabilities       string
	MysqlServerVersion            string
	Name                          string
	NetProfile                    string
	NewName                       string
	NoDefaultBindings             string
	NodegroupName                 string
	OracleServerVersion           string
	PersistenceBackup             string
	PersistenceId                 string
	PersistenceMask               string
	PersistenceType               string
	PiPolicyHits                  string
	Port                          string
	Precedence                    string
	Priority                      string
	ProbePort                     string
	ProbeProtocol                 string
	ProbeSuccessResponseCode      string
	Push                          string
	PushLabel                     string
	PushMultipleClientConnections string
	PushVserver                   string
	QuicProfileName               string
	Range                         string
	Redirect                      string
	RedirectFromPort              string
	RedirectPortRewrite           string
	RedirectUrl                   string
	RouteHealthInjectionState     string
	RtspNat                       string
	RuleType                      string
	ServiceName                   string
	ServiceType                   string
	SiteDomainTimeToLive          string
	SpilloverBackupAction         string
	SpilloverMethod               string
	SpilloverPersistence          string
	SpilloverPersistenceTimeout   string
	SpilloverThreshold            string
	State                         string
	StateChangeTimeMilliSecond    string
	StateChangeTimeSeconds        string
	StateUpdate                   string
	Status                        string
	TargetLbVserver               string
	TargetType                    string
	TargetVserver                 string
	TcpProbePort                  string
	TcpProfileName                string
	TicksSinceLastStateChange     string
	Timeout                       string
	TimeToLive                    string
	TrafficDomain                 string
	Type                          string
	Url                           string
	Value                         string
	Version                       string
	VipHeader                     string
	Weight                        string
}{
	AppflowLog:                    "appflowlog",
	Authentication:                "authentication",
	Authentication401:             "authn401",
	AuthenticationHost:            "authenticationhost",
	AuthenticationProfile:         "authnprofile",
	AuthenticationVserverName:     "authnvsname",
	BackupIp:                      "backupip",
	BackupPersistenceTimeout:      "backuppersistencetimeout",
	BackupVserver:                 "backupvserver",
	Bindpoint:                     "bindpoint",
	Cacheable:                     "cacheable",
	CacheType:                     "cachetype",
	CacheVserver:                  "cachevserver",
	CaseSensitive:                 "casesensitive",
	ClientTimeout:                 "clttimeout",
	Comment:                       "comment",
	CookieDomain:                  "cookiedomain",
	CookieName:                    "cookiename",
	CookieTimeout:                 "cookietimeout",
	Count:                         "__count",
	CurrentState:                  "curstate",
	DatabaseProfileName:           "dbprofilename",
	DisablePrimaryOnDown:          "disableprimaryondown",
	DnsOverHttps:                  "dnsoverhttps",
	DnsProfileName:                "dnsprofilename",
	DnsRecordType:                 "dnsrecordtype",
	DnsVserverName:                "dnsvservername",
	Domain:                        "domain",
	DomainName:                    "domainname",
	DownstateFlush:                "downstateflush",
	Dtls:                          "dtls",
	GreaterThan2GBTransactions:    "gt2gb",
	Hits:                          "hits",
	Homepage:                      "homepage",
	HttpProfileName:               "httpprofilename",
	HttpRedirectUrl:               "httpsredirecturl",
	IcmpVserverResponse:           "icmpvsrresponse",
	InsertVserverIpPort:           "insertvserveripport",
	Invoke:                        "invoke",
	Ip:                            "ip",
	IpMask:                        "ipmask",
	IpPattern:                     "ippattern",
	IpSet:                         "ipset",
	Ipv46:                         "ipv46",
	Ipv6PersistenceMaskLength:     "v6persistmasklen",
	L2ConnectionParameters:        "l2conn",
	LbVserver:                     "lbvserver",
	ListenPolicy:                  "listenpolicy",
	ListenPriority:                "listenpriority",
	MssqlServerVersion:            "mssqlserverversion",
	MysqlCharacterSet:             "mysqlcharacterset",
	MysqlProtocolVersion:          "mysqlprotocolversion",
	MysqlServerCapabilities:       "mysqlservercapabilities",
	MysqlServerVersion:            "mysqlserverversion",
	Name:                          "name",
	NetProfile:                    "netprofile",
	NewName:                       "newname",
	NoDefaultBindings:             "nodefaultbindings",
	NodegroupName:                 "ngname",
	OracleServerVersion:           "oracleserverversion",
	PersistenceBackup:             "persistencebackup",
	PersistenceId:                 "persistenceid",
	PersistenceMask:               "persistmask",
	PersistenceType:               "persistencetype",
	PiPolicyHits:                  "pipolicyhits",
	Port:                          "port",
	Precedence:                    "precedence",
	Priority:                      "priority",
	ProbePort:                     "probeport",
	ProbeProtocol:                 "probeprotocol",
	ProbeSuccessResponseCode:      "probesuccessresponsecode",
	Push:                          "push",
	PushLabel:                     "pushlabel",
	PushMultipleClientConnections: "pushmulticlients",
	PushVserver:                   "pushvserver",
	QuicProfileName:               "quicprofilename",
	Range:                         "range",
	Redirect:                      "redirect",
	RedirectFromPort:              "redirectfromport",
	RedirectPortRewrite:           "redirectportrewrite",
	RedirectUrl:                   "redirecturl",
	RouteHealthInjectionState:     "rhistate",
	RtspNat:                       "rtspnat",
	RuleType:                      "ruletype",
	ServiceName:                   "servicename",
	ServiceType:                   "servicetype",
	SiteDomainTimeToLive:          "sitedomainttl",
	SpilloverBackupAction:         "sobackupaction",
	SpilloverMethod:               "somethod",
	SpilloverPersistence:          "sopersistence",
	SpilloverPersistenceTimeout:   "sopersistencetimeout",
	SpilloverThreshold:            "sothreshold",
	State:                         "state",
	StateChangeTimeMilliSecond:    "statechangetimemsec",
	StateChangeTimeSeconds:        "statechangetimesec",
	StateUpdate:                   "stateupdate",
	Status:                        "status",
	TargetLbVserver:               "targetlbvserver",
	TargetType:                    "targettype",
	TargetVserver:                 "targetvserver",
	TcpProbePort:                  "tcpprobeport",
	TcpProfileName:                "tcpprofilename",
	TicksSinceLastStateChange:     "tickssincelaststatechange",
	Timeout:                       "timeout",
	TimeToLive:                    "ttl",
	TrafficDomain:                 "td",
	Type:                          "type",
	Url:                           "url",
	Value:                         "value",
	Version:                       "version",
	VipHeader:                     "vipheader",
	Weight:                        "weight",
}

type CsVserver struct {
	AppflowLog                    string  `json:"appflowlog,omitempty" nitro:"permission=readwrite"`
	Authentication                string  `json:"authentication,omitempty" nitro:"permission=readwrite"`
	Authentication401             string  `json:"authn401,omitempty" nitro:"permission=readwrite"`
	AuthenticationHost            string  `json:"authenticationhost,omitempty" nitro:"permission=readwrite"`
	AuthenticationProfile         string  `json:"authnprofile,omitempty" nitro:"permission=readwrite"`
	AuthenticationVserverName     string  `json:"authnvsname,omitempty" nitro:"permission=readwrite"`
	BackupIp                      string  `json:"backupip,omitempty" nitro:"permission=readwrite"`
	BackupPersistenceTimeout      float64 `json:"backuppersistencetimeout,omitempty" nitro:"permission=readwrite"`
	BackupVserver                 string  `json:"backupvserver,omitempty" nitro:"permission=readwrite"`
	Bindpoint                     string  `json:"bindpoint,omitempty" nitro:"permission=readonly"`
	Cacheable                     string  `json:"cacheable,omitempty" nitro:"permission=readwrite"`
	CacheType                     string  `json:"cachetype,omitempty" nitro:"permission=readonly"`
	CacheVserver                  string  `json:"cachevserver,omitempty" nitro:"permission=readonly"`
	CaseSensitive                 string  `json:"casesensitive,omitempty" nitro:"permission=readwrite"`
	ClientTimeout                 string  `json:"clttimeout,omitempty" nitro:"permission=readwrite"`
	Comment                       string  `json:"comment,omitempty" nitro:"permission=readwrite"`
	CookieDomain                  string  `json:"cookiedomain,omitempty" nitro:"permission=readwrite"`
	CookieName                    string  `json:"cookiename,omitempty" nitro:"permission=readwrite"`
	CookieTimeout                 float64 `json:"cookietimeout,omitempty" nitro:"permission=readwrite"`
	Count                         float64 `json:"__count,omitempty" nitro:"permission=readonly"`
	CurrentState                  string  `json:"curstate,omitempty" nitro:"permission=readonly"`
	DatabaseProfileName           string  `json:"dbprofilename,omitempty" nitro:"permission=readwrite"`
	DisablePrimaryOnDown          string  `json:"disableprimaryondown,omitempty" nitro:"permission=readwrite"`
	DnsOverHttps                  string  `json:"dnsoverhttps,omitempty" nitro:"permission=readwrite"`
	DnsProfileName                string  `json:"dnsprofilename,omitempty" nitro:"permission=readwrite"`
	DnsRecordType                 string  `json:"dnsrecordtype,omitempty" nitro:"permission=readwrite"`
	DnsVserverName                string  `json:"dnsvservername,omitempty" nitro:"permission=readonly"`
	Domain                        string  `json:"domain,omitempty" nitro:"permission=readonly"`
	DomainName                    string  `json:"domainname,omitempty" nitro:"permission=readwrite"`
	DownstateFlush                string  `json:"downstateflush,omitempty" nitro:"permission=readwrite"`
	Dtls                          string  `json:"dtls,omitempty" nitro:"permission=readwrite"`
	GreaterThan2GBTransactions    string  `json:"gt2gb,omitempty" nitro:"permission=readonly"`
	Hits                          string  `json:"hits,omitempty" nitro:"permission=readwrite"`
	Homepage                      string  `json:"homepage,omitempty" nitro:"permission=readonly"`
	HttpProfileName               string  `json:"httpprofilename,omitempty" nitro:"permission=readwrite"`
	HttpRedirectUrl               string  `json:"httpsredirecturl,omitempty" nitro:"permission=readwrite"`
	IcmpVserverResponse           string  `json:"icmpvsrresponse,omitempty" nitro:"permission=readwrite"`
	InsertVserverIpPort           string  `json:"insertvserveripport,omitempty" nitro:"permission=readwrite"`
	Invoke                        bool    `json:"invoke,omitempty" nitro:"permission=readwrite"`
	Ip                            string  `json:"ip,omitempty" nitro:"permission=readonly"`
	IpMask                        string  `json:"ipmask,omitempty" nitro:"permission=readwrite"`
	IpPattern                     string  `json:"ippattern,omitempty" nitro:"permission=readwrite"`
	IpSet                         string  `json:"ipset,omitempty" nitro:"permission=readwrite"`
	Ipv46                         string  `json:"ipv46,omitempty" nitro:"permission=readwrite"`
	Ipv6PersistenceMaskLength     string  `json:"v6persistmasklen,omitempty" nitro:"permission=readwrite"`
	L2ConnectionParameters        string  `json:"l2conn,omitempty" nitro:"permission=readwrite"`
	LbVserver                     string  `json:"lbvserver,omitempty" nitro:"permission=readonly"`
	ListenPolicy                  string  `json:"listenpolicy,omitempty" nitro:"permission=readwrite"`
	ListenPriority                string  `json:"listenpriority,omitempty" nitro:"permission=readwrite"`
	MssqlServerVersion            string  `json:"mssqlserverversion,omitempty" nitro:"permission=readwrite"`
	MysqlCharacterSet             string  `json:"mysqlcharacterset,omitempty" nitro:"permission=readwrite"`
	MysqlProtocolVersion          string  `json:"mysqlprotocolversion,omitempty" nitro:"permission=readwrite"`
	MysqlServerCapabilities       string  `json:"mysqlservercapabilities,omitempty" nitro:"permission=readwrite"`
	MysqlServerVersion            string  `json:"mysqlserverversion,omitempty" nitro:"permission=readwrite"`
	Name                          string  `json:"name,omitempty" nitro:"permission=readwrite"`
	NetProfile                    string  `json:"netprofile,omitempty" nitro:"permission=readwrite"`
	NewName                       string  `json:"newname,omitempty" nitro:"permission=readwrite"`
	NoDefaultBindings             string  `json:"nodefaultbindings,omitempty" nitro:"permission=readonly"`
	NodegroupName                 string  `json:"ngname,omitempty" nitro:"permission=readonly"`
	OracleServerVersion           string  `json:"oracleserverversion,omitempty" nitro:"permission=readwrite"`
	PersistenceBackup             string  `json:"persistencebackup,omitempty" nitro:"permission=readwrite"`
	PersistenceId                 float64 `json:"persistenceid,omitempty" nitro:"permission=readwrite"`
	PersistenceMask               string  `json:"persistmask,omitempty" nitro:"permission=readwrite"`
	PersistenceType               string  `json:"persistencetype,omitempty" nitro:"permission=readwrite"`
	PiPolicyHits                  string  `json:"pipolicyhits,omitempty" nitro:"permission=readwrite"`
	Port                          int     `json:"port,omitempty" nitro:"permission=readwrite"`
	Precedence                    string  `json:"precedence,omitempty" nitro:"permission=readwrite"`
	Priority                      string  `json:"priority,omitempty" nitro:"permission=readwrite"`
	ProbePort                     int     `json:"probeport,omitempty" nitro:"permission=readwrite"`
	ProbeProtocol                 string  `json:"probeprotocol,omitempty" nitro:"permission=readwrite"`
	ProbeSuccessResponseCode      string  `json:"probesuccessresponsecode,omitempty" nitro:"permission=readwrite"`
	Push                          string  `json:"push,omitempty" nitro:"permission=readwrite"`
	PushLabel                     string  `json:"pushlabel,omitempty" nitro:"permission=readwrite"`
	PushMultipleClientConnections string  `json:"pushmulticlients,omitempty" nitro:"permission=readwrite"`
	PushVserver                   string  `json:"pushvserver,omitempty" nitro:"permission=readwrite"`
	QuicProfileName               string  `json:"quicprofilename,omitempty" nitro:"permission=readwrite"`
	Range                         string  `json:"range,omitempty" nitro:"permission=readwrite"`
	Redirect                      string  `json:"redirect,omitempty" nitro:"permission=readonly"`
	RedirectFromPort              int     `json:"redirectfromport,omitempty" nitro:"permission=readwrite"`
	RedirectPortRewrite           string  `json:"redirectportrewrite,omitempty" nitro:"permission=readwrite"`
	RedirectUrl                   string  `json:"redirecturl,omitempty" nitro:"permission=readwrite"`
	RouteHealthInjectionState     string  `json:"rhistate,omitempty" nitro:"permission=readwrite"`
	RtspNat                       string  `json:"rtspnat,omitempty" nitro:"permission=readwrite"`
	RuleType                      string  `json:"ruletype,omitempty" nitro:"permission=readonly"`
	ServiceName                   string  `json:"servicename,omitempty" nitro:"permission=readonly"`
	ServiceType                   string  `json:"servicetype,omitempty" nitro:"permission=readwrite"`
	SiteDomainTimeToLive          float64 `json:"sitedomainttl,omitempty" nitro:"permission=readwrite"`
	SpilloverBackupAction         string  `json:"sobackupaction,omitempty" nitro:"permission=readwrite"`
	SpilloverMethod               string  `json:"somethod,omitempty" nitro:"permission=readwrite"`
	SpilloverPersistence          string  `json:"sopersistence,omitempty" nitro:"permission=readwrite"`
	SpilloverPersistenceTimeout   string  `json:"sopersistencetimeout,omitempty" nitro:"permission=readwrite"`
	SpilloverThreshold            float64 `json:"sothreshold,omitempty" nitro:"permission=readwrite"`
	State                         string  `json:"state,omitempty" nitro:"permission=readwrite"`
	StateChangeTimeMilliSecond    string  `json:"statechangetimemsec,omitempty" nitro:"permission=readonly"`
	StateChangeTimeSeconds        string  `json:"statechangetimesec,omitempty" nitro:"permission=readonly"`
	StateUpdate                   string  `json:"stateupdate,omitempty" nitro:"permission=readwrite"`
	Status                        int     `json:"status,omitempty" nitro:"permission=readonly"`
	TargetLbVserver               string  `json:"targetlbvserver,omitempty" nitro:"permission=readonly"`
	TargetType                    string  `json:"targettype,omitempty" nitro:"permission=readwrite"`
	TargetVserver                 string  `json:"targetvserver,omitempty" nitro:"permission=readonly"`
	TcpProbePort                  int     `json:"tcpprobeport,omitempty" nitro:"permission=readwrite"`
	TcpProfileName                string  `json:"tcpprofilename,omitempty" nitro:"permission=readwrite"`
	TicksSinceLastStateChange     string  `json:"tickssincelaststatechange,omitempty" nitro:"permission=readonly"`
	Timeout                       float64 `json:"timeout,omitempty" nitro:"permission=readwrite"`
	TimeToLive                    float64 `json:"ttl,omitempty" nitro:"permission=readwrite"`
	TrafficDomain                 string  `json:"td,omitempty" nitro:"permission=readwrite"`
	Type                          string  `json:"type,omitempty" nitro:"permission=readonly"`
	Url                           string  `json:"url,omitempty" nitro:"permission=readonly"`
	Value                         string  `json:"value,omitempty" nitro:"permission=readonly"`
	Version                       int     `json:"version,omitempty" nitro:"permission=readonly"`
	VipHeader                     string  `json:"vipheader,omitempty" nitro:"permission=readwrite"`
	Weight                        string  `json:"weight,omitempty" nitro:"permission=readonly"`
}

func (r CsVserver) GetTypeName() string {
	return "csvserver"
}

func (r CsVserver) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("resource_type", r.GetTypeName()),
		slog.String("name", r.Name),
	)
}
