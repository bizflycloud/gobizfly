package gobizfly

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// Agents contains agent information.
type Agents struct {
	Created   string `json:"_created"`
	Hostname  string `json:"hostname"`
	ID        string `json:"_id"`
	Name      string `json:"name"`
	Online    bool   `json:"online"`
	ProjectID string `json:"project_id"`
	Runtime   string `json:"runtime"`
	UserID    string `json:"user_id"`
}

func (a *agents) ResourcePath() string {
	return strings.Join([]string{agentsResourcePath}, "/")
}

func (a *agents) ItemPath(id string) string {
	return strings.Join([]string{agentsResourcePath, id}, "/")
}

// List returns a list of agents.
func (a *agents) List(ctx context.Context, filters *string) ([]*Agents, error) {
	req, err := a.client.NewRequest(ctx, http.MethodGet, cloudwatcherServiceName, a.ResourcePath(), nil)
	if err != nil {
		return nil, err
	}

	if filters != nil {
		q := req.URL.Query()
		q.Add("where", *filters)
		req.URL.RawQuery = q.Encode()
	}

	resp, err := a.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var data struct {
		Agents []*Agents `json:"_items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Agents, nil
}

// Get an agent
func (a *agents) Get(ctx context.Context, id string) (*Agents, error) {
	req, err := a.client.NewRequest(ctx, http.MethodGet, cloudwatcherServiceName, a.ItemPath(id), nil)
	if err != nil {
		return nil, err
	}
	resp, err := a.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	agent := &Agents{}
	if err := json.NewDecoder(resp.Body).Decode(agent); err != nil {
		return nil, err
	}

	return agent, nil
}

// Delete an agent
func (a *agents) Delete(ctx context.Context, id string) error {
	req, err := a.client.NewRequest(ctx, http.MethodDelete, cloudwatcherServiceName, a.ItemPath(id), nil)
	if err != nil {
		return err
	}
	resp, err := a.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(io.Discard, resp.Body)

	return resp.Body.Close()
}
