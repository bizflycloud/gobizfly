// This file is part of gobizfly

package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	cloudDatabaseAutoScalingsResourcePath = "/autoscaling"
)

func (db *cloudDatabaseService) AutoScalings() *cloudDatabaseAutoScalings {
	return &cloudDatabaseAutoScalings{client: db.client}
}

type cloudDatabaseAutoScalings struct {
	client *Client
}

// CloudDatabase AutoScaling Resource Path
func (au *cloudDatabaseAutoScalings) resourcePath(resourceID string) string {
	return cloudDatabaseInstancesResourcePath + "/" + resourceID + cloudDatabaseAutoScalingsResourcePath
}

// CloudDatabaseAutoScalingVolume contains volume information of autoscaling.
type CloudDatabaseAutoScalingVolume struct {
	Limited   int `json:"limited"`
	Threshold int `json:"threshold"`
}

// CloudDatabaseAutoScalingReceivers contains information about receivers of autoscaling
type CloudDatabaseAutoScalingReceivers struct {
	Action string `json:"action"`
	ID     string `json:"id"`
	Name   string `json:"name"`
}

// CloudDatabaseAutoScalingAlarms contains information about alarms of autoscaling
type CloudDatabaseAutoScalingAlarms struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ReceiverID string `json:"receiver_id"`
}

// CloudDatabaseAutoScaling contains autoscaling information of a database instance.
type CloudDatabaseAutoScaling struct {
	Alarms    []CloudDatabaseAutoScalingAlarms    `json:"alarms,omitempty"`
	Enable    bool                                `json:"enable"`
	Receivers []CloudDatabaseAutoScalingReceivers `json:"receivers,omitempty"`
	Volume    CloudDatabaseAutoScalingVolume      `json:"volume"`
}

// Create a new autoscaling.
func (au *cloudDatabaseAutoScalings) Create(ctx context.Context, instanceID string, option *CloudDatabaseAutoScaling) (*CloudDatabaseMessageResponse, error) {
	req, err := au.client.NewRequest(ctx, http.MethodPost, databaseServiceName, au.resourcePath(instanceID), option)
	if err != nil {
		return nil, err
	}

	resp, err := au.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	var dmr *CloudDatabaseMessageResponse
	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}
	return dmr, nil
}

// Update a autoscaling.
func (au *cloudDatabaseAutoScalings) Update(ctx context.Context, instanceID string, option *CloudDatabaseAutoScaling) (*CloudDatabaseMessageResponse, error) {
	req, err := au.client.NewRequest(ctx, http.MethodPut, databaseServiceName, au.resourcePath(instanceID), option)
	if err != nil {
		return nil, err
	}

	resp, err := au.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}

// Delete a autoscaling.
func (au *cloudDatabaseAutoScalings) Delete(ctx context.Context, instanceID string) (*CloudDatabaseMessageResponse, error) {
	req, err := au.client.NewRequest(ctx, http.MethodDelete, databaseServiceName, au.resourcePath(instanceID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := au.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}
