package config

import "log/slog"

var ServiceGroupFieldNames = struct {
	AppFlowLog                 string
	AutoDelayedTrofs           string
	AutoDisableDelay           string
	AutoDisableGraceful        string
	Autoscale                  string
	Cacheable                  string
	CacheType                  string
	ClientIpInsertion          string
	ClientIpInsertionHeader    string
	ClientKeepAlive            string
	ClientTimeout              string
	Clmonowner                 string
	Clmonview                  string
	Compression                string
	Comment                    string
	Count                      string
	CustomServerId             string
	DbsTtl                     string
	Delay                      string
	DownstateFlush             string
	Graceful                   string
	GroupCount                 string
	HashId                     string
	HealthMonitor              string
	HttpProfileName            string
	IncludeMembers             string
	Ip                         string
	MaxBandwidth               string
	MaxClient                  string
	MaxRequest                 string
	MemberPort                 string
	MonConnectionClose         string
	MonitorCurrentFailedProbes string
	MonitorState               string
	MonitorStatusCode          string
	MonitorThreshold           string
	MonitorTotalFailedProbes   string
	MonitorTotalProbes         string
	MonState                   string
	Monstatparam1              string
	Monstatparam2              string
	Monstatparam3              string
	Monuserstatusmesg          string
	MonWeight                  string
	Name                       string
	Nameserver                 string
	NetProfile                 string
	NewName                    string
	NoDefaultBindings          string
	NumberOfConnections        string
	Order                      string
	OrderString                string
	Passive                    string
	PathMonitor                string
	PathMonitorIndividual      string
	Port                       string
	ResponseTime               string
	RtspSessionIdRemap         string
	ServerId                   string
	ServerName                 string
	ServerTimeout              string
	ServiceConfigurationType   string
	ServiceGroupEffectiveState string
	ServiceIpStr               string
	ServiceType                string
	State                      string
	StateChangeTimeMsec        string
	StateChangeTimeSec         string
	StateUpdateReason          string
	SurgeProtection            string
	Svcitmactsvcs              string
	Svcitmboundsvcs            string
	Svrstate                   string
	TcpBuffering               string
	TcpProfileName             string
	TicksSinceLastStateChange  string
	TimeSinceLastStateChange   string
	TrafficDomain              string
	TrofsDelay                 string
	UseProxyPort               string
	Usip                       string
	Value                      string
	Weight                     string
}{
	AppFlowLog:                 "appflowlog",
	AutoDelayedTrofs:           "autodelayedtrofs",
	AutoDisableDelay:           "autodisabledelay",
	AutoDisableGraceful:        "autodisablegraceful",
	Autoscale:                  "autoscale",
	Cacheable:                  "cacheable",
	CacheType:                  "cachetype",
	ClientIpInsertion:          "cip",
	ClientIpInsertionHeader:    "cipheader",
	ClientKeepAlive:            "cka",
	ClientTimeout:              "clttimeout",
	Clmonowner:                 "clmonowner",
	Clmonview:                  "clmonview",
	Compression:                "cmp",
	Comment:                    "comment",
	Count:                      "__count",
	CustomServerId:             "customserverid",
	DbsTtl:                     "dbsttl",
	Delay:                      "delay",
	DownstateFlush:             "downstateflush",
	Graceful:                   "graceful",
	GroupCount:                 "groupcount",
	HashId:                     "hashid",
	HealthMonitor:              "healthmonitor",
	HttpProfileName:            "httpprofilename",
	IncludeMembers:             "includemembers",
	Ip:                         "ip",
	MaxBandwidth:               "maxbandwidth",
	MaxClient:                  "maxclient",
	MaxRequest:                 "maxreq",
	MemberPort:                 "memberport",
	MonConnectionClose:         "monconnectionclose",
	MonitorCurrentFailedProbes: "monitorcurrentfailedprobes",
	MonitorState:               "monitor_state",
	MonitorStatusCode:          "monstatcode",
	MonitorThreshold:           "monthreshold",
	MonitorTotalFailedProbes:   "monitortotalfailedprobes",
	MonitorTotalProbes:         "monitortotalprobes",
	MonState:                   "monstate",
	Monstatparam1:              "monstatparam1",
	Monstatparam2:              "monstatparam2",
	Monstatparam3:              "monstatparam3",
	Monuserstatusmesg:          "monuserstatusmesg",
	MonWeight:                  "monweight",
	Name:                       "servicegroupname",
	Nameserver:                 "nameserver",
	NetProfile:                 "netprofile",
	NewName:                    "newname",
	NoDefaultBindings:          "nodefaultbindings",
	NumberOfConnections:        "numofconnections",
	Order:                      "order",
	OrderString:                "orderstr",
	Passive:                    "passive",
	PathMonitor:                "pathmonitor",
	PathMonitorIndividual:      "pathmonitorindv",
	Port:                       "port",
	ResponseTime:               "responsetime",
	RtspSessionIdRemap:         "rtspsessionidremap",
	ServerId:                   "serverid",
	ServerName:                 "servername",
	ServerTimeout:              "svrtimeout",
	ServiceConfigurationType:   "serviceconftype",
	ServiceGroupEffectiveState: "servicegroupeffectivestate",
	ServiceIpStr:               "serviceipstr",
	ServiceType:                "servicetype",
	State:                      "state",
	StateChangeTimeMsec:        "statechangetimemsec",
	StateChangeTimeSec:         "statechangetimesec",
	StateUpdateReason:          "stateupdatereason",
	SurgeProtection:            "sp",
	Svcitmactsvcs:              "svcitmactsvcs",
	Svcitmboundsvcs:            "svcitmboundsvcs",
	Svrstate:                   "svrstate",
	TcpBuffering:               "tcpb",
	TcpProfileName:             "tcpprofilename",
	TicksSinceLastStateChange:  "tickssincelaststatechange",
	TimeSinceLastStateChange:   "timesincelaststatechange",
	TrafficDomain:              "td",
	TrofsDelay:                 "trofsdelay",
	UseProxyPort:               "useproxyport",
	Usip:                       "usip",
	Value:                      "value",
	Weight:                     "weight",
}

