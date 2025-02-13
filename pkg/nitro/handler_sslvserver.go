package nitro

import (
	"context"

	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/config"
	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/stat"
)

type SslVserverHandler handler

func (h SslVserverHandler) Count(ctx context.Context) (float64, error) {
	var err error
	var r config.SslVserver
	if r, err = countResource[config.SslVserver](ctx, h.client); err != nil {
		return 0, err
	}
	return r.Count, nil
}

func (h SslVserverHandler) Get(ctx context.Context, sslVserverName string, attributes []string) (config.SslVserver, error) {
	return getResource[config.SslVserver](ctx, h.client, sslVserverName, attributes)
}

func (h SslVserverHandler) List(ctx context.Context, attributes []string, filter map[string]string) ([]config.SslVserver, error) {
	return listResource[config.SslVserver](ctx, h.client, attributes, filter)
}

func (h SslVserverHandler) Stats(ctx context.Context, sslVserverName string, attributes []string) (stat.SslVserver, error) {
	return stats[stat.SslVserver](ctx, h.client, sslVserverName, attributes)
}

func (h SslVserverHandler) Unset(ctx context.Context, r config.SslVserver) error {
	return OperationNotImplementedError
}

func (h SslVserverHandler) Update(ctx context.Context, r config.SslVserver) error {
	return OperationNotImplementedError
}
