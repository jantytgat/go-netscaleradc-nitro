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
	"github.com/corelayer/netscaleradc-nitro-go/pkg/nitro/resource/config"
)

type SslCertKeyController struct {
	client *nitro.Client
}

func NewSslCertKeyController(client *nitro.Client) *SslCertKeyController {
	c := SslCertKeyController{
		client: client,
	}
	return &c
}

func (c *SslCertKeyController) Add(name string, cer string, key string) (*nitro.Response[config.SslCertKey], error) {
	r := nitro.Request[config.SslCertKey]{
		Method: http.MethodPost,
		Data: []config.SslCertKey{
			config.NewSslCertKeyAddRequest(name, cer, key),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) Bind(vserver string, certkey string, sni bool) (*nitro.Response[config.SslVserver_SslCertKey_Binding], error) {
	r := nitro.Request[config.SslVserver_SslCertKey_Binding]{
		Method: http.MethodPut,
		Data: []config.SslVserver_SslCertKey_Binding{
			config.NewSslVserverCertificateBinding(vserver, certkey, sni),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) Update(name string, cer string, key string) (*nitro.Response[config.SslCertKey], error) {
	r := nitro.Request[config.SslCertKey]{
		Method: http.MethodPost,
		Action: nitro.ActionUpdate,
		Data: []config.SslCertKey{
			config.NewSslCertKeyUpdateRequest(name, cer, key),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) Count() (*nitro.Response[config.SslCertKey], error) {
	r := nitro.Request[config.SslCertKey]{
		Method: http.MethodGet,
		Action: nitro.ActionCount,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) ClearOcspStaplingCache(name string) (*nitro.Response[config.SslCertKey], error) {
	r := nitro.Request[config.SslCertKey]{
		Method: http.MethodPost,
		Action: nitro.ActionClear,
		Data: []config.SslCertKey{
			config.NewSslCertKeyClearOcspStaplingCacheRequest(name),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) Delete(name string) (*nitro.Response[config.SslCertKey], error) {
	r := nitro.Request[config.SslCertKey]{
		Method:       http.MethodPost,
		ResourceName: name,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) Get(name string, attributes []string) (*nitro.Response[config.SslCertKey], error) {
	r := nitro.Request[config.SslCertKey]{
		Method:       http.MethodGet,
		ResourceName: name,
		Attributes:   attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) Link(name string, caName string) (*nitro.Response[config.SslCertKey], error) {
	r := nitro.Request[config.SslCertKey]{
		Method: http.MethodPost,
		Action: nitro.ActionLink,
		Data: []config.SslCertKey{
			config.NewSslCertKeyLinkRequest(name, caName),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) List(filter map[string]string, attributes []string) (*nitro.Response[config.SslCertKey], error) {
	r := nitro.Request[config.SslCertKey]{
		Method:     http.MethodGet,
		Filter:     filter,
		Attributes: attributes,
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) Reload(name string, monitor bool, period string) (*nitro.Response[config.SslCertKey], error) {
	r := nitro.Request[config.SslCertKey]{
		Method: http.MethodPut,
		Data: []config.SslCertKey{
			config.NewSslCertKeyReloadRequest(name, monitor, period),
		},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) Unbind(vserver string, certkey string, sni bool) (*nitro.Response[config.SslVserver_SslCertKey_Binding], error) {
	r := nitro.Request[config.SslVserver_SslCertKey_Binding]{
		Method:       http.MethodDelete,
		ResourceName: vserver,
		Arguments:    map[string]string{"certkeyname": certkey, "snicert": strconv.FormatBool(sni)},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) Unlink(name string) (*nitro.Response[config.SslCertKey], error) {
	r := nitro.Request[config.SslCertKey]{
		Method: http.MethodPost,
		Action: nitro.ActionUnlink,
		Data:   []config.SslCertKey{{Name: name}},
	}
	return nitro.ExecuteNitroRequest(c.client, &r)
}
