package nitro

import (
	"context"
	"net/http"
	"strconv"

	"github.com/corelayer/go-netscaleradc-nitro/pkg/nitro/resource/config"
)

type SslCertKeyHandler handler

func (h SslCertKeyHandler) Add(ctx context.Context, sslCertKey config.SslCertKey) error {
	return addResource[config.SslCertKey](ctx, h.client, sslCertKey)
}

func (h SslCertKeyHandler) BindSslService(ctx context.Context, binding config.SslServiceSslCertKeyBinding) error {
	return bindResource[config.SslServiceSslCertKeyBinding](ctx, h.client, binding)
}

func (h SslCertKeyHandler) BindSslVserver(ctx context.Context, binding config.SslVserverSslCertKeyBinding) error {
	return bindResource[config.SslVserverSslCertKeyBinding](ctx, h.client, binding)
}

func (h SslCertKeyHandler) ClearOcspStaplingCache(ctx context.Context, sslCertKeyName string) error {
	return clearResource[config.SslCertKey](ctx, h.client, config.NewSslCertKeyClearOcspStaplingCacheRequest(sslCertKeyName))
}

func (h SslCertKeyHandler) Count(ctx context.Context) (float64, error) {
	var err error
	var r config.SslCertKey
	if r, err = countResource[config.SslCertKey](ctx, h.client); err != nil {
		return 0, err
	}
	return r.Count, nil
}

func (h SslCertKeyHandler) CountSslServiceBindings(ctx context.Context) (float64, error) {
	var err error
	var r config.SslCertKeyServiceBinding
	if r, err = countResource[config.SslCertKeyServiceBinding](ctx, h.client); err != nil {
		return 0, err
	}
	return r.Count, nil
}

func (h SslCertKeyHandler) CountSslVserverBindings(ctx context.Context) (float64, error) {
	var err error
	var r config.SslCertKeySslVserverBinding
	if r, err = countResource[config.SslCertKeySslVserverBinding](ctx, h.client); err != nil {
		return 0, err
	}
	return r.Count, nil
}

func (h SslCertKeyHandler) Delete(ctx context.Context, sslCertKeyName string) error {
	return deleteResource[config.SslCertKey](ctx, h.client, sslCertKeyName, nil)
}

func (h SslCertKeyHandler) Get(ctx context.Context, sslCertKeyName string, attributes []string) (config.SslCertKey, error) {
	return getResource[config.SslCertKey](ctx, h.client, sslCertKeyName, attributes)
}

func (h SslCertKeyHandler) GetSslServiceBinding(ctx context.Context, sslServiceName string, attributes []string) (config.SslCertKeyServiceBinding, error) {
	return getResource[config.SslCertKeyServiceBinding](ctx, h.client, sslServiceName, attributes)
}

func (h SslCertKeyHandler) GetSslVserverBinding(ctx context.Context, sslVserverName string, attributes []string) (config.SslCertKeySslVserverBinding, error) {
	return getResource[config.SslCertKeySslVserverBinding](ctx, h.client, sslVserverName, attributes)
}

func (h SslCertKeyHandler) Link(ctx context.Context, sslCertKeyName, caSslCertKeyName string) error {
	return linkResource[config.SslCertKey](ctx, h.client, config.NewSslCertKeyLinkRequest(sslCertKeyName, caSslCertKeyName))
}

func (h SslCertKeyHandler) List(ctx context.Context, attributes []string, filter map[string]string) ([]config.SslCertKey, error) {
	return listResource[config.SslCertKey](ctx, h.client, attributes, filter)
}

func (h SslCertKeyHandler) Reload(ctx context.Context, sslCertKey config.SslCertKey) error {
	return putResource[config.SslCertKey](ctx, h.client, sslCertKey)
}

func (h SslCertKeyHandler) UnbindSslService(ctx context.Context, serviceName string, sslCertKeyName string, isSniCertBinding bool) error {
	return unbindResource[config.SslServiceSslCertKeyBinding](ctx, h.client, serviceName, map[string]string{"certkeyname": sslCertKeyName, "snicert": strconv.FormatBool(isSniCertBinding)})
}

func (h SslCertKeyHandler) UnbindSslVserver(ctx context.Context, sslVserverName string, sslCertKeyName string, isSniCertBinding bool) error {
	return unbindResource[config.SslVserverSslCertKeyBinding](ctx, h.client, sslVserverName, map[string]string{"certkeyname": sslCertKeyName, "snicert": strconv.FormatBool(isSniCertBinding)})
}

func (h SslCertKeyHandler) Unlink(ctx context.Context, sslCertKeyName string) error {
	return unlinkResource[config.SslCertKey](ctx, h.client, config.NewSslCertKeyUnlinkRequest(sslCertKeyName))
}

func (h SslCertKeyHandler) Update(ctx context.Context, r config.SslCertKey) error {
	req := Request[config.SslCertKey]{
		Method: http.MethodPost,
		Action: ActionUpdate,
		Data:   []config.SslCertKey{r},
	}

	if _, err := executeNitroRequest[config.SslCertKey](ctx, h.client, &req); err != nil {
		return err
	}
	return nil
}
