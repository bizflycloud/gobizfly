// This file is part of gobizfly

package gobizfly

import (
	"context"
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
	UpdateRecord(ctx context.Context, recordID string, urpl interface{}) (*Record, error)
	DeleteRecord(ctx context.Context, recordID string) error
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
