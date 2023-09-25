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

type DnsAddRecController struct {
	client *nitro.Client
}

func NewDnsAddRecController(client *nitro.Client) *DnsAddRecController {
	c := DnsAddRecController{
		client: client,
	}
	return &c
}

func (c *DnsAddRecController) Add(hostname string, ipaddress string, ttl float64) (*nitro.Response[config.DnsAddRec], error) {
	r := nitro.Request[config.DnsAddRec]{
		Method: http.MethodPost,
		Data: []config.DnsAddRec{
			config.NewDnsAddRecAddRequest(hostname, ipaddress, ttl),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *DnsAddRecController) Count() (*nitro.Response[config.DnsAddRec], error) {
	r := nitro.Request[config.DnsAddRec]{
		Method: http.MethodGet,
		Action: nitro.ActionCount,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *DnsAddRecController) Delete(hostname string, ipaddress string) (*nitro.Response[config.DnsAddRec], error) {
	r := nitro.Request[config.DnsAddRec]{
		Method:       http.MethodDelete,
		ResourceName: hostname,
		Arguments: map[string]string{
			"ipaddress": ipaddress,
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *DnsAddRecController) Get(name string, attributes []string) (*nitro.Response[config.DnsAddRec], error) {
	r := nitro.Request[config.DnsAddRec]{
		Method:       http.MethodGet,
		ResourceName: name,
		Attributes:   attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *DnsAddRecController) List(filter map[string]string, attributes []string) (*nitro.Response[config.DnsAddRec], error) {
	r := nitro.Request[config.DnsAddRec]{
		Method:     http.MethodGet,
		Filter:     filter,
		Attributes: attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}
