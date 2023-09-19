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
	"time"

	"github.com/corelayer/netscaleradc-nitro-go/pkg/nitro/resource/config"
	"github.com/corelayer/netscaleradc-nitro-go/pkg/nitro/resource/stat"
)

const (
	NSGO_CLIENT_DEFAULT_ACCEPT_HEADER      = "application/json"
	NSGO_CLIENT_DEFAULT_CONTENTTYPE_HEADER = "application/json"
)

type Client struct {
	Name        string
	client      *http.Client
	address     string
	credentials Credentials
	settings    ConnectionSettings
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
		return ClientLoginError.WithMessage(fmt.Sprintf(NSGO_CLIENT_LOGIN_ERROR_MESSAGE + " while creating http request")).WithError(err)
	}

	res, err = c.client.Do(req)
	// TODO CHECK RESPONSE
	if err != nil {
		return ClientLoginError.WithMessage(fmt.Sprintf(NSGO_CLIENT_LOGIN_ERROR_MESSAGE + " while executing http request")).WithError(err)
	}

	_, err = DeserializeResponse[config.Login](res)
	if err != nil {
		return ClientLoginError.WithMessage(fmt.Sprintf(NSGO_CLIENT_LOGIN_ERROR_MESSAGE + " while deserializing response")).WithError(err)
	}

	if c.client.Jar != nil {
		c.client.Jar.SetCookies(req.URL, res.Cookies())
	} else {
		var jar *cookiejar.Jar
		jar, err = cookiejar.New(nil)
		if err != nil {
			return ClientLoginError.WithMessage(fmt.Sprintf(NSGO_CLIENT_LOGIN_ERROR_MESSAGE + " while creating cookie jar")).WithError(err)
		}
		jar.SetCookies(req.URL, res.Cookies())
		c.client.Jar = jar

	}

	c.isLoggedIn = true
	return nil

}

