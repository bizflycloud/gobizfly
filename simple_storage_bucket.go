package gobizfly

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

const (
	simpleStoragePath    = "/_/"
	simpleStorageKeyPath = "/key"
)

var _ SimpleStorageService = (*cloudSimpleStorageService)(nil)

type SimpleStorageService interface {
	Create(ctx context.Context, s3cr *BucketCreateRequest) (*Bucket, error)
	List(ctx context.Context, opts *ListOptions) ([]*Bucket, error)
	ListWithBucketNameInfo(ctx context.Context, paramBucket ParamListWithBucketNameInfo) (*ResponseListBucketWithName, error)
	Delete(ctx context.Context, id string) error

	UpdateAcl(ctx context.Context, acl, bucketName string) (*ResponseUpdateACL, error)
	UpdateVersioning(ctx context.Context, versioning bool, bucketName string) (*ResponseVersioning, error)
	UpdateCors(ctx context.Context, paramUpdateCors *ParamUpdateCors) (*ResponseCors, error)
	UpdateWebsiteConfig(ctx context.Context, paramUpdateWebsiteConfig *ParamUpdateWebsiteConfig) (*ResponseWebsiteConfig, error)
	SimpleStorageKey() *cloudSimpleStorageKeyResource
}
type cloudSimpleStorageService struct {
	client *Client
}

type BucketCreateRequest struct {
	Name                string `json:"name"`
	Location            string `json:"location"`
	Acl                 string `json:"acl"`
	DefaultStorageClass string `json:"default_storage_class"`
}

type Bucket struct {
	Name                string  `json:"name"`
	CreatedAt           string  `json:"created_at"`
	Location            string  `json:"location"`
	SizeKb              float64 `json:"size_kb"`
	NumObjects          int     `json:"num_objects"`
	DefaultStorageClass string  `json:"default_storage_class"`
}

type ResponseAcl struct {
	Message string `json:"message"`
	Owner   struct {
		Id          string `json:"id"`
		DisplayName string `json:"display_name"`
	} `json:"owner"`
	Grants []struct {
		Permission string `json:"permission"`
		Grantee    struct {
			Type        string      `json:"type"`
			Id          string      `json:"id"`
			DisplayName string      `json:"display_name"`
			Email       interface{} `json:"email"`
			Uri         interface{} `json:"uri"`
		} `json:"grantee"`
	} `json:"grants"`
}

type ResponseUpdateACL struct {
	Acl ResponseAcl `json:"acl"`
}

type Rule struct {
	AllowedOrigin  string   `json:"allowed_origin"`
	AllowedMethods []string `json:"allowed_methods"`
	AllowedHeaders []string `json:"allowed_headers"`
	MaxAgeSeconds  int      `json:"max_age_seconds"`
}

type ParamUpdateCorsRead struct {
	Rules []Rule `json:"rules"`
}

type ParamUpdateCors struct {
	Rules      []Rule `json:"rules"`
	BucketName string `json:"bucket_name"`
}

type ResponseCors struct {
	Message string `json:"message"`
	Rules   []Rule `json:"rules"`
}

type ParamUpdateWebsiteConfig struct {
	Index      string `json:"index"`
	Error      string `json:"error"`
	BucketName string `json:"bucket_name"`
}

type ParamUpdateWebsiteConfigSend struct {
	Index string `json:"index"`
	Error string `json:"error"`
}

type ResponseWebsiteConfig struct {
	Message    string `json:"message"`
	WebsiteUrl string `json:"website_url"`
	Index      string `json:"index"`
	Error      string `json:"error"`
}

type ResponseListBucketWithName struct {
	Bucket Bucket      `json:"bucket"`
	Acl    ResponseAcl `json:"acl"`
	Cors   struct {
		Rules []Rule `json:"rules"`
	} `json:"cors"`
	Versioning struct {
		Status string `json:"status"`
	} `json:"versioning"`
	WebsiteConfig struct {
		WebsiteUrl string `json:"website_url"`
		Index      string `json:"index"`
		Error      string `json:"error"`
	} `json:"website_config"`
}

