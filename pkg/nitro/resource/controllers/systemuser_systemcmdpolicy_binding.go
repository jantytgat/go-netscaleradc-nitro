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

type SystemUserSystemCmdPolicyBindingController struct {
	client *nitro.Client
}

func NewSystemUserSystemCmdPolicyBindingController(client *nitro.Client) *SystemUserSystemCmdPolicyBindingController {
	c := SystemUserSystemCmdPolicyBindingController{
		client: client,
	}
	return &c
}

func (c *SystemUserSystemCmdPolicyBindingController) Add(username string, policyname string, priority float64) (*nitro.Response[config.SystemUserSystemCmdPolicyBinding], error) {
	r := nitro.Request[config.SystemUserSystemCmdPolicyBinding]{
		Method: http.MethodPost,
		Data: []config.SystemUserSystemCmdPolicyBinding{
			config.NewSystemUserSystemCmdPolicyBindingAddRequest(username, policyname, priority),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SystemUserSystemCmdPolicyBindingController) AddRequest(r nitro.Request[config.SystemUserSystemCmdPolicyBinding]) (*nitro.Response[config.SystemUserSystemCmdPolicyBinding], error) {
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SystemUserSystemCmdPolicyBindingController) Count() (*nitro.Response[config.SystemUserSystemCmdPolicyBinding], error) {
	r := nitro.Request[config.SystemUserSystemCmdPolicyBinding]{
		Method: http.MethodGet,
		Action: nitro.ActionCount,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SystemUserSystemCmdPolicyBindingController) Delete(username string, policyname string) (*nitro.Response[config.SystemUserSystemCmdPolicyBinding], error) {
	r := nitro.Request[config.SystemUserSystemCmdPolicyBinding]{
		Method:       http.MethodDelete,
		ResourceName: username,
		Arguments: map[string]string{
			"policyname": policyname,
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SystemUserSystemCmdPolicyBindingController) Get(name string, attributes []string) (*nitro.Response[config.SystemUserSystemCmdPolicyBinding], error) {
	r := nitro.Request[config.SystemUserSystemCmdPolicyBinding]{
		Method:       http.MethodGet,
		ResourceName: name,
		Attributes:   attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SystemUserSystemCmdPolicyBindingController) List(filter map[string]string, attributes []string) (*nitro.Response[config.SystemUserSystemCmdPolicyBinding], error) {
	r := nitro.Request[config.SystemUserSystemCmdPolicyBinding]{
		Method:     http.MethodGet,
		Filter:     filter,
		Attributes: attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}
