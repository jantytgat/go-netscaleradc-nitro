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

type Server struct {
	Name                      string  `json:"name,omitempty" nitro:"permission=readwrite"`
	Internal                  bool    `json:"internal,omitempty" nitro:"permission=readwrite"`
	IpAddress                 string  `json:"ipaddress,omitempty" nitro:"permission=readwrite"`
	State                     string  `json:"state,omitempty" nitro:"permission=readwrite"`
	Domain                    string  `json:"domain,omitempty" nitro:"permission=readwrite"`
	DomainResolveRetry        int     `json:"domainresolveretry,omitempty" nitro:"permission=readwrite"`
	TranslationIp             string  `json:"translationip,omitempty" nitro:"permission=readwrite"`
	TranslationMask           string  `json:"translationmask,omitempty" nitro:"permission=readwrite"`
	Comment                   string  `json:"comment,omitempty" nitro:"permission=readwrite"`
	StateChangeTimeSec        string  `json:"statechangetimesec,omitempty" nitro:"permission=readonly"`
	TicksSinceLastStateChange string  `json:"tickssincelaststatechange,omitempty" nitro:"permission=readonly"`
	Ipv6Address               string  `json:"ipv6address,omitempty" nitro:"permission=readwrite"`
	Td                        string  `json:"td,omitempty" nitro:"permission=readwrite"`
	Autoscale                 string  `json:"autoscale,omitempty" nitro:"permission=readonly"`
	Usip                      string  `json:"usip,omitempty" nitro:"permission=readonly"`
	ClientKeepAlive           string  `json:"cka,omitempty" nitro:"permission=readonly"`
	TcpBuffering              string  `json:"tcpb,omitempty" nitro:"permission=readonly"`
	Compression               string  `json:"cmp,omitempty" nitro:"permission=readonly"`
	Cacheable                 string  `json:"cacheable,omitempty" nitro:"permission=readonly"`
	SurgeProtection           string  `json:"sp,omitempty" nitro:"permission=readonly"`
	QueryType                 string  `json:"querytype,omitempty" nitro:"permission=readwrite"`
	DomainResolveNow          bool    `json:"domainresolvenow,omitempty" nitro:"permission=readwrite"`
	Delay                     float64 `json:"delay,omitempty" nitro:"permission=readwrite"`
	Graceful                  string  `json:"graceful,omitempty" nitro:"permission=readwrite"`
	NewName                   string  `json:"newname,omitempty" nitro:"permission=readwrite"`
	Count                     float64 `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r Server) GetTypeName() string {
	return "server"
}

func NewServerCreateByIpv4Request(name string, ip string) Server {
	return Server{
		Name:      name,
		IpAddress: ip,
	}
}

func NewServerCreateByIpv6Request(name string, ip string) Server {
	return Server{
		Name:        name,
		Ipv6Address: ip,
	}
}

func NewServerCreateByDomainRequest(name string, domain string) Server {
	return Server{
		Name:   name,
		Domain: domain,
	}
}

func NewServerDisableRequest(name string, delay float64, graceful bool) Server {
	r := Server{
		Name:  name,
		Delay: delay,
	}

	switch graceful {
	case true:
		r.Graceful = "YES"
	case false:
		r.Graceful = "NO"
	}

	return r
}

func NewServerRenameRequest(oldName string, newName string) Server {
	return Server{
		Name:    oldName,
		NewName: newName,
	}
}
