package gobizfly

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// AutoScalingSize - size of auto scaling group
type AutoScalingSize struct {
	DesiredCapacity int `json:"desired_capacity"`
	MaxSize         int `json:"max_size"`
	MinSize         int `json:"min_size"`
}

// AutoScalingScheduleValid - represents for a validation time of cron triggers
type autoScalingScheduleValid struct {
	From string  `json:"_from"`
	To   *string `json:"_to,omitempty"`
}

// AutoScalingScheduleInputs - represents for a input of cron triggers
type autoScalingScheduleInputs struct {
	CronPattern string          `json:"cron_pattern"`
	Inputs      AutoScalingSize `json:"inputs"`
}

// AutoScalingScheduleSizing - represents for phase time of cron triggers
type autoScalingScheduleSizing struct {
	From autoScalingScheduleInputs `json:"_from"`
	To   autoScalingScheduleInputs `json:"_to,omitempty"`
	Type string                    `json:"_type"`
}

// AutoScalingScheduleCreateRequest - payload use create a scheduler (cron trigger)
type autoScalingScheduleCreateRequest struct {
	Name   string                    `json:"name"`
	Sizing autoScalingScheduleSizing `json:"sizing"`
	Valid  autoScalingScheduleValid  `json:"valid"`
}

// AutoScalingSchedule - cron triggers to do time-based scale
type AutoScalingSchedule struct {
	ClusterID         string                    `json:"cluster_id"`
	Created           string                    `json:"created_at"`
	ID                string                    `json:"_id"`
	Name              string                    `json:"name"`
	Sizing            autoScalingScheduleSizing `json:"sizing"`
	Status            string                    `json:"status"`
	TaskID            string                    `json:"task_id"`
	Valid             autoScalingScheduleValid  `json:"valid"`
	NextExecutionTime string                    `json:"next_execution_time"`
}

// List - list all cron triggers of a cluster
func (s *schedule) List(ctx context.Context, clusterID string) ([]*AutoScalingSchedule, error) {
	if clusterID == "" {
		return nil, errors.New("auto scaling group ID is required")
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, autoScalingServiceName, s.resourcePath(clusterID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var data struct {
		AutoScalingSchdeules []*AutoScalingSchedule `json:"cron_triggers"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.AutoScalingSchdeules, nil
}

// Get - get a cron trigger of a cluster
func (s *schedule) Get(ctx context.Context, clusterID, scheduleID string) (*AutoScalingSchedule, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, autoScalingServiceName, s.itemPath(clusterID, scheduleID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var data = &AutoScalingSchedule{}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

// Delete - delete a cron trigger of a cluster
func (s *schedule) Delete(ctx context.Context, clusterID, scheduleID string) error {
	req, err := s.client.NewRequest(ctx, http.MethodDelete, autoScalingServiceName, s.itemPath(clusterID, scheduleID), nil)
	if err != nil {
		return err
	}

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(io.Discard, resp.Body)

	return resp.Body.Close()
}

// Create - create a cron trigger of a cluster
func (s *schedule) Create(ctx context.Context, clusterID string, asscr *autoScalingScheduleCreateRequest) (*TaskResponses, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, autoScalingServiceName, s.resourcePath(clusterID), &asscr)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	data := &TaskResponses{}

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}
	return data, nil
}
