package nitro

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/config"
	"github.com/jantytgat/go-netscaleradc-nitro/pkg/nitro/resource/stat"
)

const (
	NSGO_CLIENT_DEFAULT_ACCEPT_HEADER      = "application/json"
	NSGO_CLIENT_DEFAULT_CONTENTTYPE_HEADER = "application/json"
)

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

	client.initialize()

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

type Client struct {
	Name        string
	client      *http.Client
	address     string
	credentials Credentials
	settings    ConnectionSettings
	isLoggedIn  bool

	// Resource Handlers
	common                    handler
	CsVserver                 *CsVserverHandler
	DnsAddRec                 *DnsAddRecHandler
	DnsTxtRec                 *DnsTxtRecHandler
	HaFailOver                *HaFailOverHandler
	HaNode                    *HaNodeHandler
	LbVserver                 *LbVserverHandler
	NsConfig                  *NsConfigHandler
	NsVersion                 *NsVersionHandler
	PolicyStringmap           *PolicyStringmapHandler
	ResponderAction           *ResponderActionHandler
	ResponderPolicy           *ResponderPolicyHandler
	Server                    *ServerHandler
	SslVserver                *SslVserverHandler
	SystemBackup              *SystemBackupHandler
	SystemCmdPolicy           *SystemCmdPolicyHandler
	SystemFile                *SystemFileHandler
	SystemUser                *SystemUserHandler
	SystemUserSystemCmdPolicy *SystemUserSystemCmdPolicyBindingHandler
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

	req, err = createHttpRequest[config.Login](c.BaseUrl(), &nitroReq)
	if err != nil {
		return ClientLoginError.WithMessage(fmt.Sprintf(NSGO_CLIENT_LOGIN_ERROR_MESSAGE + " while creating http request")).WithError(err)
	}

	res, err = c.client.Do(req)
	if err != nil {
		return ClientLoginError.WithMessage(fmt.Sprintf(NSGO_CLIENT_LOGIN_ERROR_MESSAGE + " while executing http request")).WithError(err)
	}

	_, err = deserializeResponse[config.Login](res)
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

	req, err = createHttpRequest[config.Logout](c.BaseUrl(), &nitroReq)
	if err != nil {
		return ClientLogoutError.WithMessage(fmt.Sprintf(NSGO_CLIENT_LOGOUT_ERROR_MESSAGE + " while creating http request")).WithError(err)
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

func (c *Client) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("name", c.Name),
		slog.String("address", c.BaseUrl()),
		slog.Any("credentials", c.credentials))
}

func (c *Client) IsPrimaryNode(ctx context.Context) (bool, error) {
	var err error
	var res stat.HaNode

	if res, err = c.HaNode.Stats(ctx, []string{"hacurmasterstate"}); err != nil {
		return false, err
	}
	return res.CurrentMasterState == "Primary", nil
}

func (c *Client) SaveConfig(ctx context.Context) error {
	return c.NsConfig.Save(ctx)
}

func (c *Client) initialize() {
	c.common = handler{client: c}

	c.CsVserver = (*CsVserverHandler)(&c.common)
	c.DnsAddRec = (*DnsAddRecHandler)(&c.common)
	c.DnsTxtRec = (*DnsTxtRecHandler)(&c.common)
	c.HaFailOver = (*HaFailOverHandler)(&c.common)
	c.HaNode = (*HaNodeHandler)(&c.common)
	c.LbVserver = (*LbVserverHandler)(&c.common)
	c.NsConfig = (*NsConfigHandler)(&c.common)
	c.NsVersion = (*NsVersionHandler)(&c.common)
	c.PolicyStringmap = (*PolicyStringmapHandler)(&c.common)
	c.ResponderAction = (*ResponderActionHandler)(&c.common)
	c.ResponderPolicy = (*ResponderPolicyHandler)(&c.common)
	c.Server = (*ServerHandler)(&c.common)
	c.SslVserver = (*SslVserverHandler)(&c.common)
	c.SystemBackup = (*SystemBackupHandler)(&c.common)
	c.SystemCmdPolicy = (*SystemCmdPolicyHandler)(&c.common)
	c.SystemFile = (*SystemFileHandler)(&c.common)
	c.SystemUser = (*SystemUserHandler)(&c.common)
	c.SystemUserSystemCmdPolicy = (*SystemUserSystemCmdPolicyBindingHandler)(&c.common)
}

func (c *Client) setHeadersOnRequest(resourceName string, r *http.Request) {
	r.Header.Set("Accept", NSGO_CLIENT_DEFAULT_ACCEPT_HEADER)
	r.Header.Set("Content-Type", NSGO_CLIENT_DEFAULT_CONTENTTYPE_HEADER)
	r.Header.Set("User-Agent", c.settings.GetUserAgent())

	if !c.isLoggedIn && resourceName != "login" {
		r.Header.Set("X-NITRO-USER", c.credentials.Username)
		r.Header.Set("X-NITRO-PASS", c.credentials.Password)
	}
}
