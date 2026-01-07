// This file is part of gobizfly

package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	cloudDatabaseTasksResourcePath = "/tasks"
)

type cloudDatabaseTasks struct {
	client *Client
}

// CloudDatabaseTaskResult contains task result detail.
type CloudDatabaseTaskResult struct {
	Action   string                 `json:"action"`
	Data     map[string]interface{} `json:"data"`
	Message  string                 `json:"message"`
	Progress int                    `json:"progress"`
}

// CloudDatabaseTask contains task result.
type CloudDatabaseTask struct {
	Ready  bool                    `json:"ready"`
	Result CloudDatabaseTaskResult `json:"result"`
}

func (db *cloudDatabaseService) Tasks() *cloudDatabaseTasks {
	return &cloudDatabaseTasks{client: db.client}
}

// CloudDatabase Task Resource Path
func (ta *cloudDatabaseTasks) resourcePath(taskID string) string {
	return cloudDatabaseTasksResourcePath + "/" + taskID + "/status"
}

// Get a task status.
func (ta *cloudDatabaseTasks) Get(ctx context.Context, taskID string) (*CloudDatabaseTask, error) {
	req, err := ta.client.NewRequest(ctx, http.MethodGet, databaseServiceName, ta.resourcePath(taskID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := ta.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()
	var task *CloudDatabaseTask

	if err := json.NewDecoder(resp.Body).Decode(&task); err != nil {
		return nil, err
	}

	return task, nil
}
