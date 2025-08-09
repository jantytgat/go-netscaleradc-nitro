package config

var LbVserverFieldNames = struct {
	ActiveServices                     string
	AdfsProxyProfile                   string
	AppFlowLog                         string
	Authentication                     string
	Authentication401                  string
	AuthenticationHost                 string
	AuthenticationProfile              string
	AuthenticationVserverName          string
	BackupLbMethod                     string
	BackupPersistenceTimeout           string
	BackupVserver                      string
	BackupVserverStatus                string
	Bindpoint                          string
	ByPassAaaaQueries                  string
	Cacheable                          string
	CacheType                          string
	CacheVserver                       string
	ClientTimeout                      string
	Comment                            string
	ConnectionFailover                 string
	ConsolidatedLeastConnectionGlobal  string
	ConsolidatedLeastConnectionStats   string
	CookieDomain                       string
	CookieName                         string
	Count                              string
	CurrentActiveOrder                 string
	CurrentState                       string
	DatabaseProfileName                string
	DatabaseSpecificLb                 string
	DataLength                         string
	DataOffset                         string
	DisablePrimaryOnDown               string
	Dns64                              string
	DnsOverHttps                       string
	DnsProfileName                     string
	DnsVserverName                     string
	Domain                             string
	DownstateFlush                     string
	DynamicWeight                      string
	EffectiveState                     string
	GreaterThan2GBTransactions         string
	GroupName                          string
	HashLength                         string
	Health                             string
	HealthThreshold                    string
	Hits                               string
	Homepage                           string
	HttpProfileName                    string
	HttpRedirectUrl                    string
	IcmpVserverResponse                string
	InsertVserverIpPort                string
	Invoke                             string
	IpMapping                          string
	IpMask                             string
	IpPattern                          string
	IpSet                              string
	Ipv46                              string
	Ipv6NetmaskLength                  string
	Ipv6PersistenceMaskLength          string
	IsGslb                             string
	L2ConnectionParamenters            string
	LbMethod                           string
	LbProfileName                      string
	LbRrReason                         string
	ListenPolicy                       string
	ListenPriority                     string
	MacModeRetainVlan                  string
	Map                                string
	MaximumAutoscaleMembers            string
	MinimumAutoscaleMembers            string
	MssqlServerVersion                 string
	MysqlCharacterSet                  string
	MySqlProtocalVersion               string
	MysqlServerVersion                 string
	Name                               string
	Netmask                            string
	NetProfile                         string
	NewName                            string
	NewServiceRequest                  string
	NewServiceRequestIncrementInterval string
	NewServiceRequestUnit              string
	NoDefaultBindings                  string
	NodegroupName                      string
	OracleServerVersion                string
	OrderThreshold                     string
	PersistAvpNumber                   string
	PersistenceBackup                  string
	PersistenceMask                    string
	PersistenceRule                    string
	PersistenceType                    string
	PiPolicyHits                       string
	PolicySubType                      string
	Port                               string
	// Precedence                          string
	Priority                            string
	ProbePort                           string
	ProbeProtocol                       string
	ProbeSuccessResponseCode            string
	ProcessLocal                        string
	Push                                string
	PushLabel                           string
	PushMultipleClientConnections       string
	PushVserver                         string
	QuicBridgeProfileName               string
	QuicProfileName                     string
	Range                               string
	RecursionAvailable                  string
	Redirect                            string
	RedirectFromPort                    string
	RedirectionMode                     string
	RedirectPortRewrite                 string
	RedirectUrl                         string
	RedirectUrlFlags                    string
	RetainConnectionsOnCluster          string
	RouthHealthInjectionState           string
	RtspNat                             string
	Rule                                string
	RuleType                            string
	ServiceName                         string
	ServiceType                         string
	Sessionless                         string
	SkipPersistency                     string
	SpilloverBackupAction               string
	SpilloverDynamicConnectionThreshold string
	SpilloverMethod                     string
	SpilloverPersistence                string
	SpilloverPersistenceTimeout         string
	SpilloverThreshold                  string
	State                               string
	StateChangeTimeMilliSeconds         string
	StateChangeTimeSec                  string
	StateChangeTimeSeconds              string
	Status                              string
	TcpProbePort                        string
	TcpProfileName                      string
	ThresholdValue                      string
	TicksSinceLastStateChange           string
	Timeout                             string
	ToggleOrder                         string
	TosId                               string
	TotalServices                       string
	TrafficDomain                       string
	TrofsPersistence                    string
	Type                                string
	Version                             string
	VipHeader                           string
	VsvrBindSvcIp                       string
	VsvrBindSvcPort                     string
	Weight                              string
}{
	ActiveServices:                     "activeservices",
	AdfsProxyProfile:                   "adfsproxyprofile",
	AppFlowLog:                         "appflowlog",
	Authentication:                     "authentication",
	Authentication401:                  "authn401",
	AuthenticationHost:                 "authenticationhost",
	AuthenticationProfile:              "authnprofile",
	AuthenticationVserverName:          "authnvsname",
	BackupLbMethod:                     "backuplbmethod",
	BackupPersistenceTimeout:           "backuppersistencetimeout",
	BackupVserver:                      "backupvserver",
	BackupVserverStatus:                "backupvserverstatus",
	Bindpoint:                          "bindpoint",
	ByPassAaaaQueries:                  "bypassaaaa",
	Cacheable:                          "cacheable",
	CacheType:                          "cachetype",
	CacheVserver:                       "cachevserver",
	ClientTimeout:                      "clttimeout",
	Comment:                            "comment",
	ConnectionFailover:                 "connfailover",
	ConsolidatedLeastConnectionGlobal:  "consolidatedlconngbl",
	ConsolidatedLeastConnectionStats:   "consolidatedlconn",
	CookieDomain:                       "cookiedomain",
	CookieName:                         "cookiename",
	Count:                              "__count",
	CurrentActiveOrder:                 "currentactiveorder",
	CurrentState:                       "curstate",
	DatabaseProfileName:                "dbprofilename",
	DatabaseSpecificLb:                 "dbslb",
	DataLength:                         "datalength",
	DataOffset:                         "dataoffset",
	DisablePrimaryOnDown:               "disableprimaryondown",
	Dns64:                              "dns64",
	DnsOverHttps:                       "dnsoverhttps",
	DnsProfileName:                     "dnsprofilename",
	DnsVserverName:                     "dnsvservername",
	Domain:                             "domain",
	DownstateFlush:                     "downstateflush",
	DynamicWeight:                      "dynamicweight",
	EffectiveState:                     "effectivestate",
	GreaterThan2GBTransactions:         "gt2gb",
	GroupName:                          "groupname",
	HashLength:                         "hashlength",
	Health:                             "health",
	HealthThreshold:                    "healththreshold",
	Hits:                               "hits",
	Homepage:                           "homepage",
	HttpProfileName:                    "httpprofilename",
	HttpRedirectUrl:                    "httpsredirecturl",
	IcmpVserverResponse:                "icmpvsrresponse",
	InsertVserverIpPort:                "insertvserveripport",
	Invoke:                             "invoke",
	IpMapping:                          "ipmapping",
	IpMask:                             "ipmask",
	IpPattern:                          "ippattern",
	IpSet:                              "ipset",
	Ipv46:                              "ipv46",
	Ipv6NetmaskLength:                  "v6netmasklen",
	Ipv6PersistenceMaskLength:          "v6persistmasklen",
	IsGslb:                             "isgslb",
	L2ConnectionParamenters:            "l2conn",
	LbMethod:                           "lbmethod",
	LbProfileName:                      "lbprofilename",
	LbRrReason:                         "lbrrreason",
	ListenPolicy:                       "listenpolicy",
	ListenPriority:                     "listenpriority",
	MacModeRetainVlan:                  "macmoderetainvlan",
	Map:                                "map",
	MaximumAutoscaleMembers:            "maxautoscalemembers",
	MinimumAutoscaleMembers:            "minautoscalemembers",
	MssqlServerVersion:                 "mssqlserverversion",
	MysqlCharacterSet:                  "mysqlcharacterset",
	MySqlProtocalVersion:               "mysqlprotocolversion",
	MysqlServerVersion:                 "mysqlserverversion",
	Name:                               "name",
	Netmask:                            "netmask",
	NetProfile:                         "netprofile",
	NewName:                            "newname",
	NewServiceRequest:                  "newservicerequest",
	NewServiceRequestIncrementInterval: "newservicerequestincrementinterval",
	NewServiceRequestUnit:              "newservicerequestunit",
	NoDefaultBindings:                  "nodefaultbindings",
	NodegroupName:                      "ngname",
	OracleServerVersion:                "oracleserverversion",
	OrderThreshold:                     "orderthreshold",
	PersistAvpNumber:                   "persistavpno",
	PersistenceBackup:                  "persistencebackup",
	PersistenceMask:                    "persistmask",
	PersistenceRule:                    "resrule",
	PersistenceType:                    "persistencetype",
	PiPolicyHits:                       "pipolicyhits",
	PolicySubType:                      "policysubtype",
	Port:                               "port",
	// Precedence:                          "precedence",
	Priority:                            "priority",
	ProbePort:                           "probeport",
	ProbeProtocol:                       "probeprotocol",
	ProbeSuccessResponseCode:            "probesuccessresponsecode",
	ProcessLocal:                        "processlocal",
	Push:                                "push",
	PushLabel:                           "pushlabel",
	PushMultipleClientConnections:       "pushmulticlients",
	PushVserver:                         "pushvserver",
	QuicBridgeProfileName:               "quicbridgeprofilename",
	QuicProfileName:                     "quicprofilename",
	Range:                               "range",
	RecursionAvailable:                  "recursionavailable",
	Redirect:                            "redirect",
	RedirectFromPort:                    "redirectfromport",
	RedirectionMode:                     "m",
	RedirectPortRewrite:                 "redirectportrewrite",
	RedirectUrl:                         "redirurl",
	RedirectUrlFlags:                    "redirurlflags",
	RetainConnectionsOnCluster:          "retainconnectionsoncluster",
	RouthHealthInjectionState:           "rhistate",
	RtspNat:                             "rtspnat",
	Rule:                                "rule",
	RuleType:                            "ruletype",
	ServiceName:                         "servicename",
	ServiceType:                         "servicetype",
	Sessionless:                         "sessionless",
	SkipPersistency:                     "skippersistency",
	SpilloverBackupAction:               "sobackupaction",
	SpilloverDynamicConnectionThreshold: "vsvrdynconnsothreshold",
	SpilloverMethod:                     "somethod",
	SpilloverPersistence:                "sopersistence",
	SpilloverPersistenceTimeout:         "sopersistencetimeout",
	SpilloverThreshold:                  "sothreshold",
	State:                               "state",
	StateChangeTimeMilliSeconds:         "statechangetimemsec",
	StateChangeTimeSec:                  "statechangetimesec",
	StateChangeTimeSeconds:              "statechangetimeseconds",
	Status:                              "status",
	TcpProbePort:                        "tcpprobeport",
	TcpProfileName:                      "tcpprofilename",
	ThresholdValue:                      "thresholdvalue",
	TicksSinceLastStateChange:           "tickssincelaststatechange",
	Timeout:                             "timeout",
	ToggleOrder:                         "toggleorder",
	TosId:                               "tosid",
	TotalServices:                       "totalservices",
	TrafficDomain:                       "td",
	TrofsPersistence:                    "trofspersistence",
	Type:                                "type",
	Version:                             "version",
	VipHeader:                           "vipheader",
	VsvrBindSvcIp:                       "vsvrbindsvcip",
	VsvrBindSvcPort:                     "vsvrbindsvcport",
	Weight:                              "weight",
}

