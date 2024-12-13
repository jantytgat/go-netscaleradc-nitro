package nitro

import (
	"context"

	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/config"
)

type DnsTxtRecHandler handler

func (h DnsTxtRecHandler) Add(ctx context.Context, dnsTxtRecord config.DnsTxtRec) error {
	return addResource[config.DnsTxtRec](ctx, h.client, dnsTxtRecord)
}

func (h DnsTxtRecHandler) Count(ctx context.Context) (float64, error) {
	var err error
	var r config.DnsTxtRec
	if r, err = countResource[config.DnsTxtRec](ctx, h.client); err != nil {
		return 0, err
	}
	return r.Count, nil
}

func (h DnsTxtRecHandler) Delete(ctx context.Context, hostname string, recordId string) error {
	return deleteResource[config.DnsTxtRec](ctx, h.client, hostname, map[string]string{"recordid": recordId})
}

func (h DnsTxtRecHandler) Get(ctx context.Context, hostname string, attributes []string) (config.DnsTxtRec, error) {
	return getResource[config.DnsTxtRec](ctx, h.client, hostname, attributes)
}

func (h DnsTxtRecHandler) List(ctx context.Context, attributes []string, filter map[string]string) ([]config.DnsTxtRec, error) {
	return listResource[config.DnsTxtRec](ctx, h.client, attributes, filter)
}
