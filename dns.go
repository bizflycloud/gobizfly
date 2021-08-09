// This file is part of gobizfly

package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

const (
	zonesPath  = "/zones"
	zonePath   = "/zone"
	recordPath = "/record"
)

type dnsService struct {
	client *Client
}

var _ DNSService = (*dnsService)(nil)

type DNSService interface {
	ListZones(ctx context.Context, opts *ListOptions) (*ListZoneResp, error)
	CreateZone(ctx context.Context, czpl *CreateZonePayload) (*ExtendedZone, error)
	GetZone(ctx context.Context, zoneID string) (*ExtendedZone, error)
	DeleteZone(ctx context.Context, zoneID string) error
	CreateRecord(ctx context.Context, zoneID string, crpl interface{}) (*Record, error)
	GetRecord(ctx context.Context, recordID string) (*Record, error)
	UpdateRecord(ctx context.Context, recordID string, urpl *UpdateRecordPayload) (*ExtendedRecord, error)
	DeleteRecord(ctx context.Context, recordID string) error
}

type Zone struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Deleted    int      `json:"deleted"`
	CreatedAt  string   `json:"created_at"`
	UpdatedAt  string   `json:"updated_at"`
	TenantId   string   `json:"tenant_id"`
	NameServer []string `json:"nameserver"`
	TTL        int      `json:"ttl"`
	Active     bool     `json:"active"`
}

type WrappedZonePayload struct {
	Zones *CreateZonePayload `json:"zones"`
}

type WrappedRecordPayload struct {
	Record interface{} `json:"record"`
}

type ExtendedZone struct {
	Zone
	RecordsSet []RecordSet `json:"record_set"`
}

type RecordSet struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	TTL  string `json:"ttl"`
}

type Meta struct {
	MaxResults int `json:"max_results"`
	Total      int `json:"total"`
	Page       int `json:"page"`
}

type ListZoneResp struct {
	Zones []Zone `json:"zones"`
	Meta  Meta   `json:"_meta"`
}

type CreateZonePayload struct {
	Name        string `json:"name"`
	Required    bool   `json:"required,omitempty"`
	Description string `json:"description,omitempty"`
}

type Addrs struct {
	HN  []string `json:"HN"`
	HCM []string `json:"HCM"`
	SG  []string `json:"SG"`
	USA []string `json:"USA"`
}
type RoutingData struct {
	AddrsV4 Addrs `json:"addrs_v4,omitempty"`
	AddrsV6 Addrs `json:"addrs_v6,omitempty"`
}

type RoutingPolicyData struct {
	RoutingData RoutingData `json:"routing_data,omitempty"`
	HealthCheck HealthCheck `json:"healthcheck,omitempty"`
}

type HealthCheck struct {
	TCPConnect TCPHealthCheck  `json:"tcp_connect,omitempty"`
	HTTPStatus HTTPHealthCheck `json:"http_status,omitempty"`
}

type TCPHealthCheck struct {
	TCPPort int `json:"tcp_port,omitempty"`
}

type HTTPHealthCheck struct {
	HTTPPort int    `json:"http_port,omitempty"`
	URLPath  string `json:"url_path,omitempty"`
	VHost    string `json:"vhost,omitempty"`
	OkCodes  []int  `json:"ok_codes,omitempty"`
	Interval int    `json:"interval,omitempty"`
}

type BaseCreateRecordPayload struct {
	Name string `json:"name"`
	Type string `json:"type"`
	TTL  int    `json:"ttl"`
}

type CreateNormalRecordPayload struct {
	BaseCreateRecordPayload
	Data []string `json:"data"`
}

type CreatePolicyRecordPayload struct {
	BaseCreateRecordPayload
	RoutingPolicyData RoutingPolicyData `json:"routing_policy_data"`
}

type CreateMXRecordPayload struct {
	BaseCreateRecordPayload
	Data []MXData `json:"data"`
}

type MXData struct {
	Value    string `json:"value"`
	Priority int    `json:"priority"`
}

type RecordData struct {
	Value    string `json:"value"`
	Priority int    `json:"priority"`
}

