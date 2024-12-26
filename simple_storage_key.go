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
	simpleStorageKeyPath = "/simple-storage/key"
)

var _ SimpleStoreKey = (*cloudSimpleStoreKeyService)(nil)

type SimpleStoreKey interface {
	Create(ctx context.Context, s3cr *KeyCreateRequest) (*ResponseKeyCreate, error)
	Get(ctx context.Context, id string) (*dataKeys, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, opts *ListOptions) ([]*key, error)
}
type cloudSimpleStoreKeyService struct {
	client *Client
}

type KeyCreateRequest struct {
	SubuserId string `json:"subuser_id"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

type Key struct {
	User      string `json:"user"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

type ResponseKeyCreate struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

type dataKeys struct {
	Keys []struct {
		User      string `json:"user"`
		AccessKey string `json:"access_key"`
	} `json:"keys"`
}
type key struct {
	User      string `json:"user"`
	AccessKey string `json:"access_key"`
}

func (c *cloudSimpleStoreService) SimpleStoreKey() *cloudSimpleStoreKeyService {
	return &cloudSimpleStoreKeyService{client: c.client}
}

func (c cloudSimpleStoreKeyService) Create(ctx context.Context, dataCreatekey *KeyCreateRequest) (*ResponseKeyCreate, error) {
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
		Message string             `json:"message"`
		Key     *ResponseKeyCreate `json:"Key"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.Key, nil
}

func (c *cloudSimpleStoreKeyService) resourcePath() string {
	return simpleStorageKeyPath
}

func (c *cloudSimpleStoreKeyService) itemPath(id string) string {
	if id == "" {
		return simpleStorageKeyPath
	}
	return strings.Join([]string{simpleStorageKeyPath, id}, "/")
}

func (c *cloudSimpleStoreKeyService) Delete(ctx context.Context, id string) error {
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

func (c *cloudSimpleStoreKeyService) Get(ctx context.Context, id string) (*dataKeys, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, simpleStorageServiceName, c.itemPath(id), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	key := &dataKeys{}
	if err := json.NewDecoder(resp.Body).Decode(key); err != nil {
		return nil, err
	}
	return key, nil
}

func (c *cloudSimpleStoreKeyService) List(ctx context.Context, opts *ListOptions) ([]*key, error) {
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
		Keys []*key `json:"keys"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Keys, nil
}
