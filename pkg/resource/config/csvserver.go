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

type CsVserver struct {
	Name                          string  `json:"name,omitempty" nitro:"permission=readwrite"`
	Td                            string  `json:"td,omitempty" nitro:"permission=readwrite"`
	ServiceType                   string  `json:"servicetype,omitempty" nitro:"permission=readwrite"`
	Ipv46                         string  `json:"ipv46,omitempty" nitro:"permission=readwrite"`
	TargetType                    string  `json:"targettype,omitempty" nitro:"permission=readwrite"`
	DnsRecordType                 string  `json:"dnsrecordtype,omitempty" nitro:"permission=readwrite"`
	PersistenceId                 float64 `json:"persistenceid,omitempty" nitro:"permission=readwrite"`
	IpPattern                     string  `json:"ippattern,omitempty" nitro:"permission=readwrite"`
	IpMask                        string  `json:"ipmask,omitempty" nitro:"permission=readwrite"`
	Range                         string  `json:"range,omitempty" nitro:"permission=readwrite"`
	Port                          int     `json:"port,omitempty" nitro:"permission=readwrite"`
	IpSet                         string  `json:"ipset,omitempty" nitro:"permission=readwrite"`
	State                         string  `json:"state,omitempty" nitro:"permission=readwrite"`
	StateUpdate                   string  `json:"stateupdate,omitempty" nitro:"permission=readwrite"`
	Cacheable                     string  `json:"cacheable,omitempty" nitro:"permission=readwrite"`
	RedirectUrl                   string  `json:"redirecturl,omitempty" nitro:"permission=readwrite"`
	ClientTimeout                 string  `json:"clttimeout,omitempty" nitro:"permission=readwrite"`
	Precedence                    string  `json:"precedence,omitempty" nitro:"permission=readwrite"`
	CaseSensitive                 string  `json:"casesensitive,omitempty" nitro:"permission=readwrite"`
	SpilloverMethod               string  `json:"somethod,omitempty" nitro:"permission=readwrite"`
	SpilloverPersistence          string  `json:"sopersistence,omitempty" nitro:"permission=readwrite"`
	SpilloverPersistenceTimeout   string  `json:"sopersistencetimeout,omitempty" nitro:"permission=readwrite"`
	SpilloverThreshold            float64 `json:"sothreshold,omitempty" nitro:"permission=readwrite"`
	SpilloverBackupAction         string  `json:"sobackupaction,omitempty" nitro:"permission=readwrite"`
	RedirectPortRewrite           string  `json:"redirectportrewrite,omitempty" nitro:"permission=readwrite"`
	DownstateFlush                string  `json:"downstateflush,omitempty" nitro:"permission=readwrite"`
	BackupVserver                 string  `json:"backupvserver,omitempty" nitro:"permission=readwrite"`
	DisablePrimaryOnDown          string  `json:"disableprimaryondown,omitempty" nitro:"permission=readwrite"`
	InsertVserverIpPort           string  `json:"insertvserveripport,omitempty" nitro:"permission=readwrite"`
	VipHeader                     string  `json:"vipheader,omitempty" nitro:"permission=readwrite"`
	RtspNat                       string  `json:"rtspnat,omitempty" nitro:"permission=readwrite"`
	AuthenticationHost            string  `json:"authenticationhost,omitempty" nitro:"permission=readwrite"`
	Authentication                string  `json:"authentication,omitempty" nitro:"permission=readwrite"`
	ListenPolicy                  string  `json:"listenpolicy,omitempty" nitro:"permission=readwrite"`
	ListenPriority                float64 `json:"listenpriority,omitempty" nitro:"permission=readwrite"`
	Authentication401             string  `json:"authn401,omitempty" nitro:"permission=readwrite"`
	AuthenticationVserverName     string  `json:"authnvsname,omitempty" nitro:"permission=readwrite"`
	Push                          string  `json:"push,omitempty" nitro:"permission=readwrite"`
	PushVserver                   string  `json:"pushvserver,omitempty" nitro:"permission=readwrite"`
	PushLabel                     string  `json:"pushlabel,omitempty" nitro:"permission=readwrite"`
	PushMultipleClientConnections string  `json:"pushmulticlients,omitempty" nitro:"permission=readwrite"`
	TcpProfileName                string  `json:"tcpprofilename,omitempty" nitro:"permission=readwrite"`
	HttpProfileName               string  `json:"httpprofilename,omitempty" nitro:"permission=readwrite"`
	DatabaseProfileName           string  `json:"dbprofilename,omitempty" nitro:"permission=readwrite"`
	OracleServerVersion           string  `json:"oracleserverversion,omitempty" nitro:"permission=readwrite"`
	Comment                       string  `json:"comment,omitempty" nitro:"permission=readwrite"`
	MssqlServerVersion            string  `json:"mssqlserverversion,omitempty" nitro:"permission=readwrite"`
	L2ConnectionParameters        string  `json:"l2conn,omitempty" nitro:"permission=readwrite"`
	MysqlProtocolVersion          string  `json:"mysqlprotocolversion,omitempty" nitro:"permission=readwrite"`
	MysqlServerVersion            string  `json:"mysqlserverversion,omitempty" nitro:"permission=readwrite"`
	MysqlCharacterSet             string  `json:"mysqlcharacterset,omitempty" nitro:"permission=readwrite"`
	MysqlServerCapabilities       string  `json:"mysqlservercapabilities,omitempty" nitro:"permission=readwrite"`
	AppflowLog                    string  `json:"appflowlog,omitempty" nitro:"permission=readwrite"`
	NetProfile                    string  `json:"netprofile,omitempty" nitro:"permission=readwrite"`
	IcmpVserverResponse           string  `json:"icmpvsrresponse,omitempty" nitro:"permission=readwrite"`
	RouteHealthInjectionState     string  `json:"rhistate,omitempty" nitro:"permission=readwrite"`
	AuthenticationProfile         string  `json:"authnprofile,omitempty" nitro:"permission=readwrite"`
	DnsProfileName                string  `json:"dnsprofilename,omitempty" nitro:"permission=readwrite"`
	Dtls                          string  `json:"dtls,omitempty" nitro:"permission=readwrite"`
	PersistenceType               string  `json:"persistencetype,omitempty" nitro:"permission=readwrite"`
	PersistenceMask               string  `json:"persistmask,omitempty" nitro:"permission=readwrite"`
	Ipv6PersistenceMaskLength     string  `json:"v6persistmasklen,omitempty" nitro:"permission=readwrite"`
	Timeout                       float64 `json:"timeout,omitempty" nitro:"permission=readwrite"`
	CookieName                    string  `json:"cookiename,omitempty" nitro:"permission=readwrite"`
	PersistenceBackup             string  `json:"persistencebackup,omitempty" nitro:"permission=readwrite"`
	BackupPersistenceTimeout      float64 `json:"backuppersistencetimeout,omitempty" nitro:"permission=readwrite"`
	TcpProbePort                  int     `json:"tcpprobeport,omitempty" nitro:"permission=readwrite"`
	ProbeProtocol                 string  `json:"probeprotocol,omitempty" nitro:"permission=readwrite"`
	ProbeSuccessResponseCode      string  `json:"probesuccessresponsecode,omitempty" nitro:"permission=readwrite"`
	ProbePort                     int     `json:"probeport,omitempty" nitro:"permission=readwrite"`
	QuicProfileName               string  `json:"quicprofilename,omitempty" nitro:"permission=readwrite"`
	RedirectFromPort              int     `json:"redirectfromport,omitempty" nitro:"permission=readwrite"`
	HttpRedirectUrl               string  `json:"httpsredirecturl,omitempty" nitro:"permission=readwrite"`
	DomainName                    string  `json:"domainname,omitempty" nitro:"permission=readwrite"`
	TimeToLive                    float64 `json:"ttl,omitempty" nitro:"permission=readwrite"`
	BackupIp                      string  `json:"backupip,omitempty" nitro:"permission=readwrite"`
	CookieDomain                  string  `json:"cookiedomain,omitempty" nitro:"permission=readwrite"`
	CookieTimeout                 float64 `json:"cookietimeout,omitempty" nitro:"permission=readwrite"`
	SiteDomainTimeToLive          float64 `json:"sitedomainttl,omitempty" nitro:"permission=readwrite"`
	NewName                       string  `json:"newname,omitempty" nitro:"permission=readwrite"`
	Ip                            string  `json:"ip,omitempty" nitro:"permission=readonly"`
	Value                         string  `json:"value,omitempty" nitro:"permission=readonly"`
	NodegroupName                 string  `json:"ngname,omitempty" nitro:"permission=readonly"`
	Type                          string  `json:"type,omitempty" nitro:"permission=readonly"`
	CurrentState                  string  `json:"curstate,omitempty" nitro:"permission=readonly"`
	Status                        int     `json:"status,omitempty" nitro:"permission=readonly"`
	CacheType                     string  `json:"casetype,omitempty" nitro:"permission=readonly"`
	Redirect                      string  `json:"redirect,omitempty" nitro:"permission=readonly"`
	Homepage                      string  `json:"homepage,omitempty" nitro:"permission=readonly"`
	DnsVserverName                string  `json:"dnsvservername,omitempty" nitro:"permission=readonly"`
	Domain                        string  `json:"domain,omitempty" nitro:"permission=readonly"`
	ServiceName                   string  `json:"servicename,omitempty" nitro:"permission=readonly"`
	Weight                        string  `json:"weight,omitempty" nitro:"permission=readonly"`
	CacheVserver                  string  `json:"cachevserver,omitempty" nitro:"permission=readonly"`
	TargetVserver                 string  `json:"targetvserver,omitempty" nitro:"permission=readonly"`
	Url                           string  `json:"url,omitempty" nitro:"permission=readonly"`
	Bindpoint                     string  `json:"bindpoint,omitempty" nitro:"permission=readonly"`
	GreaterThan2GBTransactions    string  `json:"gt2gb,omitempty" nitro:"permission=readonly"`
	StateChangeTimeSeconds        string  `json:"statechangetimesec,omitempty" nitro:"permission=readonly"`
	StateChangeTimeMilliSecond    string  `json:"statechangetimsec,omitempty" nitro:"permission=readonly"`
	RuleType                      string  `json:"ruletype,omitempty" nitro:"permission=readonly"`
	LbVserver                     string  `json:"lbvserver,omitempty" nitro:"permission=readonly"`
	TargetLbVserver               string  `json:"targetlbvserver,omitempty" nitro:"permission=readonly"`
	NoDefaultBindings             string  `json:"nodefaultbindings,omitempty" nitro:"permission=readonly"`
	Version                       int     `json:"version,omitempty" nitro:"permission=readonly"`
	Count                         float64 `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r CsVserver) GetTypeName() string {
	return "csvserver"
}

func NewCsVserverAddRequest(name string, servicetype string, ipaddress string, port int) CsVserver {
	return CsVserver{
		Name:        name,
		ServiceType: servicetype,
		Ipv46:       ipaddress,
		Port:        port,
	}
}

func NewCsVserverDisableRequest(name string) CsVserver {
	return CsVserver{
		Name: name,
	}
}

func NewCsVserverEnableRequest(name string) CsVserver {
	return CsVserver{
		Name: name,
	}
}

func NewCsVserverRenameRequest(oldName string, newName string) CsVserver {
	return CsVserver{
		Name:    oldName,
		NewName: newName,
	}
}
