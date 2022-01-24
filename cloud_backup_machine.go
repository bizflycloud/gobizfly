package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type CloudBackupListMachineParams struct {
	IncludeDeleted bool
}
type CloudBackupMachine struct {
	Id                 string `json:"id"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
	Name               string `json:"name,omitempty"`
	HostName           string `json:"host_name,omitempty"`
	IpAddress          string `json:"ip_address,omitempty"`
	OsVersion          string `json:"os_version,omitempty"`
	OsVersionId        string `json:"os_version_id,omitempty"`
	AgentVersion       string `json:"agent_version,omitempty"`
	TenantId           string `json:"tenant_id"`
	Encryption         bool   `json:"encryption,omitempty"`
	OperationStatus    bool   `json:"operation_status,omitempty"`
	Status             string `json:"status,omitempty"`
	MachineStorageSize int    `json:"machine_storage_size,omitempty"`
}

type CloudBackupCreateMachinePayload struct {
	Name         string `json:"name"`
	HostName     string `json:"host_name"`
	IpAddress    string `json:"ip_address"`
	OsVersion    string `json:"os_version"`
	AgentVersion string `json:"agent_version"`
}

type CloudBackupFileContent struct {
	AccessKey string `json:"access_key"`
	ApiUrl    string `json:"api_url"`
	BrokerUrl string `json:"broker_url"`
	MachineId string `json:"machine_id"`
	SecretKey string `json:"secret_key"`
}
type CloudBackupExtendedMachine struct {
	AccessKey    string                 `json:"access_key"`
	AgentVersion string                 `json:"agent_version"`
	CreatedAt    string                 `json:"created_at"`
	HostName     string                 `json:"host_name "`
	Id           string                 `json:"id"`
	IpAddress    string                 `json:"ip_address"`
	Name         string                 `json:"name"`
	OsVersion    string                 `json:"os_version"`
	SecretKey    string                 `json:"secret_key"`
	OsMachineId  string                 `json:"os_machine_id"`
	Encryption   bool                   `json:"encryption"`
	TenantId     string                 `json:"tenant_id"`
	UpdatedAt    string                 `json:"updated_at"`
	FileContent  CloudBackupFileContent `json:"file_content"`
}

type CloudBackupPatchMachinePayload struct {
	Name         string `json:"name"`
	HostName     string `json:"host_name"`
	IpAddress    string `json:"ip_address"`
	OsVersion    string `json:"os_version"`
	AgentVersion string `json:"agent_version"`
	OsMachineID  string `json:"os_machine_id"`
}

type CloudBackupDeleteMachinePayload struct {
	Keep         bool     `json:"keep"`
	DirectoryIds []string `json:"directory_ids"`
}

type CloudBackupActionMachinePayload struct {
	Action string `json:"action"`
}

func (cb *cloudBackupService) ListTenantMachines(ctx context.Context, listOption *CloudBackupListMachineParams) ([]*CloudBackupMachine, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName,
		cb.machinesPath(), nil)
	if err != nil {
		return nil, err
	}
	if listOption.IncludeDeleted {
		req.URL.RawQuery = "include_deleted=true"
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var machines []*CloudBackupMachine
	if err := json.NewDecoder(resp.Body).Decode(&machines); err != nil {
		return nil, err
	}
	return machines, nil
}

func (cb *cloudBackupService) CreateMachine(ctx context.Context, payload *CloudBackupCreateMachinePayload) (*CloudBackupExtendedMachine, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodPost, cloudBackupServiceName,
		cb.machinesPath(), payload)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var machine *CloudBackupExtendedMachine
	if err := json.NewDecoder(resp.Body).Decode(&machine); err != nil {
		return nil, err
	}
	return machine, nil
}

func (cb *cloudBackupService) GetMachine(ctx context.Context, machineId string) (*CloudBackupMachine, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName,
		cb.itemMachinePath(machineId), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var machine *CloudBackupMachine
	if err := json.NewDecoder(resp.Body).Decode(&machine); err != nil {
		return nil, err
	}
	return machine, nil
}

func (cb *cloudBackupService) PatchMachine(ctx context.Context, machineId string, payload *CloudBackupPatchMachinePayload) (*CloudBackupMachine, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodPatch, cloudBackupServiceName,
		cb.itemMachinePath(machineId), payload)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var machine *CloudBackupMachine
	if err := json.NewDecoder(resp.Body).Decode(&machine); err != nil {
		return nil, err
	}
	return machine, nil
}

func (cb *cloudBackupService) DeleteMachine(ctx context.Context, machineId string, payload *CloudBackupDeleteMachinePayload) error {
	req, err := cb.client.NewRequest(ctx, http.MethodDelete, cloudBackupServiceName,
		cb.itemMachinePath(machineId), payload)
	if err != nil {
		return err
	}
	_, err = cb.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (cb *cloudBackupService) ActionMachine(ctx context.Context, machineId string, payload *CloudBackupActionMachinePayload) error {
	req, err := cb.client.NewRequest(ctx, http.MethodPost, cloudBackupServiceName,
		strings.Join([]string{cb.itemMachinePath(machineId), "action"}, "/"), payload)
	if err != nil {
		return err
	}
	_, err = cb.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (cb *cloudBackupService) ResetMachineSecretKey(ctx context.Context, machineId string) (*CloudBackupExtendedMachine, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodPost, cloudBackupServiceName,
		strings.Join([]string{cb.itemMachinePath(machineId), "reset-secret-key"}, "/"), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	var extendedMachine *CloudBackupExtendedMachine
	if err := json.NewDecoder(resp.Body).Decode(&extendedMachine); err != nil {
		return nil, err
	}
	return extendedMachine, err
}

func (cb *cloudBackupService) ListTenantPolicies(ctx context.Context) ([]*CloudBackupPolicy, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName,
		cb.policyPath(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data struct {
		Policies []*CloudBackupPolicy `json:"policies"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Policies, nil
}
