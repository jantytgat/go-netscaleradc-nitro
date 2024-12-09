package nitro

import (
	"context"

	"github.com/corelayer/go-netscaleradc-nitro/pkg/nitro/resource/config"
)

type SystemUserHandler handler

func (h SystemUserHandler) Add(ctx context.Context, systemUser config.SystemUser) error {
	return addResource[config.SystemUser](ctx, h.client, systemUser)
}

func (h SystemUserHandler) Count(ctx context.Context) (float64, error) {
	var err error
	var r config.SystemUser
	if r, err = countResource[config.SystemUser](ctx, h.client); err != nil {
		return 0, err
	}
	return r.Count, nil
}

func (h SystemUserHandler) Delete(ctx context.Context, username string) error {
	return deleteResource[config.SystemUser](ctx, h.client, username, nil)
}

func (h SystemUserHandler) Get(ctx context.Context, username string, attributes []string) (config.SystemUser, error) {
	return getResource[config.SystemUser](ctx, h.client, username, attributes)
}

func (h SystemUserHandler) List(ctx context.Context, attributes []string, filter map[string]string) ([]config.SystemUser, error) {
	return listResource[config.SystemUser](ctx, h.client, attributes, filter)
}

func (h SystemUserHandler) Unset(ctx context.Context, r config.SystemUser) error {
	return OperationNotImplementedError
}

func (h SystemUserHandler) Update(ctx context.Context, r config.SystemUser) error {
	return OperationNotImplementedError
}
