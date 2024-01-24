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

type DnsAddRec struct {
	Hostname    string  `json:"hostname,omitempty" nitro:"permission=readwrite"`
	IpAddress   string  `json:"ipaddress,omitempty" nitro:"permission=readwrite"`
	Ttl         float64 `json:"ttl,omitempty" nitro:"permission=readwrite"`
	EcsSubnet   string  `json:"ecssubnet,omitempty" nitro:"permission=readwrite"`
	Type        string  `json:"type,omitempty" nitro:"permission=readwrite"`
	NodeId      float64 `json:"nodeid,omitempty" nitro:"permission=readwrite"`
	Vservername string  `json:"vservername,omitempty" nitro:"permission=readonly"`
	Authtype    string  `json:"authtype,omitempty" nitro:"permission=readonly"`
	Count       float64 `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r DnsAddRec) GetTypeName() string {
	return "dnsaddrec"
}

func NewDnsAddRecAddRequest(hostname string, ipaddress string, ttl float64) DnsAddRec {
	return DnsAddRec{
		Hostname:  hostname,
		IpAddress: ipaddress,
		Ttl:       ttl,
	}
}
