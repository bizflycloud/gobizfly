package gobizfly

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

const (
	simpleStorageBucketPath = "/simple-storage/_/"
)

var _ CloudSimpleStoreBucket = (CloudSimpleStoreBucket)(nil)

type CloudSimpleStoreBucket interface {
	Create(ctx context.Context, s3cr *BucketCreateRequest) (*Bucket, error)
	List(ctx context.Context, opts *ListOptions) ([]*Bucket, error)
	ListWithBucketNameInfo(ctx context.Context, paramBucket ParamListWithBucketNameInfo) (*ResponseListBucketWithName, error)
	Delete(ctx context.Context, id string) error

	UpdateAcl(ctx context.Context, acl, bucketName string) (*ResponseUpdateACL, error)
	UpdateVersioning(ctx context.Context, versioning bool, bucketName string) (*Versioning, error)
	UpdateCors(ctx context.Context, paramUpdateCors *ParamUpdateCors) (*Cors, error)
	UpdateWebsiteConfig(ctx context.Context, paramUpdateWebsiteConfig *ParamUpdateWebsiteConfig) (*ResponseWebsiteConfig, error)

	SimpleStoreKey() *cloudSimpleStoreKeyService
}
type cloudSimpleStoreBucketService struct {
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

type Statement struct {
	User     string   `json:"user"`
	TenantId string   `json:"tenant_id"`
	Effect   string   `json:"effect"`
	Actions  []string `json:"actions"`
}

type Policy struct {
	Statements []Statement `json:"statements"`
}

type Acl struct {
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
	Acl Acl `json:"acl"`
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

type Cors struct {
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
	Bucket Bucket `json:"bucket"`
	Policy struct {
	} `json:"policy"`
	Acl  Acl `json:"acl"`
	Cors struct {
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
	LifecycleConfig struct {
		Rules []struct {
			ID         string `json:"ID"`
			Status     string `json:"Status"`
			Prefix     string `json:"Prefix"`
			Expiration struct {
				Date                      time.Time `json:"Date"`
				Days                      int       `json:"Days"`
				ExpiredObjectDeleteMarker bool      `json:"ExpiredObjectDeleteMarker"`
			} `json:"Expiration"`
			Transitions []struct {
				Date         time.Time `json:"Date"`
				Days         int       `json:"Days"`
				StorageClass string    `json:"StorageClass"`
			} `json:"Transitions"`
			NoncurrentVersionTransitions []struct {
				NoncurrentDays int    `json:"NoncurrentDays"`
				StorageClass   string `json:"StorageClass"`
			} `json:"NoncurrentVersionTransitions"`
			NoncurrentVersionExpiration struct {
				NoncurrentDays int `json:"NoncurrentDays"`
			} `json:"NoncurrentVersionExpiration"`
			AbortIncompleteMultipartUpload struct {
				DaysAfterInitiation int `json:"DaysAfterInitiation"`
			} `json:"AbortIncompleteMultipartUpload"`
		} `json:"Rules"`
	} `json:"lifecycle_config"`
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

type Versioning struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type ParamObject struct {
	BucketName string `json:"bucket_name"`
	ObjectPath string `json:"object_path"`
}

type ResponseCreateObject struct {
	Message string `json:"message"`
	Object  struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}
}

func getQueryPaths(param interface{}) []string {
	// Slice để chứa các phần của query string
	queryParts := []string{}
	// Sử dụng reflect để duyệt qua các trường trong struct
	v := reflect.ValueOf(param)
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		if field.Tag.Get("json") == "bucket_name" {
			continue
		}

		// Chỉ thêm vào query nếu giá trị không rỗng
		if value.Kind() == reflect.String && value.String() != "" {
			queryParts = append(queryParts, fmt.Sprintf("%s=%s", field.Tag.Get("json"), value.String()))
		}
	}
	return queryParts

}

func (c *cloudSimpleStoreBucketService) resourcePathWithBucketInfo(param ParamListWithBucketNameInfo) string {
	// Slice để chứa các phần của query string
	queryParts := getQueryPaths(param)

	// Kết hợp query string
	query := strings.Join(queryParts, "&")
	bucketName := param.BucketName
	// Kết hợp với bucketName để tạo đường dẫn đầy đủ

	return fmt.Sprintf("%s%s?%s", simpleStorageBucketPath, bucketName, query)
}

func (c *cloudSimpleStoreBucketService) resourcePathObject(param ParamObject) string {

	bucketName := param.BucketName
	nameFolder := url.QueryEscape(param.ObjectPath)
	return fmt.Sprintf("%s%s/%s", simpleStorageBucketPath, bucketName, nameFolder)
}

func (c *cloudSimpleStoreBucketService) resourcePathUpdateOption(bucketName, path string) string {
	return fmt.Sprintf("%s%s?%s", simpleStorageBucketPath, bucketName, path)
}

func (c *cloudSimpleStoreBucketService) resourcePath() string {
	return simpleStorageBucketPath
}

func (l *cloudSimpleStoreBucketService) itemPath(id string) string {
	return strings.Join([]string{simpleStorageBucketPath, id}, "/")
}

func (c cloudSimpleStoreBucketService) Create(ctx context.Context, bucket *BucketCreateRequest) (*Bucket, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, simpleStorage, c.resourcePath(), &bucket)
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

func (c *cloudSimpleStoreBucketService) List(ctx context.Context, opts *ListOptions) ([]*Bucket, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, simpleStorage, c.resourcePath(), nil)
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

func (c *cloudSimpleStoreBucketService) ListWithBucketNameInfo(ctx context.Context, paramBucket ParamListWithBucketNameInfo) (*ResponseListBucketWithName, error) {
	if paramBucket.BucketName == "" {
		return nil, errors.New("InvalidBucketName")
	}
	req, err := c.client.NewRequest(ctx, http.MethodGet, simpleStorage, c.resourcePathWithBucketInfo(paramBucket), nil)
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

func (c *cloudSimpleStoreBucketService) Delete(ctx context.Context, id string) error {
	req, err := c.client.NewRequest(ctx, http.MethodDelete, simpleStorage, c.itemPath(id), nil)
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

func (c *cloudSimpleStoreBucketService) UpdateAcl(ctx context.Context, acl, bucketName string) (*ResponseUpdateACL, error) {
	if bucketName == "" {
		return nil, errors.New("InvalidBucketName")
	}
	var paramUpdateAcl = struct {
		Acl string `json:"acl"`
	}{
		Acl: acl,
	}
	req, err := c.client.NewRequest(ctx, http.MethodPatch, simpleStorage, c.resourcePathUpdateOption(bucketName, "acl"), &paramUpdateAcl)
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

func (c *cloudSimpleStoreBucketService) UpdateVersioning(ctx context.Context, versioning bool, bucketName string) (*Versioning, error) {
	if bucketName == "" {
		return nil, errors.New("InvalidBucketName")
	}
	var paramUpdateVersion = struct {
		Versioning bool `json:"versioning"`
	}{
		Versioning: versioning,
	}
	req, err := c.client.NewRequest(ctx, http.MethodPatch, simpleStorage, c.resourcePathUpdateOption(bucketName, "versioning"), &paramUpdateVersion)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data *Versioning

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *cloudSimpleStoreBucketService) UpdateCors(ctx context.Context, paramUpdateCors *ParamUpdateCors) (*Cors, error) {
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

	req, err := c.client.NewRequest(ctx, http.MethodPatch, simpleStorage, c.resourcePathUpdateOption(paramUpdateCors.BucketName, "cors"), &paramCors)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Cors *Cors `json:"cors"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Cors, nil
}

func (c *cloudSimpleStoreBucketService) UpdateWebsiteConfig(ctx context.Context, paramUpdateWebsiteConfig *ParamUpdateWebsiteConfig) (*ResponseWebsiteConfig, error) {
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

	req, err := c.client.NewRequest(ctx, http.MethodPatch, simpleStorage, c.resourcePathUpdateOption(paramUpdateWebsiteConfig.BucketName, "website_config"), &paramWeb)
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
