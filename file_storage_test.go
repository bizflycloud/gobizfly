// This file is part of gobizfly

package gobizfly

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/bizflycloud/gobizfly/testlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShareList(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.FileStorageURL("/_"), func(writer http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `{"filestorages": [
  {
    "id": "share-1111-2222-3333",
    "name": "my-share",
    "size": 100,
    "share_protocol": "NFS",
    "description": "test share",
    "network_id": "net-1111",
    "subnet_id": "subnet-1111",
    "share_type": "default",
    "zone": "HN1",
    "status": "available",
    "export_locations": ["10.0.0.1:/share-1111"],
    "created_at": "2025-01-01 00:00:00",
    "updated_at": "2025-01-01 00:00:00"
  }
]}`
		_, _ = fmt.Fprint(writer, resp)
	})
	shares, err := client.FileStorage.List(ctx)
	require.NoError(t, err)
	assert.Len(t, shares, 1)
	assert.Equal(t, "my-share", shares[0].Name)
	assert.Equal(t, 100, shares[0].Size)
	assert.Equal(t, "NFS", shares[0].ShareProtocol)
}

func TestShareCreate(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.FileStorageURL("/_"), func(writer http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var payload *CreateShareRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "new-share", payload.Name)
		assert.Equal(t, 50, payload.Size)
		writer.WriteHeader(http.StatusAccepted)
		resp := `{"filestorage": {
  "id": "share-new-1111-2222",
  "name": "new-share",
  "size": 50,
  "share_protocol": "NFS",
  "zone": "HN1",
  "status": "creating",
  "created_at": "2025-01-02 00:00:00",
  "updated_at": "2025-01-02 00:00:00"
}}`
		_, _ = fmt.Fprint(writer, resp)
	})
	share, err := client.FileStorage.Create(ctx, &CreateShareRequest{
		Name: "new-share",
		Size: 50,
		Zone: "HN1",
	})
	require.NoError(t, err)
	assert.Equal(t, "share-new-1111-2222", share.ID)
	assert.Equal(t, "creating", share.Status)
}

func TestShareGet(t *testing.T) {
	setup()
	defer teardown()
	var fs fileStorageService
	mux.HandleFunc(testlib.FileStorageURL(fs.shareItemPath("share-1111-2222-3333")), func(writer http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `{"filestorage": {
  "id": "share-1111-2222-3333",
  "name": "my-share",
  "size": 100,
  "share_protocol": "NFS",
  "description": "test share",
  "zone": "HN1",
  "status": "available",
  "export_locations": ["10.0.0.1:/share-1111"],
  "created_at": "2025-01-01 00:00:00",
  "updated_at": "2025-01-01 00:00:00"
}}`
		_, _ = fmt.Fprint(writer, resp)
	})
	share, err := client.FileStorage.Get(ctx, "share-1111-2222-3333")
	require.NoError(t, err)
	assert.Equal(t, "share-1111-2222-3333", share.ID)
	assert.Equal(t, "available", share.Status)
}

func TestShareDelete(t *testing.T) {
	setup()
	defer teardown()
	var fs fileStorageService
	mux.HandleFunc(testlib.FileStorageURL(fs.shareItemPath("share-1111-2222-3333")), func(writer http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		writer.WriteHeader(http.StatusAccepted)
	})
	require.NoError(t, client.FileStorage.Delete(ctx, "share-1111-2222-3333", false))
}

func TestShareDeleteForce(t *testing.T) {
	setup()
	defer teardown()
	var fs fileStorageService
	mux.HandleFunc(testlib.FileStorageURL(fs.shareItemPath("share-1111-2222-3333")), func(writer http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "true", r.URL.Query().Get("force"))
		writer.WriteHeader(http.StatusAccepted)
	})
	require.NoError(t, client.FileStorage.Delete(ctx, "share-1111-2222-3333", true))
}

func TestShareResize(t *testing.T) {
	setup()
	defer teardown()
	var fs fileStorageService
	mux.HandleFunc(testlib.FileStorageURL(fs.shareResizePath("share-1111-2222-3333")), func(writer http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var payload *ResizeShareRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, 200, payload.NewSize)
		writer.WriteHeader(http.StatusAccepted)
		resp := `{"filestorage": {
  "id": "share-1111-2222-3333",
  "name": "my-share",
  "size": 200,
  "status": "resizing",
  "created_at": "2025-01-01 00:00:00",
  "updated_at": "2025-01-03 00:00:00"
}}`
		_, _ = fmt.Fprint(writer, resp)
	})
	share, err := client.FileStorage.Resize(ctx, "share-1111-2222-3333", &ResizeShareRequest{
		NewSize: 200,
	})
	require.NoError(t, err)
	assert.Equal(t, 200, share.Size)
	assert.Equal(t, "resizing", share.Status)
}

