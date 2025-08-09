package config

import "log/slog"

var ServiceGroupServiceGroupMemberBindingFieldNames = struct {
	Count                            string
	CustomServerId                   string
	DbsTtl                           string
	Delay                            string
	Delay1                           string
	Graceful                         string
	HashId                           string
	Ip                               string
	Name                             string
	NameServer                       string
	Order                            string
	OrderString                      string
	Port                             string
	ServerBindingFqdnServicePriority string
	ServerId                         string
	ServerName                       string
	ServerSTate                      string
	State                            string
	StateChangeTimeSeconds           string
	TicksSinceLastStateChange        string
	TrofsDelay                       string
	TrofsReason                      string
	Weight                           string
}{
	Count:                            "__count",
	CustomServerId:                   "customserverid",
	DbsTtl:                           "dbsttl",
	Delay:                            "delay",
	Delay1:                           "delay1",
	Graceful:                         "graceful",
	HashId:                           "hashid",
	Ip:                               "ip",
	Name:                             "servicegroupname",
	NameServer:                       "nameserver",
	Order:                            "order",
	OrderString:                      "orderstr",
	Port:                             "port",
	ServerBindingFqdnServicePriority: "svcitmpriority",
	ServerId:                         "serverid",
	ServerName:                       "servername",
	ServerSTate:                      "svrstate",
	State:                            "state",
	StateChangeTimeSeconds:           "statechangetimesec",
	TicksSinceLastStateChange:        "tickssincelaststatechange",
	TrofsDelay:                       "trofsdelay",
	TrofsReason:                      "trofsreason",
	Weight:                           "weight",
}

type ServiceGroupServiceGroupMemberBinding struct {
	Count                            float64 `json:"__count,omitempty" nitro:"permission=readonly"`
	CustomServerId                   string  `json:"customserverid,omitempty" nitro:"permission=readwrite"`
	DbsTtl                           float64 `json:"dbsttl,omitempty" nitro:"permission=readwrite"`
	Delay                            float64 `json:"delay,omitempty" nitro:"permission=readonly"`
	Delay1                           float64 `json:"delay1,omitempty" nitro:"permission=readonly"`
	Graceful                         string  `json:"graceful,omitempty" nitro:"permission=readonly"`
	HashId                           string  `json:"hashid,omitempty" nitro:"permission=readwrite"`
	Ip                               string  `json:"ip,omitempty" nitro:"permission=readwrite"`
	Name                             string  `json:"servicegroupname,omitempty" nitro:"permission=readwrite"`
	NameServer                       string  `json:"nameserver,omitempty" nitro:"permission=readwrite"`
	Order                            float64 `json:"order,omitempty" nitro:"permission=readwrite"`
	OrderString                      string  `json:"orderstr,omitempty" nitro:"permission=readonly"`
	Port                             int     `json:"port,omitempty" nitro:"permission=readwrite"`
	ServerBindingFqdnServicePriority float64 `json:"svcitmpriority,omitempty" nitro:"permission=readonly"`
	ServerId                         string  `json:"serverid,omitempty" nitro:"permission=readwrite"`
	ServerName                       string  `json:"servername,omitempty" nitro:"permission=readwrite"`
	ServerSTate                      string  `json:"svrstate,omitempty" nitro:"permission=readonly"`
	State                            string  `json:"state,omitempty" nitro:"permission=readwrite"`
	StateChangeTimeSeconds           string  `json:"statechangetimesec,omitempty" nitro:"permission=readonly"`
	TicksSinceLastStateChange        string  `json:"tickssincelaststatechange,omitempty" nitro:"permission=readonly"`
	TrofsDelay                       string  `json:"trofsdelay,omitempty" nitro:"permission=readonly"`
	TrofsReason                      string  `json:"trofsreason,omitempty" nitro:"permission=readonly"`
	Weight                           string  `json:"weight,omitempty" nitro:"permission=readwrite"`
}

func (r ServiceGroupServiceGroupMemberBinding) GetTypeName() string {
	return "servicegroup_servicegroupmember_binding"
}

func (r ServiceGroupServiceGroupMemberBinding) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("resource_type", r.GetTypeName()),
		slog.String("name", r.Name),
	)
}
