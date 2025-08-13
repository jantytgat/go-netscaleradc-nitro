package config

var LbVserverServiceBindingFieldNames = struct {
	CookieIpPort       string
	Count              string
	CurrentState       string
	DynamicWeight      string
	Ipv46              string
	Name               string
	Order              string
	OrderStr           string
	Port               string
	PreferredLocation  string
	ServiceGroupName   string
	ServiceName        string
	ServiceType        string
	VserverBindSvcIp   string
	VserverBindSvcPort string
	VserverId          string
	Weight             string
}{
	CookieIpPort:       "cookieipport",
	Count:              "count",
	CurrentState:       "curstate",
	DynamicWeight:      "dynamicweight",
	Ipv46:              "ipv46",
	Name:               "name",
	Order:              "order",
	OrderStr:           "orderstr",
	Port:               "port",
	PreferredLocation:  "preferredlocation",
	ServiceGroupName:   "servicegroupname",
	ServiceName:        "servicename",
	ServiceType:        "servicetype",
	VserverBindSvcIp:   "vsvrbindsvcip",
	VserverBindSvcPort: "vsvrbindsvcport",
	VserverId:          "vserverid",
	Weight:             "weight",
}

type LbVserverServiceBinding []struct {
	CookieIpPort       string  `json:"cookieipport,omitempty" nitro:"permission=readonly"`
	Count              float64 `json:"count,omitempty" nitro:"permission=readonly"`
	CurrentState       string  `json:"curstate,omitempty" nitro:"permission=readonly"`
	DynamicWeight      float64 `json:"dynamicweight,omitempty" nitro:"permission=readonly"`
	Ipv46              string  `json:"ipv46,omitempty" nitro:"permission=readonly"`
	Name               string  `json:"name,omitempty" nitro:"permission=readwrite"`
	Order              float64 `json:"order,omitempty" nitro:"permission=readwrite"`
	OrderStr           string  `json:"orderstr,omitempty" nitro:"permission=readonly"`
	Port               int     `json:"port,omitempty" nitro:"permission=readonly"`
	PreferredLocation  string  `json:"preferredlocation,omitempty" nitro:"permission=readonly"`
	ServiceGroupName   string  `json:"servicegroupname,omitempty" nitro:"permission=readwrite"`
	ServiceName        string  `json:"servicename,omitempty" nitro:"permission=readwrite"`
	ServiceType        string  `json:"servicetype,omitempty" nitro:"permission=readonly"`
	VserverBindSvcIp   string  `json:"vsvrbindsvcip,omitempty" nitro:"permission=readonly"`
	VserverBindSvcPort int     `json:"vsvrbindsvcport,omitempty" nitro:"permission=readonly"`
	VserverId          string  `json:"vserverid,omitempty" nitro:"permission=readonly"`
	Weight             float64 `json:"weight,omitempty" nitro:"permission=readwrite"`
}

func (r LbVserverServiceBinding) GetTypeName() string {
	return "lbvserver_service_binding"
}
