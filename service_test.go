// This file is part of gobizfly

package gobizfly

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServiceCatalogList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(serviceUrl, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
  "services": [
    {
      "canonical_name": "cloud_server", 
      "code": "CS", 
      "description": "Cloud Server", 
      "enabled": true, 
      "icon": "https://manage.bizflycloud.vn/iaas-cloud/api", 
      "id": 2, 
      "name": "Cloud Server", 
      "region": "HN", 
      "service_url": "https://manage.bizflycloud.vn/iaas-cloud/api"
    }
  ]
}`
		_, _ = fmt.Fprint(w, resp)
	})

	resp, err := client.Service.List(ctx)
	require.NoError(t, err)
	assert.Equal(t, "cloud_server", resp[0].CanonicalName)
	//assert.Equal(t, "894f0e66-4571-4fea-9766-5fc615aec4a5", resp.VolumeDetail.ID)
	//assert.Equal(t, "Detach successfully", resp.Message)
}
