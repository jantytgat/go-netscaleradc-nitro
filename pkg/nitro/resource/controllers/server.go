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

	"github.com/corelayer/go-netscaleradc-nitro/pkg/nitro"
	"github.com/corelayer/go-netscaleradc-nitro/pkg/nitro/resource/config"
)

type ServerController struct {
	client *nitro.Client
}

func NewServerController(client *nitro.Client) *ServerController {
	c := ServerController{
		client: client,
	}
	return &c
}

func (c *ServerController) AddByDomain(name string, domain string) (*nitro.Response[config.Server], error) {
	r := nitro.Request[config.Server]{
		Method: http.MethodPost,
		Data: []config.Server{
			config.NewServerCreateByDomainRequest(name, domain),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *ServerController) AddByIpv4Address(name string, ip string) (*nitro.Response[config.Server], error) {
	r := nitro.Request[config.Server]{
		Method: http.MethodPost,
		Data: []config.Server{
			config.NewServerCreateByIpv4Request(name, ip),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *ServerController) AddByIpv6Address(name string, ip string) (*nitro.Response[config.Server], error) {
	r := nitro.Request[config.Server]{
		Method: http.MethodPost,
		Data: []config.Server{
			config.NewServerCreateByIpv6Request(name, ip),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *ServerController) Count() (*nitro.Response[config.Server], error) {
	r := nitro.Request[config.Server]{
		Method: http.MethodGet,
		Action: nitro.ActionCount,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *ServerController) Delete(name string) (*nitro.Response[config.Server], error) {
	r := nitro.Request[config.Server]{
		Method:       http.MethodDelete,
		ResourceName: name,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *ServerController) Disable(name string, delay float64, graceful bool) (*nitro.Response[config.Server], error) {
	r := nitro.Request[config.Server]{
		Method: http.MethodPost,
		Action: nitro.ActionDisable,
		Data: []config.Server{
			config.NewServerDisableRequest(name, delay, graceful),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *ServerController) Enable(name string) (*nitro.Response[config.Server], error) {
	r := nitro.Request[config.Server]{
		Method: http.MethodPost,
		Action: nitro.ActionEnable,
		Data:   []config.Server{{Name: name}},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *ServerController) Get(name string, attributes []string) (*nitro.Response[config.Server], error) {
	r := nitro.Request[config.Server]{
		Method:       http.MethodGet,
		ResourceName: name,
		Attributes:   attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *ServerController) List(filter map[string]string, attributes []string) (*nitro.Response[config.Server], error) {
	r := nitro.Request[config.Server]{
		Method:     http.MethodGet,
		Filter:     filter,
		Attributes: attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *ServerController) Rename(oldName string, newName string) (*nitro.Response[config.Server], error) {
	r := nitro.Request[config.Server]{
		Method: http.MethodPost,
		Action: nitro.ActionRename,
		Data: []config.Server{
			config.NewServerRenameRequest(oldName, newName),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

// Update TODO SERVER
func (c *ServerController) Update() (*nitro.Response[config.Server], error) {
	return nil, nitro.ControllerOperationNotImplementedError
}

// Unset TODO SERVER
func (c *ServerController) Unset() (*nitro.Response[config.Server], error) {
	return nil, nitro.ControllerOperationNotImplementedError
}
