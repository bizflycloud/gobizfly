// Copyright 2019 The BizFly Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gobizfly

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"io"
	"net/http"
	"fmt"
)

const (
	volumeBasePath	=	"/iaas-cloud/api/volumes"
)

var _ VolumeService = (*volume)(nil)

type VolumeService interface {
	List(ctx context.Context, opts *ListOptions) ([]*Volume, error)
	Create(ctx context.Context, req *VolumeCreateRequest) (*Volume, error)
	Get(ctx context.Context, id string) (*Volume, error)
	Delete(ctx context.Context, id string) error
}

type VolumeCreateRequest struct {
	Name 				string		`json:"name"`
	Size				int			`json:"size"`
	VolumeType			string		`json:"volume_type"`
	AvailabilityZone	string	`json:"availability_zone"`
	SnapshotID			string			`json:"snapshot_id,omitempty"`
}

type VolumeAttachment struct {
	ServerID	string	`json:"server_id"`
	AttachmentID	string	`json:"attachment_id"`
	VolumeID	string	`json:"volume_id"`
	Device	string	`json:"device"`
	ID	string	`json:"id"`
}

type Volume struct {
	ID					string 	`json:"id"`
	Size				int		`json:"size"`
	AttachedType		string	`json:"attached_type"`
	Name				string	`json:"name"`
	VolumeType			string	`json:"volume_type"`
	Description			string	`json:"description"`
	SnapshotID			string	`json:"snapshot_id"`
	Bootable			string	`json:"bootable"`
	AvailabilityZone	string	`json:"availability_zone"`
	Status				string	`json:"status"`
	UserID				string	`json:"user_id"`
	ProjectID			string	`json:"os-vol-tenant-attr:tenant_id"`
	CreatedAt			string	`json:"created_at"`
	UpdatedAt			string	`json:"updated_at"`
	Metadata			map[string]string	`json:"metadata"`
	Attachments			[]Server	`json:"attachments"`
}


type volume struct {
	client *Client
}

func (v *volume) List(ctx context.Context, opts *ListOptions) ([]*Volume, error) {
	req, err := v.client.NewRequest(ctx, http.MethodGet, volumeBasePath, nil)
	if err != nil {
		return nil, err
	}
	resp, err := v.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var volumes []*Volume

	if err := json.NewDecoder(resp.Body).Decode(&volumes); err != nil {
		return nil, err
	}
	return volumes, nil
}


func (v *volume) Create(ctx context.Context, vcr *VolumeCreateRequest) (*Volume, error) {
	req, err :=  v.client.NewRequest(ctx, http.MethodPost, volumeBasePath, &vcr)
	if err != nil {
		return nil, err
	}
	resp, err := v.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	
	var volumeRespData *Volume

	if err := json.NewDecoder(resp.Body).Decode(&volumeRespData); err != nil {
		return nil, err
	}
	return volumeRespData, nil
}


func (v *volume) Get(ctx context.Context, id string) (*Volume, error) {
	req, err := v.client.NewRequest(ctx, http.MethodGet, volumeBasePath + "/" + id, nil)
	if err != nil {
		return nil, err
	}

	resp, err := v.client.Do(ctx, req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var volume *Volume

	if err := json.NewDecoder(resp.Body).Decode(&volume); err != nil {
		return nil, err
	}
	return volume, nil
}

func (v *volume) Delete(ctx context.Context, id string) error {
	req, err := v.client.NewRequest(ctx, http.MethodDelete, volumeBasePath + "/" + id, nil)

	if err != nil {
		return err
	}

	resp, err := v.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body)

	return resp.Body.Close()
}


func (v *volume) ExtendVolume(ctx context.Context, id string, newsize int) (*Volume, error) {
	return nil, nil
}

func (v *volume) Attach(ctx context.Context, id string, serverID string) (*Volume, error) {
	return nil, nil
}

func (v *volume) Detach(ctx context.Context, id string, serverID string) (*Volume, error) {
	return nil, nil
}
