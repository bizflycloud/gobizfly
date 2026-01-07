package gobizfly

import (
    "context"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)


type ClusterJoinEverywhereRequest struct {
	PoolID          string                 `json:"pool_id" yaml:"pool_id"`
	PublicIP        string                 `json:"public_ip" yaml:"public_ip"`
	Hostname        string                 `json:"hostname" yaml:"hostname"`
	HasWanInterface bool                   `json:"has_wan_interface" yaml:"has_wan_interface"`
	IPAddresses     []string               `json:"ip_addresses" yaml:"ip_addresses"`
	Capacity        EverywhereNodeCapacity `json:"capacity" yaml:"capacity"`
	Annotation      *EverywhereAnnotation  `json:"annotation,omitempty" yaml:"annotation"`
}

type EverywhereAnnotation struct {
	ForceEndpoint string `json:"kilo.squat.ai/force-endpoint" yaml:"kilo.squat.ai/force-endpoint"`
}

type EverywhereNodeCapacity struct {
	Cores    int `json:"cores" yaml:"cores"`
	MemoryKB int `json:"memory_kb" yaml:"memory_kb"`
}

type ClusterCertificate struct {
	CaCert     string `json:"ca.pem" yaml:"ca.pem"`
	ClientKey  string `json:"client-key.pem" yaml:"client-key.pem"`
	ClientCert string `json:"client.pem" yaml:"client.pem"`
}

type ClusterReserved struct {
	SystemReserved ClusterSystemReserved `json:"system_reserved" yaml:"system_reserved"`
	KubeReserved   ClusterKubeReserved   `json:"kube_reserved" yaml:"kube_reserved"`
}

type ClusterSystemReserved struct {
	CPU    string `json:"cpu" yaml:"cpu"`
	Memory string `json:"memory" yaml:"memory"`
}

type ClusterKubeReserved struct {
	CPU    string `json:"cpu" yaml:"cpu"`
	Memory string `json:"memory" yaml:"memory"`
}

type ClusterJoinEverywhereResponse struct {
	APIServer     string             `json:"apiserver" yaml:"apiserver"`
	ClusterDNS    string             `json:"cluster_dns" yaml:"cluster_dns"`
	ClusterCIDR   string             `json:"cluster_cidr" yaml:"cluster_cidr"`
	CloudProvider string             `json:"cloud_provider" yaml:"cloud_provider"`
	Certificate   ClusterCertificate `json:"certificate" yaml:"certificate"`
	UUID          string             `json:"uuid" yaml:"uuid"`
	MaxPods       int                `json:"max_pods" yaml:"max_pods"`
	Reserved      ClusterReserved    `json:"reserved" yaml:"reserved"`
}

func (c *kubernetesEngineService) AddClusterEverywhere(ctx context.Context, id string, cjer *ClusterJoinEverywhereRequest) (*ClusterJoinEverywhereResponse, error) {
	var joinEverywhereResponse ClusterJoinEverywhereResponse
	req, err := c.client.NewRequest(ctx, http.MethodPost, kubernetesServiceName, clusterJoinEverywhere+"/"+id, &cjer)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	err = json.Unmarshal(body, &joinEverywhereResponse)
	if err != nil {
		return nil, err
	}
	return &joinEverywhereResponse, nil
}

type WorkerConfig struct {
	ID                string `json:"id" yaml:"id"`
	Version           string `json:"version" yaml:"version"`
	Everywhere        bool   `json:"everywhere" yaml:"everywhere"`
	Nvidiadevice      bool   `json:"nvidiadevice" yaml:"nvidiadevice"`
	CniVersion        string `json:"CNI_VERSION" yaml:"CNI_VERSION"`
	RuncVersion       string `json:"RUNC_VERSION" yaml:"RUNC_VERSION"`
	ContainerdVersion string `json:"CONTAINERD_VERSION" yaml:"CONTAINERD_VERSION"`
	KubeVersion       string `json:"KUBE_VERSION" yaml:"KUBE_VERSION"`
}

type ClusterInfoResponse struct {
	WorkerConfig WorkerConfig `json:"worker_config" yaml:"worker_configs"`
	K8sVersion   string       `json:"k8s_version" yaml:"k8s_version"`
	ShootUid     string       `json:"shoot_uid" yaml:"shoot_uid"`
	PoolName     string       `json:"pool_name" yaml:"pool_name"`
}

func (c *kubernetesEngineService) GetClusterInfo(ctx context.Context, pool_id string) (*ClusterInfoResponse, error) {
	var clusterInfoResponse ClusterInfoResponse
	req, err := c.client.NewRequest(ctx, http.MethodGet, kubernetesServiceName, clusterInfo+"/"+pool_id, nil)
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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &clusterInfoResponse)
	if err != nil {
		return nil, err
	}
	return &clusterInfoResponse, nil
}

// ClusterLeave handles worker node leaving from a cluster
// ClusterLeave handles worker node leaving from a cluster
func (k *kubernetesEngineService) ClusterLeave(ctx context.Context, clusterUID string, clusterToken string, req *ClusterLeaveRequest) (*ClusterLeaveResponse, error) {
    // Input validation
    if clusterUID == "" {
        return nil, fmt.Errorf("cluster UID cannot be empty")
    }
    if clusterToken == "" {
        return nil, fmt.Errorf("cluster token cannot be empty")
    }
    if req == nil {
        return nil, fmt.Errorf("request cannot be nil")
    }
    if req.NodeName == "" {
        return nil, fmt.Errorf("node name cannot be empty")
    }

    // Construct path using constant
    path := fmt.Sprintf("%s/%s", clusterLeave, clusterUID)
    
    // Create HTTP request using the library's standard pattern
    httpReq, err := k.client.NewRequest(ctx, http.MethodPost, kubernetesServiceName, path, req)
    if err != nil {
        return nil, fmt.Errorf("failed to create request: %w", err)
    }

    // Set required authentication header
    httpReq.Header.Set("X-Cluster-Token", clusterToken)

    // Execute request
    resp, err := k.client.Do(ctx, httpReq)
    if err != nil {
        return nil, fmt.Errorf("cluster leave request failed: %w", err)
    }
    defer func() {
        _ = resp.Body.Close()
    }()

    // Read response body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("failed to read response body: %w", err)
    }

    // Handle different HTTP status codes
    if resp.StatusCode >= 400 {
        return nil, fmt.Errorf("cluster leave failed with status %d: %s", resp.StatusCode, string(body))
    }

    // Unmarshal successful response
    var clusterLeaveResponse ClusterLeaveResponse
    if err := json.Unmarshal(body, &clusterLeaveResponse); err != nil {
        return nil, fmt.Errorf("failed to unmarshal response: %w", err)
    }

    return &clusterLeaveResponse, nil
}


