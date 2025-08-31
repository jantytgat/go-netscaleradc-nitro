package nitro

import (
	"context"

	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/config"
)

type NsModeHandler handler

func (h NsModeHandler) Disable(ctx context.Context, mode config.NsMode) error {
	return disableResource[config.NsMode](ctx, h.client, mode)
}

func (h NsModeHandler) Enable(ctx context.Context, mode config.NsMode) error {
	return enableResource[config.NsMode](ctx, h.client, mode)
}

func (h NsModeHandler) Get(ctx context.Context) (config.NsMode, error) {
	var err error
	var modes config.NsMode
	if modes, err = getResource[config.NsMode](ctx, h.client, nil); err != nil {
		return config.NsMode{}, err
	}
	return modes, nil
}
