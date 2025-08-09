package config

var SslVserverFieldNames = struct {
	CACertificate                               string
	CipherRedirect                              string
	CipherUrl                                   string
	ClearTextPort                               string
	ClientAuthentication                        string
	ClientCertificate                           string
	Count                                       string
	CrlCheck                                    string
	DheKeyExchangeWithPsk                       string
	DiffieHellmanCount                          string
	DiffieHellmanFile                           string
	DiffieHellmanKeyExchange                    string
	DiffieHellmanPrivateKeySizeLimit            string
	Dtls1                                       string
	Dtls12                                      string
	DtlsFlag                                    string
	DtlsProfileName                             string
	EphemeralRSA                                string
	EphemeralRSACount                           string
	Hsts                                        string
	IncludeSubdomains                           string
	MaxAge                                      string
	Name                                        string
	NonFipsCiphers                              string
	OcspCheck                                   string
	OcspStapling                                string
	Preload                                     string
	PushEncryptionTrigger                       string
	RedirectPortRewrite                         string
	SendCloseNotify                             string
	Service                                     string
	SessionReuse                                string
	SessionTimeout                              string
	SkipCaName                                  string
	SniCertificate                              string
	SniEnable                                   string
	Ssl2                                        string
	Ssl3                                        string
	SslProfile                                  string
	SslRedirect                                 string
	Sslv2Redirect                               string
	Sslv2Url                                    string
	StrictSignedDigestCheck                     string
	Tls1                                        string
	Tls11                                       string
	Tls12                                       string
	Tls13                                       string
	Tls13SessionTicketsPerAuthenticationContext string
	ZeroRoundTripTimeEarlyData                  string
}{
	CACertificate:                    "ca",
	CipherRedirect:                   "cipherredirect",
	CipherUrl:                        "cipherurl",
	ClearTextPort:                    "cleartextport",
	ClientAuthentication:             "clientauth",
	ClientCertificate:                "clientcert",
	Count:                            "__count",
	CrlCheck:                         "crlcheck",
	DheKeyExchangeWithPsk:            "dhekeyexchangewithpsk",
	DiffieHellmanCount:               "dhcount",
	DiffieHellmanFile:                "dhfile",
	DiffieHellmanKeyExchange:         "dh",
	DiffieHellmanPrivateKeySizeLimit: "dhkeyexpsizelimit",
	Dtls1:                            "dtls1",
	Dtls12:                           "dtls12",
	DtlsFlag:                         "dtlsflag",
	DtlsProfileName:                  "dtlsprofilename",
	EphemeralRSA:                     "ersa",
	EphemeralRSACount:                "ersacount",
	Hsts:                             "hsts",
	IncludeSubdomains:                "includesubdomains",
	MaxAge:                           "maxage",
	Name:                             "vservername",
	NonFipsCiphers:                   "nonfipsciphers",
	OcspCheck:                        "ocspcheck",
	OcspStapling:                     "ocspstapling",
	Preload:                          "preload",
	PushEncryptionTrigger:            "pushenctrigger",
	RedirectPortRewrite:              "redirectportrewrite",
	SendCloseNotify:                  "sendclosenotify",
	Service:                          "service",
	SessionReuse:                     "sessreuse",
	SessionTimeout:                   "sesstimeout",
	SkipCaName:                       "skipcaname",
	SniCertificate:                   "snicert",
	SniEnable:                        "snienable",
	Ssl2:                             "ssl2",
	Ssl3:                             "ssl3",
	SslProfile:                       "sslprofile",
	SslRedirect:                      "sslredirect",
	Sslv2Redirect:                    "sslv2redirect",
	Sslv2Url:                         "sslv2url",
	StrictSignedDigestCheck:          "strictsigdigestcheck",
	Tls1:                             "tls1",
	Tls11:                            "tls11",
	Tls12:                            "tls12",
	Tls13:                            "tls13",
	Tls13SessionTicketsPerAuthenticationContext: "tls13sessionticketsperauthcontext",
	ZeroRoundTripTimeEarlyData:                  "zerorttearlydata",
}

