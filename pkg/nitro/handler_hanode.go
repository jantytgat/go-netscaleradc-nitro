package nitro

import (
	"context"
	"net/http"
	"strconv"

	"github.com/corelayer/go-netscaleradc-nitro/pkg/nitro/resource/config"
	"github.com/corelayer/go-netscaleradc-nitro/pkg/nitro/resource/stat"
)

type HaNodeHandler handler

func (h HaNodeHandler) Add(ctx context.Context, haNode config.HaNode) error {
	return addResource[config.HaNode](ctx, h.client, haNode)
}

func (h HaNodeHandler) Count(ctx context.Context) (float64, error) {
	var err error
	var r config.HaNode
	if r, err = countResource[config.HaNode](ctx, h.client); err != nil {

		return 0, err
	}
	return r.Count, nil
}

func (h HaNodeHandler) Delete(ctx context.Context, haNodeId int) error {
	return deleteResource[config.HaNode](ctx, h.client, strconv.Itoa(haNodeId), nil)
}

func (h HaNodeHandler) FailOver(ctx context.Context, force bool) (config.HaFailover, error) {
	req := Request[config.HaFailover]{
		Method: http.MethodPost,
		Action: ActionForce,
		Data:   []config.HaFailover{{Force: force}},
	}

	var res *Response[config.HaFailover]
	var err error
	if res, err = executeNitroRequest(ctx, h.client, &req); err != nil {
		return config.HaFailover{}, err
	}
	return res.Data[0], nil
}

func (h HaNodeHandler) Get(ctx context.Context, haNodeId int, attributes []string) (config.HaNode, error) {
	return getResource[config.HaNode](ctx, h.client, strconv.Itoa(haNodeId), attributes)
}

func (h HaNodeHandler) List(ctx context.Context, attributes []string, filter map[string]string) ([]config.HaNode, error) {
	return listResource[config.HaNode](ctx, h.client, attributes, filter)
}

func (h HaNodeHandler) Stats(ctx context.Context, attributes []string) (stat.HaNode, error) {
	return stats[stat.HaNode](ctx, h.client, "", attributes)
}

func (h HaNodeHandler) Unset(ctx context.Context, r config.HaNode) error {
	return OperationNotImplementedError
}

func (h HaNodeHandler) Update(ctx context.Context, r config.HaNode) error {
	return OperationNotImplementedError
}
