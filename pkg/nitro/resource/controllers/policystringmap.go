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

type PolicyStringmapController struct {
	client *nitro.Client
}

func NewPolicyStringmapController(client *nitro.Client) *PolicyStringmapController {
	c := PolicyStringmapController{
		client: client,
	}
	return &c
}

func (c *PolicyStringmapController) Add(name string) (*nitro.Response[config.PolicyStringmap], error) {
	r := nitro.Request[config.PolicyStringmap]{
		Method: http.MethodPost,
		Data:   []config.PolicyStringmap{{Name: name}},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *PolicyStringmapController) Bind(name string, key string, value string) (*nitro.Response[config.PolicyStringmapPatternBinding], error) {
	r := nitro.Request[config.PolicyStringmapPatternBinding]{
		Method: http.MethodPut,
		Data: []config.PolicyStringmapPatternBinding{
			config.NewPolicyStringmapPatternBindingRequest(name, key, value),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *PolicyStringmapController) Count() (*nitro.Response[config.PolicyStringmap], error) {
	r := nitro.Request[config.PolicyStringmap]{
		Method: http.MethodGet,
		Action: nitro.ActionCount,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *PolicyStringmapController) Delete(name string) (*nitro.Response[config.PolicyStringmap], error) {
	r := nitro.Request[config.PolicyStringmap]{
		Method:       http.MethodDelete,
		ResourceName: name,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *PolicyStringmapController) Get(name string, attributes []string) (*nitro.Response[config.PolicyStringmap], error) {
	r := nitro.Request[config.PolicyStringmap]{
		Method:       http.MethodGet,
		ResourceName: name,
		Attributes:   attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

// GetBindings TODO
func (c *PolicyStringmapController) GetBindings(name string, filter map[string]string) (*nitro.Response[config.PolicyStringmapPatternBinding], error) {
	r := nitro.Request[config.PolicyStringmapPatternBinding]{
		Method:       http.MethodGet,
		ResourceName: name,
		Filter:       filter,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *PolicyStringmapController) List(attributes []string) (*nitro.Response[config.PolicyStringmap], error) {
	r := nitro.Request[config.PolicyStringmap]{
		Method:     http.MethodGet,
		Attributes: attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *PolicyStringmapController) Unbind(name string, key string) (*nitro.Response[config.PolicyStringmapPatternBinding], error) {
	r := nitro.Request[config.PolicyStringmapPatternBinding]{
		Method:       http.MethodDelete,
		ResourceName: name,
		Arguments:    map[string]string{"key": key},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}
