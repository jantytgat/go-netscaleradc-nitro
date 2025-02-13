package nitro

import (
	"context"

	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/config"
)

type ResponderActionHandler handler

func (h ResponderActionHandler) Add(ctx context.Context, c config.ResponderAction) error {
	return addResource[config.ResponderAction](ctx, h.client, c)
}

func (h ResponderActionHandler) Count(ctx context.Context) (float64, error) {
	var err error
	var r config.ResponderAction
	if r, err = countResource[config.ResponderAction](ctx, h.client); err != nil {
		return 0, err
	}
	return r.Count, nil
}

func (h ResponderActionHandler) Delete(ctx context.Context, actionName string) error {
	return deleteResource[config.ResponderAction](ctx, h.client, actionName, nil)
}

func (h ResponderActionHandler) Get(ctx context.Context, actionName string, attributes []string) (config.ResponderAction, error) {
	return getResource[config.ResponderAction](ctx, h.client, actionName, attributes)
}

func (h ResponderActionHandler) List(ctx context.Context, attributes []string, filter map[string]string) ([]config.ResponderAction, error) {
	return listResource[config.ResponderAction](ctx, h.client, attributes, filter)
}

func (h ResponderActionHandler) Rename(ctx context.Context, oldActionName string, newActionName string) error {
	return renameResource[config.ResponderAction](ctx, h.client, config.ResponderAction{Name: oldActionName, NewName: newActionName})
}

func (h ResponderActionHandler) Unset(ctx context.Context, r config.ResponderAction) error {
	return OperationNotImplementedError
}

func (h ResponderActionHandler) Update(ctx context.Context, r config.ResponderAction) error {
	return OperationNotImplementedError
}
