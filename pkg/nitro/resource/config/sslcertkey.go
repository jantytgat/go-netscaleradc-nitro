package config

var SslCertKeyFieldNames = struct {
	Builtin            string
	Bundle             string
	Certificate        string
	CertificateType    string
	Count              string
	Data               string
	DateNotAfter       string
	DateNotBefore      string
	DaysToExpiration   string
	DeleteFromDevice   string
	ExpiryMonitor      string
	Feature            string
	FipsKey            string
	HsmKey             string
	Inform             string
	Issuer             string
	LinkCertKeyName    string
	Name               string
	NoDomainCheck      string
	NotificationPeriod string
	OcspResponseStatus string
	OcspStaplingCache  string
	Passcrypt          string
	Passplain          string
	Password           string
	Priority           string
	PrivateKey         string
	PublicKey          string
	PublicKeySize      string
	SanDomains         string
	SanIpAddresses     string
	Serial             string
	ServiceName        string
	SignatureAlgorithm string
	Status             string
	Subject            string
	Version            string
}{
	Builtin:            "builtin",
	Bundle:             "bundle",
	Certificate:        "cert",
	CertificateType:    "certificatetype",
	Count:              "__count",
	Data:               "data",
	DateNotAfter:       "clientcertnotafter",
	DateNotBefore:      "clientcertnotbefore",
	DaysToExpiration:   "daystoexpiration",
	DeleteFromDevice:   "deletefromdevice",
	ExpiryMonitor:      "expirymonitor",
	Feature:            "feature",
	FipsKey:            "fipskey",
	HsmKey:             "hsmkey",
	Inform:             "inform",
	Issuer:             "issuer",
	LinkCertKeyName:    "linkcertkeyname",
	Name:               "certkey",
	NoDomainCheck:      "nodomaincheck",
	NotificationPeriod: "notificationperiod",
	OcspResponseStatus: "ocspresponsestatus",
	OcspStaplingCache:  "ocspstaplingcache",
	Passcrypt:          "passcrypt",
	Passplain:          "passplain",
	Password:           "password",
	Priority:           "priority",
	PrivateKey:         "key",
	PublicKey:          "publickey",
	PublicKeySize:      "publickeysize",
	SanDomains:         "sandns",
	SanIpAddresses:     "sanipadd",
	Serial:             "serial",
	ServiceName:        "servicename",
	SignatureAlgorithm: "signaturealg",
	Status:             "status",
	Subject:            "subject",
	Version:            "version",
}

type SslCertKey struct {
	Builtin            []string `json:"builtin,omitempty" nitro:"permission=readonly"`
	Bundle             string   `json:"bundle,omitempty"  nitro:"permission=readwrite"`
	Certificate        string   `json:"cert,omitempty" nitro:"permission=readwrite"`
	CertificateType    []string `json:"certificatetype,omitempty" nitro:"permission=readonly"`
	Count              float64  `json:"__count,omitempty" nitro:"permission=readonly"`
	Data               string   `json:"data,omitempty" nitro:"permission=readonly"`
	DateNotAfter       string   `json:"clientcertnotafter,omitempty" nitro:"permission=readonly"`
	DateNotBefore      string   `json:"clientcertnotbefore,omitempty" nitro:"permission=readonly"`
	DaysToExpiration   int      `json:"daystoexpiration,omitempty" nitro:"permission=readonly"`
	DeleteFromDevice   bool     `json:"deletefromdevice,omitempty" nitro:"permission=readwrite"`
	ExpiryMonitor      string   `json:"expirymonitor,omitempty" nitro:"permission=readwrite"`
	Feature            string   `json:"feature,omitempty" nitro:"permission=readonly"`
	FipsKey            string   `json:"fipskey,omitempty" nitro:"permission=readwrite"`
	HsmKey             string   `json:"hsmkey,omitempty" nitro:"permission=readwrite"`
	Inform             string   `json:"inform,omitempty" nitro:"permission=readwrite"`
	Issuer             string   `json:"issuer,omitempty" nitro:"permission=readonly"`
	LinkCertKeyName    string   `json:"linkcertkeyname,omitempty" nitro:"permission=readwrite"`
	Name               string   `json:"certkey,omitempty" nitro:"permission=readwrite"`
	NoDomainCheck      bool     `json:"nodomaincheck,omitempty" nitro:"permission=readwrite"`
	NotificationPeriod string   `json:"notificationperiod,omitempty" nitro:"permission=readwrite"`
	OcspResponseStatus string   `json:"ocspresponsestatus,omitempty" nitro:"permission=readonly"`
	OcspStaplingCache  bool     `json:"ocspstaplingcache,omitempty" nitro:"permission=readwrite"`
	Passcrypt          string   `json:"passcrypt,omitempty" nitro:"permission=readonly"`
	Passplain          string   `json:"passplain,omitempty" nitro:"permission=readwrite"`
	Password           string   `json:"password,omitempty" nitro:"permission=readwrite"`
	Priority           string   `json:"priority,omitempty" nitro:"permission=readonly"`
	PrivateKey         string   `json:"key,omitempty" nitro:"permission=readwrite"`
	PublicKey          string   `json:"publickey,omitempty" nitro:"permission=readonly"`
	PublicKeySize      int      `json:"publickeysize,omitempty" nitro:"permission=readonly"`
	SanDomains         string   `json:"sandns,omitempty" nitro:"permission=readonly"`
	SanIpAddresses     string   `json:"sanipadd,omitempty" nitro:"permission=readonly"`
	Serial             string   `json:"serial,omitempty" nitro:"permission=readonly"`
	ServiceName        string   `json:"servicename,omitempty" nitro:"permission=readonly"`
	SignatureAlgorithm string   `json:"signaturealg,omitempty" nitro:"permission=readonly"`
	Status             string   `json:"status,omitempty" nitro:"permission=readonly"`
	Subject            string   `json:"subject,omitempty" nitro:"permission=readonly"`
	Version            int      `json:"version,omitempty" nitro:"permission=readonly"`
}

