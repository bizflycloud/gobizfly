// This file is part of gobizfly

package gobizfly

import (
	"context"
	"strings"
)

const (
	fileStorageSharesPath       = "/_"
	fileStorageResizePath       = "/resize"
	fileStorageAccessPath       = "/access"
	fileStorageAccessStatusPath = "/access/status"
	fileStorageQuotaPath        = "/manage/quotas"
	fileStorageRegionsPath      = "/regions"
)

var _ FileStorageService = (*fileStorageService)(nil)

type fileStorageService struct {
	client *Client
}

// FileStorageService is an interface to interact with Bizfly Cloud File Storage API.
type FileStorageService interface {
	// Share operations
	List(ctx context.Context) ([]*Share, error)
	Create(ctx context.Context, req *CreateShareRequest) (*Share, error)
	Get(ctx context.Context, shareID string) (*Share, error)
	Delete(ctx context.Context, shareID string, force bool) error
	Resize(ctx context.Context, shareID string, req *ResizeShareRequest) (*Share, error)

	// Access rule operations
	GetAccessRules(ctx context.Context, shareID string) ([]*FileStorageAccessRule, error)
	ManageAccessRules(ctx context.Context, shareID string, req *ManageAccessRequest) ([]*FileStorageAccessRule, error)
	DeleteAccessRule(ctx context.Context, shareID string, ruleID string) error
	GetAccessStatus(ctx context.Context, shareID string) (*FileStorageAccessStatus, error)

	// Quota operations (admin)
	GetQuota(ctx context.Context, projectID string) (*QuotaResponse, error)
	UpdateQuota(ctx context.Context, req *QuotaRequest) (*QuotaResponse, error)

	// Regions
	ListRegions(ctx context.Context) ([]*FileStorageRegion, error)
}

func (fs *fileStorageService) resourcePath() string {
	return fileStorageSharesPath
}

func (fs *fileStorageService) shareItemPath(id string) string {
	return strings.Join([]string{fileStorageSharesPath, id}, "/")
}

func (fs *fileStorageService) shareResizePath(id string) string {
	return strings.Join([]string{fileStorageSharesPath, id, "resize"}, "/")
}

func (fs *fileStorageService) shareAccessPath(id string) string {
	return strings.Join([]string{fileStorageSharesPath, id, "access"}, "/")
}

func (fs *fileStorageService) shareAccessRulePath(id string, ruleID string) string {
	return strings.Join([]string{fileStorageSharesPath, id, "access", ruleID}, "/")
}

func (fs *fileStorageService) shareAccessStatusPath(id string) string {
	return strings.Join([]string{fileStorageSharesPath, id, "access", "status"}, "/")
}

func (fs *fileStorageService) quotaPath() string {
	return fileStorageQuotaPath
}

func (fs *fileStorageService) quotaProjectPath(projectID string) string {
	return strings.Join([]string{fileStorageQuotaPath, projectID}, "/")
}

func (fs *fileStorageService) regionsPath() string {
	return fileStorageRegionsPath
}
