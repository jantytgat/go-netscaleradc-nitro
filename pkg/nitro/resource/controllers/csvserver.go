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
	"github.com/corelayer/netscaleradc-nitro-go/pkg/nitro/resource/stat"
)

type CsVserverController struct {
	client *nitro.Client
}

func NewCsVserverController(client *nitro.Client) *CsVserverController {
	c := CsVserverController{
		client: client,
	}
	return &c
}

func (c *CsVserverController) Add(name string, servicetype string, ipaddress string, port int) (*nitro.Response[config.CsVserver], error) {
	r := nitro.Request[config.CsVserver]{
		Method: http.MethodPost,
		Data: []config.CsVserver{
			config.NewCsVserverAddRequest(name, servicetype, ipaddress, port),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *CsVserverController) Count() (*nitro.Response[config.CsVserver], error) {
	r := nitro.Request[config.CsVserver]{
		Method: http.MethodGet,
		Action: nitro.ActionCount,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *CsVserverController) Delete(name string) (*nitro.Response[config.CsVserver], error) {
	r := nitro.Request[config.CsVserver]{
		Method:       http.MethodDelete,
		ResourceName: name,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *CsVserverController) Disable(name string) (*nitro.Response[config.CsVserver], error) {
	r := nitro.Request[config.CsVserver]{
		Method: http.MethodPost,
		Action: nitro.ActionDisable,
		Data: []config.CsVserver{
			config.NewCsVserverDisableRequest(name),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *CsVserverController) Enable(name string) (*nitro.Response[config.CsVserver], error) {
	r := nitro.Request[config.CsVserver]{
		Method: http.MethodPost,
		Action: nitro.ActionEnable,
		Data: []config.CsVserver{
			config.NewCsVserverEnableRequest(name),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *CsVserverController) Get(name string, attributes []string) (*nitro.Response[config.CsVserver], error) {
	r := nitro.Request[config.CsVserver]{
		Method:       http.MethodGet,
		ResourceName: name,
		Attributes:   attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *CsVserverController) List(filter map[string]string, attributes []string) (*nitro.Response[config.CsVserver], error) {
	r := nitro.Request[config.CsVserver]{
		Method:     http.MethodGet,
		Filter:     filter,
		Attributes: attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *CsVserverController) Rename(oldname string, newname string) (*nitro.Response[config.CsVserver], error) {
	r := nitro.Request[config.CsVserver]{
		Method: http.MethodPost,
		Action: nitro.ActionRename,
		Data: []config.CsVserver{
			config.NewCsVserverRenameRequest(oldname, newname),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *CsVserverController) Stats(name string, attributes []string) (*nitro.Response[stat.CsVserver], error) {
	r := nitro.Request[stat.CsVserver]{
		Method:       http.MethodGet,
		ResourceName: name,
		Attributes:   attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}
