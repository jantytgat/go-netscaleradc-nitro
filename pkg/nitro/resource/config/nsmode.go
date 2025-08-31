package config

var NsModeFieldNames = struct {
	BpduBridgingMode             string
	ClientKeepAlive              string
	DirectRouteAdvertisement     string
	EdgeConfiguration            string
	EnabledModes                 string
	FastRamp                     string
	IntranetRouteAdvertisement   string
	Ipv6DirectRouteAdvertisement string
	Ipv6StaticRouteAdvertisement string
	Layer2Mode                   string
	Layer3Mode                   string
	MacBasedForwarding           string
	MediaClassification          string
	PathMtuDiscovery             string
	SingleIpAddress              string
	StaticRouteAdvertisement     string
	TcpBuffering                 string
	UnifiedLoggingModeFramework  string
	UseSourceIpAddress           string
	UseSubnetIpAddress           string
}{
	BpduBridgingMode:             "bridgebpdus",
	ClientKeepAlive:              "cka",
	DirectRouteAdvertisement:     "dradv",
	EdgeConfiguration:            "edge",
	EnabledModes:                 "mode",
	FastRamp:                     "fr",
	IntranetRouteAdvertisement:   "iradv",
	Ipv6DirectRouteAdvertisement: "dradv6",
	Ipv6StaticRouteAdvertisement: "sradv6",
	Layer2Mode:                   "l2",
	Layer3Mode:                   "l3",
	MacBasedForwarding:           "mbf",
	MediaClassification:          "mediaclassification",
	PathMtuDiscovery:             "pmtud",
	SingleIpAddress:              "single_ip",
	StaticRouteAdvertisement:     "sradv",
	TcpBuffering:                 "tcpb",
	UnifiedLoggingModeFramework:  "ulfd",
	UseSourceIpAddress:           "usip",
	UseSubnetIpAddress:           "usnip",
}

type NsMode struct {
	BpduBridgingMode             bool     `json:"bridgebpdus,omitempty" nitro:"permission=readonly"`
	ClientKeepAlive              bool     `json:"cka,omitempty" nitro:"permission=readonly"`
	DirectRouteAdvertisement     bool     `json:"dradv,omitempty" nitro:"permission=readonly"`
	EdgeConfiguration            bool     `json:"edge,omitempty" nitro:"permission=readonly"`
	EnabledModes                 []string `json:"mode,omitempty" nitro:"permission=readwrite"`
	FastRamp                     bool     `json:"fr,omitempty" nitro:"permission=readonly"`
	IntranetRouteAdvertisement   bool     `json:"iradv,omitempty" nitro:"permission=readonly"`
	Ipv6DirectRouteAdvertisement bool     `json:"dradv6,omitempty" nitro:"permission=readonly"`
	Ipv6StaticRouteAdvertisement bool     `json:"sradv6,omitempty" nitro:"permission=readonly"`
	Layer2Mode                   bool     `json:"l2,omitempty" nitro:"permission=readonly"`
	Layer3Mode                   bool     `json:"l3,omitempty" nitro:"permission=readonly"`
	MacBasedForwarding           bool     `json:"mbf,omitempty" nitro:"permission=readonly"`
	MediaClassification          bool     `json:"mediaclassification,omitempty" nitro:"permission=readonly"`
	PathMtuDiscovery             bool     `json:"pmtud,omitempty" nitro:"permission=readonly"`
	SingleIpAddress              bool     `json:"single_ip,omitempty" nitro:"permission=readonly"`
	StaticRouteAdvertisement     bool     `json:"sradv,omitempty" nitro:"permission=readonly"`
	TcpBuffering                 bool     `json:"tcpb,omitempty" nitro:"permission=readonly"`
	UnifiedLoggingModeFramework  bool     `json:"ulfd,omitempty" nitro:"permission=readonly"`
	UseSourceIpAddress           bool     `json:"usip,omitempty" nitro:"permission=readonly"`
	UseSubnetIpAddress           bool     `json:"usnip,omitempty" nitro:"permission=readonly"`
}

func (r NsMode) GetTypeName() string {
	return "nsmode"
}
