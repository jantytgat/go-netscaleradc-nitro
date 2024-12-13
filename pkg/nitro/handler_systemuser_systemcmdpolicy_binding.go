package nitro

import (
	"context"

	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/config"
)

type SystemUserSystemCmdPolicyBindingHandler handler

func (h SystemUserSystemCmdPolicyBindingHandler) Add(ctx context.Context, binding config.SystemUserSystemCmdPolicyBinding) error {
	return addResource[config.SystemUserSystemCmdPolicyBinding](ctx, h.client, binding)
}

func (h SystemUserSystemCmdPolicyBindingHandler) Count(ctx context.Context) (float64, error) {
	var err error
	var r config.SystemUserSystemCmdPolicyBinding
	if r, err = countResource[config.SystemUserSystemCmdPolicyBinding](ctx, h.client); err != nil {
		return 0, err
	}
	return r.Count, nil
}

func (h SystemUserSystemCmdPolicyBindingHandler) Delete(ctx context.Context, username string, policyName string) error {
	return deleteResource[config.SystemUserSystemCmdPolicyBinding](ctx, h.client, username, map[string]string{"policyname": policyName})
}

func (h SystemUserSystemCmdPolicyBindingHandler) Get(ctx context.Context, username string, attributes []string) (config.SystemUserSystemCmdPolicyBinding, error) {
	return getResource[config.SystemUserSystemCmdPolicyBinding](ctx, h.client, username, attributes)
}

func (h SystemUserSystemCmdPolicyBindingHandler) List(ctx context.Context, attributes []string, filter map[string]string) ([]config.SystemUserSystemCmdPolicyBinding, error) {
	return listResource[config.SystemUserSystemCmdPolicyBinding](ctx, h.client, attributes, filter)
}

func (h SystemUserSystemCmdPolicyBindingHandler) Unset(ctx context.Context, r config.SystemUserSystemCmdPolicyBinding) error {
	return OperationNotImplementedError
}

func (h SystemUserSystemCmdPolicyBindingHandler) Update(ctx context.Context, r config.SystemUserSystemCmdPolicyBinding) error {
	return OperationNotImplementedError
}
