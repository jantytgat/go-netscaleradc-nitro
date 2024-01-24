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

type ResponderActionController struct {
	client *nitro.Client
}

func NewResponderActionController(client *nitro.Client) *ResponderActionController {
	c := ResponderActionController{
		client: client,
	}
	return &c
}

func (c *ResponderActionController) Add(name string, respondertype string, target string) (*nitro.Response[config.ResponderAction], error) {
	r := nitro.Request[config.ResponderAction]{
		Method: http.MethodPost,
		Data: []config.ResponderAction{
			config.NewResponderActionAddRequest(name, respondertype, target),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *ResponderActionController) AddRequest(r nitro.Request[config.ResponderAction]) (*nitro.Response[config.ResponderAction], error) {
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *ResponderActionController) Count() (*nitro.Response[config.ResponderAction], error) {
	r := nitro.Request[config.ResponderAction]{
		Method: http.MethodGet,
		Action: nitro.ActionCount,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *ResponderActionController) Delete(name string) (*nitro.Response[config.ResponderAction], error) {
	r := nitro.Request[config.ResponderAction]{
		Method:       http.MethodDelete,
		ResourceName: name,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *ResponderActionController) Get(name string, attributes []string) (*nitro.Response[config.ResponderAction], error) {
	r := nitro.Request[config.ResponderAction]{
		Method:       http.MethodGet,
		ResourceName: name,
		Attributes:   attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *ResponderActionController) List(filter map[string]string, attributes []string) (*nitro.Response[config.ResponderAction], error) {
	r := nitro.Request[config.ResponderAction]{
		Method:     http.MethodGet,
		Filter:     filter,
		Attributes: attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *ResponderActionController) Rename(oldname string, newname string) (*nitro.Response[config.ResponderAction], error) {
	r := nitro.Request[config.ResponderAction]{
		Method: http.MethodPost,
		Action: nitro.ActionRename,
		Data: []config.ResponderAction{
			config.NewResponderActionRenameRequest(oldname, newname),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

// func (c *ResponderActionController) Unset(name string) {}
// func (c *ResponderActionController) Update(name string) {}
