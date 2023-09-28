package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
)

type clusterJoinEverywhereRequest struct {
	Hostname    string                   `json:"hostname" yaml:"hostname"`
	IPAddresses []string                 `json:"ip_addresses" yaml:"ip_addresses"`
	Capacity    []everywhereNodeCapacity `json:"capacity" yaml:"capacity"`
}

type everywhereNodeCapacity struct {
	Cores    int `json:"cores" yaml:"cores"`
	MemoryKB int `json:"memory_kb" yaml:"memory_kb"`
}

type clusterCertificate struct {
	CaCert     string `json:"ca.pem" yaml:"ca.pem"`
	ClientKey  string `json:"client-key.pem" yaml:"client-key.pem"`
	ClientCert string `json:"client.pem" yaml:"client.pem"`
}

type clusterReserved struct {
	SystemReserved			clusterSystemReserved  `json:"system_reserved" yaml:"system_reserved"`
	KubeReserved 			clusterKubeReserved    `json:"kube_reserved" yaml:"kube_reserved"`
}

type clusterSystemReserved struct {
	CPU			string 	 `json:"cpu" yaml:"cpu"`
	Memory 		string 	 `json:"memory" yaml:"memory"`
}

type clusterKubeReserved struct {
	CPU			string 	 `json:"cpu" yaml:"cpu"`
	Memory 		string 	 `json:"memory" yaml:"memory"`
}

type clusterJoinEverywhereResponse struct {
	APIServer		string					`json:"apiserver" yaml:"apiserver"`
	ClusterDNS		string					`json:"cluster_dns" yaml:"cluster_dns"`
	ClusterCIDR		string					`json:"cluster_cidr" yaml:"cluster_cidr"`
	CloudProvider	string					`json:"cloud_provider" yaml:"cloud_provider"`
	Certificate 	clusterCertificate		`json:"certificate" yaml:"certificate"`
	UUID			string					`json:"uuid" yaml:"uuid"`
	MaxPods			int						`json:"max_pods" yaml:"max_pods"`
	Reserved 		clusterReserved			`json:"reserved" yaml:"reserved"`
}

func (c *kubernetesEngineService) AddClusterEverywhere(ctx context.Context, id string, cjer *clusterJoinEverywhereRequest) (*clusterJoinEverywhereResponse, error) {
	var joinEverywhereResponse *clusterJoinEverywhereResponse
	req, err := c.client.NewRequest(ctx, http.MethodPost, kubernetesServiceName, clusterJoinEverywhere, &cjer)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(joinEverywhereResponse); err != nil {
		return nil, err
	}
	return joinEverywhereResponse, nil
}
