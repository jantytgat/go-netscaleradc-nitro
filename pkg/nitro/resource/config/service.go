package config

import "log/slog"

var ServiceFieldNames = struct {
	AccessDown                   string
	All                          string
	AppFlowLogging               string
	Builtin                      string
	Cacheable                    string
	CacheType                    string
	ClearTextPort                string
	ClientIpInsertion            string
	ClientIpInsertionHeader      string
	ClientKeepAlive              string
	ClientMonitorOwner           string
	ClientMonitorView            string
	ClientTimeout                string
	Comment                      string
	Compression                  string
	ContentInspectionProfileName string
	Count                        string
	CustomServerId               string
	DbsServiceIp                 string
	Delay                        string
	DnsProfileName               string
	DownStateFlush               string
	DupState                     string
	DupWeight                    string
	FailedProbes                 string
	Feature                      string
	Graceful                     string
	GslbOption                   string
	HashId                       string
	HealthMonitor                string
	HttpProfileName              string
	Internal                     string
	IpAddress                    string
	MaxBandwidth                 string
	MaxClient                    string
	MaxRequests                  string
	MonitorConnectionClose       string
	MonitorCurrentFailedProbes   string
	MonitorLastResponse          string
	MonitorResponseTime          string
	MonitorState                 string
	MonitorStatusCode            string
	MonitorThreshold             string
	MonitorTotalFailedProbes     string
	MonitorTotalProbes           string
	Monstatparam1                string
	Monstatparam2                string
	Monstatparam3                string
	Name                         string
	NetProfile                   string
	NewName                      string
	NoDefaultBindings            string
	NumberOfConnections          string
	OracleServerVersion          string
	Passive                      string
	PathMonitor                  string
	PathMonitorIndividual        string
	PolicyName                   string
	Port                         string
	ProcessLocal                 string
	PublicIp                     string
	PublicPort                   string
	RtspSessionIdMapping         string
	ServerId                     string
	ServerName                   string
	ServerTimeout                string
	ServiceConfigurationType     string
	ServiceConfigurationType2    string
	ServiceState                 string
	ServiceType                  string
	SslStatus                    string
	StateChangeTimeMilliseconds  string
	StateChangeTimeSeconds       string
	StateUpdateReason            string
	SurgeProtection              string
	TcpBuffering                 string
	TcpProfileName               string
	TicksSinceLastStateChange    string
	TimeSinceLastStateChange     string
	TotalFailedProbes            string
	TotalProbes                  string
	TrafficDomain                string
	UseProxyPort                 string
	UserMonitorStatusMessage     string
	Usip                         string
	Weight                       string
}{
	AccessDown:                   "accessdown",
	All:                          "all",
	AppFlowLogging:               "appflowlog",
	Builtin:                      "builtin",
	Cacheable:                    "cacheable",
	CacheType:                    "cachetype",
	ClearTextPort:                "cleartextport",
	ClientIpInsertion:            "cip",
	ClientIpInsertionHeader:      "cipheader",
	ClientKeepAlive:              "cka",
	ClientMonitorOwner:           "clmonowner",
	ClientMonitorView:            "clmonview",
	ClientTimeout:                "clttimeout",
	Comment:                      "comment",
	Compression:                  "cmp",
	ContentInspectionProfileName: "contentinspectionprofilename",
	Count:                        "__count",
	CustomServerId:               "customserverid",
	DbsServiceIp:                 "serviceipstr",
	Delay:                        "delay",
	DnsProfileName:               "dnsprofilename",
	DownStateFlush:               "downstateflush",
	DupState:                     "dup_state",
	DupWeight:                    "dup_weight",
	FailedProbes:                 "failedprobes",
	Feature:                      "feature",
	Graceful:                     "graceful",
	GslbOption:                   "gslb",
	HashId:                       "hashid",
	HealthMonitor:                "healthmonitor",
	HttpProfileName:              "httpprofilename",
	Internal:                     "Internal",
	IpAddress:                    "ipaddress",
	MaxBandwidth:                 "maxbandwidth",
	MaxClient:                    "maxclient",
	MaxRequests:                  "maxreq",
	MonitorConnectionClose:       "monconnectionclose",
	MonitorCurrentFailedProbes:   "monitorcurrentfailedprobes",
	MonitorLastResponse:          "lastresponse",
	MonitorResponseTime:          "responsetime",
	MonitorState:                 "monitor_state",
	MonitorStatusCode:            "monstatcode",
	MonitorThreshold:             "monthreshold",
	MonitorTotalFailedProbes:     "monitortotalfailedprobes",
	MonitorTotalProbes:           "monitortotalprobes",
	Monstatparam1:                "monstatparam1",
	Monstatparam2:                "monstatparam2",
	Monstatparam3:                "monstatparam3",
	Name:                         "name",
	NetProfile:                   "netprofile",
	NewName:                      "newname",
	NoDefaultBindings:            "nodefaultbindings",
	NumberOfConnections:          "numofconnections",
	OracleServerVersion:          "oracleserverversion",
	Passive:                      "passive",
	PathMonitor:                  "pathmonitor",
	PathMonitorIndividual:        "pathmonitorindv",
	PolicyName:                   "policyname",
	Port:                         "port",
	ProcessLocal:                 "processlocal",
	PublicIp:                     "publicip",
	PublicPort:                   "publicport",
	RtspSessionIdMapping:         "rtspsessionidremap",
	ServerId:                     "serverid",
	ServerName:                   "servername",
	ServerTimeout:                "svrtimeout",
	ServiceConfigurationType:     "serviceconftype",
	ServiceConfigurationType2:    "serviceconftype2",
	ServiceState:                 "svrstate",
	ServiceType:                  "servicetype",
	SslStatus:                    "value",
	StateChangeTimeMilliseconds:  "statechangetimemsec",
	StateChangeTimeSeconds:       "statechangetimesec",
	StateUpdateReason:            "stateupdatereason",
	SurgeProtection:              "sp",
	TcpBuffering:                 "tcpb",
	TcpProfileName:               "tcpprofilename",
	TicksSinceLastStateChange:    "tickssincelaststatechange",
	TimeSinceLastStateChange:     "timesincelaststatechange",
	TotalFailedProbes:            "totalfailedprobes",
	TotalProbes:                  "totalprobes",
	TrafficDomain:                "td",
	UseProxyPort:                 "useproxyport",
	UserMonitorStatusMessage:     "monuserstatusmesg",
	Usip:                         "usip",
	Weight:                       "weight",
}