type ServiceGroup struct {
	AppFlowLog                 string  `json:"appflowlog,omitempty" nitro:"permission=readwrite"`
	AutoDelayedTrofs           string  `json:"autodelayedtrofs,omitempty" nitro:"permission=readwrite"`
	AutoDisableDelay           float64 `json:"autodisabledelay,omitempty" nitro:"permission=readwrite"`
	AutoDisableGraceful        string  `json:"autodisablegraceful,omitempty" nitro:"permission=readwrite"`
	Autoscale                  string  `json:"autoscale,omitempty" nitro:"permission=readwrite"`
	Cacheable                  string  `json:"cacheable,omitempty" nitro:"permission=readwrite"`
	CacheType                  string  `json:"cachetype,omitempty" nitro:"permission=readwrite"`
	ClientIpInsertion          string  `json:"cip,omitempty" nitro:"permission=readwrite"`
	ClientIpInsertionHeader    string  `json:"cipheader,omitempty" nitro:"permission=readwrite"`
	ClientKeepAlive            string  `json:"cka,omitempty" nitro:"permission=readwrite"`
	ClientTimeout              float64 `json:"clttimeout,omitempty" nitro:"permission=readwrite"`
	Clmonowner                 string  `json:"clmonowner,omitempty" nitro:"permission=readonly"`
	Clmonview                  string  `json:"clmonview,omitempty" nitro:"permission=readonly"`
	Compression                string  `json:"cmp,omitempty" nitro:"permission=readwrite"`
	Comment                    string  `json:"comment,omitempty" nitro:"permission=readwrite"`
	Count                      float64 `json:"__count,omitempty" nitro:"permission=readonly"`
	CustomServerId             string  `json:"customserverid,omitempty" nitro:"permission=readwrite"`
	DbsTtl                     float64 `json:"dbsttl,omitempty" nitro:"permission=readwrite"`
	Delay                      float64 `json:"delay,omitempty" nitro:"permission=readwrite"`
	DownstateFlush             string  `json:"downstateflush,omitempty" nitro:"permission=readwrite"`
	Graceful                   string  `json:"graceful,omitempty" nitro:"permission=readwrite"`
	GroupCount                 string  `json:"groupcount,omitempty" nitro:"permission=readonly"`
	HashId                     string  `json:"hashid,omitempty" nitro:"permission=readwrite"`
	HealthMonitor              string  `json:"healthmonitor,omitempty" nitro:"permission=readwrite"`
	HttpProfileName            string  `json:"httpprofilename,omitempty" nitro:"permission=readwrite"`
	IncludeMembers             bool    `json:"includemembers,omitempty" nitro:"permission=readwrite"`
	Ip                         string  `json:"ip,omitempty" nitro:"permission=readonly"`
	MaxBandwidth               string  `json:"maxbandwidth,omitempty" nitro:"permission=readwrite"`
	MaxClient                  string  `json:"maxclient,omitempty" nitro:"permission=readwrite"`
	MaxRequest                 string  `json:"maxreq,omitempty" nitro:"permission=readwrite"`
	MemberPort                 int     `json:"memberport,omitempty" nitro:"permission=readwrite"`
	MonConnectionClose         string  `json:"monconnectionclose,omitempty" nitro:"permission=readwrite"`
	MonitorCurrentFailedProbes string  `json:"monitorcurrentfailedprobes,omitempty" nitro:"permission=readonly"`
	MonitorState               string  `json:"monitor_state,omitempty" nitro:"permission=readonly"`
	MonitorStatusCode          int     `json:"monstatcode,omitempty" nitro:"permission=readonly"`
	MonitorThreshold           string  `json:"monthreshold,omitempty" nitro:"permission=readwrite"`
	MonitorTotalFailedProbes   string  `json:"monitortotalfailedprobes,omitempty" nitro:"permission=readonly"`
	MonitorTotalProbes         string  `json:"monitortotalprobes,omitempty" nitro:"permission=readonly"`
	MonState                   string  `json:"monstate,omitempty" nitro:"permission=readonly"`
	Monstatparam1              int     `json:"monstatparam1,omitempty" nitro:"permission=readonly"`
	Monstatparam2              int     `json:"monstatparam2,omitempty" nitro:"permission=readonly"`
	Monstatparam3              int     `json:"monstatparam3,omitempty" nitro:"permission=readonly"`
	Monuserstatusmesg          string  `json:"monuserstatusmesg,omitempty" nitro:"permission=readonly"`
	MonWeight                  string  `json:"monweight,omitempty" nitro:"permission=readonly"`
	Name                       string  `json:"servicegroupname,omitempty" nitro:"permission=readwrite"`
	Nameserver                 string  `json:"nameserver,omitempty" nitro:"permission=readwrite"`
	NetProfile                 string  `json:"netprofile,omitempty" nitro:"permission=readwrite"`
	NewName                    string  `json:"newname,omitempty" nitro:"permission=readwrite"`
	NoDefaultBindings          string  `json:"nodefaultbindings,omitempty" nitro:"permission=readonly"`
	NumberOfConnections        int     `json:"numofconnections,omitempty" nitro:"permission=readonly"`
	Order                      float64 `json:"order,omitempty" nitro:"permission=readwrite"`
	OrderString                string  `json:"orderstr,omitempty" nitro:"permission=readonly"`
	Passive                    bool    `json:"passive,omitempty" nitro:"permission=readonly"`
	PathMonitor                string  `json:"pathmonitor,omitempty" nitro:"permission=readwrite"`
	PathMonitorIndividual      string  `json:"pathmonitorindv,omitempty" nitro:"permission=readwrite"`
	Port                       int     `json:"port,omitempty" nitro:"permission=readwrite"`
	ResponseTime               string  `json:"responsetime,omitempty" nitro:"permission=readonly"`
	RtspSessionIdRemap         string  `json:"rtspsessionidremap,omitempty" nitro:"permission=readwrite"`
	ServerId                   float64 `json:"serverid,omitempty" nitro:"permission=readwrite"`
	ServerName                 string  `json:"servername,omitempty" nitro:"permission=readwrite"`
	ServerTimeout              float64 `json:"svrtimeout,omitempty" nitro:"permission=readwrite"`
	ServiceConfigurationType   bool    `json:"serviceconftype,omitempty" nitro:"permission=readonly"`
	ServiceGroupEffectiveState string  `json:"servicegroupeffectivestate,omitempty" nitro:"readonly=readwrite"`
	ServiceIpStr               string  `json:"serviceipstr,omitempty" nitro:"permission=readonly"`
	ServiceType                string  `json:"servicetype,omitempty" nitro:"permission=readwrite"`
	State                      string  `json:"state,omitempty" nitro:"permission=readwrite"`
	StateChangeTimeMsec        string  `json:"statechangetimemsec,omitempty" nitro:"permission=readonly"`
	StateChangeTimeSec         string  `json:"statechangetimesec,omitempty" nitro:"permission=readonly"`
	StateUpdateReason          string  `json:"stateupdatereason,omitempty" nitro:"permission=readonly"`
	SurgeProtection            string  `json:"sp,omitempty" nitro:"permission=readwrite"`
	Svcitmactsvcs              string  `json:"svcitmactsvcs,omitempty" nitro:"permission=readonly"`
	Svcitmboundsvcs            string  `json:"svcitmboundsvcs,omitempty" nitro:"permission=readonly"`
	Svrstate                   string  `json:"svrstate,omitempty" nitro:"permission=readonly"`
	TcpBuffering               string  `json:"tcpb,omitempty" nitro:"permission=readwrite"`
	TcpProfileName             string  `json:"tcpprofilename,omitempty" nitro:"permission=readwrite"`
	TicksSinceLastStateChange  string  `json:"tickssincelaststatechange,omitempty" nitro:"permission=readonly"`
	TimeSinceLastStateChange   float64 `json:"timesincelaststatechange,omitempty" nitro:"permission=readonly"`
	TrafficDomain              string  `json:"td,omitempty" nitro:"permission=readwrite"`
	TrofsDelay                 string  `json:"trofsdelay,omitempty" nitro:"permission=readonly"`
	UseProxyPort               string  `json:"useproxyport,omitempty" nitro:"permission=readwrite"`
	Usip                       string  `json:"usip,omitempty" nitro:"permission=readwrite"`
	Value                      string  `json:"value,omitempty" nitro:"permission=readonly"`
	Weight                     float64 `json:"weight,omitempty" nitro:"permission=readwrite"`
}

func (r ServiceGroup) GetTypeName() string {
	return "servicegroup"
}

func (r ServiceGroup) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("resource_type", r.GetTypeName()),
		slog.String("name", r.Name),
	)
}
