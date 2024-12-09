package config

type SslCertKeySslVserverBinding struct {
	VirtualServerName string  `json:"vservername,omitempty" nitro:"permission=readwrite"`
	ServerName        string  `json:"servername,omitempty" nitro:"permission=readwrite"`
	CertKey           string  `json:"certkey,omitempty" nitro:"permission=readwrite"`
	VirtualServer     bool    `json:"vserver,omitempty" nitro:"permission=readwrite"`
	Ca                bool    `json:"ca,omitempty" nitro:"permission=readwrite"`
	Version           int     `json:"version,omitempty" nitro:"permission=readonly"`
	Data              string  `json:"data,omitempty" nitro:"permission=readonly"`
	StateFlag         string  `json:"stateflag,omitempty" nitro:"permission=readonly"`
	Count             float64 `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r SslCertKeySslVserverBinding) GetTypeName() string {
	return "sslcertkey_sslvserver_binding"
}
