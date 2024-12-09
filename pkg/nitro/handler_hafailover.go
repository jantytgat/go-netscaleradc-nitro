package nitro

import (
	"context"
	"net/http"

	"github.com/corelayer/go-netscaleradc-nitro/pkg/nitro/resource/config"
)

type HaFailOverHandler handler

func (h HaFailOverHandler) Force(ctx context.Context) error {
	req := Request[config.HaFailover]{
		Method: http.MethodPost,
		Action: ActionForce,
		Data:   []config.HaFailover{{Force: true}},
	}

	var err error
	if _, err = executeNitroRequest(ctx, h.client, &req); err != nil {
		return err
	}
	return nil
}