func (c *Client) Logout() error {
	var err error
	var req *http.Request

	nitroReq := Request[config.Logout]{
		Method: http.MethodPost,
		Data:   []config.Logout{{}},
	}

	req, err = CreateHttpRequest[config.Logout](c, &nitroReq)
	if err != nil {
		return ClientLogoutError.WithMessage(fmt.Sprintf(NSGO_CLIENT_LOGOUT_ERROR_MESSAGE + " while executing http request")).WithError(err)
	}
	_, err = c.client.Do(req)
	if err != nil {
		return ClientLogoutError.WithMessage(fmt.Sprintf(NSGO_CLIENT_LOGOUT_ERROR_MESSAGE + " while executing http request")).WithError(err)
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
	// TODO ERROR HANDLING
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

func (c *Client) SaveConfig() error {
	var err error
	var req *http.Request

	nitroReq := Request[config.NsConfig]{
		Method: http.MethodPost,
		Data:   []config.NsConfig{{}},
	}

	req, err = CreateHttpRequest[config.NsConfig](c, &nitroReq)
	if err != nil {
		return ClientLogoutError.WithMessage(fmt.Sprintf(NSGO_CLIENT_SAVECONFIG_ERROR_MESSAGE + " while executing http request")).WithError(err)
	}
	_, err = c.client.Do(req)
	if err != nil {
		return ClientLogoutError.WithMessage(fmt.Sprintf(NSGO_CLIENT_SAVECONFIG_ERROR_MESSAGE + " while executing http request")).WithError(err)
	}

	// Reset http Client CookieJar
	c.client.Jar = nil
	c.isLoggedIn = false
	return nil

}

func (c *Client) addNitroRequestHeaders(resourceName string, r *http.Request) {
	r.Header.Set("Accept", NSGO_CLIENT_DEFAULT_ACCEPT_HEADER)
	r.Header.Set("Content-Type", NSGO_CLIENT_DEFAULT_CONTENTTYPE_HEADER)
	r.Header.Set("User-Agent", c.settings.GetUserAgent())

	if !c.isLoggedIn && resourceName != "login" {
		r.Header.Set("X-NITRO-USER", c.credentials.Username)
		r.Header.Set("X-NITRO-PASS", c.credentials.Password)
	}
}

func NewClient(name string, address string, credentials Credentials, settings ConnectionSettings) (*Client, error) {
	var (
		err     error
		client  *Client
		tlsLog  io.Writer
		timeout time.Duration
	)

	tlsLog, err = settings.GetTlsSecretLogWriter()
	if err != nil {
		return nil, ClientCreateError.WithError(err)
	}

	timeout, err = settings.GetTimeoutDuration()
	client = &Client{
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
		return client, nil
	}

	// Perform auto login
	err = client.Login()
	if err != nil {
		return client, ClientCreateError.WithMessage(fmt.Sprintf(NSGO_CLIENT_CREATE_ERROR_MESSAGE + " using auto login")).WithError(err)
	}
	return client, nil
}

func CreateHttpRequest[T ResourceReader](client *Client, req *Request[T]) (*http.Request, error) {
	var (
		err  error
		body io.Reader
		buf  = &bytes.Buffer{} // Used to validate data before creating the request
	)

	// Do not serialize body when there are no items to serialize
	if len(req.Data) > 0 {
		body, err = req.serializeBody()
		if err != nil {
			return nil, ClientCreateHttpRequestError.WithMessage(fmt.Sprintf(NSGO_CLIENT_CREATEHTTPREQUEST_ERROR_MESSAGE+" for %s", req.GetResourceTypeName())).WithError(err)
		}

		tee := io.TeeReader(body, buf)

		err = req.ValidateData(tee)
		if err != nil {
			return nil, ClientCreateHttpRequestError.WithMessage(fmt.Sprintf(NSGO_CLIENT_CREATEHTTPREQUEST_ERROR_MESSAGE+" for %s", req.GetResourceTypeName())).WithError(err)
		}
	}

	// Build URL Path and Query
	var query string
	query, err = req.GetUrlPathAndQuery()
	if err != nil {
		return nil, ClientCreateHttpRequestError.WithMessage(fmt.Sprintf(NSGO_CLIENT_CREATEHTTPREQUEST_ERROR_MESSAGE+" for %s", req.GetResourceTypeName())).WithError(err)
	}
	url := client.BaseUrl() + query

	var r *http.Request
	r, err = http.NewRequest(
		req.Method,
		url,
		bytes.NewReader(buf.Bytes()))
	if err != nil {
		return nil, ClientCreateHttpRequestError.WithMessage(fmt.Sprintf(NSGO_CLIENT_CREATEHTTPREQUEST_ERROR_MESSAGE+" for %s", req.GetResourceTypeName())).WithError(err)
	}
	client.addNitroRequestHeaders(req.GetResourceTypeName(), r)

	return r, nil
}

func ExecuteNitroRequest[T ResourceReader](c *Client, r *Request[T]) (*Response[T], error) {
	var (
		err error
		req *http.Request
		res *http.Response
	)

	req, err = CreateHttpRequest[T](c, r)
	if err != nil {
		return nil, ClientExecuteRequestError.WithMessage(fmt.Sprintf(NSGO_CLIENT_EXECUTEREQUEST_ERROR_MESSAGE + " while creating http request")).WithError(err)
	}

	res, err = c.client.Do(req)
	if err != nil {
		return nil, ClientExecuteRequestError.WithMessage(fmt.Sprintf(NSGO_CLIENT_EXECUTEREQUEST_ERROR_MESSAGE + " using configured http client")).WithError(err)
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
		return nil, ClientExecuteRequestError.WithMessage(fmt.Sprintf(NSGO_CLIENT_EXECUTEREQUEST_ERROR_MESSAGE + " while deserializing response")).WithError(err)
	}

	return nitroRes, nil
}

func DeserializeResponse[T ResourceReader](res *http.Response) (*Response[T], error) {
	var (
		err error
		r   Response[T]
	)

	defer func() {
		e := res.Body.Close()
		if e != nil {
			err = e
		}
	}()

	var rawBody []byte
	rawBody, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, ResourceDeserializationError.WithMessage(fmt.Sprintf(NSGO_RESOURCE_DESERIALIZATION_ERROR_MESSAGE + " while reading response body from http response")).WithError(err)
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
		return &r, ResourceDeserializationError.WithMessage(fmt.Sprintf(NSGO_RESOURCE_DESERIALIZATION_ERROR_MESSAGE + " while unmarshalling data from response body")).WithError(err)
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
			return &r, ApiError.WithMessage(r.Message).WithCode(r.ErrorCode)
		}
	}

	err = r.ExtractData(bodyMap[t.GetTypeName()])
	if err != nil {
		return &r, ResourceDeserializationError.WithMessage(fmt.Sprintf(NSGO_RESOURCE_DESERIALIZATION_ERROR_MESSAGE + " while extracting data for resource")).WithError(err)
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
		return ResourceSerializationError.WithMessage(fmt.Sprintf("Failed to marshal map of type %s to json", rt.Name())).WithError(err)
	}

	// Convert json to struct of type T
	err = json.Unmarshal(jsonData, t)
	if err != nil {
		return ResourceDeserializationError.WithMessage(fmt.Sprintf("Failed to unmarshal json to struct of type %s", rt.Name())).WithError(err)
	}

	return nil
}
