package config

type HaNode struct {
	Id                              string  `json:"id,omitempty" nitro:"permission=readwrite"`
	Name                            string  `json:"name,omitempty" nitro:"permission=readonly"`
	IpAddress                       string  `json:"ipaddress,omitempty" nitro:"permission=readwrite"`
	IndependentNetworkConfiguration string  `json:"inc,omitempty" nitro:"permission=readwrite"`
	HaStatus                        string  `json:"hastatus,omitempty" nitro:"permission=readwrite"`
	HaSynchronization               string  `json:"hasync,omitempty" nitro:"permission=readwrite"`
	HaPropagation                   string  `json:"haprop,omitempty" nitro:"permission=readwrite"`
	HelloInterval                   float64 `json:"hellointerval,omitempty" nitro:"permission=readwrite"`
	DeadInterval                    float64 `json:"deadinterval,omitempty" nitro:"permission=readwrite"`
	Failsafe                        string  `json:"failsafe,omitempty" nitro:"permission=readwrite"`
	MaxFlips                        string  `json:"maxflips,omitempty" nitro:"permission=readwrite"`
	MaxFliptime                     string  `json:"maxfliptime,omitempty" nitro:"permission=readwrite"`
	SyncVlan                        float64 `json:"syncvlan,omitempty" nitro:"permission=readwrite"`
	SyncStatusStrictMode            string  `json:"syncstatusstrictmode,omitempty" nitro:"permission=readwrite"`
	Flags                           string  `json:"flags,omitempty" nitro:"permission=readonly"`
	State                           string  `json:"state,omitempty" nitro:"permission=readonly"`
	EnabledInterfaces               string  `json:"enaifaces,omitempty" nitro:"permission=readonly"`
	DisabledInterfaces              string  `json:"disifaces,omitempty" nitro:"permission=readonly"`
	HaMonInterfaces                 string  `json:"hamonifaces,omitempty" nitro:"permission=readonly"`
	HaHeartbeatInterfaces           string  `json:"haheartbeatifaces,omitempty" nitro:"permission=readonly"`
	PartialFailureInterfaces        string  `json:"pfifaces,omitempty" nitro:"permission=readonly"`
	MulticastOnlyInterfaces         string  `json:"ifaces,omitempty" nitro:"permission=readonly"`
	Netmask                         string  `json:"netmask,omitempty" nitro:"permission=readonly"`
	SslCardStatus                   string  `json:"ssl2,omitempty" nitro:"permission=readonly"`
	MasterStateTime                 float64 `json:"masterstatetime,omitempty" nitro:"permission=readonly"`
	RouteMonitor                    string  `json:"routemonitor,omitempty" nitro:"permission=readonly"`
	CurrentFlips                    string  `json:"curflips,omitempty" nitro:"permission=readonly"`
	CompletedFlipTime               string  `json:"completedfliptime,omitempty" nitro:"permission=readonly"`
	RouteMonitorState               string  `json:"routemonitorstate,omitempty" nitro:"permission=readonly"`
	HaSyncFailureReason             string  `json:"hasyncfailurereason,omitempty" nitro:"permission=readonly"`
	Count                           float64 `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r HaNode) GetTypeName() string {
	return "hanode"
}
