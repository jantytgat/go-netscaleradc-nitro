package config

type SslServiceSslCertKeyBinding struct {
	ServiceName   string  `json:"servicename,omitempty" nitro:"permission=readwrite"`
	CertKeyName   string  `json:"certkeyname,omitempty" nitro:"permission=readwrite"`
	Snicert       bool    `json:"snicert,omitempty" nitro:"permission=readwrite"`
	SkipCaName    bool    `json:"skipcaname,omitempty" nitro:"permission=readwrite"`
	IsCa          bool    `json:"ca,omitempty" nitro:"permission=readwrite"`
	CrlCheck      string  `json:"crlcheck,omitempty" nitro:"permission=readwrite"`
	OcspCheck     string  `json:"ocspcheck,omitempty" nitro:"permission=readwrite"`
	CleartextPort int     `json:"cleartextport,omitempty" nitro:"permission=readonly"`
	Count         float64 `json:"__count,omitempty" nitro:"permission=readonly"`
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