type LbVserver struct {
	ActiveServices                     string    `json:"activeservices,omitempty" nitro:"permission=readonly"`
	AdfsProxyProfile                   string    `json:"adfsproxyprofile,omitempty" nitro:"permission=readwrite"`
	AppFlowLog                         string    `json:"appflowlog,omitempty" nitro:"permission=readwrite"`
	Authentication                     string    `json:"authentication,omitempty" nitro:"permission=readwrite"`
	Authentication401                  string    `json:"authn401,omitempty" nitro:"permission=readwrite"`
	AuthenticationHost                 string    `json:"authenticationhost,omitempty" nitro:"permission=readwrite"`
	AuthenticationProfile              string    `json:"authnprofile,omitempty" nitro:"permission=readwrite"`
	AuthenticationVserverName          string    `json:"authnvsname,omitempty" nitro:"permission=readwrite"`
	BackupLbMethod                     string    `json:"backuplbmethod,omitempty" nitro:"permission=readwrite"`
	BackupPersistenceTimeout           float64   `json:"backuppersistencetimeout,omitempty" nitro:"permission=readwrite"`
	BackupVserver                      string    `json:"backupvserver,omitempty" nitro:"permission=readwrite"`
	BackupVserverStatus                string    `json:"backupvserverstatus,omitempty" nitro:"permission=readonly"`
	Bindpoint                          string    `json:"bindpoint,omitempty" nitro:"permission=readonly"`
	ByPassAaaaQueries                  string    `json:"bypassaaaa,omitempty" nitro:"permission=readwrite"`
	Cacheable                          string    `json:"cacheable,omitempty" nitro:"permission=readwrite"`
	CacheType                          string    `json:"cachetype,omitempty" nitro:"permission=readonly"`
	CacheVserver                       string    `json:"cachevserver,omitempty" nitro:"permission=readonly"`
	ClientTimeout                      string    `json:"clttimeout,omitempty" nitro:"permission=readwrite"`
	Comment                            string    `json:"comment,omitempty" nitro:"permission=readwrite"`
	ConnectionFailover                 string    `json:"connfailover,omitempty" nitro:"permission=readwrite"`
	ConsolidatedLeastConnectionGlobal  string    `json:"consolidatedlconngbl,omitempty" nitro:"permission=readonly"`
	ConsolidatedLeastConnectionStats   string    `json:"consolidatedlconn,omitempty" nitro:"permission=readonly"`
	CookieDomain                       string    `json:"cookiedomain,omitempty" nitro:"permission=readonly"`
	CookieName                         string    `json:"cookiename,omitempty" nitro:"permission=readwrite"`
	Count                              float64   `json:"__count,omitempty" nitro:"permission=readonly"`
	CurrentActiveOrder                 string    `json:"currentactiveorder,omitempty" nitro:"permission=readonly"`
	CurrentState                       string    `json:"curstate,omitempty" nitro:"permission=readonly"`
	DatabaseProfileName                string    `json:"dbprofilename,omitempty" nitro:"permission=readwrite"`
	DatabaseSpecificLb                 string    `json:"dbslb,omitempty" nitro:"permission=readwrite"`
	DataLength                         string    `json:"datalength,omitempty" nitro:"permission=readwrite"`
	DataOffset                         string    `json:"dataoffset,omitempty" nitro:"permission=readwrite"`
	DisablePrimaryOnDown               string    `json:"disableprimaryondown,omitempty" nitro:"permission=readwrite"`
	Dns64                              string    `json:"dns64,omitempty" nitro:"permission=readwrite"`
	DnsOverHttps                       string    `json:"dnsoverhttps,omitempty" nitro:"permission=readwrite"`
	DnsProfileName                     string    `json:"dnsprofilename,omitempty" nitro:"permission=readwrite"`
	DnsVserverName                     string    `json:"dnsvservername,omitempty" nitro:"permission=readonly"`
	Domain                             string    `json:"domain,omitempty" nitro:"permission=readonly"`
	DownstateFlush                     string    `json:"downstateflush,omitempty" nitro:"permission=readwrite"`
	DynamicWeight                      string    `json:"dynamicweight,omitempty" nitro:"permission=readonly"`
	EffectiveState                     string    `json:"effectivestate,omitempty" nitro:"permission=readonly"`
	GreaterThan2GBTransactions         string    `json:"gt2gb,omitempty" nitro:"permission=readonly"`
	GroupName                          string    `json:"groupname,omitempty" nitro:"permission=readonly"`
	HashLength                         float64   `json:"hashlength,omitempty" nitro:"permission=readwrite"`
	Health                             string    `json:"health,omitempty" nitro:"permission=readonly"`
	HealthThreshold                    string    `json:"healththreshold,omitempty" nitro:"permission=readwrite"`
	Hits                               string    `json:"hits,omitempty" nitro:"permission=readonly"`
	Homepage                           string    `json:"homepage,omitempty" nitro:"permission=readonly"`
	HttpProfileName                    string    `json:"httpprofilename,omitempty" nitro:"permission=readwrite"`
	HttpRedirectUrl                    string    `json:"httpsredirecturl,omitempty" nitro:"permission=readwrite"`
	IcmpVserverResponse                string    `json:"icmpvsrresponse,omitempty" nitro:"permission=readwrite"`
	InsertVserverIpPort                string    `json:"insertvserveripport,omitempty" nitro:"permission=readwrite"`
	Invoke                             bool      `json:"invoke,omitempty" nitro:"permission=readwrite"`
	IpMapping                          string    `json:"ipmapping,omitempty" nitro:"permission=readonly"`
	IpMask                             string    `json:"ipmask,omitempty" nitro:"permission=readwrite"`
	IpPattern                          string    `json:"ippattern,omitempty" nitro:"permission=readwrite"`
	IpSet                              string    `json:"ipset,omitempty" nitro:"permission=readwrite"`
	Ipv46                              string    `json:"ipv46,omitempty" nitro:"permission=readwrite"`
	Ipv6NetmaskLength                  float64   `json:"v6netmasklen,omitempty" nitro:"permission=readwrite"`
	Ipv6PersistenceMaskLength          string    `json:"v6persistmasklen,omitempty" nitro:"permission=readwrite"`
	IsGslb                             bool      `json:"isgslb,omitempty" nitro:"permission=readonly"`
	L2ConnectionParamenters            string    `json:"l2conn,omitempty" nitro:"permission=readwrite"`
	LbMethod                           string    `json:"lbmethod,omitempty" nitro:"permission=readwrite"`
	LbProfileName                      string    `json:"lbprofilename,omitempty" nitro:"permission=readwrite"`
	LbRrReason                         int       `json:"lbrrreason,omitempty" nitro:"permission=readonly"`
	ListenPolicy                       string    `json:"listenpolicy,omitempty" nitro:"permission=readwrite"`
	ListenPriority                     float64   `json:"listenpriority,omitempty" nitro:"permission=readwrite"`
	MacModeRetainVlan                  string    `json:"macmoderetainvlan,omitempty" nitro:"permission=readwrite"`
	Map                                string    `json:"map,omitempty" nitro:"permission=readonly"`
	MaximumAutoscaleMembers            string    `json:"maxautoscalemembers,omitempty" nitro:"permission=readwrite"`
	MinimumAutoscaleMembers            string    `json:"minautoscalemembers,omitempty" nitro:"permission=readwrite"`
	MssqlServerVersion                 string    `json:"mssqlserverversion,omitempty" nitro:"permission=readwrite"`
	MysqlCharacterSet                  string    `json:"mysqlcharacterset,omitempty" nitro:"permission=readwrite"`
	MySqlProtocalVersion               string    `json:"mysqlprotocolversion,omitempty" nitro:"permission=readwrite"`
	MysqlServerVersion                 string    `json:"mysqlserverversion,omitempty" nitro:"permission=readwrite"`
	Name                               string    `json:"name,omitempty" nitro:"permission=readwrite"`
	Netmask                            string    `json:"netmask,omitempty" nitro:"permission=readwrite"`
	NetProfile                         string    `json:"netprofile,omitempty" nitro:"permission=readwrite"`
	NewName                            string    `json:"newname,omitempty" nitro:"permission=readwrite"`
	NewServiceRequest                  float64   `json:"newservicerequest,omitempty" nitro:"permission=readwrite"`
	NewServiceRequestIncrementInterval float64   `json:"newservicerequestincrementinterval,omitempty" nitro:"permission=readwrite"`
	NewServiceRequestUnit              string    `json:"newservicerequestunit,omitempty" nitro:"permission=readwrite"`
	NoDefaultBindings                  string    `json:"nodefaultbindings,omitempty" nitro:"permission=readonly"`
	NodegroupName                      string    `json:"ngname,omitempty" nitro:"permission=readonly"`
	OracleServerVersion                string    `json:"oracleserverversion,omitempty" nitro:"permission=readwrite"`
	OrderThreshold                     string    `json:"orderthreshold,omitempty" nitro:"permission=readonly"`
	PersistAvpNumber                   []float64 `json:"persistavpno,omitempty" nitro:"permission=readwrite"`
	PersistenceBackup                  string    `json:"persistencebackup,omitempty" nitro:"permission=readwrite"`
	PersistenceMask                    string    `json:"persistmask,omitempty" nitro:"permission=readwrite"`
	PersistenceRule                    string    `json:"resrule,omitempty" nitro:"permission=readwrite"`
	PersistenceType                    string    `json:"persistencetype,omitempty" nitro:"permission=readwrite"`
	PiPolicyHits                       string    `json:"pipolicyhits,omitempty" nitro:"permission=readonly"`
	PolicySubType                      string    `json:"policysubtype,omitempty" nitro:"permission=readonly"`
	Port                               int       `json:"port" nitro:"permission=readwrite"`
	// Precedence                          string    `json:"precedence,omitempty" nitro:"permission=readonly"`
	Priority                            string  `json:"priority,omitempty" nitro:"permission=readwrite"`
	ProbePort                           int     `json:"probeport,omitempty" nitro:"permission=readwrite"`
	ProbeProtocol                       string  `json:"probeprotocol,omitempty" nitro:"permission=readwrite"`
	ProbeSuccessResponseCode            string  `json:"probesuccessresponsecode,omitempty" nitro:"permission=readwrite"`
	ProcessLocal                        string  `json:"processlocal,omitempty" nitro:"permission=readwrite"`
	Push                                string  `json:"push,omitempty" nitro:"permission=readwrite"`
	PushLabel                           string  `json:"pushlabel,omitempty" nitro:"permission=readwrite"`
	PushMultipleClientConnections       string  `json:"pushmulticlients,omitempty" nitro:"permission=readwrite"`
	PushVserver                         string  `json:"pushvserver,omitempty" nitro:"permission=readwrite"`
	QuicBridgeProfileName               string  `json:"quicbridgeprofilename,omitempty" nitro:"permission=readwrite"`
	QuicProfileName                     string  `json:"quicprofilename,omitempty" nitro:"permission=readwrite"`
	Range                               string  `json:"range,omitempty" nitro:"permission=readwrite"`
	RecursionAvailable                  string  `json:"recursionavailable,omitempty" nitro:"permission=readwrite"`
	Redirect                            string  `json:"redirect,omitempty" nitro:"permission=readonly"`
	RedirectFromPort                    int     `json:"redirectfromport,omitempty" nitro:"permission=readwrite"`
	RedirectionMode                     string  `json:"m,omitempty" nitro:"permission=readwrite"`
	RedirectPortRewrite                 string  `json:"redirectportrewrite,omitempty" nitro:"permission=readwrite"`
	RedirectUrl                         string  `json:"redirurl,omitempty" nitro:"permission=readwrite"`
	RedirectUrlFlags                    bool    `json:"redirurlflags,omitempty" nitro:"permission=readwrite"`
	RetainConnectionsOnCluster          string  `json:"retainconnectionsoncluster,omitempty" nitro:"permission=readwrite"`
	RouthHealthInjectionState           string  `json:"rhistate,omitempty" nitro:"permission=readwrite"`
	RtspNat                             string  `json:"rtspnat,omitempty" nitro:"permission=readwrite"`
	Rule                                string  `json:"rule,omitempty" nitro:"permission=readwrite"`
	RuleType                            string  `json:"ruletype,omitempty" nitro:"permission=readonly"`
	ServiceName                         string  `json:"servicename,omitempty" nitro:"permission=readwrite"`
	ServiceType                         string  `json:"servicetype,omitempty" nitro:"permission=readwrite"`
	Sessionless                         string  `json:"sessionless,omitempty" nitro:"permission=readwrite"`
	SkipPersistency                     string  `json:"skippersistency,omitempty" nitro:"permission=readwrite"`
	SpilloverBackupAction               string  `json:"sobackupaction,omitempty" nitro:"permission=readwrite"`
	SpilloverDynamicConnectionThreshold string  `json:"vsvrdynconnsothreshold,omitempty" nitro:"permission=readonly"`
	SpilloverMethod                     string  `json:"somethod,omitempty" nitro:"permission=readwrite"`
	SpilloverPersistence                string  `json:"sopersistence,omitempty" nitro:"permission=readwrite"`
	SpilloverPersistenceTimeout         string  `json:"sopersistencetimeout,omitempty" nitro:"permission=readwrite"`
	SpilloverThreshold                  float64 `json:"sothreshold,omitempty" nitro:"permission=readwrite"`
	State                               string  `json:"state,omitempty" nitro:"permission=readwrite"`
	StateChangeTimeMilliSeconds         string  `json:"statechangetimemsec,omitempty" nitro:"permission=readonly"`
	StateChangeTimeSec                  string  `json:"statechangetimesec,omitempty" nitro:"permission=readonly"`
	StateChangeTimeSeconds              string  `json:"statechangetimeseconds,omitempty" nitro:"permission=readonly"`
	Status                              int     `json:"status,omitempty" nitro:"permission=readonly"`
	TcpProbePort                        int     `json:"tcpprobeport,omitempty" nitro:"permission=readwrite"`
	TcpProfileName                      string  `json:"tcpprofilename,omitempty" nitro:"permission=readwrite"`
	ThresholdValue                      int     `json:"thresholdvalue,omitempty" nitro:"permission=readonly"`
	TicksSinceLastStateChange           string  `json:"tickssincelaststatechange,omitempty" nitro:"permission=readonly"`
	Timeout                             float64 `json:"timeout,omitempty" nitro:"permission=readwrite"`
	ToggleOrder                         string  `json:"toggleorder,omitempty" nitro:"permission=readonly"`
	TosId                               string  `json:"tosid,omitempty" nitro:"permission=readwrite"`
	TotalServices                       string  `json:"totalservices,omitempty" nitro:"permission=readonly"`
	TrafficDomain                       string  `json:"td,omitempty" nitro:"permission=readwrite"`
	TrofsPersistence                    string  `json:"trofspersistence,omitempty" nitro:"permission=readwrite"`
	Type                                string  `json:"type,omitempty" nitro:"permission=readonly"`
	Version                             int     `json:"version,omitempty" nitro:"permission=readonly"`
	VipHeader                           string  `json:"vipheader,omitempty" nitro:"permission=readwrite"`
	VsvrBindSvcIp                       string  `json:"vsvrbindsvcip,omitempty" nitro:"permission=readonly"`
	VsvrBindSvcPort                     int     `json:"vsvrbindsvcport,omitempty" nitro:"permission=readonly"`
	Weight                              float64 `json:"weight,omitempty" nitro:"permission=readwrite"`
}

func (r LbVserver) GetTypeName() string {
	return "lbvserver"
}
