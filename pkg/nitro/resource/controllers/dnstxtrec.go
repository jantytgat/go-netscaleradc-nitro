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
	"net/http"

	"github.com/corelayer/netscaleradc-nitro-go/pkg/nitro"
	"github.com/corelayer/netscaleradc-nitro-go/pkg/nitro/resource/config"
)

type DnsTxtRecController struct {
	client *nitro.Client
}

func NewDnsTxtRecController(client *nitro.Client) *DnsTxtRecController {
	c := DnsTxtRecController{
		client: client,
	}
	return &c
}

func (c *DnsTxtRecController) Add(domain string, data []string, ttl float64) (*nitro.Response[config.DnsTxtRec], error) {
	r := nitro.Request[config.DnsTxtRec]{
		Method: http.MethodPost,
		Data: []config.DnsTxtRec{
			config.NewDnsTxtRecAddRequest(domain, data, ttl),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *DnsTxtRecController) Count() (*nitro.Response[config.DnsTxtRec], error) {
	r := nitro.Request[config.DnsTxtRec]{
		Method: http.MethodGet,
		Action: nitro.ActionCount,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *DnsTxtRecController) Delete(domain string, recordId string) (*nitro.Response[config.DnsTxtRec], error) {
	r := nitro.Request[config.DnsTxtRec]{
		Method:       http.MethodDelete,
		ResourceName: domain,
		Arguments: map[string]string{
			"recordid": recordId,
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *DnsTxtRecController) Get(name string, attributes []string) (*nitro.Response[config.DnsTxtRec], error) {
	r := nitro.Request[config.DnsTxtRec]{
		Method:       http.MethodGet,
		ResourceName: name,
		Attributes:   attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *DnsTxtRecController) List(filter map[string]string, attributes []string) (*nitro.Response[config.DnsTxtRec], error) {
	r := nitro.Request[config.DnsTxtRec]{
		Method:     http.MethodGet,
		Filter:     filter,
		Attributes: attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}
