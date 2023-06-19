/*
 * Copyright 2023 CoreLayer BV
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package nitro

import (
	"io"
	"os"
	"strconv"
	"time"
)

type Settings struct {
	UseSsl                    bool
	Timeout                   int
	UserAgent                 string
	ValidateServerCertificate bool
	LogTlsSecrets             bool
	LogTlsSecretsDestination  string
	AutoLogin                 bool
}

func (n *Settings) GetUrlScheme() string {
	switch n.UseSsl {
	case false:
		return "http://"
	default:
		return "https://"
	}
}

func (n *Settings) GetTlsSecretLogWriter() (io.Writer, error) {
	if !n.LogTlsSecrets {
		return nil, nil
	}

	tlsLog, err := os.OpenFile(n.LogTlsSecretsDestination, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return nil, err
	}
	return tlsLog, nil

}

// GetTimeoutDuration Returns a time.Duration based on the set timout in milliseconds
func (n *Settings) GetTimeoutDuration() (time.Duration, error) {
	// TODO add default timeout
	return time.ParseDuration(strconv.Itoa(n.Timeout) + "ms")
}
