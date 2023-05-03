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

type SystemBackup struct {
	Filename         string  `json:"filename,omitempty" nitro:"permission=readwrite"`
	Level            string  `json:"level,omitempty" nitro:"permission=readwrite"`
	IncludeKernel    string  `json:"includekernel,omitempty" nitro:"permission=readwrite"`
	Comment          string  `json:"comment,omitempty" nitro:"permission=readwrite"`
	Size             float64 `json:"size,omitempty" nitro:"permission=readonly"`
	UseLocalTimezone bool    `json:"uselocaltimezone,omitempty" nitro:"permission=readonly"`
	CreationTime     string  `json:"creationtime,omitempty" nitro:"permission=readonly"`
	Version          string  `json:"version,omitempty" nitro:"permission=readonly"`
	CreatedBy        string  `json:"createdby,omitempty" nitro:"permission=readonly"`
	IpAddress        string  `json:"ipaddress,omitempty" nitro:"permission=readonly"`
	SkipBackup       bool    `json:"skipbackup,omitempty" nitro:"permission=readwrite"`
	Count            float64 `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r SystemBackup) GetTypeName() string {
	return "systembackup"
}

func NewSystemBackupCreateRequest(name string, level SystemBackupLevel) SystemBackup {
	return SystemBackup{
		Filename: name,
		Level:    level.String(),
	}
}
