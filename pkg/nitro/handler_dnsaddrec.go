package nitro

import (
	"context"

	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/config"
)

type DnsAddRecHandler handler

func (h DnsAddRecHandler) Add(ctx context.Context, dnsARecord config.DnsAddRec) error {
	return addResource[config.DnsAddRec](ctx, h.client, dnsARecord)
}

func (h DnsAddRecHandler) Count(ctx context.Context) (float64, error) {
	var err error
	var r config.DnsAddRec
	if r, err = countResource[config.DnsAddRec](ctx, h.client); err != nil {
		return 0, err
	}
	return r.Count, nil
}

func (h DnsAddRecHandler) Delete(ctx context.Context, hostname string, ipaddress string) error {
	return deleteResource[config.DnsAddRec](ctx, h.client, hostname, map[string]string{"ipaddress": ipaddress})
}

func (h DnsAddRecHandler) Get(ctx context.Context, hostname string, attributes []string) (config.DnsAddRec, error) {
	return getResource[config.DnsAddRec](ctx, h.client, hostname, attributes)
}

func (h DnsAddRecHandler) List(ctx context.Context, attributes []string, filter map[string]string) ([]config.DnsAddRec, error) {
	return listResource[config.DnsAddRec](ctx, h.client, attributes, filter)
}
