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

type PolicyStringmapPatternBinding struct {
	Name    string  `json:"name,omitempty" nitro:"permission=readwrite"`
	Key     string  `json:"key,omitempty" nitro:"permission=readwrite"`
	Value   string  `json:"value,omitempty" nitro:"permission=readwrite"`
	Comment string  `json:"comment,omitempty" nitro:"permission=readwrite"`
	Count   float64 `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r PolicyStringmapPatternBinding) GetTypeName() string {
	return "policystringmap_pattern_binding"
}

func NewPolicyStringmapPatternBindingRequest(name string, key string, value string) PolicyStringmapPatternBinding {
	return PolicyStringmapPatternBinding{
		Name:  name,
		Key:   key,
		Value: value,
	}
}
