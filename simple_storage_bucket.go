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
	"time"
)

const (
	simpleStoragePath = "/simple-storage/_/"
)

var _ SimpleStorageService = (*cloudSimpleStoreService)(nil)

type SimpleStorageService interface {
	Create(ctx context.Context, s3cr *BucketCreateRequest) (*Bucket, error)
	List(ctx context.Context, opts *ListOptions) ([]*Bucket, error)
	ListWithBucketNameInfo(ctx context.Context, paramBucket ParamListWithBucketNameInfo) (*ResponseListBucketWithName, error)
	Delete(ctx context.Context, id string) error

	UpdateAcl(ctx context.Context, acl, bucketName string) (*ResponseUpdateACL, error)
	UpdateVersioning(ctx context.Context, versioning bool, bucketName string) (*ResponseVersioning, error)
	UpdateCors(ctx context.Context, paramUpdateCors *ParamUpdateCors) (*ResponseCors, error)
	UpdateWebsiteConfig(ctx context.Context, paramUpdateWebsiteConfig *ParamUpdateWebsiteConfig) (*ResponseWebsiteConfig, error)

	SimpleStoreKey() *cloudSimpleStoreKeyService
}
type cloudSimpleStoreService struct {
	client *Client
}

type BucketCreateRequest struct {
	Name                string `json:"name"`
	Location            string `json:"location"`
	Acl                 string `json:"acl"`
	DefaultStorageClass string `json:"default_storage_class"`
}

type Bucket struct {
	Name                string    `json:"name"`
	CreatedAt           time.Time `json:"created_at"`
	Location            string    `json:"location"`
	SizeKb              int       `json:"size_kb"`
	NumObjects          int       `json:"num_objects"`
	DefaultStorageClass string    `json:"default_storage_class"`
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

type Rules struct {
	AllowedOrigin  string   `json:"allowed_origin"`
	AllowedMethods []string `json:"allowed_methods"`
	AllowedHeaders []string `json:"allowed_headers"`
	MaxAgeSeconds  int      `json:"max_age_seconds"`
}

type ParamUpdateCorsRead struct {
	Rules []Rules `json:"rules"`
}

type ParamUpdateCors struct {
	Rules      []Rules `json:"rules"`
	BucketName string  `json:"bucket_name"`
}

type ResponseCors struct {
	Message string `json:"message"`
	Rules   []struct {
		AllowedOrigin  string        `json:"allowed_origin"`
		AllowedMethods []string      `json:"allowed_methods"`
		AllowedHeaders []interface{} `json:"allowed_headers"`
		ExposedHeaders []interface{} `json:"exposed_headers"`
		MaxAgeSeconds  int           `json:"max_age_seconds"`
	} `json:"rules"`
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
		Rules []struct {
			AllowedOrigin  string   `json:"allowed_origin"`
			AllowedMethods []string `json:"allowed_methods"`
			AllowedHeaders []string `json:"allowed_headers"`
			ExposedHeaders []string `json:"exposed_headers"`
			MaxAgeSeconds  int      `json:"max_age_seconds"`
		} `json:"rules"`
	} `json:"cors"`
	Versioning struct {
		Status string `json:"status"`
	} `json:"versioning"`
	WebsiteConfig struct {
		WebsiteUrl string `json:"website_url"`
		Index      string `json:"index"`
		Error      string `json:"error"`
	} `json:"website_config"`
	Quota struct {
		Enabled    bool `json:"enabled"`
		MaxObjects int  `json:"max_objects"`
		MaxSize    int  `json:"max_size"`
		MaxSizeKb  int  `json:"max_size_kb"`
	} `json:"quota"`
	Tags []string `json:"tags"`
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

func (c *cloudSimpleStoreService) resourcePathWithBucketInfo(param ParamListWithBucketNameInfo) string {
	queryParts := getQueryPaths(param)
	query := strings.Join(queryParts, "&")
	bucketName := param.BucketName
	return fmt.Sprintf("%s%s?%s", simpleStoragePath, bucketName, query)
}

func (c *cloudSimpleStoreService) resourcePathUpdateOption(bucketName, path string) string {
	return fmt.Sprintf("%s%s?%s", simpleStoragePath, bucketName, path)
}

func (c *cloudSimpleStoreService) resourcePath() string {
	return simpleStoragePath
}

func (c *cloudSimpleStoreService) itemPath(id string) string {
	return strings.Join([]string{simpleStoragePath, id}, "/")
}

func (c cloudSimpleStoreService) Create(ctx context.Context, bucket *BucketCreateRequest) (*Bucket, error) {
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

func (c *cloudSimpleStoreService) List(ctx context.Context, opts *ListOptions) ([]*Bucket, error) {
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

func (c *cloudSimpleStoreService) ListWithBucketNameInfo(ctx context.Context, paramBucket ParamListWithBucketNameInfo) (*ResponseListBucketWithName, error) {
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

func (c *cloudSimpleStoreService) Delete(ctx context.Context, id string) error {
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

func (c *cloudSimpleStoreService) UpdateAcl(ctx context.Context, acl, bucketName string) (*ResponseUpdateACL, error) {
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

func (c *cloudSimpleStoreService) UpdateVersioning(ctx context.Context, versioning bool, bucketName string) (*ResponseVersioning, error) {
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

	var data *ResponseVersioning

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *cloudSimpleStoreService) UpdateCors(ctx context.Context, paramUpdateCors *ParamUpdateCors) (*ResponseCors, error) {
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

func (c *cloudSimpleStoreService) UpdateWebsiteConfig(ctx context.Context, paramUpdateWebsiteConfig *ParamUpdateWebsiteConfig) (*ResponseWebsiteConfig, error) {
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
