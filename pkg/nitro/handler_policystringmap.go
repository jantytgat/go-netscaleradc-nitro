package nitro

import (
	"context"
	"net/http"

	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/config"
)

type PolicyStringmapHandler handler

func (h PolicyStringmapHandler) Add(ctx context.Context, policyStringmap config.PolicyStringmap) error {
	return addResource[config.PolicyStringmap](ctx, h.client, policyStringmap)
}

func (h PolicyStringmapHandler) Bind(ctx context.Context, stringmapName string, key string, value string) error {
	return putResource[config.PolicyStringmapPatternBinding](ctx, h.client, config.PolicyStringmapPatternBinding{
		Name:  stringmapName,
		Key:   key,
		Value: value,
	})
}

func (h PolicyStringmapHandler) Count(ctx context.Context) (float64, error) {
	var err error
	var r config.PolicyStringmap
	if r, err = countResource[config.PolicyStringmap](ctx, h.client); err != nil {
		return 0, err
	}
	return r.Count, nil
}

func (h PolicyStringmapHandler) CountBindings(ctx context.Context, stringmapName string) (float64, error) {
	var err error
	var bindings []config.PolicyStringmapPatternBinding

	if bindings, err = h.GetBindings(ctx, stringmapName, []string{"name"}, nil); err != nil {
		return 0, err
	}
	return float64(len(bindings)), nil
}

func (h PolicyStringmapHandler) Delete(ctx context.Context, stringmapName string) error {
	return deleteResource[config.PolicyStringmap](ctx, h.client, stringmapName, nil)
}

func (h PolicyStringmapHandler) Get(ctx context.Context, stringmapName string, attributes []string) (config.PolicyStringmap, error) {
	return getResource[config.PolicyStringmap](ctx, h.client, stringmapName, attributes)
}

func (h PolicyStringmapHandler) GetBindings(ctx context.Context, stringmapName string, attributes []string, filter map[string]string) ([]config.PolicyStringmapPatternBinding, error) {
	return listResourceWithName[config.PolicyStringmapPatternBinding](ctx, h.client, stringmapName, attributes, filter)
}

func (h PolicyStringmapHandler) List(ctx context.Context, attributes []string, filter map[string]string) ([]config.PolicyStringmap, error) {
	return listResource[config.PolicyStringmap](ctx, h.client, attributes, filter)
}

func (h PolicyStringmapHandler) Unbind(ctx context.Context, stringmapName string, key string) error {
	var err error

	req := Request[config.PolicyStringmapPatternBinding]{
		Method:       http.MethodDelete,
		ResourceName: stringmapName,
		Arguments:    map[string]string{"key": key},
	}

	if _, err = executeNitroRequest[config.PolicyStringmapPatternBinding](ctx, h.client, &req); err != nil {
		return err
	}
	return nil
}

func (h PolicyStringmapHandler) Unset(ctx context.Context, r config.PolicyStringmap) error {
	return OperationNotImplementedError
}

func (h PolicyStringmapHandler) Update(ctx context.Context, r config.PolicyStringmap) error {
	return OperationNotImplementedError
}
