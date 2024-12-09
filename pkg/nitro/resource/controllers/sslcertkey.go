package controllers

import (
	"net/http"
	"strconv"

	"github.com/corelayer/go-netscaleradc-nitro/pkg/nitro"
	"github.com/corelayer/go-netscaleradc-nitro/pkg/nitro/resource/config"
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
	return nitro.executeNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) AddWithPassphrase(name string, cer string, key string, passphrase string) (*nitro.Response[config.SslCertKey], error) {
	r := nitro.Request[config.SslCertKey]{
		Method: http.MethodPost,
		Data: []config.SslCertKey{
			config.NewSslCertKeyWithPassphraseRequest(name, cer, key, passphrase),
		},
	}
	return nitro.executeNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) AddBundle(name string, cer string, key string) (*nitro.Response[config.SslCertKey], error) {
	r := nitro.Request[config.SslCertKey]{
		Method: http.MethodPost,
		Data: []config.SslCertKey{
			config.NewSslCertKeyBundleAddRequest(name, cer, key),
		},
	}
	return nitro.executeNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) AddBundleWithPassphrase(name string, cer string, key string, passphrase string) (*nitro.Response[config.SslCertKey], error) {
	r := nitro.Request[config.SslCertKey]{
		Method: http.MethodPost,
		Data: []config.SslCertKey{
			config.NewSslCertKeyBundleWithPassphraseAddRequest(name, cer, key, passphrase),
		},
	}
	return nitro.executeNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) BindSslService(service string, certkey string, sni bool) (*nitro.Response[config.SslServiceSslCertKeyBinding], error) {
	r := nitro.Request[config.SslServiceSslCertKeyBinding]{
		Method: http.MethodPut,
		Data: []config.SslServiceSslCertKeyBinding{
			config.NewSslServiceCertificateBindingAddRequest(service, certkey, sni),
		},
	}
	return nitro.executeNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) BindSslVserver(vserver string, certkey string, sni bool) (*nitro.Response[config.SslVserverSslCertKeyBinding], error) {
	r := nitro.Request[config.SslVserverSslCertKeyBinding]{
		Method: http.MethodPut,
		Data: []config.SslVserverSslCertKeyBinding{
			config.NewSslVserverCertificateBindingAddRequest(vserver, certkey, sni),
		},
	}
	return nitro.executeNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) Count() (*nitro.Response[config.SslCertKey], error) {
	r := nitro.Request[config.SslCertKey]{
		Method: http.MethodGet,
		Action: nitro.ActionCount,
	}
	return nitro.executeNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) CountServiceBindings() (*nitro.Response[config.SslCertKeyServiceBinding], error) {
	r := nitro.Request[config.SslCertKeyServiceBinding]{
		Method: http.MethodGet,
		Action: nitro.ActionCount,
	}
	return nitro.executeNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) CountSslVserverBindings() (*nitro.Response[config.SslCertKeySslVserverBinding], error) {
	r := nitro.Request[config.SslCertKeySslVserverBinding]{
		Method: http.MethodGet,
		Action: nitro.ActionCount,
	}
	return nitro.executeNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) ClearOcspStaplingCache(name string) (*nitro.Response[config.SslCertKey], error) {
	r := nitro.Request[config.SslCertKey]{
		Method: http.MethodPost,
		Action: nitro.ActionClear,
		Data: []config.SslCertKey{
			config.NewSslCertKeyClearOcspStaplingCacheRequest(name),
		},
	}
	return nitro.executeNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) Delete(name string) (*nitro.Response[config.SslCertKey], error) {
	r := nitro.Request[config.SslCertKey]{
		Method:       http.MethodPost,
		ResourceName: name,
	}
	return nitro.executeNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) Get(name string, attributes []string) (*nitro.Response[config.SslCertKey], error) {
	r := nitro.Request[config.SslCertKey]{
		Method:       http.MethodGet,
		ResourceName: name,
		Attributes:   attributes,
	}
	return nitro.executeNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) GetServiceBinding(name string, attributes []string, filter map[string]string) (*nitro.Response[config.SslCertKeyServiceBinding], error) {
	r := nitro.Request[config.SslCertKeyServiceBinding]{
		Method:       http.MethodGet,
		ResourceName: name,
		Attributes:   attributes,
		Filter:       filter,
	}
	return nitro.executeNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) GetSslVserverBinding(name string, attributes []string, filter map[string]string) (*nitro.Response[config.SslCertKeySslVserverBinding], error) {
	r := nitro.Request[config.SslCertKeySslVserverBinding]{
		Method:       http.MethodGet,
		ResourceName: name,
		Attributes:   attributes,
		Filter:       filter,
	}
	return nitro.executeNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) Link(name string, caName string) (*nitro.Response[config.SslCertKey], error) {
	r := nitro.Request[config.SslCertKey]{
		Method: http.MethodPost,
		Action: nitro.ActionLink,
		Data: []config.SslCertKey{
			config.NewSslCertKeyLinkRequest(name, caName),
		},
	}
	return nitro.executeNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) List(filter map[string]string, attributes []string) (*nitro.Response[config.SslCertKey], error) {
	r := nitro.Request[config.SslCertKey]{
		Method:     http.MethodGet,
		Filter:     filter,
		Attributes: attributes,
	}
	return nitro.executeNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) Reload(name string, monitor bool, period string) (*nitro.Response[config.SslCertKey], error) {
	r := nitro.Request[config.SslCertKey]{
		Method: http.MethodPut,
		Data: []config.SslCertKey{
			config.NewSslCertKeyReloadRequest(name, monitor, period),
		},
	}
	return nitro.executeNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) UnbindSslService(service string, certkey string, sni bool) (*nitro.Response[config.SslServiceSslCertKeyBinding], error) {
	r := nitro.Request[config.SslServiceSslCertKeyBinding]{
		Method:       http.MethodDelete,
		ResourceName: service,
		Arguments:    map[string]string{"certkeyname": certkey, "snicert": strconv.FormatBool(sni)},
	}
	return nitro.executeNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) UnbindSslVserver(vserver string, certkey string, sni bool) (*nitro.Response[config.SslVserverSslCertKeyBinding], error) {
	r := nitro.Request[config.SslVserverSslCertKeyBinding]{
		Method:       http.MethodDelete,
		ResourceName: vserver,
		Arguments:    map[string]string{"certkeyname": certkey, "snicert": strconv.FormatBool(sni)},
	}
	return nitro.executeNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) Unlink(name string) (*nitro.Response[config.SslCertKey], error) {
	r := nitro.Request[config.SslCertKey]{
		Method: http.MethodPost,
		Action: nitro.ActionUnlink,
		Data:   []config.SslCertKey{{Name: name}},
	}
	return nitro.executeNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) Update(name string, cer string, key string, noDomainCheck bool) (*nitro.Response[config.SslCertKey], error) {
	r := nitro.Request[config.SslCertKey]{
		Method: http.MethodPost,
		Action: nitro.ActionUpdate,
		Data: []config.SslCertKey{
			config.NewSslCertKeyUpdateRequest(name, cer, key, noDomainCheck),
		},
	}
	return nitro.executeNitroRequest(c.client, &r)
}

func (c *SslCertKeyController) UpdateWithPassphrase(name string, cer string, key string, passphrase string, noDomainCheck bool) (*nitro.Response[config.SslCertKey], error) {
	r := nitro.Request[config.SslCertKey]{
		Method: http.MethodPost,
		Action: nitro.ActionUpdate,
		Data: []config.SslCertKey{
			config.NewSslCertKeyWithPassphraseUpdateRequest(name, cer, key, passphrase, noDomainCheck),
		},
	}
	return nitro.executeNitroRequest(c.client, &r)
}
