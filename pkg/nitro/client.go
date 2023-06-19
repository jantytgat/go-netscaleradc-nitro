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
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"reflect"

	"github.com/corelayer/netscaleradc-nitro-go/pkg/nitro/resource/config"
	"github.com/corelayer/netscaleradc-nitro-go/pkg/nitro/resource/stat"
)

type Client struct {
	client      *http.Client
	Name        string
	address     string
	credentials Credentials
	settings    Settings
	isLoggedIn  bool
}

func (c *Client) BaseUrl() string {
	return c.settings.GetUrlScheme() + c.address
}

func (c *Client) IsLoggedIn() bool {
	return c.isLoggedIn
}

// Login TODO Nextfactorauthentication
func (c *Client) Login() error {
	var (
		err error
		req *http.Request
		res *http.Response
	)

	nitroReq := Request[config.Login]{
		Method: http.MethodPost,
		Data: []config.Login{{
			Username: c.credentials.Username,
			Password: c.credentials.Password,
		}},
	}

	req, err = CreateHttpRequest[config.Login](c, &nitroReq)
	if err != nil {
		return err
	}

	res, err = c.client.Do(req)
	// TODO CHECK RESPONSE
	if err != nil {
		return err
	}

	_, err = DeserializeResponse[config.Login](res)
	if err != nil {
		return fmt.Errorf("could not log in: %w", err)
	}

	if c.client.Jar != nil {
		c.client.Jar.SetCookies(req.URL, res.Cookies())
	} else {
		var jar *cookiejar.Jar
		jar, err = cookiejar.New(nil)
		if err != nil {
			return fmt.Errorf("failed creating cookie jar: %w", err)
		}
		jar.SetCookies(req.URL, res.Cookies())
		c.client.Jar = jar

	}

	c.isLoggedIn = true
	return err

}

func (c *Client) Logout() error {
	var err error
	var req *http.Request
	// var res *http.Response

	nitroReq := Request[config.Logout]{
		Method: http.MethodPost,
		Data:   []config.Logout{{}},
	}

	req, err = CreateHttpRequest[config.Logout](c, &nitroReq)
	if err != nil {
		return fmt.Errorf("failed creating logout request: %w", err)
	}
	_, err = c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed executing logout request: %w", err)
	}

	// Reset http Client CookieJar
	c.client.Jar = nil
	c.isLoggedIn = false
	return nil

}

func (c *Client) IsPrimaryNode() (bool, error) {
	var err error

	nitroReq := Request[stat.HaNode]{
		Method:     http.MethodGet,
		Attributes: []string{"hacurmasterstate"},
	}

	var res *Response[stat.HaNode]
	res, err = ExecuteNitroRequest(c, &nitroReq)
	if err != nil {
		return false, err
	}

	// Check if Status is Primary
	d := res.Data[0]
	if d.CurrentMasterState == "Primary" {
		return true, nil
	}
	return false, nil
}

func (c *Client) addNitroRequestHeaders(resourceName string, r *http.Request) {
	r.Header.Set("Accept", "application/json")
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("User-Agent", c.getUserAgent())

	if !c.isLoggedIn && resourceName != "login" {
		r.Header.Set("X-NITRO-USER", c.credentials.Username)
		r.Header.Set("X-NITRO-PASS", c.credentials.Password)
	}
}

func (c *Client) getUserAgent() string {
	if c.settings.UserAgent != "" {
		return c.settings.UserAgent
	}
	return "go-netscaler-nitro"
}

func NewClient(name string, address string, credentials Credentials, settings Settings) (*Client, error) {
	tlsLog, err := settings.GetTlsSecretLogWriter()
	if err != nil {
		return nil, err
	}

	// if tlsLog != nil {
	//	log.Println("Exporting TLS Secrets to" + settings.LogTlsSecretsDestination)
	// }

	timeout, err := settings.GetTimeoutDuration()
	c := &Client{
		Name:        name,
		address:     address,
		credentials: credentials,
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					KeyLogWriter:       tlsLog,
					InsecureSkipVerify: !settings.ValidateServerCertificate,
				},
				Proxy: http.ProxyFromEnvironment,
			},
			Timeout: timeout,
		},
		settings:   settings,
		isLoggedIn: false,
	}

	if !settings.AutoLogin {
		return c, nil
	}

	err = c.Login()
	if err != nil {
		return c, err
	}
	return c, nil
}

