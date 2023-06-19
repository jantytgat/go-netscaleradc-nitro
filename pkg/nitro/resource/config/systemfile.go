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

import "encoding/base64"

type SystemFile struct {
	Filename     string   `json:"filename,omitempty" nitro:"permission=readwrite"`
	Content      string   `json:"filecontent,omitempty" nitro:"permission=readwrite"`
	Location     string   `json:"filelocation,omitempty" nitro:"permission=readwrite"`
	Encoding     string   `json:"fileencoding,omitempty" nitro:"permission=readwrite"`
	AccessTime   string   `json:"fileaccesstime,omitempty" nitro:"permission=readonly"`
	ModifiedTime string   `json:"filemodifiedtime,omitempty" nitro:"permission=readonly"`
	Mode         []string `json:"filemode,omitempty" nitro:"permission=readonly"`
	Size         string   `json:"filesize,omitempty" nitro:"permission=readonly"`
	Count        float64  `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r SystemFile) GetTypeName() string {
	return "systemfile"
}

func NewSystemFileAddRequest(fileName string, location string, fileContent []byte) SystemFile {
	return SystemFile{
		Filename: fileName,
		Location: location,
		Content:  base64.StdEncoding.EncodeToString(fileContent),
		Encoding: "BASE64",
	}
}
