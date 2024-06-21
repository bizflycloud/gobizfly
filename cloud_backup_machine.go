package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

// CloudBackupListMachineParams - Parameters for list cloud backup machine
type CloudBackupListMachineParams struct {
	IncludeDeleted bool
}

// CloudBackupMachine represents cloud backup machine
type CloudBackupMachine struct {
	ID                 string `json:"id"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
	Name               string `json:"name,omitempty"`
	HostName           string `json:"host_name,omitempty"`
	IPAddress          string `json:"ip_address,omitempty"`
	OsVersion          string `json:"os_version,omitempty"`
	OsVersionID        string `json:"os_version_id,omitempty"`
	AgentVersion       string `json:"agent_version,omitempty"`
	TenantID           string `json:"tenant_id"`
	Encryption         bool   `json:"encryption,omitempty"`
	OperationStatus    bool   `json:"operation_status,omitempty"`
	Status             string `json:"status,omitempty"`
	MachineStorageSize int    `json:"machine_storage_size,omitempty"`
}

// CloudBackupCreateMachinePayload represents cloud backup machine payload
type CloudBackupCreateMachinePayload struct {
	Name         string `json:"name"`
	HostName     string `json:"host_name"`
	IPAddress    string `json:"ip_address"`
	OsVersion    string `json:"os_version"`
	AgentVersion string `json:"agent_version"`
}

// CloudBackupFileContent represents cloud backup file content
type CloudBackupFileContent struct {
	AccessKey string `json:"access_key"`
	APIURL    string `json:"api_url"`
	BrokerURL string `json:"broker_url"`
	MachineID string `json:"machine_id"`
	SecretKey string `json:"secret_key"`
}

// CloudBackupExtendedMachine represents cloud backup extended machine
type CloudBackupExtendedMachine struct {
	AccessKey    string                 `json:"access_key"`
	AgentVersion string                 `json:"agent_version"`
	CreatedAt    string                 `json:"created_at"`
	HostName     string                 `json:"host_name "`
	ID           string                 `json:"id"`
	IPAddress    string                 `json:"ip_address"`
	Name         string                 `json:"name"`
	OsVersion    string                 `json:"os_version"`
	SecretKey    string                 `json:"secret_key"`
	OsMachineID  string                 `json:"os_machine_id"`
	Encryption   bool                   `json:"encryption"`
	TenantID     string                 `json:"tenant_id"`
	UpdatedAt    string                 `json:"updated_at"`
	FileContent  CloudBackupFileContent `json:"file_content"`
}

// CloudBackupPatchMachinePayload represents cloud backup patch machine payload
type CloudBackupPatchMachinePayload struct {
	Name         string `json:"name"`
	HostName     string `json:"host_name"`
	IPAddress    string `json:"ip_address"`
	OsVersion    string `json:"os_version"`
	AgentVersion string `json:"agent_version"`
	OsMachineID  string `json:"os_machine_id"`
}

// CloudBackupDeleteMachinePayload represents cloud backup delete machine payload
type CloudBackupDeleteMachinePayload struct {
	Keep         bool     `json:"keep"`
	DirectoryIDs []string `json:"directory_ids"`
}

// CloudBackupActionMachinePayload represents cloud backup action machine payload
type CloudBackupActionMachinePayload struct {
	Action string `json:"action"`
}

// ListTenantMachines - List cloud backup machine belonging to tenant
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

// CreateMachine - Create cloud backup machine
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

// GetMachine - Get cloud backup machine
func (cb *cloudBackupService) GetMachine(ctx context.Context, machineID string) (*CloudBackupMachine, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodGet, cloudBackupServiceName,
		cb.itemMachinePath(machineID), nil)
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

// PatchMachine - Patch cloud backup machine
func (cb *cloudBackupService) PatchMachine(ctx context.Context, machineID string, payload *CloudBackupPatchMachinePayload) (*CloudBackupMachine, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodPatch, cloudBackupServiceName,
		cb.itemMachinePath(machineID), payload)
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

// DeleteMachine - Delete cloud backup machine
func (cb *cloudBackupService) DeleteMachine(ctx context.Context, machineID string, payload *CloudBackupDeleteMachinePayload) error {
	req, err := cb.client.NewRequest(ctx, http.MethodDelete, cloudBackupServiceName,
		cb.itemMachinePath(machineID), payload)
	if err != nil {
		return err
	}
	_, err = cb.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

// ActionMachine - Action cloud backup machine
func (cb *cloudBackupService) ActionMachine(ctx context.Context, machineID string, payload *CloudBackupActionMachinePayload) error {
	req, err := cb.client.NewRequest(ctx, http.MethodPost, cloudBackupServiceName,
		strings.Join([]string{cb.itemMachinePath(machineID), "action"}, "/"), payload)
	if err != nil {
		return err
	}
	_, err = cb.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

// ResetMachineSecretKey - Reset cloud backup machine secret key
func (cb *cloudBackupService) ResetMachineSecretKey(ctx context.Context, machineID string) (*CloudBackupExtendedMachine, error) {
	req, err := cb.client.NewRequest(ctx, http.MethodPost, cloudBackupServiceName,
		strings.Join([]string{cb.itemMachinePath(machineID), "reset-secret-key"}, "/"), nil)
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

// ListTenantPolicies - List cloud backup tenant policy
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
