// This file is part of gobizfly

package gobizfly

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/bizflycloud/gobizfly/utils"
)

const (
	appCredentialAuthType    = "application_credential"
	accountName              = "bizfly_account"
	authServiceName          = "auth"
	autoScalingServiceName   = "auto_scaling"
	cdnName                  = "cdn"
	cloudBackupServiceName   = "cloud-backup"
	cloudwatcherServiceName  = "alert"
	containerRegistryName    = "container_registry"
	databaseServiceName      = "cloud_database"
	defaultAPIURL            = "https://manage.bizflycloud.vn/api"
	defaultAuthType          = "token"
	dnsName                  = "dns"
	iamServiceName           = "iam"
	kubernetesServiceName    = "kubernetes_engine"
	loadBalancerServiceName  = "load_balancer"
	simpleStorageServiceName = "simple_storage"
	kmsServiceName           = "key_management_service"
	mediaType                = "application/json; charset=utf-8"
	serverServiceName        = "cloud_server"
	ua                       = "bizfly-client-go/" + version
	version                  = "0.0.1"
)

var (
	// ErrNotFound for resource not found status
	ErrNotFound = errors.New("resource not found")
	// ErrPermissionDenied for permission denied
	ErrPermissionDenied = errors.New("you are not allowed to do this action")
	// ErrCommon for common error
	ErrCommon = errors.New("error")
)

// Client represents Bizfly API client.
type Client struct {
	appCredID     string
	appCredSecret string
	authMethod    string
	authType      string
	basicAuth     string
	keystoneToken string
	password      string
	projectID     string
	regionName    string
	userAgent     string
	username      string

	apiURL     *url.URL
	httpClient *http.Client
	services   []*Service

	Account            AccountService
	AutoScaling        AutoScalingService
	CDN                CDNService
	CloudBackup        CloudBackupService
	CloudDatabase      CloudDatabaseService
	CloudLoadBalancer  LoadBalancerService
	CloudSimpleStorage SimpleStorageService
	CloudServer        CloudServerService
	CloudWatcher       CloudWatcherService
	ContainerRegistry  ContainerRegistryService
	DNS                DNSService
	IAM                IAMService
	KubernetesEngine   KubernetesEngineService
	Service            ServiceInterface
	Token              TokenService
	KMS                KMSService
}

// Option set Client specific attributes
type Option func(c *Client) error

// WithAPIURL sets the API url option for Bizfly client.
func WithAPIURL(u string) Option {
	return func(c *Client) error {
		var err error
		c.apiURL, err = url.Parse(u)
		return err
	}
}

// WithHTTPClient sets the client option for Bizfly client.
func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) error {
		if client == nil {
			return errors.New("client is nil")
		}

		c.httpClient = client

		return nil
	}
}

// WithRegionName sets the client region for Bizfly client.
func WithRegionName(region string) Option {
	return func(c *Client) error {
		regionName, err := utils.ParseRegionName(region)
		if err != nil {
			return err
		}
		c.regionName = regionName
		return nil
	}
}

func WithProjectID(id string) Option {
	return func(c *Client) error {
		c.projectID = id
		return nil
	}
}

func WithBasicAuth(basicAuth string) Option {
	return func(c *Client) error {
		c.basicAuth = basicAuth
		return nil
	}
}

// NewClient creates new Bizfly client.
func NewClient(options ...Option) (*Client, error) {
	c := &Client{
		httpClient: http.DefaultClient,
		userAgent:  ua,
	}

	err := WithAPIURL(defaultAPIURL)(c)
	if err != nil {
		return nil, err
	}

	for _, option := range options {
		if err := option(c); err != nil {
			return nil, err
		}
	}

	c.Account = &accountService{client: c}
	c.AutoScaling = &autoscalingService{client: c}
	c.CDN = &cdnService{client: c}
	c.CloudBackup = &cloudBackupService{client: c}
	c.CloudDatabase = &cloudDatabaseService{client: c}
	c.CloudServer = &cloudServerService{client: c}
	c.CloudWatcher = &cloudwatcherService{client: c}
	c.ContainerRegistry = &containerRegistry{client: c}
	c.DNS = &dnsService{client: c}
	c.IAM = &iamService{client: c}
	c.KubernetesEngine = &kubernetesEngineService{client: c}
	c.CloudLoadBalancer = &cloudLoadBalancerService{client: c}
	c.CloudSimpleStorage = &cloudSimpleStorageService{client: c}
	c.Service = &service{client: c}
	c.Token = &token{client: c}
	c.KMS = &kmsService{client: c}
	return c, nil
}