type ParamListWithBucketNameInfo struct {
	Policy          string `json:"policy"`
	Acl             string `json:"acl"`
	Cors            string `json:"cors"`
	Versioning      string `json:"versioning"`
	WebsiteConfig   string `json:"website_config"`
	LifecycleConfig string `json:"lifecycle_config"`
	Tags            string `json:"tags"`
	BucketName      string `json:"bucket_name"`
}

type ResponseVersioning struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func getQueryPaths(param interface{}) []string {
	queryParts := []string{}
	// Using reflect to iterate through fields in a struct
	v := reflect.ValueOf(param)
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		if field.Tag.Get("json") == "bucket_name" {
			continue
		}

		// Only add to the query if the value is not empty
		if value.Kind() == reflect.String && value.String() != "" {
			queryParts = append(queryParts, fmt.Sprintf("%s=%s", field.Tag.Get("json"), value.String()))
		}
	}
	return queryParts

}

func (c *cloudSimpleStorageService) resourcePathWithBucketInfo(param ParamListWithBucketNameInfo) string {
	queryParts := getQueryPaths(param)
	query := strings.Join(queryParts, "&")
	bucketName := param.BucketName
	return fmt.Sprintf("%s%s?%s", simpleStoragePath, bucketName, query)
}

func (c *cloudSimpleStorageService) resourcePathUpdateOption(bucketName, path string) string {
	return fmt.Sprintf("%s%s?%s", simpleStoragePath, bucketName, path)
}

func (c *cloudSimpleStorageService) resourcePath() string {
	return simpleStoragePath
}

func (c *cloudSimpleStorageService) itemPath(id string) string {
	return strings.Join([]string{simpleStoragePath, id}, "")
}

func (c cloudSimpleStorageService) Create(ctx context.Context, bucket *BucketCreateRequest) (*Bucket, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, simpleStorageServiceName, c.resourcePath(), &bucket)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respData struct {
		Message string  `json:"message"`
		Buckets *Bucket `json:"bucket"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.Buckets, nil
}

func (c *cloudSimpleStorageService) List(ctx context.Context, opts *ListOptions) ([]*Bucket, error) {
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
		Buckets []*Bucket `json:"buckets"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Buckets, nil
}

