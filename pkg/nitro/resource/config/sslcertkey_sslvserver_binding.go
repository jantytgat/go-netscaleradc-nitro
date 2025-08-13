package config

var SslCertKeySslVserverBindingFieldNames = struct {
	Ca                string
	CertKey           string
	Count             string
	Data              string
	ServerName        string
	StateFlag         string
	Version           string
	VirtualServer     string
	VirtualServerName string
}{
	Ca:                "ca",
	CertKey:           "certkey",
	Count:             "__count",
	Data:              "data",
	ServerName:        "servername",
	StateFlag:         "stateflag",
	Version:           "version",
	VirtualServer:     "vserver",
	VirtualServerName: "vservername",
}

type SslCertKeySslVserverBinding struct {
	Ca                bool    `json:"ca,omitempty" nitro:"permission=readwrite"`
	CertKey           string  `json:"certkey,omitempty" nitro:"permission=readwrite"`
	Count             float64 `json:"__count,omitempty" nitro:"permission=readonly"`
	Data              string  `json:"data,omitempty" nitro:"permission=readonly"`
	ServerName        string  `json:"servername,omitempty" nitro:"permission=readwrite"`
	StateFlag         string  `json:"stateflag,omitempty" nitro:"permission=readonly"`
	Version           int     `json:"version,omitempty" nitro:"permission=readonly"`
	VirtualServer     bool    `json:"vserver,omitempty" nitro:"permission=readwrite"`
	VirtualServerName string  `json:"vservername,omitempty" nitro:"permission=readwrite"`
}

func (r SslCertKeySslVserverBinding) GetTypeName() string {
	return "sslcertkey_sslvserver_binding"
}
