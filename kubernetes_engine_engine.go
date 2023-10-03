package gobizfly

import (
	"io"
	"context"
	"encoding/json"
	"net/http"
)

type ClusterJoinEverywhereRequest struct {
	Hostname    string                   `json:"hostname" yaml:"hostname"`
	IPAddresses []string                 `json:"ip_addresses" yaml:"ip_addresses"`
	Capacity    EverywhereNodeCapacity `json:"capacity" yaml:"capacity"`
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
	SystemReserved		ClusterSystemReserved  `json:"system_reserved" yaml:"system_reserved"`
	KubeReserved 			ClusterKubeReserved    `json:"kube_reserved" yaml:"kube_reserved"`
}

type ClusterSystemReserved struct {
	CPU				string 	 `json:"cpu" yaml:"cpu"`
	Memory 		string 	 `json:"memory" yaml:"memory"`
}

type ClusterKubeReserved struct {
	CPU				string 	 `json:"cpu" yaml:"cpu"`
	Memory 		string 	 `json:"memory" yaml:"memory"`
}

type ClusterJoinEverywhereResponse struct {
	APIServer			string					`json:"apiserver" yaml:"apiserver"`
	ClusterDNS		string					`json:"cluster_dns" yaml:"cluster_dns"`
	ClusterCIDR		string					`json:"cluster_cidr" yaml:"cluster_cidr"`
	CloudProvider	string					`json:"cloud_provider" yaml:"cloud_provider"`
	Certificate 	ClusterCertificate		`json:"certificate" yaml:"certificate"`
	UUID					string					`json:"uuid" yaml:"uuid"`
	MaxPods				int						`json:"max_pods" yaml:"max_pods"`
	Reserved 			ClusterReserved			`json:"reserved" yaml:"reserved"`
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
	defer resp.Body.Close()

	err = json.Unmarshal(body, &joinEverywhereResponse)
	if err != nil {
		return nil, err
	}
	return &joinEverywhereResponse, nil
}

type WorkerConfig struct {
	ID               	string		`json:"id" yaml:"id"`
	Version          	string 		`json:"version" yaml:"version"`
	Everywhere       	bool      `json:"everywhere" yaml:"everywhere"`
	Nvidiadevice			bool			`json:"nvidiadevice" yaml:"nvidiadevice"`
	CniVersion        string    `json:"CNI_VERSION" yaml:"CNI_VERSION"`
	RuncVersion       string    `json:"RUNC_VERSION" yaml:"RUNC_VERSION"`
	ContainerdVersion string    `json:"CONTAINERD_VERSION" yaml:"CONTAINERD_VERSION"`
	KubeVersion    		string   	`json:"KUBE_VERSION" yaml:"KUBE_VERSION"`

}

type ClusterInfoResponse struct {
	WorkerConfig 	WorkerConfig 	`json:"worker_config" yaml:"worker_configs"`
	K8sVersion		string			 	`json:"k8s_version" yaml:"k8s_version"`
	ShootUid			string 				`json:"shoot_uid" yaml:"shoot_uid"`
	PoolName			string 				`json:"pool_name" yaml:"pool_name"`
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
	defer resp.Body.Close()
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

