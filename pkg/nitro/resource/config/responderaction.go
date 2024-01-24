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

type ResponderAction struct {
	Name               string   `json:"name,omitempty" nitro:"permission=readwrite"`
	Type               string   `json:"type,omitempty" nitro:"permission=readwrite"`
	Target             string   `json:"target,omitempty" nitro:"permission=readwrite"`
	HtmlPage           string   `json:"htmlpage,omitempty" nitro:"permission=readwrite"`
	BypassSafetyCheck  string   `json:"bypasssafetycheck,omitempty" nitro:"permission=readwrite"`
	Comment            string   `json:"comment,omitempty" nitro:"permission=readwrite"`
	ResponseStatusCode float64  `json:"responsestatuscode,omitempty" nitro:"permission=readwrite"`
	ReasonPhrase       string   `json:"reasonphrase,omitempty" nitro:"permission=readwrite"`
	Headers            []string `json:"headers,omitempty" nitro:"permission=readwrite"`
	NewName            string   `json:"newName,omitempty" nitro:"permission=readwrite"`
	Hits               float64  `json:"hits,omitempty" nitro:"permission=readonly"`
	ReferenceCount     float64  `json:"referencecount,omitempty" nitro:"permission=readonly"`
	UndefinedHits      float64  `json:"undefinedhits,omitempty" nitro:"permission=readonly"`
	Builtin            []string `json:"builtin,omitempty" nitro:"permission=readonly"`
	Feature            string   `json:"feature,omitempty" nitro:"permission=readonly"`
	Count              float64  `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r ResponderAction) GetTypeName() string {
	return "responderaction"
}

func NewResponderActionAddRequest(name string, respondertype string, target string) ResponderAction {
	return ResponderAction{
		Name:   name,
		Type:   respondertype,
		Target: target,
	}
}

func NewResponderActionRenameRequest(oldName string, newName string) ResponderAction {
	return ResponderAction{
		Name:    oldName,
		NewName: newName,
	}
}
