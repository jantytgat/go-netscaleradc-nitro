package nitro

import (
	"context"

	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/config"
	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/stat"
)

type ServiceGroupHandler handler

func (h ServiceGroupHandler) Add(ctx context.Context, serviceGroup config.ServiceGroup) error {
	return addResource[config.ServiceGroup](ctx, h.client, serviceGroup)
}

func (h ServiceGroupHandler) Count(ctx context.Context) (float64, error) {
	var err error
	var r config.ServiceGroup
	if r, err = countResource[config.ServiceGroup](ctx, h.client); err != nil {
		return 0, err
	}
	return r.Count, nil
}

func (h ServiceGroupHandler) CountServiceGroupMemberBindings(ctx context.Context, serviceGroupName string) (float64, error) {
	var err error
	var r config.ServiceGroupServiceGroupMemberBinding
	if r, err = countResourceWithName[config.ServiceGroupServiceGroupMemberBinding](ctx, h.client, serviceGroupName); err != nil {
		return 0, err
	}
	return r.Count, nil
}

func (h ServiceGroupHandler) Delete(ctx context.Context, serviceGroupName string) error {
	return deleteResource[config.ServiceGroup](ctx, h.client, serviceGroupName, nil)
}

func (h ServiceGroupHandler) Disable(ctx context.Context, serviceGroupName string) error {
	return disableResource[config.ServiceGroup](ctx, h.client, config.ServiceGroup{Name: serviceGroupName})
}

func (h ServiceGroupHandler) Enable(ctx context.Context, serviceGroupName string) error {
	return enableResource[config.ServiceGroup](ctx, h.client, config.ServiceGroup{Name: serviceGroupName})
}

func (h ServiceGroupHandler) Get(ctx context.Context, serviceGroupName string, attributes []string) (config.ServiceGroup, error) {
	return getResourceWithName[config.ServiceGroup](ctx, h.client, serviceGroupName, attributes)
}

func (h ServiceGroupHandler) GetServiceGroupMemberBindings(ctx context.Context, serviceGroupName string, attributes []string, filter map[string]string) ([]config.ServiceGroupServiceGroupMemberBinding, error) {
	return listResourceWithName[config.ServiceGroupServiceGroupMemberBinding](ctx, h.client, serviceGroupName, attributes, filter)
}

func (h ServiceGroupHandler) List(ctx context.Context, attributes []string, filter map[string]string) ([]config.ServiceGroup, error) {
	return listResource[config.ServiceGroup](ctx, h.client, attributes, filter)
}

func (h ServiceGroupHandler) Rename(ctx context.Context, oldName string, newName string) error {
	return renameResource[config.ServiceGroup](ctx, h.client, config.ServiceGroup{Name: oldName, NewName: newName})
}

func (h ServiceGroupHandler) Stats(ctx context.Context, serviceGroupName string, attributes []string) (stat.ServiceGroup, error) {
	return stats[stat.ServiceGroup](ctx, h.client, serviceGroupName, attributes)
}

func (h ServiceGroupHandler) Unset(ctx context.Context, r config.ServiceGroup) error {
	return OperationNotImplementedError
}

func (h ServiceGroupHandler) Update(ctx context.Context, r config.ServiceGroup) error {
	return OperationNotImplementedError
}
