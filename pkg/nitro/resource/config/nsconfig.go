package config

var NsConfigFieldNames = struct {
	All                             string
	CacheRedirectionPortRange       string
	ChangedPassword                 string
	ClientIpAddressHeader           string
	Config                          string
	Config1                         string
	Config2                         string
	ConfigChanged                   string
	CookieVersion                   string
	CurrentSystemTime               string
	DiffOutputType                  string
	Flags                           string
	Force                           string
	FtpPortRange                    string
	HttpPort                        string
	IfNum                           string
	IgnoreDeviceSpecificDifferences string
	InjectClientIpAddress           string
	IpAddress                       string
	LastConfigChangeTime            string
	LastConfigSaveTime              string
	Level                           string
	MappedIpAddress                 string
	MaxClientExclusivePercentage    string
	MaxClientSharedQuota            string
	MaxConnections                  string
	MaxRequests                     string
	Message                         string
	MinimumPathMTU                  string
	Netmask                         string
	NsVlan                          string
	PathMTUTimeout                  string
	PrimaryIp                       string
	PrimaryIpv6                     string
	Range                           string
	RbaConfig                       string
	Response                        string
	SecureCookie                    string
	SpilloverExclusivePercentage    string
	SpilloverSharedQuota            string
	SvmCmd                          string
	SystemTime                      string
	SystemType                      string
	Tagged                          string
	Template                        string
	Timezone                        string
	WeakPassword                    string
}{
	All:                             "all",
	CacheRedirectionPortRange:       "crportrange",
	ChangedPassword:                 "changedpassword",
	ClientIpAddressHeader:           "cipheader",
	Config:                          "config",
	Config1:                         "config1",
	Config2:                         "config2",
	ConfigChanged:                   "configchanged",
	CookieVersion:                   "cookieversion",
	CurrentSystemTime:               "currentsytemtime",
	DiffOutputType:                  "outtype",
	Flags:                           "flags",
	Force:                           "force",
	FtpPortRange:                    "ftpportrange",
	HttpPort:                        "httpport",
	IfNum:                           "ifnum",
	IgnoreDeviceSpecificDifferences: "ignoredevicespecific",
	InjectClientIpAddress:           "cip",
	IpAddress:                       "ipaddress",
	LastConfigChangeTime:            "lastconfigchangedtime",
	LastConfigSaveTime:              "lastconfigsavetime",
	Level:                           "level",
	MappedIpAddress:                 "mappedip",
	MaxClientExclusivePercentage:    "exclusivequotamaxclient",
	MaxClientSharedQuota:            "grantquotamaxclient",
	MaxConnections:                  "maxconn",
	MaxRequests:                     "maxreq",
	Message:                         "message",
	MinimumPathMTU:                  "pmtumin",
	Netmask:                         "netmask",
	NsVlan:                          "nsvlan",
	PathMTUTimeout:                  "pmtutimeout",
	PrimaryIp:                       "primaryip",
	PrimaryIpv6:                     "primaryip6",
	Range:                           "range",
	RbaConfig:                       "rbaconfig",
	Response:                        "response",
	SecureCookie:                    "securecookie",
	SpilloverExclusivePercentage:    "exclusivequotaspillover",
	SpilloverSharedQuota:            "grantquotaspillover",
	SvmCmd:                          "svmcmd",
	SystemTime:                      "systemtime",
	SystemType:                      "systemtype",
	Tagged:                          "tagged",
	Template:                        "template",
	Timezone:                        "timezone",
	WeakPassword:                    "weakpassword",
}