func (r SslCertKey) GetTypeName() string {
	return "sslcertkey"
}

func NewSslCertKeyAddRequest(name string, cer string, key string) SslCertKey {
	return SslCertKey{
		Name:        name,
		Certificate: cer,
		PrivateKey:  key,
	}
}

func NewSslCertKeyWithPassphraseRequest(name string, cer string, key string, passphrase string) SslCertKey {
	return SslCertKey{
		Name:        name,
		Certificate: cer,
		PrivateKey:  key,
		Passplain:   passphrase,
	}
}

func NewSslCertKeyBundleAddRequest(name string, cer string, key string) SslCertKey {
	return SslCertKey{
		Name:        name,
		Certificate: cer,
		PrivateKey:  key,
		Bundle:      "yes",
	}
}

func NewSslCertKeyBundleWithPassphraseAddRequest(name string, cer string, key string, passphrase string) SslCertKey {
	return SslCertKey{
		Name:        name,
		Certificate: cer,
		PrivateKey:  key,
		Passplain:   passphrase,
		Bundle:      "yes",
	}
}

func NewSslCertKeyUpdateRequest(name string, cer string, key string, noDomainCheck bool) SslCertKey {
	return SslCertKey{
		Name:          name,
		Certificate:   cer,
		PrivateKey:    key,
		NoDomainCheck: noDomainCheck,
	}
}

func NewSslCertKeyWithPassphraseUpdateRequest(name string, cer string, key string, passphrase string, noDomainCheck bool) SslCertKey {
	return SslCertKey{
		Name:          name,
		Certificate:   cer,
		PrivateKey:    key,
		Passplain:     passphrase,
		NoDomainCheck: noDomainCheck,
	}
}

func NewSslCertKeyClearOcspStaplingCacheRequest(name string) SslCertKey {
	return SslCertKey{
		Name:              name,
		OcspStaplingCache: true,
	}
}

func NewSslCertKeyLinkRequest(name string, caName string) SslCertKey {
	return SslCertKey{
		Name:            name,
		LinkCertKeyName: caName,
	}
}

func NewSslCertKeyReloadRequest(name string, monitor bool, period string) SslCertKey {
	r := SslCertKey{
		Name:               name,
		NotificationPeriod: period,
	}

	switch monitor {
	case true:
		r.ExpiryMonitor = "ENABLED"
	case false:
		r.ExpiryMonitor = "DISABLED"
	}
	return r
}

func NewSslCertKeyUnlinkRequest(name string) SslCertKey {
	return SslCertKey{
		Name: name,
	}
}
