package nitro

import (
	"context"

	"github.com/corelayer/go-netscaleradc-nitro/pkg/nitro/resource/config"
)

type ServerHandler handler

func (h ServerHandler) Add(ctx context.Context, c config.Server) error {
	return addResource[config.Server](ctx, h.client, c)
}

func (h ServerHandler) Count(ctx context.Context) (float64, error) {
	var err error
	var r config.Server
	if r, err = countResource[config.Server](ctx, h.client); err != nil {
		return 0, err
	}
	return r.Count, nil
}

func (h ServerHandler) Delete(ctx context.Context, serverName string) error {
	return deleteResource[config.Server](ctx, h.client, serverName, nil)
}

func (h ServerHandler) Disable(ctx context.Context, serverName string) error {
	return disableResource[config.Server](ctx, h.client, config.Server{Name: serverName})
}

func (h ServerHandler) Enable(ctx context.Context, serverName string) error {
	return enableResource[config.Server](ctx, h.client, config.Server{Name: serverName})
}

func (h ServerHandler) Get(ctx context.Context, serverName string, attributes []string) (config.Server, error) {
	return getResource[config.Server](ctx, h.client, serverName, attributes)
}

func (h ServerHandler) List(ctx context.Context, attributes []string, filter map[string]string) ([]config.Server, error) {
	return listResource[config.Server](ctx, h.client, attributes, filter)
}

func (h ServerHandler) Rename(ctx context.Context, oldServerName string, newServerName string) error {
	return renameResource[config.Server](ctx, h.client, config.Server{Name: oldServerName, NewName: newServerName})
}

func (h ServerHandler) Update() (config.Server, error) {
	return config.Server{}, OperationNotImplementedError
}

func (h ServerHandler) Unset() (config.Server, error) {
	return config.Server{}, OperationNotImplementedError
}
