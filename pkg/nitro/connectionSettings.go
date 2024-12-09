package nitro

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

const (
	NSGO_CONNECTIONSETTINGS_DEFAULT_TIMEOUT   = 5
	NSGO_CONNECTIONSETTINGS_DEFAULT_USERAGENT = "netscaleradc-nitro-go"
)

type ConnectionSettings struct {
	UseSsl                    bool   `json:"useSsl" yaml:"useSsl" mapstructure:"useSsl"`
	Timeout                   int    `json:"timeout" yaml:"timeout" mapstructure:"timeout"`
	UserAgent                 string `json:"userAgent" yaml:"userAgent" mapstructure:"userAgent"`
	ValidateServerCertificate bool   `json:"validateServerCertificate" yaml:"validateServerCertificate" mapstructure:"validateServerCertificate"`
	LogTlsSecrets             bool   `json:"logTlsSecrets" yaml:"logTlsSecrets" mapstructure:"logTlsSecrets"`
	LogTlsSecretsDestination  string `json:"logTlsSecretsDestination" yaml:"logTlsSecretsDestination" mapstructure:"logTlsSecretsDestination"`
	AutoLogin                 bool   `json:"autoLogin" yaml:"autoLogin" mapstructure:"autoLogin"`
}

func (n *ConnectionSettings) GetUrlScheme() string {
	switch n.UseSsl {
	case false:
		return "http://"
	default:
		return "https://"
	}
}

func (n *ConnectionSettings) GetTlsSecretLogWriter() (io.Writer, error) {
	if !n.LogTlsSecrets {
		return nil, nil
	}

	tlsLog, err := os.OpenFile(n.LogTlsSecretsDestination, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return nil, ClientConnectionSettingsError.WithMessage(fmt.Sprintf(NSGO_CLIENT_CONNECTIONSETTINGS_ERROR_MESSAGE + " tls secrets log destination")).WithError(err)
	}
	return tlsLog, nil

}

// GetTimeoutDuration Returns a time.Duration based on the set timout in milliseconds
func (n *ConnectionSettings) GetTimeoutDuration() (time.Duration, error) {
	var (
		err     error
		timeout time.Duration
	)

	timeout, err = time.ParseDuration(strconv.Itoa(n.Timeout) + "s")
	if err != nil {
		return NSGO_CONNECTIONSETTINGS_DEFAULT_TIMEOUT * time.Second, ClientConnectionSettingsError.WithMessage(fmt.Sprintf(NSGO_CLIENT_CONNECTIONSETTINGS_ERROR_MESSAGE + " timeout duration parsing")).WithError(err)
	}
	return timeout, nil
}

func (n *ConnectionSettings) GetUserAgent() string {
	if n.UserAgent != "" {
		return n.UserAgent
	}
	return NSGO_CONNECTIONSETTINGS_DEFAULT_USERAGENT
}
