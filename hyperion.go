package hyperion

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	m "github.com/denwwer/hyperion-ng/internal/model"
)

// Basic configuration.
const (
	clientName   = "hyperion-ng"
	clientHeader = "X-Client"
	authHeader   = "Authorization"

	attemptDelay = 5 * time.Second
	attemptCount = 5
)

// ClientOption available options.
type ClientOption func(c *Client)

// WithLogger set custom logger.
func WithLogger(l Logger) ClientOption {
	return func(c *Client) {
		c.logger = l
	}
}

// WithHeader set custom headers.
func WithHeader(headers map[string]string) ClientOption {
	return func(c *Client) {
		c.headers = headers
	}
}

// Client instance for Hyperion.
type Client struct {
	cl         *http.Client
	url        string
	verboseLog bool
	logger     Logger
	headers    map[string]string
	token      string
}

// NewClient creates new client.
func NewClient(conf Config, opt ...ClientOption) *Client {
	c := &Client{
		cl: &http.Client{
			Timeout: conf.GetTimeout(),
		},
		url:        getURL(conf),
		verboseLog: conf.VerboseLog,
		token:      conf.Connection.Token,
	}

	// apply options
	for _, o := range opt {
		o(c)
	}

	if c.logger == nil {
		c.logger = &StdLogger{} // default logger
	}

	return c
}

func (c *Client) send(req interface{}, respInfo interface{}) error {
	reqData, err := json.Marshal(req)
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequest(http.MethodPost, c.url, bytes.NewReader(reqData))
	if err != nil {
		return err
	}

	c.setHeaders(httpReq)
	c.logRequest(httpReq)

	var resp *http.Response
	var respErr error

	// process request and retry if failed
	for i := 1; i <= attemptCount; i++ {
		resp, respErr = c.cl.Do(httpReq)

		c.logResponse(resp)

		if respErr == nil {
			break // success
		}

		// retry
		c.logger.Warn(fmt.Sprintf("[WARN] could not connect to Hyperion [%s] (attem %d) becouse of error: %s", c.url, i, respErr))
		time.Sleep(attemptDelay)
	}

	if respErr != nil {
		return respErr // request is failed, return last error
	}

	dec := json.NewDecoder(resp.Body)
	respData := m.Response{}

	if respInfo != nil {
		respData.Info = respInfo
	}

	err = dec.Decode(&respData)
	if err != nil {
		return err
	}

	if !respData.Success {
		if c.token == "" && strings.ToLower(respData.Error) == m.AuthError {
			return errors.New(m.TokenRequire)
		}
		return errors.New(respData.Error) // request processed with error
	}

	return nil
}

func (c Client) setHeaders(req *http.Request) {
	if c.token != "" {
		req.Header.Set(authHeader, "token "+c.token)
	}
	req.Header.Set(clientHeader, clientName)

	for key, val := range c.headers {
		req.Header.Set(key, val)
	}
}

func (c Client) logRequest(req *http.Request) {
	if !c.verboseLog {
		return
	}

	reqLog, err := httputil.DumpRequest(req, true)
	if err != nil {
		c.logger.Warn("log error: " + err.Error())
	}

	c.logger.Info(">>>\n" + string(reqLog) + "\n")
}

func (c Client) logResponse(resp *http.Response) {
	if !c.verboseLog {
		return
	}

	if resp == nil {
		return // empty response
	}

	respLog, err := httputil.DumpResponse(resp, true)
	if err != nil {
		c.logger.Warn("log error: " + err.Error())
		return
	}

	c.logger.Info("<<<\n" + string(respLog))
}

func getURL(conf Config) string {
	if conf.Connection.Type == ConnectHTTP {
		schema := "http"
		if conf.Connection.SSL {
			schema = "https"
		}

		host := conf.Connection.Host
		if conf.Connection.Port > 0 {
			host = fmt.Sprintf("%s:%d", host, conf.Connection.Port)
		}

		return fmt.Sprintf("%s://%s/json-rpc", schema, host)
	}

	return ""
}
