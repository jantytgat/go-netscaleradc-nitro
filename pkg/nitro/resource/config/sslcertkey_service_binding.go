package config

type SslCertKeyServiceBinding struct {
	Service          bool    `json:"service,omitempty" nitro:"permission=readwrite"`
	ServiceName      string  `json:"servicename,omitempty" nitro:"permission=readwrite"`
	CertKey          string  `json:"certkey,omitempty" nitro:"permission=readwrite"`
	ServiceGroupName string  `json:"servicegroupname,omitempty" nitro:"permission=readwrite"`
	Ca               bool    `json:"ca,omitempty" nitro:"permission=readwrite"`
	Version          int     `json:"version,omitempty" nitro:"permission=readonly"`
	Data             string  `json:"data,omitempty" nitro:"permission=readonly"`
	StateFlag        string  `json:"stateflag,omitempty" nitro:"permission=readonly"`
	Count            float64 `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r SslCertKeyServiceBinding) GetTypeName() string {
	return "sslcertkey_service_binding"
}
