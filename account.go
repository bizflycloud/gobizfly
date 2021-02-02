package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

const (
	regionsPath = "/regions"
)

type accountService struct {
	client *Client
}

var _ AccountService = (*accountService)(nil)

type AccountService interface {
	ListRegion(ctx context.Context) (*Regions, error)
	GetRegion(ctx context.Context, regionName string) (*Region, error)
}

type Region struct {
	Active     bool          `json:"active"`
	Icon       string        `json:"icon"`
	Name       string        `json:"name"`
	Order      int           `json:"order"`
	RegionName string        `json:"region_name"`
	ShortName  string        `json:"short_name"`
	Zones      []AccountZone `json:"zones"`
}

type AccountZone struct {
	Active    bool   `json:"active"`
	Icon      string `json:"icon"`
	Name      string `json:"name"`
	Order     int    `json:"order"`
	ShortName string `json:"short_name"`
}

type Regions struct {
	HN  Region `json:"HN"`
	HCM Region `json:"HCM"`
}

func (a accountService) resourcePath() string {
	return regionsPath
}

func (a accountService) itemPath(name string) string {
	return strings.Join([]string{regionsPath, name}, "/")
}

func (a accountService) ListRegion(ctx context.Context) (*Regions, error) {
	req, err := a.client.NewRequest(ctx, http.MethodGet, accountName, a.resourcePath(), nil)
	if err != nil {
		return nil, err
	}
	var regions *Regions
	resp, err := a.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&regions); err != nil {
		return nil, err
	}
	return regions, nil
}

func (a accountService) GetRegion(ctx context.Context, regionName string) (*Region, error) {
	req, err := a.client.NewRequest(ctx, http.MethodGet, accountName, a.itemPath(regionName), nil)
	if err != nil {
		return nil, err
	}
	var region *Region
	resp, err := a.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&region); err != nil {
		return nil, err
	}
	return region, nil
}
