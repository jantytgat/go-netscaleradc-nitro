package nitro

import (
	"context"

	"github.com/corelayer/go-netscaleradc-nitro/pkg/nitro/resource/config"
)

type ResponderPolicyHandler handler

func (h ResponderPolicyHandler) Add(ctx context.Context, c config.ResponderPolicy) error {
	return addResource[config.ResponderPolicy](ctx, h.client, c)
}

func (h ResponderPolicyHandler) Count(ctx context.Context) (float64, error) {
	var err error
	var r config.ResponderPolicy
	if r, err = countResource[config.ResponderPolicy](ctx, h.client); err != nil {
		return 0, err
	}
	return r.Count, nil
}

func (h ResponderPolicyHandler) Delete(ctx context.Context, policyName string) error {
	return deleteResource[config.ResponderPolicy](ctx, h.client, policyName, nil)
}

func (h ResponderPolicyHandler) Get(ctx context.Context, name string, attributes []string) (config.ResponderPolicy, error) {
	return getResource[config.ResponderPolicy](ctx, h.client, name, attributes)
}

func (h ResponderPolicyHandler) List(ctx context.Context, attributes []string, filter map[string]string) ([]config.ResponderPolicy, error) {
	return listResource[config.ResponderPolicy](ctx, h.client, attributes, filter)
}

func (h ResponderPolicyHandler) Rename(ctx context.Context, oldName string, newName string) error {
	return renameResource[config.ResponderPolicy](ctx, h.client, config.ResponderPolicy{Name: oldName, NewName: newName})
}

func (h ResponderPolicyHandler) Unset(ctx context.Context, r config.ResponderPolicy) error {
	return OperationNotImplementedError
}

func (h ResponderPolicyHandler) Update(ctx context.Context, r config.ResponderPolicy) error {
	return OperationNotImplementedError
}