type NsConfig struct {
	All                             bool     `json:"all,omitempty" nitro:"permission=readwrite"`
	CacheRedirectionPortRange       string   `json:"crportrange,omitempty" nitro:"permission=readwrite"`
	ChangedPassword                 string   `json:"changedpassword,omitempty" nitro:"permission=readwrite"`
	ClientIpAddressHeader           string   `json:"cipheader,omitempty" nitro:"permission=readwrite"`
	Config                          string   `json:"config,omitempty" nitro:"permission=readwrite"`
	Config1                         string   `json:"config1,omitempty" nitro:"permission=readwrite"`
	Config2                         string   `json:"config2,omitempty" nitro:"permission=readwrite"`
	ConfigChanged                   bool     `json:"configchanged,omitempty" nitro:"permission=readonly"`
	CookieVersion                   string   `json:"cookieversion,omitempty" nitro:"permission=readwrite"`
	CurrentSystemTime               string   `json:"currentsytemtime,omitempty" nitro:"permission=readonly"`
	DiffOutputType                  string   `json:"outtype,omitempty" nitro:"permission=readwrite"`
	Flags                           string   `json:"flags,omitempty" nitro:"permission=readonly"`
	Force                           bool     `json:"force,omitempty" nitro:"permission=readwrite"`
	FtpPortRange                    string   `json:"ftpportrange,omitempty" nitro:"permission=readwrite"`
	HttpPort                        []string `json:"httpport,omitempty" nitro:"permission=readwrite"`
	IfNum                           []string `json:"ifnum,omitempty" nitro:"permission=readwrite"`
	IgnoreDeviceSpecificDifferences bool     `json:"ignoredevicespecific,omitempty" nitro:"permission=readwrite"`
	InjectClientIpAddress           string   `json:"cip,omitempty" nitro:"permission=readwrite"`
	IpAddress                       string   `json:"ipaddress,omitempty" nitro:"permission=readwrite"`
	LastConfigChangeTime            string   `json:"lastconfigchangedtime,omitempty" nitro:"permission=readonly"`
	LastConfigSaveTime              string   `json:"lastconfigsavetime,omitempty" nitro:"permission=readonly"`
	Level                           string   `json:"level,omitempty" nitro:"permission=readwrite"`
	MappedIpAddress                 string   `json:"mappedip,omitempty" nitro:"permission=readonly"`
	MaxClientExclusivePercentage    string   `json:"exclusivequotamaxclient,omitempty" nitro:"permission=readwrite"`
	MaxClientSharedQuota            string   `json:"grantquotamaxclient,omitempty" nitro:"permission=readwrite"`
	MaxConnections                  string   `json:"maxconn,omitempty" nitro:"permission=readwrite"`
	MaxRequests                     string   `json:"maxreq,omitempty" nitro:"permission=readwrite"`
	Message                         string   `json:"message,omitempty" nitro:"permission=readonly"`
	MinimumPathMTU                  string   `json:"pmtumin,omitempty" nitro:"permission=readwrite"`
	Netmask                         string   `json:"netmask,omitempty" nitro:"permission=readwrite"`
	NsVlan                          string   `json:"nsvlan,omitempty" nitro:"permission=readwrite"`
	PathMTUTimeout                  float64  `json:"pmtutimeout,omitempty" nitro:"permission=readwrite"`
	PrimaryIp                       string   `json:"primaryip,omitempty" nitro:"permission=readonly"`
	PrimaryIpv6                     string   `json:"primaryip6,omitempty" nitro:"permission=readonly"`
	Range                           string   `json:"range,omitempty" nitro:"permission=readonly"`
	RbaConfig                       bool     `json:"rbaconfig,omitempty" nitro:"permission=readwrite"`
	Response                        string   `json:"response,omitempty" nitro:"permission=readwrite"`
	SecureCookie                    string   `json:"securecookie,omitempty" nitro:"permission=readwrite"`
	SpilloverExclusivePercentage    string   `json:"exclusivequotaspillover,omitempty" nitro:"permission=readwrite"`
	SpilloverSharedQuota            string   `json:"grantquotaspillover,omitempty" nitro:"permission=readwrite"`
	SvmCmd                          string   `json:"svmcmd,omitempty" nitro:"permission=readonly"`
	SystemTime                      string   `json:"systemtime,omitempty" nitro:"permission=readonly"`
	SystemType                      string   `json:"systemtype,omitempty" nitro:"permission=readonly"`
	Tagged                          string   `json:"tagged,omitempty" nitro:"permission=readwrite"`
	Template                        bool     `json:"template,omitempty" nitro:"permission=readwrite"`
	Timezone                        string   `json:"timezone,omitempty" nitro:"permission=readwrite"`
	WeakPassword                    string   `json:"weakpassword,omitempty" nitro:"permission=readwrite"`
}

func (r NsConfig) GetTypeName() string {
	return "nsconfig"
}
