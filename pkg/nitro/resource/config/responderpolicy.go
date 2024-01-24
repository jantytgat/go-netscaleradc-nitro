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

type ResponderPolicy struct {
	Name            string   `json:"name,omitempty" nitro:"permission=readwrite"`
	Rule            string   `json:"rule,omitempty" nitro:"permission=readwrite"`
	Action          string   `json:"action,omitempty" nitro:"permission=readwrite"`
	UndefinedAction string   `json:"undefaction,omitempty" nitro:"permission=readwrite"`
	Comment         string   `json:"comment,omitempty" nitro:"permission=readwrite"`
	LogAction       string   `json:"logaction,omitempty" nitro:"permission=readwrite"`
	AppFlowAction   string   `json:"appflowaction,omitempty" nitro:"permission=readwrite"`
	NewName         string   `json:"newname,omitempty" nitro:"permission=readwrite"`
	Hits            float64  `json:"hits,omitempty" nitro:"permission=readonly"`
	UndefinedHits   float64  `json:"undefhits,omitempty" nitro:"permission=readonly"`
	Builtin         []string `json:"builtin,omitempty" nitro:"permission=readonly"`
	Feature         string   `json:"feature,omitempty" nitro:"permission=readonly"`
	Count           float64  `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r ResponderPolicy) GetTypeName() string {
	return "responderpolicy"
}

func NewResponderPolicyAddRequest(name string, rule string, action string, undefinedAction string) ResponderPolicy {
	return ResponderPolicy{
		Name:            name,
		Rule:            rule,
		Action:          action,
		UndefinedAction: undefinedAction,
	}
}

func NewResponderPolicyRenameRequest(oldName string, newName string) ResponderPolicy {
	return ResponderPolicy{
		Name:    oldName,
		NewName: newName,
	}
}
