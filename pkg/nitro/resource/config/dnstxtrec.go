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

type DnsTxtRec struct {
	Domain     string   `json:"domain,omitempty" nitro:"permission=readwrite"`
	Data       []string `json:"string,omitempty" nitro:"permission=readwrite"`
	Ttl        float64  `json:"ttl,omitempty" nitro:"permission=readwrite"`
	RecordId   string   `json:"recordid,omitempty" nitro:"permission=readwrite"`
	EcsSubnet  string   `json:"ecssubnet,omitempty" nitro:"permission=readwrite"`
	RecordType string   `json:"type,omitempty" nitro:"permission=readwrite"`
	NodeId     string   `json:"nodeid,omitempty" nitro:"permission=readwrite"`
	AuthType   string   `json:"authtype,omitempty" nitro:"permission=readonly"`
	Count      float64  `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r DnsTxtRec) GetTypeName() string {
	return "dnstxtrec"
}

func NewDnsTxtRecAddRequest(domain string, data []string, ttl float64) DnsTxtRec {
	return DnsTxtRec{
		Domain: domain,
		Data:   data,
		Ttl:    ttl,
	}
}
