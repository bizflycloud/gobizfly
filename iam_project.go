package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
)

// IAMProject - contains the information of a IAM project
type IAMProject struct {
	IsActive    bool   `json:"is_active"`
	CreatedAt   string `json:"created_at"`
	OriginName  string `json:"origin_name"`
	AliasName   string `json:"alias_name"`
	Description string `json:"description"`
	UpdatedAt   string `json:"updated_at"`
	UUID        string `json:"uuid"`
	ShortUUID   string `json:"short_uuid"`
}

func (i *iamService) ListProjects(ctx context.Context) ([]*IAMProject, error) {
	req, err := i.client.NewRequest(ctx, http.MethodGet, iamServiceName, i.projectResourcePath(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := i.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var respData struct {
		Data []*IAMProject `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.Data, nil
}