func (c *cloudSimpleStorageService) ListWithBucketNameInfo(ctx context.Context, paramBucket ParamListWithBucketNameInfo) (*ResponseListBucketWithName, error) {
	if paramBucket.BucketName == "" {
		return nil, errors.New("InvalidBucketName")
	}
	req, err := c.client.NewRequest(ctx, http.MethodGet, simpleStorageServiceName, c.resourcePathWithBucketInfo(paramBucket), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data = ResponseListBucketWithName{}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *cloudSimpleStorageService) Delete(ctx context.Context, id string) error {
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

func (c *cloudSimpleStorageService) UpdateAcl(ctx context.Context, acl, bucketName string) (*ResponseUpdateACL, error) {
	if bucketName == "" {
		return nil, errors.New("InvalidBucketName")
	}
	var paramUpdateAcl = struct {
		Acl string `json:"acl"`
	}{
		Acl: acl,
	}
	req, err := c.client.NewRequest(ctx, http.MethodPatch, simpleStorageServiceName, c.resourcePathUpdateOption(bucketName, "acl"), &paramUpdateAcl)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data *ResponseUpdateACL

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *cloudSimpleStorageService) UpdateVersioning(ctx context.Context, versioning bool, bucketName string) (*ResponseVersioning, error) {
	if bucketName == "" {
		return nil, errors.New("InvalidBucketName")
	}
	var paramUpdateVersion = struct {
		Versioning bool `json:"versioning"`
	}{
		Versioning: versioning,
	}
	req, err := c.client.NewRequest(ctx, http.MethodPatch, simpleStorageServiceName, c.resourcePathUpdateOption(bucketName, "versioning"), &paramUpdateVersion)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Versioning *ResponseVersioning `json:"versioning"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Versioning, nil
}

func (c *cloudSimpleStorageService) UpdateCors(ctx context.Context, paramUpdateCors *ParamUpdateCors) (*ResponseCors, error) {
	if paramUpdateCors.BucketName == "" {
		return nil, errors.New("InvalidBucketName")
	}
	var paramCors = struct {
		ParamCorsRead ParamUpdateCorsRead `json:"cors"`
	}{
		ParamCorsRead: ParamUpdateCorsRead{
			Rules: paramUpdateCors.Rules,
		},
	}

	req, err := c.client.NewRequest(ctx, http.MethodPatch, simpleStorageServiceName, c.resourcePathUpdateOption(paramUpdateCors.BucketName, "cors"), &paramCors)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Cors *ResponseCors `json:"cors"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Cors, nil
}

func (c *cloudSimpleStorageService) UpdateWebsiteConfig(ctx context.Context, paramUpdateWebsiteConfig *ParamUpdateWebsiteConfig) (*ResponseWebsiteConfig, error) {
	if paramUpdateWebsiteConfig.BucketName == "" {
		return nil, errors.New("InvalidBucketName")
	}
	var paramWeb = struct {
		ParamWebConfigSend ParamUpdateWebsiteConfigSend `json:"website_config"`
	}{
		ParamWebConfigSend: ParamUpdateWebsiteConfigSend{
			Index: paramUpdateWebsiteConfig.Index,
			Error: paramUpdateWebsiteConfig.Error,
		},
	}

	req, err := c.client.NewRequest(ctx, http.MethodPatch, simpleStorageServiceName, c.resourcePathUpdateOption(paramUpdateWebsiteConfig.BucketName, "website_config"), &paramWeb)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		WebsiteConfig *ResponseWebsiteConfig `json:"website_config"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.WebsiteConfig, nil
}

func (c *cloudSimpleStorageService) SimpleStorageKey() *cloudSimpleStorageKeyResource {
	return &cloudSimpleStorageKeyResource{client: c.client}
}

// file simple storage key

var _ SimpleStorageKeyService = (*cloudSimpleStorageKeyResource)(nil)

type SimpleStorageKeyService interface {
	CreateAccessKey(ctx context.Context, s3cr *KeyCreateRequest) (*KeyHaveSecret, error)
	GetAccessKey(ctx context.Context, id string) (*KeyHaveSecret, error)
	DeleteAccessKey(ctx context.Context, id string) error
	ListAccessKey(ctx context.Context, opts *ListOptions) ([]*KeyInList, error)
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

func (c *cloudSimpleStorageKeyResource) CreateAccessKey(ctx context.Context, dataCreatekey *KeyCreateRequest) (*KeyHaveSecret, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, simpleStorageServiceName, c.keyResourcePath(), &dataCreatekey)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respData struct {
		Message string         `json:"message"`
		Key     *KeyHaveSecret `json:"Key"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}
	return respData.Key, nil
}

func (c *cloudSimpleStorageKeyResource) keyResourcePath() string {
	return simpleStorageKeyPath
}

func (c *cloudSimpleStorageKeyResource) keyItemPath(id string) string {
	if id == "" {
		return simpleStorageKeyPath
	}
	return strings.Join([]string{simpleStorageKeyPath, id}, "/")
}

func (c *cloudSimpleStorageKeyResource) DeleteAccessKey(ctx context.Context, id string) error {
	req, err := c.client.NewRequest(ctx, http.MethodDelete, simpleStorageServiceName, c.keyItemPath(id), nil)
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

func (c *cloudSimpleStorageKeyResource) GetAccessKey(ctx context.Context, id string) (*KeyHaveSecret, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, simpleStorageServiceName, c.keyItemPath(id), nil)
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

func (c *cloudSimpleStorageKeyResource) ListAccessKey(ctx context.Context, opts *ListOptions) ([]*KeyInList, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, simpleStorageServiceName, c.keyResourcePath(), nil)
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
