package nitro

import (
	"context"

	"github.com/corelayer/go-netscaleradc-nitro/pkg/nitro/resource/config"
)

type ResponderGlobalResponderPolicyBindingHandler handler

func (h ResponderGlobalResponderPolicyBindingHandler) Add(ctx context.Context, binding config.ResponderGlobalResponderPolicyBinding) error {
	return addResource[config.ResponderGlobalResponderPolicyBinding](ctx, h.client, binding)
}

func (h ResponderGlobalResponderPolicyBindingHandler) Count(ctx context.Context) (float64, error) {
	var err error
	var r config.ResponderGlobalResponderPolicyBinding
	if r, err = countResource[config.ResponderGlobalResponderPolicyBinding](ctx, h.client); err != nil {
		return 0, err
	}
	return r.Count, nil
}

func (h ResponderGlobalResponderPolicyBindingHandler) Delete(ctx context.Context, policyname string, bindType string) error {
	return deleteResource[config.ResponderGlobalResponderPolicyBinding](ctx, h.client, "", map[string]string{"policyname": policyname, "type": bindType})
}

func (h ResponderGlobalResponderPolicyBindingHandler) Get(ctx context.Context, binding string, attributes []string) (config.ResponderGlobalResponderPolicyBinding, error) {
	return getResource[config.ResponderGlobalResponderPolicyBinding](ctx, h.client, binding, attributes)
}

func (h ResponderGlobalResponderPolicyBindingHandler) List(ctx context.Context, attributes []string, filter map[string]string) ([]config.ResponderGlobalResponderPolicyBinding, error) {
	return listResource[config.ResponderGlobalResponderPolicyBinding](ctx, h.client, attributes, filter)
}
