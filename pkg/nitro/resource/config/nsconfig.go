package config

type NsConfig struct {
	Force                           bool     `json:"force,omitempty" nitro:"permission=readwrite"`
	Level                           string   `json:"level,omitempty" nitro:"permission=readwrite"`
	RbaConfig                       bool     `json:"rbaconfig,omitempty" nitro:"permission=readwrite"`
	IpAddress                       string   `json:"ipaddress,omitempty" nitro:"permission=readwrite"`
	Netmask                         string   `json:"netmask,omitempty" nitro:"permission=readwrite"`
	NsVlan                          string   `json:"nsvlan,omitempty" nitro:"permission=readwrite"`
	IfNum                           []string `json:"ifnum,omitempty" nitro:"permission=readwrite"`
	Tagged                          string   `json:"tagged,omitempty" nitro:"permission=readwrite"`
	HttpPort                        []string `json:"httpport,omitempty" nitro:"permission=readwrite"`
	MaxConnections                  string   `json:"maxconn,omitempty" nitro:"permission=readwrite"`
	MaxRequests                     string   `json:"maxreq,omitempty" nitro:"permission=readwrite"`
	InjectClientIpAddress           string   `json:"cip,omitempty" nitro:"permission=readwrite"`
	ClientIpAddressHeader           string   `json:"cipheader,omitempty" nitro:"permission=readwrite"`
	CookieVersion                   string   `json:"cookieversion,omitempty" nitro:"permission=readwrite"`
	SecureCookie                    string   `json:"securecookie,omitempty" nitro:"permission=readwrite"`
	MinimumPathMTU                  string   `json:"pmtumin,omitempty" nitro:"permission=readwrite"`
	PathMTUTimeout                  float64  `json:"pmtutimeout,omitempty" nitro:"permission=readwrite"`
	FtpPortRange                    string   `json:"ftpportrange,omitempty" nitro:"permission=readwrite"`
	CacheRedirectionPortRange       string   `json:"crportrange,omitempty" nitro:"permission=readwrite"`
	Timezone                        string   `json:"timezone,omitempty" nitro:"permission=readwrite"`
	MaxClientSharedQuota            string   `json:"grantquotamaxclient,omitempty" nitro:"permission=readwrite"`
	MaxClientExclusivePercentage    string   `json:"exclusivequotamaxclient,omitempty" nitro:"permission=readwrite"`
	SpilloverSharedQuota            string   `json:"grantquotaspillover,omitempty" nitro:"permission=readwrite"`
	SpilloverExclusivePercentage    string   `json:"exclusivequotaspillover,omitempty" nitro:"permission=readwrite"`
	All                             bool     `json:"all,omitempty" nitro:"permission=readwrite"`
	Config1                         string   `json:"config1,omitempty" nitro:"permission=readwrite"`
	Config2                         string   `json:"config2,omitempty" nitro:"permission=readwrite"`
	DiffOutputType                  string   `json:"outtype,omitempty" nitro:"permission=readwrite"`
	Template                        bool     `json:"template,omitempty" nitro:"permission=readwrite"`
	IgnoreDeviceSpecificDifferences bool     `json:"ignoredevicespecific,omitempty" nitro:"permission=readwrite"`
	WeakPassword                    string   `json:"weakpassword,omitempty" nitro:"permission=readwrite"`
	ChangedPassword                 string   `json:"changedpassword,omitempty" nitro:"permission=readwrite"`
	Config                          string   `json:"config,omitempty" nitro:"permission=readwrite"`
	Message                         string   `json:"message,omitempty" nitro:"permission=readonly"`
	MappedIpAddress                 string   `json:"mappedip,omitempty" nitro:"permission=readonly"`
	Range                           string   `json:"range,omitempty" nitro:"permission=readonly"`
	SvmCmd                          string   `json:"svmcmd,omitempty" nitro:"permission=readonly"`
	SystemType                      string   `json:"systemtype,omitempty" nitro:"permission=readonly"`
	PrimaryIp                       string   `json:"primaryip,omitempty" nitro:"permission=readonly"`
	PrimaryIpv6                     string   `json:"primaryip6,omitempty" nitro:"permission=readonly"`
	Flags                           string   `json:"flags,omitempty" nitro:"permission=readonly"`
	LastConfigChangeTime            string   `json:"lastconfigchangedtime,omitempty" nitro:"permission=readonly"`
	LastConfigSaveTime              string   `json:"lastconfigsavetime,omitempty" nitro:"permission=readonly"`
	CurrentSystemTime               string   `json:"currentsytemtime,omitempty" nitro:"permission=readonly"`
	SystemTime                      string   `json:"systemtime,omitempty" nitro:"permission=readonly"`
	ConfigChanged                   bool     `json:"configchanged,omitempty" nitro:"permission=readonly"`
	Response                        string   `json:"response,omitempty" nitro:"permission=readwrite"`
}

func (r NsConfig) GetTypeName() string {
	return "nsconfig"
}
