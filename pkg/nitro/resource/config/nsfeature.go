package config

var NsFeatureFieldNames = struct {
	Aaa                        string
	AdaptiveTcp                string
	ApiGateway                 string
	AppFlow                    string
	ApplicationFirewall        string
	AppQoe                     string
	BgpRouting                 string
	BotManagement              string
	CacheRedirection           string
	CallHome                   string
	CloudBridge                string
	Compression                string
	ConnectionQualityAnalytics string
	ContentAccelerator         string
	ContentFiltering           string
	ContentInspection          string
	ContentSwitching           string
	EnabledFeatures            string
	ForwardProxy               string
	FrontEndOptimization       string
	GlobalServerLoadBalancing  string
	IntegratedCaching          string
	Ipv6ProtocolTranslation    string
	Isis                       string
	LargeScaleNat              string
	LoadBalancing              string
	OspfRouting                string
	Push                       string
	RdpProxy                   string
	Reputation                 string
	Responder                  string
	Rewrite                    string
	RipRouting                 string
	Routing                    string
	Ssl                        string
	SslInterception            string
	SslVpn                     string
	SurgeProtection            string
	UrlFiltering               string
	VideoOptimization          string
	WebLogging                 string
}{
	Aaa:                        "aaa",
	AdaptiveTcp:                "adaptivetcp",
	ApiGateway:                 "apigateway",
	AppFlow:                    "appflow",
	ApplicationFirewall:        "appfw",
	AppQoe:                     "appqoe",
	BgpRouting:                 "bgp",
	BotManagement:              "bot",
	CacheRedirection:           "cr",
	CallHome:                   "ch",
	CloudBridge:                "cloudbridge",
	Compression:                "cmp",
	ConnectionQualityAnalytics: "cqa",
	ContentAccelerator:         "contentaccelerator",
	ContentFiltering:           "cf",
	ContentInspection:          "ci",
	ContentSwitching:           "cs",
	EnabledFeatures:            "feature",
	ForwardProxy:               "forwardproxy",
	FrontEndOptimization:       "feo",
	GlobalServerLoadBalancing:  "gslb",
	IntegratedCaching:          "ic",
	Ipv6ProtocolTranslation:    "ipv6Pt",
	Isis:                       "isis",
	LargeScaleNat:              "lsn",
	LoadBalancing:              "lb",
	OspfRouting:                "ospf",
	Push:                       "push",
	RdpProxy:                   "rdpproxy",
	Reputation:                 "rep",
	Responder:                  "responder",
	Rewrite:                    "rewrite",
	RipRouting:                 "rip",
	Routing:                    "routing",
	Ssl:                        "ssl",
	SslInterception:            "sslinterception",
	SslVpn:                     "sslvpn",
	SurgeProtection:            "sp",
	UrlFiltering:               "urlfiltering",
	VideoOptimization:          "videooptimization",
	WebLogging:                 "wl",
}

type NsFeature struct {
	Aaa                        bool     `json:"aaa,omitempty" nitro:"permission=readonly"`
	AdaptiveTcp                bool     `json:"adaptivetcp,omitempty" nitro:"permission=readonly"`
	ApiGateway                 bool     `json:"apigateway,omitempty" nitro:"permission=readonly"`
	AppFlow                    bool     `json:"appflow,omitempty" nitro:"permission=readonly"`
	ApplicationFirewall        bool     `json:"appfw,omitempty" nitro:"permission=readonly"`
	AppQoe                     bool     `json:"appqoe,omitempty" nitro:"permission=readonly"`
	BgpRouting                 bool     `json:"bgp,omitempty" nitro:"permission=readonly"`
	BotManagement              bool     `json:"bot,omitempty" nitro:"permission=readonly"`
	CacheRedirection           bool     `json:"cr,omitempty" nitro:"permission=readonly"`
	CallHome                   bool     `json:"ch,omitempty" nitro:"permission=readonly"`
	CloudBridge                bool     `json:"cloudbridge,omitempty" nitro:"permission=readonly"`
	Compression                bool     `json:"cmp,omitempty" nitro:"permission=readonly"`
	ConnectionQualityAnalytics bool     `json:"cqa,omitempty" nitro:"permission=readonly"`
	ContentAccelerator         bool     `json:"contentaccelerator,omitempty" nitro:"permission=readonly"`
	ContentFiltering           bool     `json:"cf,omitempty" nitro:"permission=readonly"`
	ContentInspection          bool     `json:"ci,omitempty" nitro:"permission=readonly"`
	ContentSwitching           bool     `json:"cs,omitempty" nitro:"permission=readonly"`
	EnabledFeatures            []string `json:"feature,omitempty" nitro:"permission=readonly"`
	ForwardProxy               bool     `json:"forwardproxy,omitempty" nitro:"permission=readonly"`
	FrontEndOptimization       bool     `json:"feo,omitempty" nitro:"permission=readonly"`
	GlobalServerLoadBalancing  bool     `json:"gslb,omitempty" nitro:"permission=readonly"`
	IntegratedCaching          bool     `json:"ic,omitempty" nitro:"permission=readonly"`
	Ipv6ProtocolTranslation    bool     `json:"ipv6Pt,omitempty" nitro:"permission=readonly"`
	Isis                       bool     `json:"isis,omitempty" nitro:"permission=readonly"`
	LargeScaleNat              bool     `json:"lsn,omitempty" nitro:"permission=readonly"`
	LoadBalancing              bool     `json:"lb,omitempty" nitro:"permission=readonly"`
	OspfRouting                bool     `json:"ospf,omitempty" nitro:"permission=readonly"`
	Push                       bool     `json:"push,omitempty" nitro:"permission=readonly"`
	RdpProxy                   bool     `json:"rdpproxy,omitempty" nitro:"permission=readonly"`
	Reputation                 bool     `json:"rep,omitempty" nitro:"permission=readonly"`
	Responder                  bool     `json:"responder,omitempty" nitro:"permission=readonly"`
	Rewrite                    bool     `json:"rewrite,omitempty" nitro:"permission=readonly"`
	RipRouting                 bool     `json:"rip,omitempty" nitro:"permission=readonly"`
	Routing                    bool     `json:"routing,omitempty" nitro:"permission=readonly"`
	Ssl                        bool     `json:"ssl,omitempty" nitro:"permission=readonly"`
	SslInterception            bool     `json:"sslinterception,omitempty" nitro:"permission=readonly"`
	SslVpn                     bool     `json:"sslvpn,omitempty" nitro:"permission=readonly"`
	SurgeProtection            bool     `json:"sp,omitempty" nitro:"permission=readonly"`
	UrlFiltering               bool     `json:"urlfiltering,omitempty" nitro:"permission=readonly"`
	VideoOptimization          bool     `json:"videooptimization,omitempty" nitro:"permission=readonly"`
	WebLogging                 bool     `json:"wl,omitempty" nitro:"permission=readonly"`
}

func (r NsFeature) GetTypeName() string {
	return "nsfeature"
}
