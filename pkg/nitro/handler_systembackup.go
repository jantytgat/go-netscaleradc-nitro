package nitro

import (
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/config"
)

type SystemBackupHandler handler

func (h SystemBackupHandler) Add(ctx context.Context, systemBackup config.SystemBackup) error {
	return addResource[config.SystemBackup](ctx, h.client, systemBackup)
}

func (h SystemBackupHandler) Count(ctx context.Context) (float64, error) {
	var err error
	var r config.SystemBackup
	if r, err = countResource[config.SystemBackup](ctx, h.client); err != nil {
		return 0, err
	}
	return r.Count, nil
}

func (h SystemBackupHandler) Create(ctx context.Context, systemBackupName string, level string) error {
	return createResource[config.SystemBackup](ctx, h.client, config.NewSystemBackupCreateRequest(strings.TrimSuffix(systemBackupName, ".tgz"), level))
}

func (h SystemBackupHandler) Delete(ctx context.Context, systemBackupName string) error {
	return deleteResource[config.SystemBackup](ctx, h.client, systemBackupName, nil)
}

func (h SystemBackupHandler) Download(ctx context.Context, systemBackupName string) (*io.Reader, error) {
	var err error
	var systemBackup config.SystemBackup
	if systemBackup, err = h.Get(ctx, systemBackupName, []string{"filename"}); err != nil {
		return nil, err
	}
	return h.client.SystemFile.Download(ctx, systemBackup.Filename, "/var/ns_sys_backup")
}

func (h SystemBackupHandler) Get(ctx context.Context, systemBackupName string, attributes []string) (config.SystemBackup, error) {
	return getResource[config.SystemBackup](ctx, h.client, systemBackupName, attributes)
}

func (h SystemBackupHandler) List(ctx context.Context, attributes []string, filter map[string]string) ([]config.SystemBackup, error) {
	return listResource[config.SystemBackup](ctx, h.client, attributes, filter)
}

func (h SystemBackupHandler) Restore(ctx context.Context, filename string, skipBackup bool) error {
	req := Request[config.SystemBackup]{
		Method: http.MethodPost,
		Action: ActionRestore,
		Data: []config.SystemBackup{{
			Filename:   filename,
			SkipBackup: skipBackup,
		}},
	}
	if _, err := executeNitroRequest(ctx, h.client, &req); err != nil {
		return err
	}
	return nil
}
