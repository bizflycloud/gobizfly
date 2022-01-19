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
	CloudBackupActivity(ctx context.Context) ([]*Activity, error)

	ListTenantRecoveryPoints(ctx context.Context) ([]*MachineRecoveryPoint, error)
	DeleteMultipleRecoveryPoints(ctx context.Context, payload DeleteMultipleRecoveryPointPayload) error
	ListDirectoryRecoveryPoints(ctx context.Context, machineId string, directoryId string) ([]*MachineRecoveryPoint, error)
	RecoveryPointAction(ctx context.Context, recoveryPointId string, payload *RecoveryPointActionPayload) (*MachineRecoveryPoint, error)
	ListMachineRecoveryPoints(ctx context.Context, machineId string) ([]*ExtendedRecoveryPoint, error)
	GetRecoveryPoint(ctx context.Context, recoveryPointId string) (*MachineRecoveryPoint, error)
	DeleteRecoveryPoint(ctx context.Context, recoveryPointId string) error
	ListRecoveryPointItems(ctx context.Context, recoveryPointId string) ([]*RecoveryPointItem, error)
	RestoreRecoveryPoint(ctx context.Context, recoveryPointId string, payload *RestoreRecoveryPointPayload) error

	ListStorageVaults(ctx context.Context) ([]*StorageVault, error)
	GetStorageVault(ctx context.Context, vaultId string) (*StorageVault, error)
	CreateStorageVault(ctx context.Context, payload *CreateStorageVaultPayload) (*StorageVault, error)

	ListTenantMachines(ctx context.Context, listOption *ListMachineParams) ([]*Machine, error)
	CreateMachine(ctx context.Context, payload *CreateMachinePayload) (*ExtendedMachine, error)
	GetMachine(ctx context.Context, machineId string) (*Machine, error)
	PatchMachine(ctx context.Context, machineId string, payload *PatchMachinePayload) (*Machine, error)
	DeleteMachine(ctx context.Context, machineId string, payload *DeleteMachinePayload) error
	ActionMachine(ctx context.Context, machineId string, payload *ActionMachinePayload) error
	ResetMachineSecretKey(ctx context.Context, machineId string) (*ExtendedMachine, error)

	ActionDirectory(ctx context.Context, machineId string, payload *StateDirectoryAction) error
	ListMachineBackupDirectories(ctx context.Context, machineId string) ([]*BackupDirectory, error)
	CreateBackupDirectory(ctx context.Context, machineId string, payload *CreateDirectoryPayload) (*BackupDirectory, error)
	GetBackupDirectory(ctx context.Context, machineId string, directoryId string) (*BackupDirectory, error)
	PatchBackupDirectory(ctx context.Context, machineId string, directoryId string, payload *PatchDirectoryPayload) (*BackupDirectory, error)
	DeleteBackupDirectory(ctx context.Context, machineId string, directoryId string, payload *DeleteDirectoryPayload) error
	ListTenantDirectories(ctx context.Context) ([]*BackupDirectory, error)
	ActionBackupDirectory(ctx context.Context, machineId string, directoryId string, payload *ActionDirectoryPayload) error
	DeleteMultipleDirectories(ctx context.Context, machineId string, payload *DeleteMultipleDirectoriesPayload) error
	ActionMultipleDirectories(ctx context.Context, machineId string, payload *ActionMultipleDirectoriesPayload) error

	ListTenantPolicies(ctx context.Context) ([]*Policy, error)
	CreatePolicy(ctx context.Context, payload *CreatePolicyPayload) (*Policy, error)
	GetBackupDirectoryPolicy(ctx context.Context, machineId string, directoryId string) (*Policy, error)
	GetPolicy(ctx context.Context, policyId string) (*Policy, error)
	PatchPolicy(ctx context.Context, policyId string, payload *PatchPolicyPayload) (*Policy, error)
	DeletePolicy(ctx context.Context, policyId string) error
	ListAppliedPolicyDirectories(ctx context.Context, policyId string) ([]*BackupDirectory, error)
	ActionPolicyDirectory(ctx context.Context, policyId string, payload *ActionPolicyDirectoryPayload) error
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
