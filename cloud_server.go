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

var _ CloudServerService = (*cloudServerService)(nil)

type cloudServerService struct {
	client *Client
}

// CloudServerService is an interface to interact with Bizfly Cloud API.
type CloudServerService interface {
	AddVirtualPrivateNetwork(ctx context.Context, id string, vpcs []string) (*Server, error)
	AttachPublicNetworkInterface(ctx context.Context, id string, wanIps []string) error
	ChangeCategory(ctx context.Context, id string, newCategory string) (*ServerTask, error)
	ChangeNetworkPlan(ctx context.Context, id string, newNetworkPlan string) error
	Create(ctx context.Context, scr *ServerCreateRequest) (*ServerCreateResponse, error)
	Delete(ctx context.Context, id string, deletedRootDisk []string) (*ServerTask, error)
	EnableIPv6(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*Server, error)
	GetTask(ctx context.Context, id string) (*ServerTaskResponse, error)
	GetVNC(ctx context.Context, id string) (*ServerConsoleResponse, error)
	HardReboot(ctx context.Context, id string) (*ServerMessageResponse, error)
	List(ctx context.Context, opts *ServerListOptions) ([]*Server, error)
	ListServerTypes(ctx context.Context) ([]*ServerType, error)
	Rebuild(ctx context.Context, id string, imageID string) (*ServerTask, error)
	RemoveNetworkInterface(ctx context.Context, id string, vpcs []string) (*Server, error)
	Rename(ctx context.Context, id string, newName string) error
	Resize(ctx context.Context, id string, newFlavor string) (*ServerTask, error)
	SoftReboot(ctx context.Context, id string) (*ServerMessageResponse, error)
	Start(ctx context.Context, id string) (*Server, error)
	Stop(ctx context.Context, id string) (*Server, error)
	SwitchBillingPlan(ctx context.Context, id string, newBillingPlan string) error

	CustomImages() *cloudServerCustomOSImageResource
	Firewalls() *cloudServerFirewallResource
	Flavors() *cloudServerFlavorResource
	NetworkInterfaces() *cloudServerNetworkInterfaceResource
	OSImages() *cloudServerOSImageResource
	PublicNetworkInterfaces() *cloudServerPublicNetworkInterfaceResource
	ScheduledVolumeBackups() *cloudServerScheduledVolumeBackupResource
	Snapshots() *cloudServerSnapshotResource
	SSHKeys() *cloudServerSSHKeyResource
	Volumes() *cloudServerVolumeResource
	VPCNetworks() *cloudServerVPCNetworkResource
	InternetGateways() CloudServerInternetGatewayInterface
}
