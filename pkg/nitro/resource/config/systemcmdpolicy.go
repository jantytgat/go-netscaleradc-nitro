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

var (
	CmdActionAllow = CmdPolicyAction{"ALLOW"}
	CmdActionDeny  = CmdPolicyAction{"DENY"}
)

type CmdPolicyAction struct {
	string
}

type SystemCmdPolicy struct {
	PolicyName string   `json:"policyname,omitempty" nitro:"permission=readwrite"`
	Action     string   `json:"action" nitro:"permission=readwrite"`
	CmdSpec    string   `json:"cmdspec,omitempty" nitro:"permission=readwrite"`
	Builtin    []string `json:"builtin,omitempty" nitro:"permission=readonly"`
	Feature    string   `json:"feature,omitempty" nitro:"permission=readonly"`
	Count      float64  `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r SystemCmdPolicy) GetTypeName() string {
	return "systemcmdpolicy"
}

func NewSystemCmdPolicyAddRequest(name string, action CmdPolicyAction, spec string) SystemCmdPolicy {
	return SystemCmdPolicy{
		PolicyName: name,
		Action:     action.string,
		CmdSpec:    spec,
	}
}
