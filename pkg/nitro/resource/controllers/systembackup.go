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
	"io"
	"net/http"
	"strings"

	"github.com/corelayer/go-netscaleradc-nitro/pkg/nitro"
	"github.com/corelayer/go-netscaleradc-nitro/pkg/nitro/resource/config"
)

type SystemBackupController struct {
	client *nitro.Client
}

func NewBackupController(client *nitro.Client) *SystemBackupController {
	c := SystemBackupController{
		client: client,
	}

	return &c
}

// Add TODO SYSTEMBACKUP
func (c *SystemBackupController) Add() (*nitro.Response[config.SystemBackup], error) {
	return nil, nitro.ControllerOperationNotImplementedError
}

// Create sends a request to the configured Client to create a backup with the configured name and level
func (c *SystemBackupController) Create(name string, level string) (*nitro.Response[config.SystemBackup], error) {
	r := nitro.Request[config.SystemBackup]{
		Method: http.MethodPost,
		Action: nitro.ActionCreate,
		Data: []config.SystemBackup{
			config.NewSystemBackupCreateRequest(strings.TrimSuffix(name, ".tgz"), level),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SystemBackupController) Count() (*nitro.Response[config.SystemBackup], error) {
	r := nitro.Request[config.SystemBackup]{
		Method: http.MethodGet,
		Action: nitro.ActionCount,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SystemBackupController) Delete(name string) (*nitro.Response[config.SystemBackup], error) {
	r := nitro.Request[config.SystemBackup]{
		Method:       http.MethodDelete,
		ResourceName: name,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SystemBackupController) Download(name string) (*io.Reader, error) {
	// Check if backup exists
	_, err := c.Get(name, []string{"filename"})
	if err != nil {
		return nil, err
	}
	fc := NewSystemFileController(c.client)
	return fc.Download(name, "/var/ns_sys_backup")
}

func (c *SystemBackupController) Get(name string, attributes []string) (*nitro.Response[config.SystemBackup], error) {
	r := nitro.Request[config.SystemBackup]{
		Method:       http.MethodGet,
		ResourceName: name,
		Attributes:   attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SystemBackupController) List(arguments map[string]string, filter map[string]string, attributes []string) (*nitro.Response[config.SystemBackup], error) {
	r := nitro.Request[config.SystemBackup]{
		Method:     http.MethodGet,
		Arguments:  arguments,
		Filter:     filter,
		Attributes: attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SystemBackupController) Restore(filename string, skipBackup bool) (*nitro.Response[config.SystemBackup], error) {
	r := nitro.Request[config.SystemBackup]{
		Method: http.MethodPost,
		Action: nitro.ActionRestore,
		Data: []config.SystemBackup{{
			Filename:   filename,
			SkipBackup: skipBackup,
		}},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}
