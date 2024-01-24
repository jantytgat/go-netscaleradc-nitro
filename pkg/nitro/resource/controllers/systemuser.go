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

type SystemUserController struct {
	client *nitro.Client
}

func NewSystemUserController(client *nitro.Client) *SystemUserController {
	c := SystemUserController{
		client: client,
	}
	return &c
}

func (c *SystemUserController) Add(username string, password string) (*nitro.Response[config.SystemUser], error) {
	r := nitro.Request[config.SystemUser]{
		Method: http.MethodPost,
		Data: []config.SystemUser{
			config.NewSystemUserAddRequest(username, password),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SystemUserController) Count() (*nitro.Response[config.SystemUser], error) {
	r := nitro.Request[config.SystemUser]{
		Method: http.MethodGet,
		Action: nitro.ActionCount,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SystemUserController) Delete(name string) (*nitro.Response[config.SystemUser], error) {
	r := nitro.Request[config.SystemUser]{
		Method:       http.MethodDelete,
		ResourceName: name,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SystemUserController) Get(name string, attributes []string) (*nitro.Response[config.SystemUser], error) {
	r := nitro.Request[config.SystemUser]{
		Method:       http.MethodGet,
		ResourceName: name,
		Attributes:   attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SystemUserController) List(filter map[string]string, attributes []string) (*nitro.Response[config.SystemUser], error) {
	r := nitro.Request[config.SystemUser]{
		Method:     http.MethodGet,
		Filter:     filter,
		Attributes: attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

// Update TODO SystemUser
func (c *SystemUserController) Update() (*nitro.Response[config.SystemUser], error) {
	return nil, nitro.ControllerOperationNotImplementedError
}

// Unset TODO SystemUser
func (c *SystemUserController) Unset() (*nitro.Response[config.SystemUser], error) {
	return nil, nitro.ControllerOperationNotImplementedError
}
