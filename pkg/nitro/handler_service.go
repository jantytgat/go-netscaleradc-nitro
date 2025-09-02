package nitro

import (
	"context"

	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/config"
	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/stat"
)

type ServiceHandler handler

func (h ServiceHandler) Add(ctx context.Context, service config.Service) error {
	return addResource[config.Service](ctx, h.client, service)
}

func (h ServiceHandler) Count(ctx context.Context) (float64, error) {
	var err error
	var r config.Service
	if r, err = countResource[config.Service](ctx, h.client); err != nil {
		return 0, err
	}
	return r.Count, nil
}

func (h ServiceHandler) Delete(ctx context.Context, serviceName string) error {
	return deleteResource[config.Service](ctx, h.client, serviceName, nil)
}

func (h ServiceHandler) Disable(ctx context.Context, serviceName string) error {
	return disableResource[config.Service](ctx, h.client, config.Service{Name: serviceName})
}

func (h ServiceHandler) Enable(ctx context.Context, serviceName string) error {
	return enableResource[config.Service](ctx, h.client, config.Service{Name: serviceName})
}

func (h ServiceHandler) Get(ctx context.Context, serviceName string, attributes []string) (config.Service, error) {
	return getResourceWithName[config.Service](ctx, h.client, serviceName, attributes)
}

func (h ServiceHandler) List(ctx context.Context, attributes []string, filter map[string]string) ([]config.Service, error) {
	return listResource[config.Service](ctx, h.client, attributes, filter)
}

func (h ServiceHandler) Rename(ctx context.Context, oldName string, newName string) error {
	return renameResource[config.Service](ctx, h.client, config.Service{Name: oldName, NewName: newName})
}

func (h ServiceHandler) Stats(ctx context.Context, serviceName string, attributes []string) (stat.Service, error) {
	return statResourceWithName[stat.Service](ctx, h.client, serviceName, attributes)
}

func (h ServiceHandler) Unset(ctx context.Context, r config.Service) error {
	return OperationNotImplementedError
}

func (h ServiceHandler) Update(ctx context.Context, r config.Service) error {
	return OperationNotImplementedError
}