func (c *Client) GetServiceURL(serviceName string) string {
	// If service name is auth, return auth url without checking catalog
	if serviceName == authServiceName {
		// create a copy of apiURL
		apiURL := *c.apiURL
		// if apiURL doesn't end with /api, append it
		if !strings.HasSuffix(apiURL.Path, "/api") {
			apiURL.Path = path.Join(c.apiURL.Path, "/api")
		}
		return apiURL.String()
	}
	for _, service := range c.services {
		if service.CanonicalName == serviceName && service.Region == c.regionName {
			return service.ServiceURL
		}
	}
	return defaultAPIURL
}

// NewRequest creates an API request.
func (c *Client) NewRequest(ctx context.Context, method, serviceName string, urlStr string, body interface{}) (*http.Request, error) {
	serviceURL := c.GetServiceURL(serviceName)
	url := serviceURL + urlStr
	buf := new(bytes.Buffer)
	if body != nil {
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("Accept", mediaType)
	req.Header.Add("User-Agent", c.userAgent)
	req.Header.Add("X-Project-ID", c.projectID)
	req.Header.Add("Authorization", "Basic "+c.basicAuth)

	if c.authType == "" {
		c.authType = defaultAuthType
	}

	if c.keystoneToken != "" {
		req.Header.Add("X-Auth-Token", c.keystoneToken)
	}

	req.Header.Add("X-Auth-Type", c.authType)
	if c.authType == appCredentialAuthType {
		req.Header.Add("X-App-Credential-ID", c.appCredID)
		req.Header.Add("X-App-Credential-Secret", c.appCredSecret)
	}
	return req, nil
}

func (c *Client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	return c.httpClient.Do(req)
}

func (c *Client) DoInit(ctx context.Context, req *http.Request) (resp *http.Response, err error) {

	resp, err = c.do(ctx, req)
	if err != nil {
		return
	}

	if resp.StatusCode >= http.StatusBadRequest {
		defer func() {
			_ = resp.Body.Close()
		}()
		buf, _ := io.ReadAll(resp.Body)
		err = errorFromStatus(resp.StatusCode, string(buf))

	}
	return
}

// Do sends API request.
func (c *Client) Do(ctx context.Context, req *http.Request) (resp *http.Response, err error) {

	resp, err = c.do(ctx, req)
	if err != nil {
		return
	}

	// If 401, get new token and retry one time.
	if resp.StatusCode == http.StatusUnauthorized {
		tok, tokErr := c.Token.Refresh(ctx)
		if tokErr != nil {
			buf, _ := io.ReadAll(resp.Body)
			err = fmt.Errorf("%s : %w", string(buf), tokErr)
			return
		}
		c.SetKeystoneToken(tok)
		req.Header.Set("X-Auth-Token", c.keystoneToken)
		resp, err = c.do(ctx, req)
	}
	if resp.StatusCode >= http.StatusBadRequest {
		defer func() {
			_ = resp.Body.Close()
		}()
		buf, _ := io.ReadAll(resp.Body)
		err = errorFromStatus(resp.StatusCode, string(buf))
	}
	return
}

// SetKeystoneToken sets keystone token value, which will be used for authentication.
func (c *Client) SetKeystoneToken(token *Token) {
	c.keystoneToken = token.KeystoneToken
	c.projectID = token.ProjectID
}

// ListOptions specifies the optional parameters for List method.
type ListOptions struct {
	Page  int `json:"page,omitempty"`
	Limit int `json:"limit,omitempty"`
}

func errorFromStatus(code int, msg string) error {
	switch code {
	case http.StatusNotFound:
		return fmt.Errorf("%s: %w", msg, ErrNotFound)
	case http.StatusForbidden:
		return fmt.Errorf("%s: %w", msg, ErrPermissionDenied)
	default:
		return fmt.Errorf("%s: %w", msg, ErrCommon)
	}
}
