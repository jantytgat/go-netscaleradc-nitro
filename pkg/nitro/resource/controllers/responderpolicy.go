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

type ResponderPolicyController struct {
	client *nitro.Client
}

func NewResponderPolicyController(client *nitro.Client) *ResponderPolicyController {
	c := ResponderPolicyController{
		client: client,
	}
	return &c
}

func (c *ResponderPolicyController) Add(name string, rule string, action string, undefinedAction string) (*nitro.Response[config.ResponderPolicy], error) {
	r := nitro.Request[config.ResponderPolicy]{
		Method: http.MethodPost,
		Data: []config.ResponderPolicy{
			config.NewResponderPolicyAddRequest(name, rule, action, undefinedAction),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *ResponderPolicyController) AddRequest(r nitro.Request[config.ResponderPolicy]) (*nitro.Response[config.ResponderPolicy], error) {
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *ResponderPolicyController) Count() (*nitro.Response[config.ResponderPolicy], error) {
	r := nitro.Request[config.ResponderPolicy]{
		Method: http.MethodGet,
		Action: nitro.ActionCount,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *ResponderPolicyController) Delete(name string) (*nitro.Response[config.ResponderPolicy], error) {
	r := nitro.Request[config.ResponderPolicy]{
		Method:       http.MethodDelete,
		ResourceName: name,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *ResponderPolicyController) Get(name string, attributes []string) (*nitro.Response[config.ResponderPolicy], error) {
	r := nitro.Request[config.ResponderPolicy]{
		Method:       http.MethodGet,
		ResourceName: name,
		Attributes:   attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *ResponderPolicyController) List(filter map[string]string, attributes []string) (*nitro.Response[config.ResponderPolicy], error) {
	r := nitro.Request[config.ResponderPolicy]{
		Method:     http.MethodGet,
		Filter:     filter,
		Attributes: attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *ResponderPolicyController) Rename(oldname string, newname string) (*nitro.Response[config.ResponderPolicy], error) {
	r := nitro.Request[config.ResponderPolicy]{
		Method: http.MethodPost,
		Action: nitro.ActionRename,
		Data: []config.ResponderPolicy{
			config.NewResponderPolicyRenameRequest(oldname, newname),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

// func (c *ResponderPolicyController) Update(name string) {}
