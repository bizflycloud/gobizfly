// This file is part of gobizfly

package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	cloudDatabaseBackupsResourcePath         = "/backups"
	cloudDatabaseBackupSchedulesResourcePath = "/schedules"
)

type cloudDatabaseBackups struct {
	client *Client
}

type cloudDatabaseBackupSchedules struct {
	client *Client
}

// CloudDatabaseBackup contains detail of a backup.
type CloudDatabaseBackup struct {
	Created     string                 `json:"created"`
	Datastore   CloudDatabaseDatastore `json:"datastore"`
	Description string                 `json:"description"`
	ID          string                 `json:"id"`
	Message     string                 `json:"message"`
	Name        string                 `json:"name"`
	NodeID      string                 `json:"node_id"`
	ParentID    string                 `json:"parent_id"`
	ProjectID   string                 `json:"project_id"`
	Size        float32                `json:"size"`
	Status      string                 `json:"status"`
	Type        string                 `json:"type"`
	Updated     string                 `json:"updated"`
}

// CloudDatabaseBackupCreate contains payload require create backup.
type CloudDatabaseBackupCreate struct {
	Name       string `json:"backup_name,omitempty" validate:"required"`
	NodeID     string `json:"node_id,omitempty"`
	ParentID   string `json:"parent_id,omitempty"`
	Suggestion bool   `json:"suggestion,omitempty"`
}

// CloudDatabaseBackupResource contains option list backup.
type CloudDatabaseBackupResource struct {
	ResourceID   string `json:"resource_id"`
	ResourceType string `json:"resource_type"`
}

// CloudDatabase Schedule Backup Struct

// CloudDatabaseBackupSchedule contains detail of a schedule.
type CloudDatabaseBackupSchedule struct {
	CronExpression     string `json:"pattern"`
	FirstExecutionTime string `json:"first_execution_time"`
	ID                 string `json:"id"`
	InstanceID         string `json:"instance_id"`
	InstanceName       string `json:"instance_name"`
	LimitBackup        int    `json:"limit_backup"`
	Message            string `json:"message"`
	Name               string `json:"name"`
	NextExecutionTime  string `json:"next_execution_time"`
	NodeID             string `json:"node_id"`
	NodeName           string `json:"node_name"`
	ProjectID          string `json:"project_id"`
}

// CloudDatabaseBackupScheduleCreate contains schedule create payload info.
type CloudDatabaseBackupScheduleCreate struct {
	CronExpression string `json:"pattern"`
	LimitBackup    int    `json:"limit_backup,omitempty" validate:"required"`
	Name           string `json:"schedule_name,omitempty" validate:"required"`
}

// CloudDatabaseBackupScheduleDelete contains option when delete a schedule.
type CloudDatabaseBackupScheduleDelete struct {
	PurgeBackup bool `json:"purge_backup"`
}

// CloudDatabaseBackupScheduleListResourceOption contains option when list a database schedule.
type CloudDatabaseBackupScheduleListResourceOption struct {
	All          bool   `json:"all,omitempty"`
	ListBackup   bool   `json:"list_backup,omitempty"`
	ResourceID   string `json:"resource_id,omitempty"`
	ResourceType string `json:"resource_type,omitempty"`
}

func (db *cloudDatabaseService) Backups() *cloudDatabaseBackups {
	return &cloudDatabaseBackups{client: db.client}
}

func (db *cloudDatabaseService) BackupSchedules() *cloudDatabaseBackupSchedules {
	return &cloudDatabaseBackupSchedules{client: db.client}
}

// CloudDatabase Backup Resource Path
func (bk *cloudDatabaseBackups) resourcePath(backupID string) string {
	return cloudDatabaseBackupsResourcePath + "/" + backupID
}

// CloudDatabase Backup resourceType Create Path
func (bk *cloudDatabaseBackups) resourceCreatePath(resourceType string, resourceID string) string {
	return "/" + resourceType + "/" + resourceID + cloudDatabaseBackupsResourcePath
}

// CloudDatabase Schedule Resource List Path
func (sc *cloudDatabaseBackupSchedules) resourcePath(scheduleID string) string {
	return cloudDatabaseBackupSchedulesResourcePath + "/" + scheduleID
}

// CloudDatabase Schedule List In resourceType Path
func (sc *cloudDatabaseBackupSchedules) resourceCreatePath(resourceType string, resourceID string) string {
	return "/" + resourceType + "/" + resourceID + cloudDatabaseBackupSchedulesResourcePath
}

// CloudDatabase Schedule Resource Backup Path
func (sc *cloudDatabaseBackupSchedules) resourceBackupPath(scheduleID string) string {
	return cloudDatabaseBackupSchedulesResourcePath + "/" + scheduleID + cloudDatabaseBackupsResourcePath
}

