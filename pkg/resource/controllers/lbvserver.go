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
	"github.com/corelayer/netscaleradc-nitro-go/pkg/resource/config"
	"github.com/corelayer/netscaleradc-nitro-go/pkg/resource/stat"
)

type LbVserverController struct {
	client *nitro.Client
}

func NewLbVserverController(client *nitro.Client) *LbVserverController {
	c := LbVserverController{
		client: client,
	}
	return &c
}

func (c *LbVserverController) Add(name string, servicetype string, ipaddress string, port int) (*nitro.Response[config.LbVserver], error) {
	r := nitro.Request[config.LbVserver]{
		Method: http.MethodPost,
		Data: []config.LbVserver{
			config.NewLbVserverAddRequest(name, servicetype, ipaddress, port),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *LbVserverController) Count() (*nitro.Response[config.LbVserver], error) {
	r := nitro.Request[config.LbVserver]{
		Method: http.MethodGet,
		Action: nitro.ActionCount,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *LbVserverController) Delete(name string) (*nitro.Response[config.LbVserver], error) {
	r := nitro.Request[config.LbVserver]{
		Method:       http.MethodDelete,
		ResourceName: name,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *LbVserverController) Disable(name string) (*nitro.Response[config.LbVserver], error) {
	r := nitro.Request[config.LbVserver]{
		Method: http.MethodPost,
		Action: nitro.ActionDisable,
		Data: []config.LbVserver{
			config.NewLbVserverDisableRequest(name),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *LbVserverController) Enable(name string) (*nitro.Response[config.LbVserver], error) {
	r := nitro.Request[config.LbVserver]{
		Method: http.MethodPost,
		Action: nitro.ActionEnable,
		Data: []config.LbVserver{
			config.NewLbVserverEnableRequest(name),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *LbVserverController) Get(name string, attributes []string) (*nitro.Response[config.LbVserver], error) {
	r := nitro.Request[config.LbVserver]{
		Method:       http.MethodGet,
		ResourceName: name,
		Attributes:   attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *LbVserverController) List(filter map[string]string, attributes []string) (*nitro.Response[config.LbVserver], error) {
	r := nitro.Request[config.LbVserver]{
		Method:     http.MethodGet,
		Filter:     filter,
		Attributes: attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *LbVserverController) Rename(oldname string, newname string) (*nitro.Response[config.LbVserver], error) {
	r := nitro.Request[config.LbVserver]{
		Method: http.MethodPost,
		Action: nitro.ActionRename,
		Data: []config.LbVserver{
			config.NewLbVserverRenameRequest(oldname, newname),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *LbVserverController) Stats(name string, attributes []string) (*nitro.Response[stat.LbVserver], error) {
	r := nitro.Request[stat.LbVserver]{
		Method:       http.MethodGet,
		ResourceName: name,
		Attributes:   attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}
