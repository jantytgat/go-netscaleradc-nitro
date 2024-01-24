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

type SystemUserSystemCmdPolicyBinding struct {
	Priority   float64 `json:"priority,omitempty" nitro:"permission=readwrite"`
	PolicyName string  `json:"policyname,omitempty" nitro:"permission=readwrite"`
	Username   string  `json:"username,omitempty" nitro:"permission=readwrite"`
	Count      float64 `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r SystemUserSystemCmdPolicyBinding) GetTypeName() string {
	return "systemuser_systemcmdpolicy_binding"
}

func NewSystemUserSystemCmdPolicyBindingAddRequest(username string, policyname string, priority float64) SystemUserSystemCmdPolicyBinding {
	return SystemUserSystemCmdPolicyBinding{
		Priority:   priority,
		PolicyName: policyname,
		Username:   username,
	}
}