// List all backups.
func (bk *cloudDatabaseBackups) List(ctx context.Context, resource *CloudDatabaseBackupResource, opts *CloudDatabaseListOption) ([]*CloudDatabaseBackup, error) {
	resourcePath := cloudDatabaseBackupsResourcePath
	if resource != nil {
		resourcePath = bk.resourceCreatePath(resource.ResourceType, resource.ResourceID)
	}

	req, err := bk.client.NewRequest(ctx, http.MethodGet, databaseServiceName, resourcePath, nil)
	if err != nil {
		return nil, err
	}

	AddParamsListOption(req, opts)

	resp, err := bk.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var data struct {
		Backups []*CloudDatabaseBackup `json:"backups"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Backups, nil
}

// Create a backups.
func (bk *cloudDatabaseBackups) Create(ctx context.Context, resourceType string, resourceID string, cr *CloudDatabaseBackupCreate) (*CloudDatabaseBackup, error) {
	req, err := bk.client.NewRequest(ctx, http.MethodPost, databaseServiceName, bk.resourceCreatePath(resourceType, resourceID), &cr)
	if err != nil {
		return nil, err
	}

	resp, err := bk.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var backups *CloudDatabaseBackup

	if err := json.NewDecoder(resp.Body).Decode(&backups); err != nil {
		return nil, err
	}

	return backups, nil
}

// Get a backups.
func (bk *cloudDatabaseBackups) Get(ctx context.Context, backupID string) (*CloudDatabaseBackup, error) {
	req, err := bk.client.NewRequest(ctx, http.MethodGet, databaseServiceName, bk.resourcePath(backupID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := bk.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var backups *CloudDatabaseBackup

	if err := json.NewDecoder(resp.Body).Decode(&backups); err != nil {
		return nil, err
	}

	return backups, nil
}

// Delete a backups.
func (bk *cloudDatabaseBackups) Delete(ctx context.Context, backupID string) (*CloudDatabaseMessageResponse, error) {
	req, err := bk.client.NewRequest(ctx, http.MethodDelete, databaseServiceName, bk.resourcePath(backupID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := bk.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}

// List all schedule.
func (sc *cloudDatabaseBackupSchedules) List(ctx context.Context, resource *CloudDatabaseBackupScheduleListResourceOption, opts *CloudDatabaseListOption) ([]*CloudDatabaseBackupSchedule, error) {
	var resourcePath string

	switch {
	case resource.All:
		resourcePath = cloudDatabaseBackupSchedulesResourcePath
	case resource.ListBackup && resource.ResourceID != "":
		resourcePath = sc.resourceBackupPath(resource.ResourceID)
	case resource.ResourceType != "" && resource.ResourceID != "":
		resourcePath = sc.resourceCreatePath(resource.ResourceType, resource.ResourceID)
	default:
		resourcePath = cloudDatabaseBackupSchedulesResourcePath
	}

	req, err := sc.client.NewRequest(ctx, http.MethodGet, databaseServiceName, resourcePath, nil)
	if err != nil {
		return nil, err
	}

	AddParamsListOption(req, opts)

	resp, err := sc.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var data struct {
		BackupSchedules []*CloudDatabaseBackupSchedule `json:"schedules"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.BackupSchedules, nil
}

// ListBackups list all backup in schedule.
func (sc *cloudDatabaseBackupSchedules) ListBackups(ctx context.Context, scheduleID string, opts *CloudDatabaseListOption) ([]*CloudDatabaseBackup, error) {
	req, err := sc.client.NewRequest(ctx, http.MethodGet, databaseServiceName, sc.resourceBackupPath(scheduleID), nil)
	if err != nil {
		return nil, err
	}

	AddParamsListOption(req, opts)

	resp, err := sc.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var data struct {
		Backups []*CloudDatabaseBackup `json:"backups"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Backups, nil
}

// Create new schedule.
func (sc *cloudDatabaseBackupSchedules) Create(ctx context.Context, nodeID string, scc *CloudDatabaseBackupScheduleCreate) (*CloudDatabaseBackupSchedule, error) {
	req, err := sc.client.NewRequest(ctx, http.MethodPost, databaseServiceName, sc.resourceCreatePath("nodes", nodeID), &scc)
	if err != nil {
		return nil, err
	}

	resp, err := sc.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var schedules *CloudDatabaseBackupSchedule

	if err := json.NewDecoder(resp.Body).Decode(&schedules); err != nil {
		return nil, err
	}

	return schedules, nil
}

// Get a schedule.
func (sc *cloudDatabaseBackupSchedules) Get(ctx context.Context, scheduleID string) (*CloudDatabaseBackupSchedule, error) {
	req, err := sc.client.NewRequest(ctx, http.MethodGet, databaseServiceName, sc.resourcePath(scheduleID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := sc.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var schedules *CloudDatabaseBackupSchedule

	if err := json.NewDecoder(resp.Body).Decode(&schedules); err != nil {
		return nil, err
	}

	return schedules, nil
}

// Delete a schedule.
func (sc *cloudDatabaseBackupSchedules) Delete(ctx context.Context, scheduleID string, option *CloudDatabaseBackupScheduleDelete) (*CloudDatabaseMessageResponse, error) {
	req, err := sc.client.NewRequest(ctx, http.MethodDelete, databaseServiceName, sc.resourcePath(scheduleID), option)
	if err != nil {
		return nil, err
	}

	resp, err := sc.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var dmr *CloudDatabaseMessageResponse

	if err := json.NewDecoder(resp.Body).Decode(&dmr); err != nil {
		return nil, err
	}

	return dmr, nil
}
