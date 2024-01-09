// This file is part of gobizfly

package gobizfly

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	domainPath = "/domains"
	usersPath  = "/users"
)

type cdnService struct {
	client *Client
}

var _ CDNService = (*cdnService)(nil)

type CDNService interface {
	List(ctx context.Context, opts *ListOptions) (*DomainsResp, error)
	Get(ctx context.Context, domainID string) (*ExtendedDomain, error)
	Create(ctx context.Context, cdrq *CreateDomainPayload) (*CreateDomainResp, error)
	Update(ctx context.Context, domainID string, udrq *UpdateDomainReq) (*UpdateDomainResp, error)
	Delete(ctx context.Context, domainID string) error
	DeleteCache(ctx context.Context, domainID string, files *Files) error
}

// Origin represents a CDN origin information
type Origin struct {
	Name          string `json:"name"`
	UpstreamHost  string `json:"upstream_host"`
	UpstreamAddrs string `json:"upstream_addrs"`
	UpstreamProto string `json:"upstream_proto"`
	OriginType    string `json:"origin_type"`
}

// Domain represents CDN domain information
type Domain struct {
	ID             int      `json:"id"`
	User           int      `json:"user"`
	Certificate    int      `json:"certificate"`
	CName          string   `json:"cname"`
	Origins        []Origin `json:"origins_detail"`
	Slug           string   `json:"slug"`
	PageSpeed      int      `json:"pagespeed"`
	UpstreamProto  string   `json:"upstream_proto"`
	DDOSProtection int      `json:"ddos_protection"`
	Status         string   `json:"status"`
	CreatedAt      string   `json:"created_at"`
	DomainID       string   `json:"domain_id"`
	Domain         string   `json:"domain"`
	Type           string   `json:"type"`
}

// DomainsResp represents a list of CDN domains and pagination information
type DomainsResp struct {
	Domains []Domain `json:"results"`
	Next    string   `json:"next"`
	Prev    string   `json:"prev"`
	Pages   int      `json:"pages"`
	Total   int      `json:"total"`
}

// OriginAddr represents a CDN origin address information
type OriginAddr struct {
	Type string `json:"type"`
	Host string `json:"host"`
}

// ExtendedDomain represents a CDN domain information with additional information
type ExtendedDomain struct {
	Domain
	Last24hUsage int          `json:"last_24h_usage"`
	UpstreamHost string       `json:"upstream_host"`
	Slug         string       `json:"slug"`
	SecretKey    string       `json:"secret_key"`
	DomainCDN    string       `json:"domain_cdn"`
	OriginAddrs  []OriginAddr `json:"origin_addrs"`
	HostID       string       `json:"host_id"`
}

// CreateDomainReq represents a request body for creating CDN domain
type CreateDomainPayload struct {
	Domain string  `json:"domain"`
	Type   int     `json:"domain_type"`
	Origin *Origin `json:"origin"`
}

// CreateDomainResp represents a response body when creating CDN domain
type CreateDomainResp struct {
	Message string `json:"message"`
	Domain  Domain `json:"domain"`
}

// UpdateDomainReq represents a request body for updating CDN domain
type UpdateDomainReq struct {
	UpstreamAddrs string `json:"upstream_addrs"`
	UpstreamProto string `json:"upstream_proto"`
	PageSpeed     int    `json:"pagespeed"`
	SecureLink    int    `json:"secure_link"`
}

// UpdateDomainResp represents a response body when updating CDN domain
type UpdateDomainResp struct {
	Message string         `json:"message"`
	Domain  ExtendedDomain `json:"domain"`
}

// CheckConnResp represents a response body when checking CDN connection
type CheckConnResp struct {
	Status string `json:"status"`
}

type Files struct {
	Files []string `json:"files"`
}

func (c *cdnService) resourcePath() string {
	return strings.Join([]string{usersPath, domainPath}, "")
}

func (c *cdnService) itemPath(id string) string {
	return strings.Join([]string{domainPath, id}, "/")
}

func (c *cdnService) List(ctx context.Context, opts *ListOptions) (*DomainsResp, error) {
	u, _ := url.Parse(strings.Join([]string{usersPath, domainPath}, ""))
	query := url.Values{}
	if opts.Page != 0 {
		query.Add("page", strconv.Itoa(opts.Page))
	}
	if opts.Page != 0 {
		query.Add("limit", strconv.Itoa(opts.Limit))
	}
	u.RawQuery = query.Encode()
	req, err := c.client.NewRequest(ctx, http.MethodGet, cdnName, u.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data *DomainsResp
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *cdnService) Get(ctx context.Context, domainID string) (*ExtendedDomain, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, cdnName, c.itemPath(domainID), nil)
	var data struct {
		Domain *ExtendedDomain `json:"domain"`
	}
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Domain, nil
}

func (c *cdnService) Create(ctx context.Context, cdrq *CreateDomainPayload) (*CreateDomainResp, error) {
	var data *CreateDomainResp
	req, err := c.client.NewRequest(ctx, http.MethodPost, cdnName, c.resourcePath(), &cdrq)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *cdnService) Update(ctx context.Context, domainID string, udrq *UpdateDomainReq) (*UpdateDomainResp, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPut, cdnName, c.itemPath(domainID), udrq)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	var data *UpdateDomainResp
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *cdnService) Delete(ctx context.Context, domainID string) error {
	req, err := c.client.NewRequest(ctx, http.MethodDelete, cdnName, c.itemPath(domainID), nil)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body)
	return resp.Body.Close()
}

func (c *cdnService) DeleteCache(ctx context.Context, domainID string, files *Files) error {
	req, err := c.client.NewRequest(ctx, http.MethodDelete, cdnName, c.itemPath(domainID), files)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body)
	return resp.Body.Close()
}
