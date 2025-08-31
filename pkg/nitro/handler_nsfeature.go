package nitro

import (
	"context"

	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/config"
)

type NsFeatureHandler handler

func (h NsFeatureHandler) Disable(ctx context.Context, feature config.NsFeature) error {
	return disableResource[config.NsFeature](ctx, h.client, feature)
}

func (h NsFeatureHandler) Enable(ctx context.Context, feature config.NsFeature) error {
	return enableResource[config.NsFeature](ctx, h.client, feature)
}

func (h NsFeatureHandler) Get(ctx context.Context) (config.NsFeature, error) {
	var err error
	var features config.NsFeature
	if features, err = getResource[config.NsFeature](ctx, h.client, nil); err != nil {
		return config.NsFeature{}, err
	}
	return features, nil
}
