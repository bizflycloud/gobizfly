// This file is part of gobizfly
//
// Copyright (C) 2020  BizFly Cloud
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>

package gobizfly

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
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
