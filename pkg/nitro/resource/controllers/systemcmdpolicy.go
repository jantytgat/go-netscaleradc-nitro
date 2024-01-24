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

type SystemCmdPolicyController struct {
	client *nitro.Client
}

func NewSystemCmdPolicyController(client *nitro.Client) *SystemCmdPolicyController {
	c := SystemCmdPolicyController{
		client: client,
	}
	return &c
}

func (c *SystemCmdPolicyController) Add(name string, action config.CmdPolicyAction, spec string) (*nitro.Response[config.SystemCmdPolicy], error) {
	r := nitro.Request[config.SystemCmdPolicy]{
		Method: http.MethodPost,
		Data: []config.SystemCmdPolicy{
			config.NewSystemCmdPolicyAddRequest(name, action, spec),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SystemCmdPolicyController) Count() (*nitro.Response[config.SystemCmdPolicy], error) {
	r := nitro.Request[config.SystemCmdPolicy]{
		Method: http.MethodGet,
		Action: nitro.ActionCount,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SystemCmdPolicyController) Delete(name string) (*nitro.Response[config.SystemCmdPolicy], error) {
	r := nitro.Request[config.SystemCmdPolicy]{
		Method:       http.MethodDelete,
		ResourceName: name,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SystemCmdPolicyController) Get(name string, attributes []string) (*nitro.Response[config.SystemCmdPolicy], error) {
	r := nitro.Request[config.SystemCmdPolicy]{
		Method:       http.MethodGet,
		ResourceName: name,
		Attributes:   attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SystemCmdPolicyController) List(filter map[string]string, attributes []string) (*nitro.Response[config.SystemCmdPolicy], error) {
	r := nitro.Request[config.SystemCmdPolicy]{
		Method:     http.MethodGet,
		Filter:     filter,
		Attributes: attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

// Update TODO SystemCmdPolicy
func (c *SystemCmdPolicyController) Update() (*nitro.Response[config.SystemCmdPolicy], error) {
	return nil, nitro.ControllerOperationNotImplementedError
}
