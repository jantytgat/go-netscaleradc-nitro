package config

var SslServiceSslCertKeyBindingFieldNames = struct {
	CertKeyName   string
	CleartextPort string
	Count         string
	CrlCheck      string
	IsCa          string
	OcspCheck     string
	ServiceName   string
	SkipCaName    string
	Snicert       string
}{
	CertKeyName:   "certkeyname",
	CleartextPort: "cleartextport",
	Count:         "__count",
	CrlCheck:      "crlcheck",
	IsCa:          "ca",
	OcspCheck:     "ocspcheck",
	ServiceName:   "servicename",
	SkipCaName:    "skipcaname",
	Snicert:       "snicert",
}

type SslServiceSslCertKeyBinding struct {
	CertKeyName   string  `json:"certkeyname,omitempty" nitro:"permission=readwrite"`
	CleartextPort int     `json:"cleartextport,omitempty" nitro:"permission=readonly"`
	Count         float64 `json:"__count,omitempty" nitro:"permission=readonly"`
	CrlCheck      string  `json:"crlcheck,omitempty" nitro:"permission=readwrite"`
	IsCa          bool    `json:"ca,omitempty" nitro:"permission=readwrite"`
	OcspCheck     string  `json:"ocspcheck,omitempty" nitro:"permission=readwrite"`
	ServiceName   string  `json:"servicename,omitempty" nitro:"permission=readwrite"`
	SkipCaName    bool    `json:"skipcaname,omitempty" nitro:"permission=readwrite"`
	Snicert       bool    `json:"snicert,omitempty" nitro:"permission=readwrite"`
}

func (r SslServiceSslCertKeyBinding) GetTypeName() string {
	return "sslservice_sslcertkey_binding"
}

func NewSslServiceCertificateBindingAddRequest(servicename string, certkey string, sni bool) SslServiceSslCertKeyBinding {
	return SslServiceSslCertKeyBinding{
		ServiceName: servicename,
		CertKeyName: certkey,
		Snicert:     sni,
	}
}