type Service struct {
	AccessDown                   string   `json:"accessdown,omitempty" nitro:"permission=readwrite"`
	All                          bool     `json:"all,omitempty" nitro:"permission=readwrite"`
	AppFlowLogging               string   `json:"appflowlog,omitempty" nitro:"permission=readwrite"`
	Builtin                      []string `json:"builtin,omitempty" nitro:"permission=readonly"`
	Cacheable                    string   `json:"cacheable,omitempty" nitro:"permission=readwrite"`
	CacheType                    string   `json:"cachetype,omitempty" nitro:"permission=readwrite"`
	ClearTextPort                int      `json:"cleartextport,omitempty" nitro:"permission=readwrite"`
	ClientIpInsertion            string   `json:"cip,omitempty" nitro:"permission=readwrite"`
	ClientIpInsertionHeader      string   `json:"cipheader,omitempty" nitro:"permission=readwrite"`
	ClientKeepAlive              string   `json:"cka,omitempty" nitro:"permission=readwrite"`
	ClientMonitorOwner           string   `json:"clmonowner,omitempty" nitro:"permission=readonly"`
	ClientMonitorView            string   `json:"clmonview,omitempty" nitro:"permission=readonly"`
	ClientTimeout                float64  `json:"clttimeout,omitempty" nitro:"permission=readwrite"`
	Comment                      string   `json:"comment,omitempty" nitro:"permission=readwrite"`
	Compression                  string   `json:"cmp,omitempty" nitro:"permission=readwrite"`
	ContentInspectionProfileName string   `json:"contentinspectionprofilename,omitempty" nitro:"permission=readwrite"`
	Count                        float64  `json:"__count,omitempty" nitro:"permission=readonly"`
	CustomServerId               string   `json:"customserverid,omitempty" nitro:"permission=readwrite"`
	DbsServiceIp                 string   `json:"serviceipstr,omitempty" nitro:"permission=readonly"`
	Delay                        float64  `json:"delay,omitempty" nitro:"permission=readwrite"`
	DnsProfileName               string   `json:"dnsprofilename,omitempty" nitro:"permission=readwrite"`
	DownStateFlush               string   `json:"downstateflush,omitempty" nitro:"permission=readwrite"`
	DupState                     string   `json:"dup_state,omitempty" nitro:"permission=readonly"`
	DupWeight                    string   `json:"dup_weight,omitempty" nitro:"permission=readonly"`
	FailedProbes                 string   `json:"failedprobes,omitempty" nitro:"permission=readonly"`
	Feature                      string   `json:"feature,omitempty" nitro:"permission=readonly"`
	Graceful                     string   `json:"graceful,omitempty" nitro:"permission=readwrite"`
	GslbOption                   string   `json:"gslb,omitempty" nitro:"permission=readonly"`
	HashId                       float64  `json:"hashid,omitempty" nitro:"permission=readwrite"`
	HealthMonitor                string   `json:"healthmonitor,omitempty" nitro:"permission=readwrite"`
	HttpProfileName              string   `json:"httpprofilename,omitempty" nitro:"permission=readwrite"`
	Internal                     bool     `json:"Internal,omitempty" nitro:"permission=readwrite"`
	IpAddress                    string   `json:"ipaddress,omitempty" nitro:"permission=readwrite"`
	MaxBandwidth                 string   `json:"maxbandwidth,omitempty" nitro:"permission=readwrite"`
	MaxClient                    string   `json:"maxclient,omitempty" nitro:"permission=readwrite"`
	MaxRequests                  string   `json:"maxreq,omitempty" nitro:"permission=readwrite"`
	MonitorConnectionClose       string   `json:"monconnectionclose,omitempty" nitro:"permission=readwrite"`
	MonitorCurrentFailedProbes   string   `json:"monitorcurrentfailedprobes,omitempty" nitro:"permission=readonly"`
	MonitorLastResponse          string   `json:"lastresponse,omitempty" nitro:"permission=readonly"`
	MonitorResponseTime          string   `json:"responsetime,omitempty" nitro:"permission=readonly"`
	MonitorState                 string   `json:"monitor_state,omitempty" nitro:"permission=readonly"`
	MonitorStatusCode            int      `json:"monstatcode,omitempty" nitro:"permission=readonly"`
	MonitorThreshold             string   `json:"monthreshold,omitempty" nitro:"permission=readwrite"`
	MonitorTotalFailedProbes     string   `json:"monitortotalfailedprobes,omitempty" nitro:"permission=readonly"`
	MonitorTotalProbes           string   `json:"monitortotalprobes,omitempty" nitro:"permission=readonly"`
	Monstatparam1                int      `json:"monstatparam1,omitempty" nitro:"permission=readonly"`
	Monstatparam2                int      `json:"monstatparam2,omitempty" nitro:"permission=readonly"`
	Monstatparam3                int      `json:"monstatparam3,omitempty" nitro:"permission=readonly"`
	Name                         string   `json:"name,omitempty" nitro:"permission=readwrite"`
	NetProfile                   string   `json:"netprofile,omitempty" nitro:"permission=readwrite"`
	NewName                      string   `json:"newname,omitempty" nitro:"permission=readwrite"`
	NoDefaultBindings            string   `json:"nodefaultbindings,omitempty" nitro:"permission=readonly"`
	NumberOfConnections          int      `json:"numofconnections,omitempty" nitro:"permission=readonly"`
	OracleServerVersion          string   `json:"oracleserverversion,omitempty" nitro:"permission=readonly"`
	Passive                      bool     `json:"passive,omitempty" nitro:"permission=readonly"`
	PathMonitor                  string   `json:"pathmonitor,omitempty" nitro:"permission=readwrite"`
	PathMonitorIndividual        string   `json:"pathmonitorindv,omitempty" nitro:"permission=readwrite"`
	PolicyName                   string   `json:"policyname,omitempty" nitro:"permission=readonly"`
	Port                         int      `json:"port,omitempty" nitro:"permission=readwrite"`
	ProcessLocal                 string   `json:"processlocal,omitempty" nitro:"permission=readwrite"`
	PublicIp                     string   `json:"publicip,omitempty" nitro:"permission=readonly"`
	PublicPort                   int      `json:"publicport,omitempty" nitro:"permission=readonly"`
	RtspSessionIdMapping         string   `json:"rtspsessionidremap,omitempty" nitro:"permission=readwrite"`
	ServerId                     float64  `json:"serverid,omitempty" nitro:"permission=readwrite"`
	ServerName                   string   `json:"servername,omitempty" nitro:"permission=readwrite"`
	ServerTimeout                float64  `json:"svrtimeout,omitempty" nitro:"permission=readwrite"`
	ServiceConfigurationType     bool     `json:"serviceconftype,omitempty" nitro:"permission=readwrite"`
	ServiceConfigurationType2    string   `json:"serviceconftype2,omitempty" nitro:"permission=readwrite"`
	ServiceState                 string   `json:"svrstate,omitempty" nitro:"permission=readonly"`
	ServiceType                  string   `json:"servicetype,omitempty" nitro:"permission=readwrite"`
	SslStatus                    string   `json:"value,omitempty" nitro:"permission=readonly"`
	StateChangeTimeMilliseconds  string   `json:"statechangetimemsec,omitempty" nitro:"permission=readonly"`
	StateChangeTimeSeconds       string   `json:"statechangetimesec,omitempty" nitro:"permission=readonly"`
	StateUpdateReason            string   `json:"stateupdatereason,omitempty" nitro:"permission=readonly"`
	SurgeProtection              string   `json:"sp,omitempty" nitro:"permission=readwrite"`
	TcpBuffering                 string   `json:"tcpb,omitempty" nitro:"permission=readwrite"`
	TcpProfileName               string   `json:"tcpprofilename,omitempty" nitro:"permission=readwrite"`
	TicksSinceLastStateChange    string   `json:"tickssincelaststatechange,omitempty" nitro:"permission=readonly"`
	TimeSinceLastStateChange     float64  `json:"timesincelaststatechange,omitempty" nitro:"permission=readonly"`
	TotalFailedProbes            string   `json:"totalfailedprobes,omitempty" nitro:"permission=readonly"`
	TotalProbes                  string   `json:"totalprobes,omitempty" nitro:"permission=readonly"`
	TrafficDomain                string   `json:"td,omitempty" nitro:"permission=readwrite"`
	UseProxyPort                 string   `json:"useproxyport,omitempty" nitro:"permission=readwrite"`
	UserMonitorStatusMessage     string   `json:"monuserstatusmesg,omitempty" nitro:"permission=readonly"`
	Usip                         string   `json:"usip,omitempty" nitro:"permission=readwrite"`
	Weight                       float64  `json:"weight,omitempty" nitro:"permission=readwrite"`
}

func (r Service) GetTypeName() string {
	return "service"
}

func (r Service) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("resource_type", r.GetTypeName()),
		slog.String("name", r.Name),
	)
}
