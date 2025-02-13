package nitro

import (
	"io"
	"time"
)

type SettingsReader interface {
	GetTlsSecretLogWriter() (io.Writer, error)
	GetTimeoutDuration() (time.Duration, error)
}
