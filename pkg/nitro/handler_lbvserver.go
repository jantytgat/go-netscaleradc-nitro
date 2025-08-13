package nitro

import (
	"context"

	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/config"
	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/stat"
)

type LbVserverHandler handler

func (h LbVserverHandler) Add(ctx context.Context, lbVserver config.LbVserver) error {
	return addResource[config.LbVserver](ctx, h.client, lbVserver)
}

func (h LbVserverHandler) Count(ctx context.Context) (float64, error) {
	var err error
	var r config.LbVserver
	if r, err = countResource[config.LbVserver](ctx, h.client); err != nil {
		return 0, err
	}
	return r.Count, nil
}

func (h LbVserverHandler) Delete(ctx context.Context, lbVserverName string) error {
	return deleteResource[config.LbVserver](ctx, h.client, lbVserverName, nil)
}

func (h LbVserverHandler) Disable(ctx context.Context, lbVserverName string) error {
	return disableResource[config.LbVserver](ctx, h.client, config.LbVserver{Name: lbVserverName})
}

func (h LbVserverHandler) Enable(ctx context.Context, lbVserverName string) error {
	return enableResource[config.LbVserver](ctx, h.client, config.LbVserver{Name: lbVserverName})
}

func (h LbVserverHandler) Get(ctx context.Context, lbVserverName string, attributes []string) (config.LbVserver, error) {
	return getResourceWithName[config.LbVserver](ctx, h.client, lbVserverName, attributes)
}

func (h LbVserverHandler) GetServiceBindings(ctx context.Context, lbVserverName string, attributes []string, filter map[string]string) ([]config.LbVserverServiceBinding, error) {
	return listResourceWithName[config.LbVserverServiceBinding](ctx, h.client, lbVserverName, attributes, filter)
}

func (h LbVserverHandler) GetServiceGroupBindings(ctx context.Context, lbVserverName string, attributes []string, filter map[string]string) ([]config.LbVserverServiceGroupBinding, error) {
	return listResourceWithName[config.LbVserverServiceGroupBinding](ctx, h.client, lbVserverName, attributes, filter)
}

func (h LbVserverHandler) List(ctx context.Context, attributes []string, filter map[string]string) ([]config.LbVserver, error) {
	return listResource[config.LbVserver](ctx, h.client, attributes, filter)
}

func (h LbVserverHandler) Rename(ctx context.Context, oldLbVserverName string, newLbVserverName string) error {
	return renameResource[config.LbVserver](ctx, h.client, config.LbVserver{Name: oldLbVserverName, NewName: newLbVserverName})
}

func (h LbVserverHandler) Stats(ctx context.Context, lbVserverName string, attributes []string) (stat.LbVserver, error) {
	return stats[stat.LbVserver](ctx, h.client, lbVserverName, attributes)
}

func (h LbVserverHandler) Unset(ctx context.Context, r config.LbVserver) error {
	return OperationNotImplementedError
}

func (h LbVserverHandler) Update(ctx context.Context, r config.LbVserver) error {
	return OperationNotImplementedError
}
