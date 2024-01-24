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

type SslVserverSslCertKeyBinding struct {
	VirtualServerName string  `json:"vservername,omitempty" nitro:"permission=readwrite"`
	CertKeyName       string  `json:"certkeyname,omitempty" nitro:"permission=readwrite"`
	Snicert           bool    `json:"snicert,omitempty" nitro:"permission=readwrite"`
	SkipCaName        bool    `json:"skipcaname,omitempty" nitro:"permission=readwrite"`
	IsCa              bool    `json:"ca,omitempty" nitro:"permission=readwrite"`
	CrlCheck          string  `json:"crlcheck,omitempty" nitro:"permission=readwrite"`
	OcspCheck         string  `json:"ocspcheck,omitempty" nitro:"permission=readwrite"`
	CleartextPort     int     `json:"cleartextport,omitempty" nitro:"permission=readonly"`
	Count             float64 `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r SslVserverSslCertKeyBinding) GetTypeName() string {
	return "sslvserver_sslcertkey_binding"
}

func NewSslVserverCertificateBindingAddRequest(vserver string, certkey string, sni bool) SslVserverSslCertKeyBinding {
	return SslVserverSslCertKeyBinding{
		VirtualServerName: vserver,
		CertKeyName:       certkey,
		Snicert:           sni,
	}
}
