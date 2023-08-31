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

type ResponderGlobalResponderPolicyBinding struct {
	Priority               float64 `json:"priority,omitempty" nitro:"permission=readwrite"`
	GlobalBindType         string  `json:"globalbindtype,omitempty" nitro:"permission=readwrite"`
	GotoPriorityExpression string  `json:"gotopriorityexpression,omitempty" nitro:"permission=readwrite"`
	PolicyName             string  `json:"policyname,omitempty" nitro:"permission=readwrite"`
	Type                   string  `json:"type,omitempty" nitro:"permission=readwrite"`
	LabelType              string  `json:"labeltype,omitempty" nitro:"permission=readwrite"`
	LabelName              string  `json:"labelname,omitempty" nitro:"permission=readwrite"`
	Invoke                 bool    `json:"invoke,omitempty" nitro:"permission=readwrite"`
	BoundPolicyNumber      float64 `json:"numpol,omitempty" nitro:"permission=readonly"`
	FlowType               float64 `json:"flowtype,omitempty" nitro:"permission=readonly"`
	Count                  float64 `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r ResponderGlobalResponderPolicyBinding) GetTypeName() string {
	return "responderglobal_responderpolicy_binding"
}

func NewResponderGlobalResponderPolicyBindingAddRequest(name string, bindType string, priority float64, gotoPriorityExpression string) ResponderGlobalResponderPolicyBinding {
	return ResponderGlobalResponderPolicyBinding{
		PolicyName:             name,
		Type:                   bindType,
		Priority:               priority,
		GotoPriorityExpression: gotoPriorityExpression,
	}
}
