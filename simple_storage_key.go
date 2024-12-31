package gobizfly

import "C"
import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	simpleStorageKeyPath = "/key"
)

var _ SimpleStorageKeyService = (*cloudSimpleStorageKeyResource)(nil)


type SimpleStorageKeyService interface {
	Create(ctx context.Context, s3cr *KeyCreateRequest) (*KeyHaveSecret, error)
	Get(ctx context.Context, id string) (*KeyHaveSecret, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, opts *ListOptions) ([]*KeyInList, error)
}

type KeyCreateRequest struct {
	SubuserId string `json:"subuser_id"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

type KeyHaveSecret struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

type KeyInList struct {
	User      string `json:"user"`
	AccessKey string `json:"access_key"`
}


type cloudSimpleStorageKeyResource struct {
	client *Client
}

func (c *cloudSimpleStorageKeyResource) Create(ctx context.Context, dataCreatekey *KeyCreateRequest) (*KeyHaveSecret, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, simpleStorageServiceName, c.resourcePath(), &dataCreatekey)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respData struct {
		Message string          `json:"message"`
		Key     *KeyHaveSecret `json:"Key"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.Key, nil
}

func (c *cloudSimpleStorageKeyResource) resourcePath() string {
	return simpleStorageKeyPath
}

func (c *cloudSimpleStorageKeyResource) itemPath(id string) string {
	if id == "" {
		return simpleStorageKeyPath
	}
	return strings.Join([]string{simpleStorageKeyPath, id}, "/")
}

func (c *cloudSimpleStorageKeyResource) Delete(ctx context.Context, id string) error {
	req, err := c.client.NewRequest(ctx, http.MethodDelete, simpleStorageServiceName, c.itemPath(id), nil)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body)

	return resp.Body.Close()
}

func (c *cloudSimpleStorageKeyResource) Get(ctx context.Context, id string) (*KeyHaveSecret, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, simpleStorageServiceName, c.itemPath(id), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	key := &KeyHaveSecret{}
	if err := json.NewDecoder(resp.Body).Decode(key); err != nil {
		return nil, err
	}
	return key, nil
}

func (c *cloudSimpleStorageKeyResource) List(ctx context.Context, opts *ListOptions) ([]*KeyInList, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, simpleStorageServiceName, c.resourcePath(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Keys []*KeyInList `json:"keys"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Keys, nil
}
