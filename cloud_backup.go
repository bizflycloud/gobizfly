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
	CloudBackupActivity(ctx context.Context) ([]*CloudBackupActivity, error)

	ListTenantRecoveryPoints(ctx context.Context) ([]*CloudBackupMachineRecoveryPoint, error)
	DeleteMultipleRecoveryPoints(ctx context.Context, payload CloudBackupDeleteMultipleRecoveryPointPayload) error
	ListDirectoryRecoveryPoints(ctx context.Context, machineId string, directoryId string) ([]*CloudBackupMachineRecoveryPoint, error)
	RecoveryPointAction(ctx context.Context, recoveryPointId string, payload *CloudBackupRecoveryPointActionPayload) (*CloudBackupMachineRecoveryPoint, error)
	ListMachineRecoveryPoints(ctx context.Context, machineId string) ([]*CloudBackupExtendedRecoveryPoint, error)
	GetRecoveryPoint(ctx context.Context, recoveryPointId string) (*CloudBackupMachineRecoveryPoint, error)
	DeleteRecoveryPoint(ctx context.Context, recoveryPointId string) error
	ListRecoveryPointItems(ctx context.Context, recoveryPointId string) ([]*CloudBackupRecoveryPointItem, error)
	RestoreRecoveryPoint(ctx context.Context, recoveryPointId string, payload *CloudBackupRestoreRecoveryPointPayload) error

	ListStorageVaults(ctx context.Context) ([]*CloudBackupStorageVault, error)
	GetStorageVault(ctx context.Context, vaultId string) (*CloudBackupStorageVault, error)
	CreateStorageVault(ctx context.Context, payload *CloudBackupCreateStorageVaultPayload) (*CloudBackupStorageVault, error)

	ListTenantMachines(ctx context.Context, listOption *CloudBackupListMachineParams) ([]*CloudBackupMachine, error)
	CreateMachine(ctx context.Context, payload *CloudBackupCreateMachinePayload) (*CloudBackupExtendedMachine, error)
	GetMachine(ctx context.Context, machineId string) (*CloudBackupMachine, error)
	PatchMachine(ctx context.Context, machineId string, payload *CloudBackupPatchMachinePayload) (*CloudBackupMachine, error)
	DeleteMachine(ctx context.Context, machineId string, payload *CloudBackupDeleteMachinePayload) error
	ActionMachine(ctx context.Context, machineId string, payload *CloudBackupActionMachinePayload) error
	ResetMachineSecretKey(ctx context.Context, machineId string) (*CloudBackupExtendedMachine, error)

	ActionDirectory(ctx context.Context, machineId string, payload *CloudBackupStateDirectoryAction) error
	ListMachineBackupDirectories(ctx context.Context, machineId string) ([]*CloudBackupDirectory, error)
	CreateBackupDirectory(ctx context.Context, machineId string, payload *CloudBackupCreateDirectoryPayload) (*CloudBackupDirectory, error)
	GetBackupDirectory(ctx context.Context, machineId string, directoryId string) (*CloudBackupDirectory, error)
	PatchBackupDirectory(ctx context.Context, machineId string, directoryId string, payload *CloudBackupPatchDirectoryPayload) (*CloudBackupDirectory, error)
	DeleteBackupDirectory(ctx context.Context, machineId string, directoryId string, payload *CloudBackupDeleteDirectoryPayload) error
	ListTenantDirectories(ctx context.Context) ([]*CloudBackupDirectory, error)
	ActionBackupDirectory(ctx context.Context, machineId string, directoryId string, payload *CloudBackupActionDirectoryPayload) error
	DeleteMultipleDirectories(ctx context.Context, machineId string, payload *CloudBackupDeleteMultipleDirectoriesPayload) error
	ActionMultipleDirectories(ctx context.Context, machineId string, payload *CloudBackupActionMultipleDirectoriesPayload) error

	ListTenantPolicies(ctx context.Context) ([]*CloudBackupPolicy, error)
	CreatePolicy(ctx context.Context, payload *CloudBackupCreatePolicyPayload) (*CloudBackupPolicy, error)
	GetBackupDirectoryPolicy(ctx context.Context, machineId string, directoryId string) (*CloudBackupPolicy, error)
	GetPolicy(ctx context.Context, policyId string) (*CloudBackupPolicy, error)
	PatchPolicy(ctx context.Context, policyId string, payload *CloudBackupPatchPolicyPayload) (*CloudBackupPolicy, error)
	DeletePolicy(ctx context.Context, policyId string) error
	ListAppliedPolicyDirectories(ctx context.Context, policyId string) ([]*CloudBackupDirectory, error)
	ActionPolicyDirectory(ctx context.Context, policyId string, payload *CloudBackupActionPolicyDirectoryPayload) error
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

func (cb *cloudBackupService) machineDirectoryPath(machineId string, directoryId string) string {
	return strings.Join([]string{cb.machinesPath(), machineId, "directories", directoryId}, "/")
}

func (cb *cloudBackupService) machineDirectoryRecoveryPointPath(machineId string, directoryId string) string {
	return strings.Join([]string{cb.machineDirectoryPath(machineId, directoryId), "recovery-points"}, "/")
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

func (cb *cloudBackupService) itemPolicyPath(policyId string) string {
	return strings.Join([]string{cb.policyPath(), policyId}, "/")
}

func (cb *cloudBackupService) recoveryPointPath() string {
	return strings.Join([]string{cb.dashboardPath(), "recovery-points"}, "/")
}

func (cb *cloudBackupService) itemRecoveryPointPath(recoveryPointId string) string {
	return strings.Join([]string{cb.recoveryPointPath(), recoveryPointId}, "/")
}

func (cb *cloudBackupService) storageVaultsPath() string {
	return strings.Join([]string{cb.dashboardPath(), "storage-vaults"}, "/")
}

func (cb *cloudBackupService) itemStorageVaultPath(storageVaultId string) string {
	return strings.Join([]string{cb.storageVaultsPath(), storageVaultId}, "/")
}
