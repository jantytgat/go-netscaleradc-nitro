/*
 * Copyright 2023 CoreLayer BV
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package config

type SslCertKey struct {
	Name               string   `json:"certkey,omitempty" nitro:"permission=readwrite"`
	Certificate        string   `json:"cert,omitempty" nitro:"permission=readwrite"`
	PrivateKey         string   `json:"key,omitempty" nitro:"permission=readwrite"`
	Inform             string   `json:"inform,omitempty" nitro:"permission=readwrite"`
	SignatureAlgorithm string   `json:"signaturealg,omitempty" nitro:"permission=readonly"`
	CertificateType    []string `json:"certificatetype,omitempty" nitro:"permission=readonly"`
	Serial             string   `json:"serial,omitempty" nitro:"permission=readonly"`
	Issuer             string   `json:"issuer,omitempty" nitro:"permission=readonly"`
	DateNotBefore      string   `json:"clientcertnotbefore,omitempty" nitro:"permission=readonly"`
	DateNotAfter       string   `json:"clientcertnotafter,omitempty" nitro:"permission=readonly"`
	DaysToExpiration   int      `json:"daystoexpiration,omitempty" nitro:"permission=readonly"`
	Subject            string   `json:"subject,omitempty" nitro:"permission=readonly"`
	PublicKey          string   `json:"publickey,omitempty" nitro:"permission=readonly"`
	PublicKeySize      int      `json:"publickeysize,omitempty" nitro:"permission=readonly"`
	Version            int      `json:"version,omitempty" nitro:"permission=readonly"`
	Priority           string   `json:"priority,omitempty" nitro:"permission=readonly"`
	Status             string   `json:"status,omitempty" nitro:"permission=readonly"`
	FipsKey            string   `json:"fipskey,omitempty" nitro:"permission=readwrite"`
	HsmKey             string   `json:"hsmkey,omitempty" nitro:"permission=readwrite"`
	Passcrypt          string   `json:"passcrypt,omitempty" nitro:"permission=readonly"`
	Passplain          string   `json:"passplain,omitempty" nitro:"permission=readwrite"`
	Data               string   `json:"data,omitempty" nitro:"permission=readonly"`
	ServiceName        string   `json:"servicename,omitempty" nitro:"permission=readonly"`
	ExpiryMonitor      string   `json:"expirymonitor,omitempty" nitro:"permission=readwrite"`
	NotificationPeriod string   `json:"notificationperiod,omitempty" nitro:"permission=readwrite"`
	LinkCertKeyName    string   `json:"linkcertkeyname,omitempty" nitro:"permission=readwrite"`
	SanDomains         string   `json:"sandns,omitempty" nitro:"permission=readonly"`
	SanIpAddresses     string   `json:"sanipadd,omitempty" nitro:"permission=readonly"`
	OcspResponseStatus string   `json:"ocspresponsestatus,omitempty" nitro:"permission=readonly"`
	Builtin            []string `json:"builtin,omitempty" nitro:"permission=readonly"`
	Feature            string   `json:"feature,omitempty" nitro:"permission=readonly"`
	Budle              string   `json:"bundle,omitempty"  nitro:"permission=readwrite"`
	Password           string   `json:"password,omitempty" nitro:"permission=readwrite"`
	DeleteFromDevice   bool     `json:"deletefromdevice,omitempty" nitro:"permission=readwrite"`
	NoDomainCheck      bool     `json:"nodomaincheck,omitempty" nitro:"permission=readwrite"`
	OcspStaplingCache  bool     `json:"ocspstaplingcache,omitempty" nitro:"permission=readwrite"`
	Count              float64  `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r SslCertKey) GetTypeName() string {
	return "sslcertkey"
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
