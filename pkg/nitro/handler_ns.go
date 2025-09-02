package nitro

import (
	"context"

	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/stat"
)

type NsHandler handler

func (h NsHandler) Stats(ctx context.Context, attributes []string) (stat.Ns, error) {
	return statResource[stat.Ns](ctx, h.client, attributes)
}
