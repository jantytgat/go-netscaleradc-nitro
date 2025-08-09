package config

var HaNodeFieldNames = struct {
	CompletedFlipTime               string
	Count                           string
	CurrentFlips                    string
	DeadInterval                    string
	DisabledInterfaces              string
	EnabledInterfaces               string
	Failsafe                        string
	Flags                           string
	HaHeartbeatInterfaces           string
	HaMonInterfaces                 string
	HaPropagation                   string
	HaStatus                        string
	HaSyncFailureReason             string
	HaSynchronization               string
	HelloInterval                   string
	Id                              string
	IndependentNetworkConfiguration string
	IpAddress                       string
	MasterStateTime                 string
	MaxFlips                        string
	MaxFliptime                     string
	MulticastOnlyInterfaces         string
	Name                            string
	Netmask                         string
	PartialFailureInterfaces        string
	RouteMonitor                    string
	RouteMonitorState               string
	SslCardStatus                   string
	State                           string
	SyncStatusStrictMode            string
	SyncVlan                        string
}{
	CompletedFlipTime:               "completedfliptime",
	Count:                           "__count",
	CurrentFlips:                    "curflips",
	DeadInterval:                    "deadinterval",
	DisabledInterfaces:              "disifaces",
	EnabledInterfaces:               "enaifaces",
	Failsafe:                        "failsafe",
	Flags:                           "flags",
	HaHeartbeatInterfaces:           "haheartbeatifaces",
	HaMonInterfaces:                 "hamonifaces",
	HaPropagation:                   "haprop",
	HaStatus:                        "hastatus",
	HaSyncFailureReason:             "hasyncfailurereason",
	HaSynchronization:               "hasync",
	HelloInterval:                   "hellointerval",
	Id:                              "id",
	IndependentNetworkConfiguration: "inc",
	IpAddress:                       "ipaddress",
	MasterStateTime:                 "masterstatetime",
	MaxFlips:                        "maxflips",
	MaxFliptime:                     "maxfliptime",
	MulticastOnlyInterfaces:         "ifaces",
	Name:                            "name",
	Netmask:                         "netmask",
	PartialFailureInterfaces:        "pfifaces",
	RouteMonitor:                    "routemonitor",
	RouteMonitorState:               "routemonitorstate",
	SslCardStatus:                   "ssl2",
	State:                           "state",
	SyncStatusStrictMode:            "syncstatusstrictmode",
	SyncVlan:                        "syncvlan",
}

type HaNode struct {
	CompletedFlipTime               string  `json:"completedfliptime,omitempty" nitro:"permission=readonly"`
	Count                           float64 `json:"__count,omitempty" nitro:"permission=readonly"`
	CurrentFlips                    string  `json:"curflips,omitempty" nitro:"permission=readonly"`
	DeadInterval                    float64 `json:"deadinterval,omitempty" nitro:"permission=readwrite"`
	DisabledInterfaces              string  `json:"disifaces,omitempty" nitro:"permission=readonly"`
	EnabledInterfaces               string  `json:"enaifaces,omitempty" nitro:"permission=readonly"`
	Failsafe                        string  `json:"failsafe,omitempty" nitro:"permission=readwrite"`
	Flags                           string  `json:"flags,omitempty" nitro:"permission=readonly"`
	HaHeartbeatInterfaces           string  `json:"haheartbeatifaces,omitempty" nitro:"permission=readonly"`
	HaMonInterfaces                 string  `json:"hamonifaces,omitempty" nitro:"permission=readonly"`
	HaPropagation                   string  `json:"haprop,omitempty" nitro:"permission=readwrite"`
	HaStatus                        string  `json:"hastatus,omitempty" nitro:"permission=readwrite"`
	HaSyncFailureReason             string  `json:"hasyncfailurereason,omitempty" nitro:"permission=readonly"`
	HaSynchronization               string  `json:"hasync,omitempty" nitro:"permission=readwrite"`
	HelloInterval                   float64 `json:"hellointerval,omitempty" nitro:"permission=readwrite"`
	Id                              string  `json:"id,omitempty" nitro:"permission=readwrite"`
	IndependentNetworkConfiguration string  `json:"inc,omitempty" nitro:"permission=readwrite"`
	IpAddress                       string  `json:"ipaddress,omitempty" nitro:"permission=readwrite"`
	MasterStateTime                 float64 `json:"masterstatetime,omitempty" nitro:"permission=readonly"`
	MaxFlips                        string  `json:"maxflips,omitempty" nitro:"permission=readwrite"`
	MaxFliptime                     string  `json:"maxfliptime,omitempty" nitro:"permission=readwrite"`
	MulticastOnlyInterfaces         string  `json:"ifaces,omitempty" nitro:"permission=readonly"`
	Name                            string  `json:"name,omitempty" nitro:"permission=readonly"`
	Netmask                         string  `json:"netmask,omitempty" nitro:"permission=readonly"`
	PartialFailureInterfaces        string  `json:"pfifaces,omitempty" nitro:"permission=readonly"`
	RouteMonitor                    string  `json:"routemonitor,omitempty" nitro:"permission=readonly"`
	RouteMonitorState               string  `json:"routemonitorstate,omitempty" nitro:"permission=readonly"`
	SslCardStatus                   string  `json:"ssl2,omitempty" nitro:"permission=readonly"`
	State                           string  `json:"state,omitempty" nitro:"permission=readonly"`
	SyncStatusStrictMode            string  `json:"syncstatusstrictmode,omitempty" nitro:"permission=readwrite"`
	SyncVlan                        float64 `json:"syncvlan,omitempty" nitro:"permission=readwrite"`
}

func (r HaNode) GetTypeName() string {
	return "hanode"
}