func TestGetAccessRules(t *testing.T) {
	setup()
	defer teardown()
	var fs fileStorageService
	mux.HandleFunc(testlib.FileStorageURL(fs.shareAccessPath("share-1111-2222-3333")), func(writer http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `[
  {
    "id": "rule-1111",
    "access_to": "10.0.0.0/24",
    "access_level": "rw",
    "access_type": "ip",
    "state": "active",
    "created_at": "2025-01-01 00:00:00",
    "updated_at": "2025-01-01 00:00:00"
  }
]`
		_, _ = fmt.Fprint(writer, resp)
	})
	rules, err := client.FileStorage.GetAccessRules(ctx, "share-1111-2222-3333")
	require.NoError(t, err)
	assert.Len(t, rules, 1)
	assert.Equal(t, "10.0.0.0/24", rules[0].AccessTo)
	assert.Equal(t, "rw", rules[0].AccessLevel)
}

func TestManageAccessRules(t *testing.T) {
	setup()
	defer teardown()
	var fs fileStorageService
	mux.HandleFunc(testlib.FileStorageURL(fs.shareAccessPath("share-1111-2222-3333")), func(writer http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var payload *ManageAccessRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Len(t, payload.Access, 1)
		resp := `[
  {
    "id": "rule-2222",
    "access_to": "192.168.1.0/24",
    "access_level": "ro",
    "access_type": "ip",
    "state": "active",
    "created_at": "2025-01-02 00:00:00",
    "updated_at": "2025-01-02 00:00:00"
  }
]`
		_, _ = fmt.Fprint(writer, resp)
	})
	rules, err := client.FileStorage.ManageAccessRules(ctx, "share-1111-2222-3333", &ManageAccessRequest{
		Access: []ShareAccessRule{
			{AccessTo: "192.168.1.0/24", AccessLevel: "ro", AccessType: "ip"},
		},
	})
	require.NoError(t, err)
	assert.Len(t, rules, 1)
	assert.Equal(t, "192.168.1.0/24", rules[0].AccessTo)
}

func TestDeleteAccessRule(t *testing.T) {
	setup()
	defer teardown()
	var fs fileStorageService
	mux.HandleFunc(testlib.FileStorageURL(fs.shareAccessRulePath("share-1111-2222-3333", "rule-1111")), func(writer http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
	})
	require.NoError(t, client.FileStorage.DeleteAccessRule(ctx, "share-1111-2222-3333", "rule-1111"))
}

func TestGetAccessStatus(t *testing.T) {
	setup()
	defer teardown()
	var fs fileStorageService
	mux.HandleFunc(testlib.FileStorageURL(fs.shareAccessStatusPath("share-1111-2222-3333")), func(writer http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `{"status": "active"}`
		_, _ = fmt.Fprint(writer, resp)
	})
	status, err := client.FileStorage.GetAccessStatus(ctx, "share-1111-2222-3333")
	require.NoError(t, err)
	assert.Equal(t, "active", status.Status)
}

func TestGetQuota(t *testing.T) {
	setup()
	defer teardown()
	var fs fileStorageService
	mux.HandleFunc(testlib.FileStorageURL(fs.quotaProjectPath("project-1111")), func(writer http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `{
  "project_id": "project-1111",
  "max_shares": 10,
  "max_total_size_gb": 1000
}`
		_, _ = fmt.Fprint(writer, resp)
	})
	quota, err := client.FileStorage.GetQuota(ctx, "project-1111")
	require.NoError(t, err)
	assert.Equal(t, "project-1111", quota.ProjectID)
	assert.Equal(t, 10, quota.MaxShares)
	assert.Equal(t, 1000, quota.MaxTotalSizeGB)
}

func TestUpdateQuota(t *testing.T) {
	setup()
	defer teardown()
	var fs fileStorageService
	mux.HandleFunc(testlib.FileStorageURL(fs.quotaPath()), func(writer http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var payload *QuotaRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "project-1111", payload.ProjectID)
		resp := `{
  "project_id": "project-1111",
  "max_shares": 20,
  "max_total_size_gb": 2000
}`
		_, _ = fmt.Fprint(writer, resp)
	})
	quota, err := client.FileStorage.UpdateQuota(ctx, &QuotaRequest{
		ProjectID:      "project-1111",
		MaxShares:      20,
		MaxTotalSizeGB: 2000,
	})
	require.NoError(t, err)
	assert.Equal(t, 20, quota.MaxShares)
	assert.Equal(t, 2000, quota.MaxTotalSizeGB)
}

func TestListRegions(t *testing.T) {
	setup()
	defer teardown()
	var fs fileStorageService
	mux.HandleFunc(testlib.FileStorageURL(fs.regionsPath()), func(writer http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `[
  {"name": "Ha Noi", "code": "HN1"},
  {"name": "Ho Chi Minh", "code": "HCM1"}
]`
		_, _ = fmt.Fprint(writer, resp)
	})
	regions, err := client.FileStorage.ListRegions(ctx)
	require.NoError(t, err)
	assert.Len(t, regions, 2)
	assert.Equal(t, "HN1", regions[0].Code)
}
