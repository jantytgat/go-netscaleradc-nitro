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
	"github.com/corelayer/go-netscaleradc-nitro/pkg/nitro/resource/stat"
)

type SslVserverController struct {
	client *nitro.Client
}

func NewSslVserverController(client *nitro.Client) *SslVserverController {
	c := SslVserverController{
		client: client,
	}
	return &c
}

func (c *SslVserverController) Count() (*nitro.Response[config.SslVserver], error) {
	r := nitro.Request[config.SslVserver]{
		Method: http.MethodGet,
		Action: nitro.ActionCount,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SslVserverController) Get(name string, attributes []string) (*nitro.Response[config.SslVserver], error) {
	r := nitro.Request[config.SslVserver]{
		Method:       http.MethodGet,
		ResourceName: name,
		Attributes:   attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SslVserverController) List(filter map[string]string, attributes []string) (*nitro.Response[config.SslVserver], error) {
	r := nitro.Request[config.SslVserver]{
		Method:     http.MethodGet,
		Filter:     filter,
		Attributes: attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SslVserverController) Stats(name string, attributes []string) (*nitro.Response[stat.SslVserver], error) {
	r := nitro.Request[stat.SslVserver]{
		Method:       http.MethodGet,
		ResourceName: name,
		Attributes:   attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

// Update TODO SslVserverController

// Unset TODO SslVserverController
