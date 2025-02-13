package nitro

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/config"
)

type SystemFileHandler handler

func (h SystemFileHandler) Add(ctx context.Context, systemFile config.SystemFile) error {
	return addResource[config.SystemFile](ctx, h.client, systemFile)
}

func (h SystemFileHandler) Count(ctx context.Context, location string) (float64, error) {
	var err error

	req := Request[config.SystemFile]{
		Method:    http.MethodGet,
		Arguments: map[string]string{"filelocation": location},
	}

	var res *Response[config.SystemFile]
	if res, err = executeNitroRequest[config.SystemFile](ctx, h.client, &req); err != nil {
		return 0, err
	}

	return float64(len(res.Data)), nil
}

func (h SystemFileHandler) Delete(ctx context.Context, filename string, location string) error {
	return deleteResource[config.SystemFile](ctx, h.client, filename, map[string]string{"filelocation": location})
}

func (h SystemFileHandler) Download(ctx context.Context, filename string, location string) (*io.Reader, error) {
	var err error
	var file config.SystemFile

	if file, err = h.Get(ctx, filename, location, []string{"filecontent"}); err != nil {
		return nil, err
	}

	output := base64.NewDecoder(base64.StdEncoding, strings.NewReader(file.Content))
	return &output, nil
}

func (h SystemFileHandler) Get(ctx context.Context, filename string, location string, attributes []string) (config.SystemFile, error) {
	var err error

	req := Request[config.SystemFile]{
		Method: http.MethodGet,
		Arguments: map[string]string{
			"filename":     filename,
			"filelocation": location},
		Attributes: attributes,
	}

	var res *Response[config.SystemFile]
	if res, err = executeNitroRequest[config.SystemFile](ctx, h.client, &req); err != nil {
		return config.SystemFile{}, err
	}

	if len(res.Data) >= 1 {
		return res.Data[0], nil
	}
	return config.SystemFile{}, fmt.Errorf("invalid amount of files in data")
}

func (h SystemFileHandler) List(ctx context.Context, location string, attributes []string, filter map[string]string) ([]config.SystemFile, error) {
	var err error

	req := Request[config.SystemFile]{
		Method:     http.MethodGet,
		Arguments:  map[string]string{"filelocation": location},
		Attributes: attributes,
		Filter:     filter,
	}

	var res *Response[config.SystemFile]
	if res, err = executeNitroRequest[config.SystemFile](ctx, h.client, &req); err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (h SystemFileHandler) Unset(ctx context.Context, r config.SystemFile) error {
	return OperationNotImplementedError
}

func (h SystemFileHandler) Update(ctx context.Context, r config.SystemFile) error {
	return OperationNotImplementedError
}

func (h SystemFileHandler) Upload(ctx context.Context, filename string, location string, content []byte) error {
	return h.Add(ctx, config.NewSystemFileAddRequest(filename, location, content))
}
