package nitro

import (
	"context"

	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/config"
)

type NsSavedConfigHandler handler

func (h NsSavedConfigHandler) Get(ctx context.Context) (string, error) {
	var err error
	var savedConfig config.NsSavedConfig
	if savedConfig, err = getResource[config.NsSavedConfig](ctx, h.client, nil); err != nil {
		return "", err
	}
	return savedConfig.TextBlob, nil
}
