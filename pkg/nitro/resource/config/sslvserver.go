package config

type SslVserver struct {
	Name                                        string  `json:"vservername,omitempty" nitro:"permission=readwrite"`
	ClearTextPort                               int     `json:"cleartextport,omitempty" nitro:"permission=readwrite"`
	DiffieHellmanKeyExchange                    string  `json:"dh,omitempty" nitro:"permission=readwrite"`
	DiffieHellmanFile                           string  `json:"dhfile,omitempty" nitro:"permission=readwrite"`
	DiffieHellmanCount                          float64 `json:"dhcount,omitempty" nitro:"permission=readwrite"`
	DiffieHellmanPrivateKeySizeLimit            string  `json:"dhkeyexpsizelimit,omitempty" nitro:"permission=readwrite"`
	EphemeralRSA                                string  `json:"ersa,omitempty" nitro:"permission=readwrite"`
	EphemeralRSACount                           float64 `json:"ersacount,omitempty" nitro:"permission=readwrite"`
	SessionReuse                                string  `json:"sessreuse,omitempty" nitro:"permission=readwrite"`
	SessionTimeout                              float64 `json:"sesstimeout,omitempty" nitro:"permission=readwrite"`
	CipherRedirect                              string  `json:"cipherredirect,omitempty" nitro:"permission=readwrite"`
	CipherUrl                                   string  `json:"cipherurl,omitempty" nitro:"permission=readwrite"`
	Sslv2Redirect                               string  `json:"sslv2redirect,omitempty" nitro:"permission=readwrite"`
	Sslv2Url                                    string  `json:"sslv2url,omitempty" nitro:"permission=readwrite"`
	ClientAuthentication                        string  `json:"clientauth,omitempty" nitro:"permission=readwrite"`
	ClientCertificate                           string  `json:"clientcert,omitempty" nitro:"permission=readwrite"`
	SslRedirect                                 string  `json:"sslredirect,omitempty" nitro:"permission=readwrite"`
	RedirectPortRewrite                         string  `json:"redirectportrewrite,omitempty" nitro:"permission=readwrite"`
	Ssl2                                        string  `json:"ssl2,omitempty" nitro:"permission=readwrite"`
	Ssl3                                        string  `json:"ssl3,omitempty" nitro:"permission=readwrite"`
	Tls1                                        string  `json:"tls1,omitempty" nitro:"permission=readwrite"`
	Tls11                                       string  `json:"tls11,omitempty" nitro:"permission=readwrite"`
	Tls12                                       string  `json:"tls12,omitempty" nitro:"permission=readwrite"`
	Tls13                                       string  `json:"tls13,omitempty" nitro:"permission=readwrite"`
	Dtls1                                       string  `json:"dtls1,omitempty" nitro:"permission=readwrite"`
	Dtls12                                      string  `json:"dtls12,omitempty" nitro:"permission=readwrite"`
	SniEnable                                   string  `json:"snienable,omitempty" nitro:"permission=readwrite"`
	OcspStapling                                string  `json:"ocspstapling,omitempty" nitro:"permission=readwrite"`
	PushEncryptionTrigger                       string  `json:"pushenctrigger,omitempty" nitro:"permission=readwrite"`
	SendCloseNotify                             string  `json:"sendclosenotify,omitempty" nitro:"permission=readwrite"`
	DtlsProfileName                             string  `json:"dtlsprofilename,omitempty" nitro:"permission=readwrite"`
	SslProfile                                  string  `json:"sslprofile,omitempty" nitro:"permission=readwrite"`
	Hsts                                        string  `json:"hsts,omitempty" nitro:"permission=readwrite"`
	MaxAge                                      float64 `json:"maxage,omitempty" nitro:"permission=readwrite"`
	IncludeSubdomains                           string  `json:"includesubdomains,omitempty" nitro:"permission=readwrite"`
	Preload                                     string  `json:"preload,omitempty" nitro:"permission=readwrite"`
	StrictSignedDigestCheck                     string  `json:"strictsigdigestcheck,omitempty" nitro:"permission=readwrite"`
	ZeroRoundTripTimeEarlyData                  string  `json:"zerorttearlydata,omitempty" nitro:"permission=readwrite"`
	Tls13SessionTicketsPerAuthenticationContext float64 `json:"tls13sessionticketsperauthcontext,omitempty" nitro:"permission=readwrite"`
	DheKeyExchangeWithPsk                       string  `json:"dhekeyexchangewithpsk,omitempty" nitro:"permission=readwrite"`
	CrlCheck                                    string  `json:"crlcheck,omitempty" nitro:"permission=readonly"`
	NonFipsCiphers                              string  `json:"nonfipsciphers,omitempty" nitro:"permission=readonly"`
	Service                                     float64 `json:"service,omitempty" nitro:"permission=readonly"`
	OcspCheck                                   string  `json:"ocspcheck,omitempty" nitro:"permission=readonly"`
	CACertificate                               bool    `json:"ca,omitempty" nitro:"permission=readonly"`
	SniCertificate                              bool    `json:"snicert,omitempty" nitro:"permission=readonly"`
	SkipCaName                                  bool    `json:"skipcaname,omitempty" nitro:"permission=readonly"`
	DtlsFlag                                    bool    `json:"dtlsflag,omitempty" nitro:"permission=readonly"`
	Count                                       float64 `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r SslVserver) GetTypeName() string {
	return "sslvserver"
}
