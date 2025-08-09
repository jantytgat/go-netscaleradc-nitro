package stat

var SslVserverFieldNames = struct {
	ActiveServices                          string
	ClearStats                              string
	DecryptedBytesRate                      string
	EncryptedBytesRate                      string
	HardwareDecryptedBytesRate              string
	HardwareEncryptedBytesRate              string
	NewSslSessionsRate                      string
	PrimaryIpAddress                        string
	PrimaryPort                             string
	SslClientAuthenticationSuccessRate      string
	SslSessionHitsRate                      string
	SslTotalClientAuthenticationFailure     string
	SslTotalClientAuthenticationFailureRate string
	SslTotalClientAuthenticationSuccess     string
	State                                   string
	TotalDecryptedBytes                     string
	TotalEncryptedBytes                     string
	TotalHardwareDecryptedBytes             string
	TotalHardwareEncryptedBytes             string
	TotalNewSslSessions                     string
	TotalSslSessionHits                     string
	Type                                    string
	VserverLbHealth                         string
	VserverName                             string
}{
	ActiveServices:                          "actsvcs",
	ClearStats:                              "clearstats",
	DecryptedBytesRate:                      "sslctxdecbytesrate",
	EncryptedBytesRate:                      "sslctxencbytesrate",
	HardwareDecryptedBytesRate:              "sslctxhwdec_bytesrate",
	HardwareEncryptedBytesRate:              "sslctxhwencbytesrate",
	NewSslSessionsRate:                      "sslctxsessionnewrate",
	PrimaryIpAddress:                        "primaryipaddress",
	PrimaryPort:                             "primaryport",
	SslClientAuthenticationSuccessRate:      "sslclientauthsuccessrate",
	SslSessionHitsRate:                      "sslctxsessionhitsrate",
	SslTotalClientAuthenticationFailure:     "ssltotclientauthfailure",
	SslTotalClientAuthenticationFailureRate: "sslclientauthfailurerate",
	SslTotalClientAuthenticationSuccess:     "ssltotclientauthsuccess",
	State:                                   "state",
	TotalDecryptedBytes:                     "sslctxtotdecbytes",
	TotalEncryptedBytes:                     "sslctxtotencbytes",
	TotalHardwareDecryptedBytes:             "sslctxtothwdec_bytes",
	TotalHardwareEncryptedBytes:             "sslctxtothwencbytes",
	TotalNewSslSessions:                     "sslctxtotsessionnew",
	TotalSslSessionHits:                     "sslctxtotsessionhits",
	Type:                                    "type",
	VserverLbHealth:                         "vslbhealth",
	VserverName:                             "vservername",
}

type SslVserver struct {
	ActiveServices                          float64 `json:"actsvcs,omitempty" nitro:"permission=readonly"`
	ClearStats                              string  `json:"clearstats,omitempty" nitro:"permission=readwrite"`
	DecryptedBytesRate                      float64 `json:"sslctxdecbytesrate,omitempty" nitro:"permission=readonly"`
	EncryptedBytesRate                      float64 `json:"sslctxencbytesrate,omitempty" nitro:"permission=readonly"`
	HardwareDecryptedBytesRate              float64 `json:"sslctxhwdec_bytesrate,omitempty" nitro:"permission=readonly"`
	HardwareEncryptedBytesRate              float64 `json:"sslctxhwencbytesrate,omitempty" nitro:"permission=readonly"`
	NewSslSessionsRate                      float64 `json:"sslctxsessionnewrate,omitempty" nitro:"permission=readonly"`
	PrimaryIpAddress                        string  `json:"primaryipaddress,omitempty" nitro:"permission=readonly"`
	PrimaryPort                             int     `json:"primaryport,omitempty" nitro:"permission=readonly"`
	SslClientAuthenticationSuccessRate      string  `json:"sslclientauthsuccessrate,omitempty" nitro:"permission=readonly"`
	SslSessionHitsRate                      float64 `json:"sslctxsessionhitsrate,omitempty" nitro:"permission=readonly"`
	SslTotalClientAuthenticationFailure     float64 `json:"ssltotclientauthfailure,omitempty" nitro:"permission=readonly"`
	SslTotalClientAuthenticationFailureRate float64 `json:"sslclientauthfailurerate,omitempty" nitro:"permission=readonly"`
	SslTotalClientAuthenticationSuccess     float64 `json:"ssltotclientauthsuccess,omitempty" nitro:"permission=readonly"`
	State                                   string  `json:"state,omitempty" nitro:"permission=readonly"`
	TotalDecryptedBytes                     float64 `json:"sslctxtotdecbytes,omitempty" nitro:"permission=readonly"`
	TotalEncryptedBytes                     float64 `json:"sslctxtotencbytes,omitempty" nitro:"permission=readonly"`
	TotalHardwareDecryptedBytes             float64 `json:"sslctxtothwdec_bytes,omitempty" nitro:"permission=readonly"`
	TotalHardwareEncryptedBytes             float64 `json:"sslctxtothwencbytes,omitempty" nitro:"permission=readonly"`
	TotalNewSslSessions                     float64 `json:"sslctxtotsessionnew,omitempty" nitro:"permission=readonly"`
	TotalSslSessionHits                     float64 `json:"sslctxtotsessionhits,omitempty" nitro:"permission=readonly"`
	Type                                    string  `json:"type,omitempty" nitro:"permission=readonly"`
	VserverLbHealth                         float64 `json:"vslbhealth,omitempty" nitro:"permission=readonly"`
	VserverName                             string  `json:"vservername,omitempty" nitro:"permission=readwrite"`
}

func (r SslVserver) GetTypeName() string {
	return "sslvserver"
}
