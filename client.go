// Copyright 2019 The BizFly Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gobizfly is the BizFly API client for Go.
package gobizfly

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	version       = "0.0.1"
	ua            = "bizfly-client-go/" + version
	defaultAPIURL = "https://manage.bizflycloud.vn"
	mediaType     = "application/json; charset=utf-8"
)

var (
	// ErrNotFound for resource not found status
	ErrNotFound = errors.New("Resource not found")
	// ErrPermissionDenied for permission denied
	ErrPermissionDenied = errors.New("You are not allowed to do this action")
	// ErrCommon for common error
	ErrCommon = errors.New("Error")
)

// Client represents BizFly API client.
type Client struct {
	Token        TokenService
	LoadBalancer LoadBalancerService
	Listener     ListenerService
	Pool         PoolService
	Member       MemberService

	Snapshot SnapshotService

	Volume VolumeService
	Server ServerService

	httpClient    *http.Client
	apiURL        *url.URL
	userAgent     string
	keystoneToken string
	authMethod    string
	username      string
	password      string
	appCredID     string
	appCredSecret string
	// TODO: this will be removed in near future
	tenantName string
}

// Option set Client specific attributes
type Option func(c *Client) error

// WithAPIUrl sets the API url option for BizFly client.
func WithAPIUrl(u string) Option {
	return func(c *Client) error {
		var err error
		c.apiURL, err = url.Parse(u)
		return err
	}
}

// WithHTTPClient sets the client option for BizFly client.
func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) error {
		if client == nil {
			return errors.New("client is nil")
		}

		c.httpClient = client

		return nil
	}
}

// WithTenantName sets the tenant name option for BizFly client.
//
// Deprecated: X-Tenant-Name header required will be removed in API server.
func WithTenantName(tenant string) Option {
	return func(c *Client) error {
		c.tenantName = tenant
		return nil
	}
}

// NewClient creates new BizFly client.
func NewClient(options ...Option) (*Client, error) {
	c := &Client{
		httpClient: http.DefaultClient,
		userAgent:  ua,
	}

	err := WithAPIUrl(defaultAPIURL)(c)
	if err != nil {
		return nil, err
	}

	for _, option := range options {
		if err := option(c); err != nil {
			return nil, err
		}
	}

	c.Snapshot = &snapshot{client: c}
	c.Token = &token{client: c}
	c.LoadBalancer = &loadbalancer{client: c}
	c.Listener = &listener{client: c}
	c.Pool = &pool{client: c}
	c.Member = &member{client: c}
	c.Volume = &volume{client: c}
	c.Server = &server{client: c}

	return c, nil
}

// NewRequest creates an API request.
func (c *Client) NewRequest(ctx context.Context, method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.apiURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if body != nil {
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("Accept", mediaType)
	req.Header.Add("User-Agent", c.userAgent)
	req.Header.Add("X-Tenant-Name", c.tenantName)
	if c.keystoneToken != "" {
		req.Header.Add("X-Auth-Token", c.keystoneToken)
	}
	return req, nil
}

func (c *Client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	return c.httpClient.Do(req)
}

// Do sends API request.
func (c *Client) Do(ctx context.Context, req *http.Request) (resp *http.Response, err error) {

	resp, err = c.do(ctx, req)
	if err != nil {
		return
	}

	// If 401, get new token and retry one time.
	if resp.StatusCode == http.StatusUnauthorized {
		tcr := &TokenCreateRequest{AuthMethod: c.authMethod, Username: c.username, Password: c.password, AppCredID: c.appCredID, AppCredSecret: c.appCredSecret}
		tok, tokErr := c.Token.Create(ctx, tcr)
		if tokErr != nil {
			buf, _ := ioutil.ReadAll(resp.Body)
			err = fmt.Errorf("%s : %w", string(buf), tokErr)
			return
		}
		c.SetKeystoneToken(tok.KeystoneToken)
		req.Header.Set("X-Auth-Token", c.keystoneToken)
		resp, err = c.do(ctx, req)
	}
	if resp.StatusCode >= http.StatusBadRequest {
		defer resp.Body.Close()
		buf, _ := ioutil.ReadAll(resp.Body)
		err = errorFromStatus(resp.StatusCode, string(buf))

	}
	return
}

// SetKeystoneToken sets keystone token value, which will be used for authentication.
func (c *Client) SetKeystoneToken(s string) {
	c.keystoneToken = s
}

// ListOptions specifies the optional parameters for List method.
type ListOptions struct{}

func errorFromStatus(code int, msg string) error {
	switch code {
	case http.StatusNotFound:
		return fmt.Errorf("%s: %w", msg, ErrNotFound)
	case http.StatusForbidden:
		return fmt.Errorf("%s: %w", msg, ErrPermissionDenied)
	default:
		return fmt.Errorf("%s: %w", msg, ErrCommon)
	}
}
