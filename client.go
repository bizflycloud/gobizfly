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
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	version       = "0.0.1"
	ua            = "bizfly-client-go/" + version
	defaultAPIURL = "https://manage.bizflycloud.vn"
	mediaType     = "application/json"
)

// Client represents BizFly API client.
type Client struct {
	Token        TokenService
	LoadBalancer LoadBalancerService
	Listener     ListenerService
	Pool         PoolService
	Member       MemberService

	httpClient    *http.Client
	apiURL        *url.URL
	userAgent     string
	keystoneToken string
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

	c.Token = &token{client: c}
	c.LoadBalancer = &loadbalancer{client: c}
	c.Listener = &listener{client: c}
	c.Pool = &pool{client: c}
	c.Member = &member{client: c}

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

// Do sends API request.
func (c *Client) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= http.StatusBadRequest {
		defer resp.Body.Close()
		buf, _ := ioutil.ReadAll(resp.Body)
		return nil, errors.New(string(buf))
	}
	return resp, nil
}

// SetKeystoneToken sets keystone token value, which will be used for authentication.
func (c *Client) SetKeystoneToken(s string) {
	c.keystoneToken = s
}

// ListOptions specifies the optional parameters for List method.
type ListOptions struct{}
