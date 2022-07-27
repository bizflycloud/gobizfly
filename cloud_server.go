// This file is part of gobizfly

package gobizfly

import (
	"context"
)

const (
	serverBasePath   = "/servers"
	flavorPath       = "/flavors"
	osImagePath      = "/images"
	taskPath         = "/tasks"
	customImagesPath = "/user/images"
	customImagePath  = "/user/image"
)

var _ ServerService = (*server)(nil)

// ServerService is an interface to interact with BizFly Cloud API.
type ServerService interface {
	List(ctx context.Context, opts *ServerListOptions) ([]*Server, error)
	Create(ctx context.Context, scr *ServerCreateRequest) (*ServerCreateResponse, error)
	Get(ctx context.Context, id string) (*Server, error)
	Delete(ctx context.Context, id string, deletedRootDisk []string) (*ServerTask, error)
	Resize(ctx context.Context, id string, newFlavor string) (*ServerTask, error)
	Start(ctx context.Context, id string) (*Server, error)
	Stop(ctx context.Context, id string) (*Server, error)
	SoftReboot(ctx context.Context, id string) (*ServerMessageResponse, error)
	HardReboot(ctx context.Context, id string) (*ServerMessageResponse, error)
	Rebuild(ctx context.Context, id string, imageID string) (*ServerTask, error)
	GetVNC(ctx context.Context, id string) (*ServerConsoleResponse, error)
	ListFlavors(ctx context.Context) ([]*serverFlavorResponse, error)
	ListOSImages(ctx context.Context) ([]osImageResponse, error)
	GetTask(ctx context.Context, id string) (*ServerTaskResponse, error)
	ChangeCategory(ctx context.Context, id string, newCategory string) (*ServerTask, error)
	AddVPC(ctx context.Context, id string, vpcs []string) (*Server, error)
	RemoveVPC(ctx context.Context, id string, vpcs []string) (*Server, error)
	ListCustomImages(ctx context.Context) ([]*CustomImage, error)
	CreateCustomImage(ctx context.Context, cipl *CreateCustomImagePayload) (*CreateCustomImageResp, error)
	DeleteCustomImage(ctx context.Context, imageID string) error
	GetCustomImage(ctx context.Context, imageID string) (*CustomImageGetResp, error)
	AttachWanIps(ctx context.Context, id string, wanIps []string) error
	ListServerTypes(ctx context.Context) ([]*ServerType, error)
}
