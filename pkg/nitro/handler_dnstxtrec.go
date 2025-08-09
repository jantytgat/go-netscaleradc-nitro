package nitro

import (
	"context"

	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/config"
)

type DnsTxtRecordHandler handler

func (h DnsTxtRecordHandler) Add(ctx context.Context, dnsTxtRecord config.DnsTxtRecord) error {
	return addResource[config.DnsTxtRecord](ctx, h.client, dnsTxtRecord)
}

func (h DnsTxtRecordHandler) Count(ctx context.Context) (float64, error) {
	var err error
	var r config.DnsTxtRecord
	if r, err = countResource[config.DnsTxtRecord](ctx, h.client); err != nil {
		return 0, err
	}
	return r.Count, nil
}

func (h DnsTxtRecordHandler) Delete(ctx context.Context, hostname string, recordId string) error {
	return deleteResource[config.DnsTxtRecord](ctx, h.client, hostname, map[string]string{"recordid": recordId})
}

func (h DnsTxtRecordHandler) Get(ctx context.Context, hostname string, attributes []string) (config.DnsTxtRecord, error) {
	return getResourceWithName[config.DnsTxtRecord](ctx, h.client, hostname, attributes)
}

func (h DnsTxtRecordHandler) List(ctx context.Context, attributes []string, filter map[string]string) ([]config.DnsTxtRecord, error) {
	return listResource[config.DnsTxtRecord](ctx, h.client, attributes, filter)
}