type SslVserver struct {
	CACertificate                               bool    `json:"ca,omitempty" nitro:"permission=readonly"`
	CipherRedirect                              string  `json:"cipherredirect,omitempty" nitro:"permission=readwrite"`
	CipherUrl                                   string  `json:"cipherurl,omitempty" nitro:"permission=readwrite"`
	ClearTextPort                               int     `json:"cleartextport,omitempty" nitro:"permission=readwrite"`
	ClientAuthentication                        string  `json:"clientauth,omitempty" nitro:"permission=readwrite"`
	ClientCertificate                           string  `json:"clientcert,omitempty" nitro:"permission=readwrite"`
	Count                                       float64 `json:"__count,omitempty" nitro:"permission=readonly"`
	CrlCheck                                    string  `json:"crlcheck,omitempty" nitro:"permission=readonly"`
	DheKeyExchangeWithPsk                       string  `json:"dhekeyexchangewithpsk,omitempty" nitro:"permission=readwrite"`
	DiffieHellmanCount                          float64 `json:"dhcount,omitempty" nitro:"permission=readwrite"`
	DiffieHellmanFile                           string  `json:"dhfile,omitempty" nitro:"permission=readwrite"`
	DiffieHellmanKeyExchange                    string  `json:"dh,omitempty" nitro:"permission=readwrite"`
	DiffieHellmanPrivateKeySizeLimit            string  `json:"dhkeyexpsizelimit,omitempty" nitro:"permission=readwrite"`
	Dtls1                                       string  `json:"dtls1,omitempty" nitro:"permission=readwrite"`
	Dtls12                                      string  `json:"dtls12,omitempty" nitro:"permission=readwrite"`
	DtlsFlag                                    bool    `json:"dtlsflag,omitempty" nitro:"permission=readonly"`
	DtlsProfileName                             string  `json:"dtlsprofilename,omitempty" nitro:"permission=readwrite"`
	EphemeralRSA                                string  `json:"ersa,omitempty" nitro:"permission=readwrite"`
	EphemeralRSACount                           float64 `json:"ersacount,omitempty" nitro:"permission=readwrite"`
	Hsts                                        string  `json:"hsts,omitempty" nitro:"permission=readwrite"`
	IncludeSubdomains                           string  `json:"includesubdomains,omitempty" nitro:"permission=readwrite"`
	MaxAge                                      float64 `json:"maxage,omitempty" nitro:"permission=readwrite"`
	Name                                        string  `json:"vservername,omitempty" nitro:"permission=readwrite"`
	NonFipsCiphers                              string  `json:"nonfipsciphers,omitempty" nitro:"permission=readonly"`
	OcspCheck                                   string  `json:"ocspcheck,omitempty" nitro:"permission=readonly"`
	OcspStapling                                string  `json:"ocspstapling,omitempty" nitro:"permission=readwrite"`
	Preload                                     string  `json:"preload,omitempty" nitro:"permission=readwrite"`
	PushEncryptionTrigger                       string  `json:"pushenctrigger,omitempty" nitro:"permission=readwrite"`
	RedirectPortRewrite                         string  `json:"redirectportrewrite,omitempty" nitro:"permission=readwrite"`
	SendCloseNotify                             string  `json:"sendclosenotify,omitempty" nitro:"permission=readwrite"`
	Service                                     float64 `json:"service,omitempty" nitro:"permission=readonly"`
	SessionReuse                                string  `json:"sessreuse,omitempty" nitro:"permission=readwrite"`
	SessionTimeout                              float64 `json:"sesstimeout,omitempty" nitro:"permission=readwrite"`
	SkipCaName                                  bool    `json:"skipcaname,omitempty" nitro:"permission=readonly"`
	SniCertificate                              bool    `json:"snicert,omitempty" nitro:"permission=readonly"`
	SniEnable                                   string  `json:"snienable,omitempty" nitro:"permission=readwrite"`
	Ssl2                                        string  `json:"ssl2,omitempty" nitro:"permission=readwrite"`
	Ssl3                                        string  `json:"ssl3,omitempty" nitro:"permission=readwrite"`
	SslProfile                                  string  `json:"sslprofile,omitempty" nitro:"permission=readwrite"`
	SslRedirect                                 string  `json:"sslredirect,omitempty" nitro:"permission=readwrite"`
	Sslv2Redirect                               string  `json:"sslv2redirect,omitempty" nitro:"permission=readwrite"`
	Sslv2Url                                    string  `json:"sslv2url,omitempty" nitro:"permission=readwrite"`
	StrictSignedDigestCheck                     string  `json:"strictsigdigestcheck,omitempty" nitro:"permission=readwrite"`
	Tls1                                        string  `json:"tls1,omitempty" nitro:"permission=readwrite"`
	Tls11                                       string  `json:"tls11,omitempty" nitro:"permission=readwrite"`
	Tls12                                       string  `json:"tls12,omitempty" nitro:"permission=readwrite"`
	Tls13                                       string  `json:"tls13,omitempty" nitro:"permission=readwrite"`
	Tls13SessionTicketsPerAuthenticationContext float64 `json:"tls13sessionticketsperauthcontext,omitempty" nitro:"permission=readwrite"`
	ZeroRoundTripTimeEarlyData                  string  `json:"zerorttearlydata,omitempty" nitro:"permission=readwrite"`
}

func (r SslVserver) GetTypeName() string {
	return "sslvserver"
}
