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

package controllers

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/corelayer/netscaleradc-nitro-go/pkg/nitro"
	"github.com/corelayer/netscaleradc-nitro-go/pkg/resource/config"
)

type SystemFileController struct {
	client *nitro.Client
}

func NewSystemFileController(client *nitro.Client) *SystemFileController {
	c := SystemFileController{
		client: client,
	}

	return &c
}

func (c *SystemFileController) Add(fileName string, location string, fileContent []byte) (*nitro.Response[config.SystemFile], error) {
	r := nitro.Request[config.SystemFile]{
		Method: http.MethodPost,
		Data: []config.SystemFile{
			config.NewSystemFileAddRequest(fileName, location, fileContent),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

// Count TODO Check functionality, probably needs location
func (c *SystemFileController) Count(location string) (*nitro.Response[config.SystemFile], error) {
	r := nitro.Request[config.SystemFile]{
		Method:    http.MethodGet,
		Action:    nitro.ActionCount,
		Arguments: map[string]string{"filelocation": location},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SystemFileController) Delete(filename string, location string) (*nitro.Response[config.SystemFile], error) {
	r := nitro.Request[config.SystemFile]{
		Method:       http.MethodDelete,
		ResourceName: filename,
		Arguments:    map[string]string{"filelocation": location},
	}
	return nitro.ExecuteNitroRequest[config.SystemFile](c.client, &r)
}

func (c *SystemFileController) Download(filename string, location string) (*io.Reader, error) {
	var err error

	var res *nitro.Response[config.SystemFile]
	res, err = c.Get(filename, location, nil)
	if err != nil {
		return nil, err
	}

	if len(res.Data) != 1 {
		err = fmt.Errorf("invalid amount of files in data")
		return nil, err
	}

	var output io.Reader
	output = base64.NewDecoder(base64.StdEncoding, strings.NewReader(res.Data[0].Content))

	return &output, nil
}

func (c *SystemFileController) Get(filename string, location string, attributes []string) (*nitro.Response[config.SystemFile], error) {
	r := nitro.Request[config.SystemFile]{
		Method:     http.MethodGet,
		Arguments:  map[string]string{"filename": filename, "filelocation": location},
		Attributes: attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SystemFileController) List(location string, filter map[string]string, attributes []string) (*nitro.Response[config.SystemFile], error) {
	r := nitro.Request[config.SystemFile]{
		Method:     http.MethodGet,
		Arguments:  map[string]string{"filelocation": location},
		Filter:     filter,
		Attributes: attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}