func CreateHttpRequest[T ResourceReader](c *Client, req *Request[T]) (*http.Request, error) {
	var (
		err  error
		body io.Reader
	)
	buf := &bytes.Buffer{}

	// Do not serialize body when there are no items to serialize
	if len(req.Data) > 0 {
		body, err = req.serializeBody()
		if err != nil {
			return nil, FormatCreateHttpRequestError(req.GetResourceTypeName(), err)
		}

		tee := io.TeeReader(body, buf)

		err = req.ValidateData(tee)
		if err != nil {
			return nil, FormatCreateHttpRequestError(req.GetResourceTypeName(), err)
		}
	}

	// Build URL Path and Query
	var query string
	query, err = req.GetUrlPathAndQuery()
	if err != nil {
		return nil, FormatCreateHttpRequestError(req.GetResourceTypeName(), err)
	}
	url := c.BaseUrl() + query

	var r *http.Request
	r, err = http.NewRequest(
		req.Method,
		url,
		bytes.NewReader(buf.Bytes()))
	if err != nil {
		return nil, FormatCreateHttpRequestError(req.GetResourceTypeName(), err)
	}
	c.addNitroRequestHeaders(req.GetResourceTypeName(), r)

	return r, nil
}

func ExecuteNitroRequest[T ResourceReader](c *Client, r *Request[T]) (*Response[T], error) {
	var req *http.Request
	var res *http.Response
	var err error
	req, err = CreateHttpRequest[T](c, r)
	if err != nil {
		return nil, FormatExecuteNitroRequestError(err)
	}

	res, err = c.client.Do(req)
	if err != nil {
		return nil, FormatExecuteNitroRequestError(err)
	}

	var nitroRes *Response[T]
	if res.ContentLength == 0 {
		nitroRes = &Response[T]{
			ErrorCode: 0,
			Message:   "Done",
			Severity:  "NONE",
		}
		return nitroRes, nil
	}

	nitroRes, err = DeserializeResponse[T](res)
	if err != nil {
		return nitroRes, FormatExecuteNitroRequestError(err)
	}

	return nitroRes, nil
}

func DeserializeResponse[T ResourceReader](res *http.Response) (*Response[T], error) {
	var r Response[T]
	var err error

	defer func() {
		e := res.Body.Close()
		if e != nil {
			err = e
		}
	}()

	var rawBody []byte
	rawBody, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, FormatDeserializeResponseError(err)
	}

	// Check if any data is returned in the body
	// if len(rawBody) == 0 {
	//	r.ErrorCode = 0
	//	r.Message = "Done"
	//	r.Severity = "NONE"
	//
	//	return &r, nil
	// }

	// Unmarshal JSON into map
	var bodyMap map[string]interface{}
	err = json.Unmarshal(rawBody, &bodyMap)
	if err != nil {
		return &r, FormatDeserializeResponseError(err)
	}

	// TODO --> Validate keys??
	r.ErrorCode = bodyMap["errorcode"].(float64)
	r.Message = bodyMap["message"].(string)
	r.Severity = bodyMap["severity"].(string)

	t := *new(T)
	// Return early if there is no data to be serialized
	// if res.StatusCode == 200 || res.StatusCode == 201 {
	//	return &r, nil
	// }

	if _, ok := bodyMap[t.GetTypeName()]; !ok {
		switch res.StatusCode {
		case 200:
			return &r, nil
		case 201:
			return &r, nil
		default:
			return &r, FormatSpecificNitroError(r.Message)
		}
	}

	err = r.ExtractData(bodyMap[t.GetTypeName()])
	if err != nil {
		return &r, FormatDeserializeResponseError(err)
	}

	return &r, nil
}

func mapToStruct[T any](t *T, v map[string]interface{}) error {
	var err error
	var jsonData []byte

	rt := reflect.TypeOf(t)
	// Convert map to json
	jsonData, err = json.Marshal(v)
	if err != nil {
		return fmt.Errorf("failed marshalling map of type %s to json: %w", rt.Name(), err)
	}

	// Convert json to struct of type T
	err = json.Unmarshal(jsonData, t)
	if err != nil {
		return fmt.Errorf("failed unmarshalling json to struct of type %s: %w", rt.Name(), err)
	}

	return nil
}
