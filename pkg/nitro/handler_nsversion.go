package nitro

import (
	"context"

	"github.com/corelayer/go-netscaleradc-nitro/pkg/nitro/resource/config"
)

type NsVersionHandler handler

func (h NsVersionHandler) Get(ctx context.Context) (config.NsVersionDetail, error) {
	var err error
	var version config.NsVersion
	if version, err = getResource[config.NsVersion](ctx, h.client, "", nil); err != nil {
		return config.NsVersionDetail{}, err
	}
	return version.Details()
}
