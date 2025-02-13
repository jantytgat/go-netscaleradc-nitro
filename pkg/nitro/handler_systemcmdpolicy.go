package nitro

import (
	"context"

	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/config"
)

type SystemCmdPolicyHandler handler

func (h SystemCmdPolicyHandler) Add(ctx context.Context, c config.SystemCmdPolicy) error {
	return addResource[config.SystemCmdPolicy](ctx, h.client, c)
}

func (h SystemCmdPolicyHandler) Count(ctx context.Context) (float64, error) {
	var err error
	var r config.SystemCmdPolicy
	if r, err = countResource[config.SystemCmdPolicy](ctx, h.client); err != nil {
		return 0, err
	}
	return r.Count, nil
}

func (h SystemCmdPolicyHandler) Delete(ctx context.Context, policyName string) error {
	return deleteResource[config.SystemCmdPolicy](ctx, h.client, policyName, nil)
}

func (h SystemCmdPolicyHandler) Get(ctx context.Context, policyName string, attributes []string) (config.SystemCmdPolicy, error) {
	return getResource[config.SystemCmdPolicy](ctx, h.client, policyName, attributes)
}

func (h SystemCmdPolicyHandler) List(ctx context.Context, attributes []string, filter map[string]string) ([]config.SystemCmdPolicy, error) {
	return listResource[config.SystemCmdPolicy](ctx, h.client, attributes, filter)
}

func (h SystemCmdPolicyHandler) Unset(ctx context.Context, r config.SystemCmdPolicy) error {
	return OperationNotImplementedError
}

func (h SystemCmdPolicyHandler) Update(ctx context.Context, r config.SystemCmdPolicy) error {
	return OperationNotImplementedError
}
