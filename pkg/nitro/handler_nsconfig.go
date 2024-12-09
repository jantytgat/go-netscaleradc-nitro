package nitro

import (
	"context"
	"net/http"

	"github.com/corelayer/go-netscaleradc-nitro/pkg/nitro/resource/config"
)

type NsConfigHandler handler

func (h NsConfigHandler) Clear(ctx context.Context, level string) error {
	req := Request[config.NsConfig]{
		Method: http.MethodPost,
		Action: ActionClear,
		Data: []config.NsConfig{
			{
				Level: level,
			},
		},
	}

	var err error
	if _, err = executeNitroRequest(ctx, h.client, &req); err != nil {
		return err
	}
	return nil
}

func (h NsConfigHandler) Diff(ctx context.Context, config1 string, config2 string, output string, template bool, ignoreDeviceSpecificDifferences bool) (string, error) {
	req := Request[config.NsConfig]{
		Method: http.MethodPost,
		Action: ActionDiff,
		Data: []config.NsConfig{
			{
				Config1:                         config1,
				Config2:                         config2,
				DiffOutputType:                  output,
				Template:                        template,
				IgnoreDeviceSpecificDifferences: ignoreDeviceSpecificDifferences,
			},
		},
	}

	var err error
	var res *Response[config.NsConfig]
	if res, err = executeNitroRequest(ctx, h.client, &req); err != nil {
		return "", err
	}
	return res.Data[0].Response, nil
}

func (h NsConfigHandler) Get(ctx context.Context, attributes []string) (config.NsConfig, error) {
	return getResource[config.NsConfig](ctx, h.client, "", attributes)
}

func (h NsConfigHandler) Save(ctx context.Context) error {
	req := Request[config.NsConfig]{
		Method: http.MethodPost,
		Action: ActionSave,
		Data:   []config.NsConfig{{All: true}},
	}

	var err error
	if _, err = executeNitroRequest(ctx, h.client, &req); err != nil {
		return err
	}
	return nil
}

func (h NsConfigHandler) Unset(ctx context.Context, r config.NsConfig) error {
	return OperationNotImplementedError
}

func (h NsConfigHandler) Update(ctx context.Context, r config.NsConfig) error {
	return OperationNotImplementedError
}
