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

package stat

type SslVserver struct {
	VserverName                             string  `json:"vservername,omitempty" nitro:"permission=readwrite"`
	ClearStats                              string  `json:"clearstats,omitempty" nitro:"permission=readwrite"`
	SslClientAuthenticationSuccessRate      string  `json:"sslclientauthsuccessrate,omitempty" nitro:"permission=readonly"`
	TotalEncryptedBytes                     float64 `json:"sslctxtotencbytes,omitempty" nitro:"permission=readonly"`
	SslTotalClientAuthenticationFailure     float64 `json:"ssltotclientauthfailure,omitempty" nitro:"permission=readonly"`
	DecryptedBytesRate                      float64 `json:"sslctxdecbytesrate,omitempty" nitro:"permission=readonly"`
	SslTotalClientAuthenticationFailureRate float64 `json:"sslclientauthfailurerate,omitempty" nitro:"permission=readonly"`
	Type                                    string  `json:"type,omitempty" nitro:"permission=readonly"`
	PrimaryIpAddress                        string  `json:"primaryipaddress,omitempty" nitro:"permission=readonly"`
	TotalSslSessionHits                     float64 `json:"sslctxtotsessionhits,omitempty" nitro:"permission=readonly"`
	State                                   string  `json:"state,omitempty" nitro:"permission=readonly"`
	VserverLbHealth                         float64 `json:"vslbhealth,omitempty" nitro:"permission=readonly"`
	HardwareDecryptedBytesRate              float64 `json:"sslctxhwdec_bytesrate,omitempty" nitro:"permission=readonly"`
	SslTotalClientAuthenticationSuccess     float64 `json:"ssltotclientauthsuccess,omitempty" nitro:"permission=readonly"`
	TotalHardwareDecryptedBytes             float64 `json:"sslctxtothwdec_bytes,omitempty" nitro:"permission=readonly"`
	TotalNewSslSessions                     float64 `json:"sslctxtotsessionnew,omitempty" nitro:"permission=readonly"`
	NewSslSessionsRate                      float64 `json:"sslctxsessionnewrate,omitempty" nitro:"permission=readonly"`
	EncryptedBytesRate                      float64 `json:"sslctxencbytesrate,omitempty" nitro:"permission=readonly"`
	HardwareEncryptedBytesRate              float64 `json:"sslctxhwencbytesrate,omitempty" nitro:"permission=readonly"`
	TotalDecryptedBytes                     float64 `json:"sslctxtotdecbytes,omitempty" nitro:"permission=readonly"`
	TotalHardwareEncryptedBytes             float64 `json:"sslctxtothwencbytes,omitempty" nitro:"permission=readonly"`
	SslSessionHitsRate                      float64 `json:"sslctxsessionhitsrate,omitempty" nitro:"permission=readonly"`
	ActiveServices                          float64 `json:"actsvcs,omitempty" nitro:"permission=readonly"`
	PrimaryPort                             int     `json:"primaryport,omitempty" nitro:"permission=readonly"`
}

func (r SslVserver) GetTypeName() string {
	return "sslvserver"
}
