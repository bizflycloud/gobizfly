// This file is part of gobizfly

package gobizfly

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	registryPath = "/_"
)

type containerRegistry struct {
	client *Client
}

var _ ContainerRegistryService = (*containerRegistry)(nil)

type ContainerRegistryService interface {
	List(ctx context.Context, opts *ListOptions) ([]*Repository, error)
	Create(ctx context.Context, crpl *CreateRepositoryPayload) error
	Delete(ctx context.Context, repositoryName string) error
	GetTags(ctx context.Context, repositoryName string) (*TagRepository, error)
	EditRepo(ctx context.Context, repositoryName string, erpl *EditRepositoryPayload) error
	DeleteTag(ctx context.Context, tagName string, repositoryName string) error
	GetTag(ctx context.Context, repositoryName string, tagName string, vulnerabilities string) (*Image, error)
	GenerateToken(ctx context.Context, gtpl *GenerateTokenPayload) (*TokenResp, error)
}

// Repository represents a container registry repository information
type Repository struct {
	Name      string `json:"name"`
	LastPush  string `json:"last_push"`
	Pulls     int    `json:"pulls"`
	Public    bool   `json:"public"`
	CreatedAt string `json:"created_at"`
}

// Repositories represents a list of container registry repositories
type Repositories struct {
	Repositories []Repository `json:"repositories"`
}

// CreateRepositoryPayload represents the payload for creating a container registry repository
type CreateRepositoryPayload struct {
	Name   string `json:"name"`
	Public bool   `json:"public"`
}

// RepositoryTag represents a container registry tag information
type RepositoryTag struct {
	Name            string `json:"name"`
	Author          string `json:"author"`
	LastUpdated     string `json:"last_updated"`
	CreatedAt       string `json:"created_at"`
	LastScan        string `json:"last_scan"`
	ScanStatus      string `json:"scan_status"`
	Vulnerabilities int    `json:"vulnerabilities"`
	Fixes           int    `json:"fixes"`
}

// EditRepositoryPayload represents the payload for updating a container registry repository
type EditRepositoryPayload struct {
	Public bool `json:"public"`
}

// Vulnerability represents a container registry vulnerability information
type Vulnerability struct {
	Package     string `json:"package"`
	Name        string `json:"name"`
	Namespace   string `json:"namespace"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Severity    string `json:"severity"`
	FixedBy     string `json:"fixed_by"`
}

// Image represents container registry vulnerability information
type Image struct {
	Repository      Repository      `json:"repository"`
	Tag             RepositoryTag   `json:"tag"`
	Vulnerabilities []Vulnerability `json:"vulnerabilities"`
}

// TagRepository represents a container registry repository with additional tag information
type TagRepository struct {
	Repository Repository      `json:"repository"`
	Tags       []RepositoryTag `json:"tags"`
}

// GenerateTokenPayload represents the payload for generating a container registry token
type GenerateTokenPayload struct {
	ExpiresIn int     `json:"expires_in"`
	Scopes    []Scope `json:"scopes"`
}

// Scope represents a container registry token scope information
type Scope struct {
	Action     []string `json:"actions"`
	Repository string   `json:"repository"`
}

// TokenResp represents a container registry token information
type TokenResp struct {
	Token     string `json:"token"`
	ExpiresIn int    `json:"expires_in"`
	IssuedAt  string `json:"issued_at"`
}

func (c *containerRegistry) resourcePath() string {
	return registryPath + "/"
}

func (c *containerRegistry) itemPath(id string) string {
	return strings.Join([]string{registryPath, id}, "/")
}

// List - List container registry repositories
func (c *containerRegistry) List(ctx context.Context, opts *ListOptions) ([]*Repository, error) {
	req, err := c.client.NewRequest(ctx, http.MethodGet, containerRegistryName, c.resourcePath(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	var data struct {
		Repositories []*Repository `json:"repositories"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Repositories, nil
}

// Create - Create a container registry repository
func (c *containerRegistry) Create(ctx context.Context, crpl *CreateRepositoryPayload) error {
	req, err := c.client.NewRequest(ctx, http.MethodPost, containerRegistryName, c.resourcePath(), &crpl)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}

// Delete - Delete a container registry repository
func (c *containerRegistry) Delete(ctx context.Context, repositoryName string) error {
	req, err := c.client.NewRequest(ctx, http.MethodDelete, containerRegistryName, c.itemPath(repositoryName), nil)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(io.Discard, resp.Body)
	return resp.Body.Close()
}

// GetTags - Get tags of a container registry repository
func (c *containerRegistry) GetTags(ctx context.Context, repositoryName string) (*TagRepository, error) {
	var data *TagRepository
	req, err := c.client.NewRequest(ctx, http.MethodGet, containerRegistryName, strings.Join([]string{registryPath, repositoryName}, "/"), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// EditRepo - Edit the visibility of a container registry repository
func (c *containerRegistry) EditRepo(ctx context.Context, repositoryName string, erpl *EditRepositoryPayload) error {
	req, err := c.client.NewRequest(ctx, http.MethodPatch, containerRegistryName, strings.Join([]string{registryPath, repositoryName}, "/"), erpl)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}

// DeleteTag - Delete a tag of the container registry repository
func (c *containerRegistry) DeleteTag(ctx context.Context, repositoryName string, tagName string) error {
	req, err := c.client.NewRequest(ctx, http.MethodDelete, containerRegistryName, strings.Join([]string{registryPath, repositoryName, "tag", tagName}, "/"), nil)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}

// GetTag - Get tag of the container registry repository
func (c *containerRegistry) GetTag(ctx context.Context, repositoryName string, tagName string, vulnerabilities string) (*Image, error) {
	var data *Image
	u, _ := url.Parse(strings.Join([]string{registryPath, repositoryName, "tag", tagName}, "/"))
	if vulnerabilities != "" {
		query := url.Values{
			"vulnerabilities": {vulnerabilities},
		}
		u.RawQuery = query.Encode()

	}
	req, err := c.client.NewRequest(ctx, http.MethodGet, containerRegistryName, u.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// GenerateToken - Generate token for container registry usages
func (c *containerRegistry) GenerateToken(ctx context.Context, gtpl *GenerateTokenPayload) (*TokenResp, error) {
	req, err := c.client.NewRequest(ctx, http.MethodPost, containerRegistryName, tokenPath, gtpl)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	var data *TokenResp
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}
