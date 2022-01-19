package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type ListMachineParams struct {
	IncludeDeleted bool
}
type Machine struct {
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

type CreateMachinePayload struct {
	Name         string `json:"name"`
	HostName     string `json:"host_name"`
	IpAddress    string `json:"ip_address"`
	OsVersion    string `json:"os_version"`
	AgentVersion string `json:"agent_version"`
}

type FileContent struct {
	AccessKey string `json:"access_key"`
	ApiUrl    string `json:"api_url"`
	BrokerUrl string `json:"broker_url"`
	MachineId string `json:"machine_id"`
	SecretKey string `json:"secret_key"`
}
type ExtendedMachine struct {
	AccessKey    string      `json:"access_key"`
	AgentVersion string      `json:"agent_version"`
	CreatedAt    string      `json:"created_at"`
	HostName     string      `json:"host_name "`
	Id           string      `json:"id"`
	IpAddress    string      `json:"ip_address"`
	Name         string      `json:"name"`
	OsVersion    string      `json:"os_version"`
	SecretKey    string      `json:"secret_key"`
	OsMachineId  string      `json:"os_machine_id"`
	Encryption   bool        `json:"encryption"`
	TenantId     string      `json:"tenant_id"`
	UpdatedAt    string      `json:"updated_at"`
	FileContent  FileContent `json:"file_content"`
}

type PatchMachinePayload struct {
	Name         string `json:"name"`
	HostName     string `json:"host_name"`
	IpAddress    string `json:"ip_address"`
	OsVersion    string `json:"os_version"`
	AgentVersion string `json:"agent_version"`
	OsMachineID  string `json:"os_machine_id"`
}

type DeleteMachinePayload struct {
	Keep         bool     `json:"keep"`
	DirectoryIds []string `json:"directory_ids"`
}

type ActionMachinePayload struct {
	Action string `json:"action"`
}

func (cb *cloudBackupService) ListTenantMachines(ctx context.Context, listOption *ListMachineParams) ([]*Machine, error) {
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
	var machines []*Machine
	if err := json.NewDecoder(resp.Body).Decode(&machines); err != nil {
		return nil, err
	}
	return machines, nil
}

func (cb *cloudBackupService) CreateMachine(ctx context.Context, payload *CreateMachinePayload) (*ExtendedMachine, error) {
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
	var machine *ExtendedMachine
	if err := json.NewDecoder(resp.Body).Decode(&machine); err != nil {
		return nil, err
	}
	return machine, nil
}

func (cb *cloudBackupService) GetMachine(ctx context.Context, machineId string) (*Machine, error) {
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
	var machine *Machine
	if err := json.NewDecoder(resp.Body).Decode(&machine); err != nil {
		return nil, err
	}
	return machine, nil
}

func (cb *cloudBackupService) PatchMachine(ctx context.Context, machineId string, payload *PatchMachinePayload) (*Machine, error) {
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
	var machine *Machine
	if err := json.NewDecoder(resp.Body).Decode(&machine); err != nil {
		return nil, err
	}
	return machine, nil
}

func (cb *cloudBackupService) DeleteMachine(ctx context.Context, machineId string, payload *DeleteMachinePayload) error {
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

func (cb *cloudBackupService) ActionMachine(ctx context.Context, machineId string, payload *ActionMachinePayload) error {
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

func (cb *cloudBackupService) ResetMachineSecretKey(ctx context.Context, machineId string) (*ExtendedMachine, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodPost, cloudBackupServiceName,
		strings.Join([]string{cb.itemMachinePath(machineId), "reset-secret-key"}, "/"), nil)
	if err != nil {
		return nil, err
	}
	resp, err := cb.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	var extendedMachine *ExtendedMachine
	if err := json.NewDecoder(resp.Body).Decode(&extendedMachine); err != nil {
		return nil, err
	}
	return extendedMachine, err
}

func (cb *cloudBackupService) ListTenantPolicies(ctx context.Context) ([]*Policy, error) {
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
		Policies []*Policy `json:"policies"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Policies, nil
}
