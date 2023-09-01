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

type SslCertKey_SslVserver_Binding struct {
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

func (r SslCertKey_SslVserver_Binding) GetTypeName() string {
	return "sslcertkey_sslvserver_binding"
}