type UpdateRecordPayload struct {
	Name              string            `json:"name,omitempty"`
	Type              string            `json:"type,omitempty"`
	TTL               int               `json:"ttl,omitempty"`
	Data              []MXData          `json:"data,omitempty"`
	RoutingPolicyData RoutingPolicyData `json:"routing_policy_data,omitempty"`
}

type Record struct {
	ID                string            `json:"id"`
	Name              string            `json:"name"`
	Delete            int               `json:"deleted"`
	CreatedAt         string            `json:"created_at"`
	UpdatedAt         string            `json:"updated_at"`
	TenantID          string            `json:"tenant_id"`
	ZoneID            string            `json:"zone_id"`
	Type              string            `json:"type"`
	TTL               int               `json:"ttl"`
	Data              []interface{}     `json:"data"`
	RoutingPolicyData RoutingPolicyData `json:"routing_policy_data"`
}

type ExtendedRecord struct {
	Record
	Data []RecordData `json:"data"`
}

type Records struct {
	Records []Record `json:"records"`
}

func (d dnsService) resourcePath() string {
	return zonesPath
}

func (d dnsService) zoneItemPath(id string) string {
	return strings.Join([]string{zonePath, id}, "/")
}

func (d dnsService) recordItemPath(id string) string {
	return strings.Join([]string{recordPath, id}, "/")
}

func (d *dnsService) ListZones(ctx context.Context, opts *ListOptions) (*ListZoneResp, error) {
	req, err := d.client.NewRequest(ctx, http.MethodGet, dnsName, d.resourcePath(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := d.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data *ListZoneResp
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (d *dnsService) CreateZone(ctx context.Context, czpl *CreateZonePayload) (*ExtendedZone, error) {
	payload := WrappedZonePayload{
		Zones: czpl,
	}
	req, err := d.client.NewRequest(ctx, http.MethodPost, dnsName, d.resourcePath(), payload)
	if err != nil {
		return nil, err
	}
	resp, err := d.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data *ExtendedZone
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (d *dnsService) GetZone(ctx context.Context, zoneID string) (*ExtendedZone, error) {
	req, err := d.client.NewRequest(ctx, http.MethodGet, dnsName, d.zoneItemPath(zoneID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := d.client.Do(ctx, req)
	var data *ExtendedZone
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (d *dnsService) DeleteZone(ctx context.Context, zoneID string) error {
	req, err := d.client.NewRequest(ctx, http.MethodDelete, dnsName, d.zoneItemPath(zoneID), nil)
	if err != nil {
		return err
	}
	resp, err := d.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}

func (d *dnsService) CreateRecord(ctx context.Context, zoneID string, crpl interface{}) (*Record, error) {
	payload := WrappedRecordPayload{
		Record: crpl,
	}
	req, err := d.client.NewRequest(ctx, http.MethodPost, dnsName, strings.Join([]string{d.zoneItemPath(zoneID), "record"}, "/"), &payload)
	if err != nil {
		return nil, err
	}
	resp, err := d.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data *struct {
		Record *Record `json:"record"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Record, nil
}

func (d *dnsService) GetRecord(ctx context.Context, recordID string) (*Record, error) {
	req, err := d.client.NewRequest(ctx, http.MethodGet, dnsName, d.recordItemPath(recordID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := d.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data *struct {
		Record *Record `json:"record"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Record, nil
}

func (d *dnsService) UpdateRecord(ctx context.Context, recordID string, urpl *UpdateRecordPayload) (*ExtendedRecord, error) {
	req, err := d.client.NewRequest(ctx, http.MethodPut, dnsName, d.recordItemPath(recordID), urpl)
	if err != nil {
		return nil, err
	}
	resp, err := d.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data *ExtendedRecord
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (d *dnsService) DeleteRecord(ctx context.Context, recordID string) error {
	req, err := d.client.NewRequest(ctx, http.MethodDelete, dnsName, d.recordItemPath(recordID), nil)
	if err != nil {
		return err
	}
	resp, err := d.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}
