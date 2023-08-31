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

type ResponderGlobalResponderPolicyBindingController struct {
	client *nitro.Client
}

func NewResponderGlobalResponderPolicyBindingController(client *nitro.Client) *ResponderGlobalResponderPolicyBindingController {
	c := ResponderGlobalResponderPolicyBindingController{
		client: client,
	}
	return &c
}

func (c *ResponderGlobalResponderPolicyBindingController) Add(name string, bindType string, priority string, gotoPriorityExpression string) (*nitro.Response[config.ResponderGlobalResponderPolicyBinding], error) {
	r := nitro.Request[config.ResponderGlobalResponderPolicyBinding]{
		Method: http.MethodPost,
		Data: []config.ResponderGlobalResponderPolicyBinding{
			config.NewResponderGlobalResponderPolicyBindingAddRequest(name, bindType, priority, gotoPriorityExpression),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *ResponderGlobalResponderPolicyBindingController) AddRequest(r nitro.Request[config.ResponderGlobalResponderPolicyBinding]) (*nitro.Response[config.ResponderGlobalResponderPolicyBinding], error) {
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *ResponderGlobalResponderPolicyBindingController) Count() (*nitro.Response[config.ResponderGlobalResponderPolicyBinding], error) {
	r := nitro.Request[config.ResponderGlobalResponderPolicyBinding]{
		Method: http.MethodGet,
		Action: nitro.ActionCount,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *ResponderGlobalResponderPolicyBindingController) Delete(name string, bindType string) (*nitro.Response[config.ResponderGlobalResponderPolicyBinding], error) {
	r := nitro.Request[config.ResponderGlobalResponderPolicyBinding]{
		Method: http.MethodDelete,
		Arguments: map[string]string{
			"policyname": name,
			"type":       bindType,
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *ResponderGlobalResponderPolicyBindingController) Get(name string, attributes []string) (*nitro.Response[config.ResponderGlobalResponderPolicyBinding], error) {
	r := nitro.Request[config.ResponderGlobalResponderPolicyBinding]{
		Method:       http.MethodGet,
		ResourceName: name,
		Attributes:   attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *ResponderGlobalResponderPolicyBindingController) List(filter map[string]string, attributes []string) (*nitro.Response[config.ResponderGlobalResponderPolicyBinding], error) {
	r := nitro.Request[config.ResponderGlobalResponderPolicyBinding]{
		Method:     http.MethodGet,
		Filter:     filter,
		Attributes: attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}
