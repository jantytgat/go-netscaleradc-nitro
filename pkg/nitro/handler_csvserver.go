package nitro

import (
	"context"

	"github.com/corelayer/go-netscaleradc-nitro/pkg/nitro/resource/config"
	"github.com/corelayer/go-netscaleradc-nitro/pkg/nitro/resource/stat"
)

type CsVserverHandler handler

func (h CsVserverHandler) Add(ctx context.Context, csVserver config.CsVserver) error {
	return addResource[config.CsVserver](ctx, h.client, csVserver)
}

func (h CsVserverHandler) Count(ctx context.Context) (float64, error) {
	var err error
	var r config.CsVserver
	if r, err = countResource[config.CsVserver](ctx, h.client); err != nil {
		return 0, err
	}
	return r.Count, nil
}

func (h CsVserverHandler) Delete(ctx context.Context, csVserverName string) error {
	return deleteResource[config.CsVserver](ctx, h.client, csVserverName, nil)
}

func (h CsVserverHandler) Disable(ctx context.Context, csVserverName string) error {
	return disableResource[config.CsVserver](ctx, h.client, config.CsVserver{Name: csVserverName})
}

func (h CsVserverHandler) Enable(ctx context.Context, csVserverName string) error {
	return enableResource[config.CsVserver](ctx, h.client, config.CsVserver{Name: csVserverName})
}

func (h CsVserverHandler) Get(ctx context.Context, csVserverName string, attributes []string) (config.CsVserver, error) {
	return getResource[config.CsVserver](ctx, h.client, csVserverName, attributes)
}

func (h CsVserverHandler) List(ctx context.Context, attributes []string, filter map[string]string) ([]config.CsVserver, error) {
	return listResource[config.CsVserver](ctx, h.client, attributes, filter)
}

func (h CsVserverHandler) Rename(ctx context.Context, oldName string, newName string) error {
	return renameResource[config.CsVserver](ctx, h.client, config.CsVserver{Name: oldName, NewName: newName})
}

func (h CsVserverHandler) Stats(ctx context.Context, csVserverName string, attributes []string) (stat.CsVserver, error) {
	return stats[stat.CsVserver](ctx, h.client, csVserverName, attributes)
}

func (h CsVserverHandler) Unset(ctx context.Context, r config.CsVserver) error {
	return OperationNotImplementedError
}

func (h CsVserverHandler) Update(ctx context.Context, r config.CsVserver) error {
	return OperationNotImplementedError
}
