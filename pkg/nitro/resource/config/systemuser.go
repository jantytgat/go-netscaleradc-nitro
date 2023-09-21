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

type SystemUser struct {
	Username                       string   `json:"username,omitempty" nitro:"permission=readwrite"`
	Password                       string   `json:"password,omitempty" nitro:"permission=readwrite"`
	ExternalAuth                   string   `json:"externalauth,omitempty" nitro:"permission=readwrite"`
	PromptString                   string   `json:"promptstring,omitempty" nitro:"permission=readwrite"`
	Timeout                        float64  `json:"timeout,omitempty" nitro:"permission=readwrite"`
	Logging                        string   `json:"logging,omitempty" nitro:"permission=readwrite"`
	MaxSession                     float64  `json:"maxsession,omitempty" nitro:"permission=readwrite"`
	AllowedManagementInterface     []string `json:"allowedmanagementinterface,omitempty" nitro:"permission=readwrite"`
	Encrypted                      string   `json:"encrypted,omitempty" nitro:"permission=readonly"`
	HashMethod                     string   `json:"hashmethod,omitempty" nitro:"permission=readonly"`
	PromptInheritedFrom            string   `json:"promptinheritedfrom,omitempty" nitro:"permission=readonly"`
	TimeoutKind                    string   `json:"timeoutkind,omitempty" nitro:"permission=readonly"`
	AllowedManagementInterfaceKind string   `json:"allowedmanagementinterfacekind,omitempty" nitro:"permission=readonly"`
	LastPasswordChangeTimestamp    float64  `json:"lastpwdchangetimestamp,omitempty" nitro:"permission=readonly"`
	Count                          float64  `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r SystemUser) GetTypeName() string {
	return "systemuser"
}

func NewSystemUserAddRequest(username string, password string) SystemUser {
	return SystemUser{
		Username:     username,
		Password:     password,
		ExternalAuth: "DISABLED",
	}
}
