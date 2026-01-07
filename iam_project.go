package gobizfly

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bizflycloud/gobizfly/utils"
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

// ListProjectsOpts params for list projects
type ListProjectsOpts struct {
	Limit *string `json:"limit,omitempty"`
	Page  *string `json:"page,omitempty"`
	Sort  *string `json:"sort,omitempty"`
}

func (i *iamService) ListProjects(ctx context.Context, opts ListProjectsOpts) ([]*IAMProject, error) {
	var convOpts map[string]interface{}
	req, err := i.client.NewRequest(ctx, http.MethodGet, iamServiceName, i.projectResourcePath(), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	err = utils.ConvDataWithJson(opts, &convOpts)
	if err != nil {
		return nil, err
	}
	for key, val := range convOpts {
		q.Add(key, fmt.Sprintf("%v", val))
	}
	req.URL.RawQuery = q.Encode()
	resp, err := i.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	var respData struct {
		Data []*IAMProject `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.Data, nil
}
