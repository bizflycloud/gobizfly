package gobizfly

import (
	"context"
	"encoding/json"
	"net/http"
)

type WorkerConfig struct {
	ID               	string		`json:"id" yaml:"id"`
	Version          	string 		`json:"version" yaml:"version"`
	Everywhere       	bool        `json:"everywhere" yaml:"everywhere"`
	Nvidiadevice		bool		`json:"nvidiadevice" yaml:"nvidiadevice"`
	CniVersion         string       `json:"CNI_VERSION" yaml:"CNI_VERSION"`
	RUNC_VERSION        string      `json:"RUNC_VERSION" yaml:"RUNC_VERSION"`
	CONTAINERD_VERSION  string      `json:"CONTAINERD_VERSION" yaml:"CONTAINERD_VERSION"`
	KUBE_VERSION    	string      `json:"KUBE_VERSION" yaml:"KUBE_VERSION"`

}

type WorkerConfigs struct {
	WorkerConfigs_ []WorkerConfig `json:"worker_configs" yaml:"worker_configs"`
}

func (c *kubernetesEngineService) GetAdminWorkerConfig(ctx context.Context) (*WorkerConfigs, error) {
	var workerConfigs *WorkerConfigs
	req, err := c.client.NewRequest(ctx, http.MethodGet, kubernetesServiceName, adminWorkerConfig, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(workerConfigs); err != nil {
		return nil, err
	}
	return workerConfigs, nil
}