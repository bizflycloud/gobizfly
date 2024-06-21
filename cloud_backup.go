package gobizfly

import (
	"context"
	"strings"
)

const (
	dashboardPath = "/dashboard"
)

var _ CloudBackupService = (*cloudBackupService)(nil)

type cloudBackupService struct {
	client *Client
}

type CloudBackupService interface {
	CloudBackupListActivities(ctx context.Context) ([]*CloudBackupActivity, error)

	ListTenantRecoveryPoints(ctx context.Context) ([]*CloudBackupMachineRecoveryPoint, error)
	DeleteMultipleRecoveryPoints(ctx context.Context, payload CloudBackupDeleteMultipleRecoveryPointPayload) error
	ListDirectoryRecoveryPoints(ctx context.Context, machineID string, directoryID string) ([]*CloudBackupMachineRecoveryPoint, error)
	RecoveryPointAction(ctx context.Context, recoveryPointID string, payload *CloudBackupRecoveryPointActionPayload) (*CloudBackupMachineRecoveryPoint, error)
	ListMachineRecoveryPoints(ctx context.Context, machineID string) ([]*CloudBackupExtendedRecoveryPoint, error)
	GetRecoveryPoint(ctx context.Context, recoveryPointID string) (*CloudBackupMachineRecoveryPoint, error)
	DeleteRecoveryPoint(ctx context.Context, recoveryPointID string) error
	ListRecoveryPointItems(ctx context.Context, recoveryPointID string) ([]*CloudBackupRecoveryPointItem, error)
	RestoreRecoveryPoint(ctx context.Context, recoveryPointID string, payload *CloudBackupRestoreRecoveryPointPayload) error

	ListStorageVaults(ctx context.Context) ([]*CloudBackupStorageVault, error)
	GetStorageVault(ctx context.Context, vaultID string) (*CloudBackupStorageVault, error)
	CreateStorageVault(ctx context.Context, payload *CloudBackupCreateStorageVaultPayload) (*CloudBackupStorageVault, error)

	ListTenantMachines(ctx context.Context, listOption *CloudBackupListMachineParams) ([]*CloudBackupMachine, error)
	CreateMachine(ctx context.Context, payload *CloudBackupCreateMachinePayload) (*CloudBackupExtendedMachine, error)
	GetMachine(ctx context.Context, machineID string) (*CloudBackupMachine, error)
	PatchMachine(ctx context.Context, machineID string, payload *CloudBackupPatchMachinePayload) (*CloudBackupMachine, error)
	DeleteMachine(ctx context.Context, machineID string, payload *CloudBackupDeleteMachinePayload) error
	ActionMachine(ctx context.Context, machineID string, payload *CloudBackupActionMachinePayload) error
	ResetMachineSecretKey(ctx context.Context, machineID string) (*CloudBackupExtendedMachine, error)

	ActionDirectory(ctx context.Context, machineID string, payload *CloudBackupStateDirectoryAction) error
	ListMachineBackupDirectories(ctx context.Context, machineID string) ([]*CloudBackupDirectory, error)
	CreateBackupDirectory(ctx context.Context, machineID string, payload *CloudBackupCreateDirectoryPayload) (*CloudBackupDirectory, error)
	GetBackupDirectory(ctx context.Context, machineID string, directoryID string) (*CloudBackupDirectory, error)
	PatchBackupDirectory(ctx context.Context, machineID string, directoryID string, payload *CloudBackupPatchDirectoryPayload) (*CloudBackupDirectory, error)
	DeleteBackupDirectory(ctx context.Context, machineID string, directoryID string, payload *CloudBackupDeleteDirectoryPayload) error
	ListTenantDirectories(ctx context.Context) ([]*CloudBackupDirectory, error)
	ActionBackupDirectory(ctx context.Context, machineID string, directoryID string, payload *CloudBackupActionDirectoryPayload) error
	DeleteMultipleDirectories(ctx context.Context, machineID string, payload *CloudBackupDeleteMultipleDirectoriesPayload) error
	ActionMultipleDirectories(ctx context.Context, machineID string, payload *CloudBackupActionMultipleDirectoriesPayload) error

	ListTenantPolicies(ctx context.Context) ([]*CloudBackupPolicy, error)
	CreatePolicy(ctx context.Context, payload *CloudBackupCreatePolicyPayload) (*CloudBackupPolicy, error)
	GetBackupDirectoryPolicy(ctx context.Context, machineID string, directoryID string) (*CloudBackupPolicy, error)
	GetPolicy(ctx context.Context, policyID string) (*CloudBackupPolicy, error)
	PatchPolicy(ctx context.Context, policyID string, payload *CloudBackupPatchPolicyPayload) (*CloudBackupPolicy, error)
	DeletePolicy(ctx context.Context, policyID string) error
	ListAppliedPolicyDirectories(ctx context.Context, policyID string) ([]*CloudBackupDirectory, error)
	ActionPolicyDirectory(ctx context.Context, policyID string, payload *CloudBackupActionPolicyDirectoryPayload) error
}

func (cb *cloudBackupService) dashboardPath() string {
	return dashboardPath
}

func (cb *cloudBackupService) machinesPath() string {
	return strings.Join([]string{cb.dashboardPath(), "machines"}, "/")
}

func (cb *cloudBackupService) itemMachinePath(id string) string {
	return strings.Join([]string{cb.machinesPath(), id}, "/")
}

func (cb *cloudBackupService) machineDirectoryPath(machineID string, directoryID string) string {
	return strings.Join([]string{cb.machinesPath(), machineID, "directories", directoryID}, "/")
}

func (cb *cloudBackupService) machineDirectoryRecoveryPointPath(machineID string, directoryID string) string {
	return strings.Join([]string{cb.machineDirectoryPath(machineID, directoryID), "recovery-points"}, "/")
}

func (cb *cloudBackupService) activityPath() string {
	return strings.Join([]string{cb.dashboardPath(), "activity"}, "/")
}

func (cb *cloudBackupService) directoryPath() string {
	return strings.Join([]string{cb.dashboardPath(), "backup-directories"}, "/")
}

func (cb *cloudBackupService) policyPath() string {
	return strings.Join([]string{cb.dashboardPath(), "policies"}, "/")
}

func (cb *cloudBackupService) itemPolicyPath(policyID string) string {
	return strings.Join([]string{cb.policyPath(), policyID}, "/")
}

func (cb *cloudBackupService) recoveryPointPath() string {
	return strings.Join([]string{cb.dashboardPath(), "recovery-points"}, "/")
}

func (cb *cloudBackupService) itemRecoveryPointPath(recoveryPointID string) string {
	return strings.Join([]string{cb.recoveryPointPath(), recoveryPointID}, "/")
}

func (cb *cloudBackupService) storageVaultsPath() string {
	return strings.Join([]string{cb.dashboardPath(), "storage-vaults"}, "/")
}

func (cb *cloudBackupService) itemStorageVaultPath(storageVaultID string) string {
	return strings.Join([]string{cb.storageVaultsPath(), storageVaultID}, "/")
}
