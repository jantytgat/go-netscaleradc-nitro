package nitro

import (
	"context"

	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/config"
)

type DnsAddressRecordHandler handler

func (h DnsAddressRecordHandler) Add(ctx context.Context, dnsARecord config.DnsAddressRecord) error {
	return addResource[config.DnsAddressRecord](ctx, h.client, dnsARecord)
}

func (h DnsAddressRecordHandler) Count(ctx context.Context) (float64, error) {
	var err error
	var r config.DnsAddressRecord
	if r, err = countResource[config.DnsAddressRecord](ctx, h.client); err != nil {
		return 0, err
	}
	return r.Count, nil
}

func (h DnsAddressRecordHandler) Delete(ctx context.Context, hostname string, ipaddress string) error {
	return deleteResource[config.DnsAddressRecord](ctx, h.client, hostname, map[string]string{"ipaddress": ipaddress})
}

func (h DnsAddressRecordHandler) Get(ctx context.Context, hostname string, attributes []string) (config.DnsAddressRecord, error) {
	return getResource[config.DnsAddressRecord](ctx, h.client, hostname, attributes)
}

func (h DnsAddressRecordHandler) List(ctx context.Context, attributes []string, filter map[string]string) ([]config.DnsAddressRecord, error) {
	return listResource[config.DnsAddressRecord](ctx, h.client, attributes, filter)
}
