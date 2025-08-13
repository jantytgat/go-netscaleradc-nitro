package config

var SslCertKeyServiceBindingFieldNames = struct {
	Ca               string
	CertKey          string
	Count            string
	Data             string
	Service          string
	ServiceGroupName string
	ServiceName      string
	StateFlag        string
	Version          string
}{
	Ca:               "ca",
	CertKey:          "certkey",
	Count:            "__count",
	Data:             "data",
	Service:          "service",
	ServiceGroupName: "servicegroupname",
	ServiceName:      "servicename",
	StateFlag:        "stateflag",
	Version:          "version",
}

type SslCertKeyServiceBinding struct {
	Ca               bool    `json:"ca,omitempty" nitro:"permission=readwrite"`
	CertKey          string  `json:"certkey,omitempty" nitro:"permission=readwrite"`
	Count            float64 `json:"__count,omitempty" nitro:"permission=readonly"`
	Data             string  `json:"data,omitempty" nitro:"permission=readonly"`
	Service          bool    `json:"service,omitempty" nitro:"permission=readwrite"`
	ServiceGroupName string  `json:"servicegroupname,omitempty" nitro:"permission=readwrite"`
	ServiceName      string  `json:"servicename,omitempty" nitro:"permission=readwrite"`
	StateFlag        string  `json:"stateflag,omitempty" nitro:"permission=readonly"`
	Version          int     `json:"version,omitempty" nitro:"permission=readonly"`
}

func (r SslCertKeyServiceBinding) GetTypeName() string {
	return "sslcertkey_service_binding"
}
