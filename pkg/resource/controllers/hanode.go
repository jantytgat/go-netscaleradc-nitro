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
	"strconv"

	"github.com/corelayer/netscaleradc-nitro-go/pkg/nitro"
	"github.com/corelayer/netscaleradc-nitro-go/pkg/resource/config"
	"github.com/corelayer/netscaleradc-nitro-go/pkg/resource/stat"
)

type HaNodeController struct {
	client *nitro.Client
}

func NewHaNodeController(client *nitro.Client) *HaNodeController {
	c := HaNodeController{
		client: client,
	}

	return &c
}

// Add TODO HANODE
func (c *HaNodeController) Add() (*nitro.Response[config.HaNode], error) {
	return nil, nitro.FormatNotImplementedError("Add")
}

func (c *HaNodeController) Count() (*nitro.Response[config.HaNode], error) {
	r := nitro.Request[config.HaNode]{
		Method: http.MethodGet,
		Action: nitro.ActionCount,
	}

	return nitro.ExecuteNitroRequest(c.client, &r)
}

// Delete TODO HANODE
func (c *HaNodeController) Delete() (*nitro.Response[config.HaNode], error) {
	return nil, nitro.FormatNotImplementedError("Delete")
}

func (c *HaNodeController) Failover(force bool) (*nitro.Response[config.HaFailover], error) {
	r := nitro.Request[config.HaFailover]{
		Method: http.MethodPost,
		Action: nitro.ActionForce,
		Data:   []config.HaFailover{{Force: force}},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *HaNodeController) Get(id int, attributes []string) (*nitro.Response[config.HaNode], error) {
	r := nitro.Request[config.HaNode]{
		Method:       http.MethodGet,
		ResourceName: strconv.Itoa(id),
		Attributes:   attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *HaNodeController) List(attributes []string) (*nitro.Response[config.HaNode], error) {
	r := nitro.Request[config.HaNode]{
		Method:     http.MethodGet,
		Attributes: attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *HaNodeController) Stats(attributes []string) (*nitro.Response[stat.HaNode], error) {
	r := nitro.Request[stat.HaNode]{
		Method:     http.MethodGet,
		Attributes: attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

// Update TODO HANODE
func (c *HaNodeController) Update() (*nitro.Response[config.HaNode], error) {
	return nil, nitro.FormatNotImplementedError("Update")
}
