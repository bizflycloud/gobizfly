// This file is part of gobizfly

package gobizfly

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/bizflycloud/gobizfly/testlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReposList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(testlib.RegistryURL(registryPath), func(writer http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
  "repositories": [
    {
      "name": "string",
      "last_push": "2021-01-22T07:06:51.597Z",
      "pulls": 0,
      "public": true,
      "created_at": "2021-01-22T07:06:51.597Z"
    }
  ]
}`
		_, _ = fmt.Fprint(writer, resp)
	})
	repos, err := client.ContainerRegistry.List(ctx, &ListOptions{})
	require.NoError(t, err)
	assert.Len(t, repos, 1)
	assert.Equal(t, "string", repos[0].Name)
	assert.Equal(t, 0, repos[0].Pulls)
}

func TestRepoCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(testlib.RegistryURL(registryPath), func(writer http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var payload *CreateRepositoryPayload
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, false, payload.Public)
	})
	err := client.ContainerRegistry.Create(ctx, &CreateRepositoryPayload{
		Name:   "abc",
		Public: false,
	})
	require.NoError(t, err)
}

func TestRepoDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(strings.Join([]string{testlib.RegistryURL(registryPath), "ji84wqtzr77ogo6b"}, "/"),
		func(writer http.ResponseWriter, r *http.Request) {
			require.Equal(t, http.MethodDelete, r.Method)
		})
	require.NoError(t, client.ContainerRegistry.Delete(ctx, "ji84wqtzr77ogo6b"))
}

func TestGetRepoTags(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(strings.Join([]string{testlib.RegistryURL(registryPath), "ji84wqtzr77ogo6b"}, "/"),
		func(writer http.ResponseWriter, r *http.Request) {
			require.Equal(t, http.MethodGet, r.Method)
			resp := `{
  "repository": {
    "name": "string",
    "last_push": "2021-01-22T07:56:49.839Z",
    "pulls": 0,
    "public": true,
    "created_at": "2021-01-22T07:56:49.839Z"
  },
  "tags": [
    {
      "name": "string",
      "author": "string",
      "last_updated": "2021-01-22T07:56:49.839Z",
      "created_at": "2021-01-22T07:56:49.839Z",
      "last_scan": "2021-01-22T07:56:49.839Z",
      "scan_status": "string",
      "vulnerabilities": 0,
      "fixes": 0
    }
  ]
}`
			_, _ = fmt.Fprint(writer, resp)
		})
	resp, err := client.ContainerRegistry.GetTags(ctx, "ji84wqtzr77ogo6b")
	require.NoError(t, err)
	assert.Equal(t, "string", resp.Repository.Name)
	assert.Equal(t, 0, resp.Tags[0].Fixes)
}

func TestEditRepo(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(strings.Join([]string{testlib.RegistryURL(registryPath), "ji84wqtzr77ogo6b"}, "/"),
		func(writer http.ResponseWriter, r *http.Request) {
			require.Equal(t, http.MethodPatch, r.Method)
			var payload *EditRepositoryPayload
			require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
			assert.Equal(t, false, payload.Public)
		})
	err := client.ContainerRegistry.EditRepo(ctx, "ji84wqtzr77ogo6b", &EditRepositoryPayload{
		Public: false,
	})
	require.NoError(t, err)
}

func TestDeleteImageTag(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(strings.Join([]string{testlib.RegistryURL(registryPath), "ji84wqtzr77ogo6b", "tag", "tag"}, "/"),
		func(writer http.ResponseWriter, r *http.Request) {
			require.Equal(t, http.MethodDelete, r.Method)
		})
	err := client.ContainerRegistry.DeleteTag(ctx, "ji84wqtzr77ogo6b", "tag")
	require.NoError(t, err)
}

func TestGetImage(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(strings.Join([]string{testlib.RegistryURL(registryPath), "ji84wqtzr77ogo6b", "tag", "tag"}, "/"),
		func(writer http.ResponseWriter, r *http.Request) {
			require.Equal(t, http.MethodGet, r.Method)
			resp := `{
  "repository": {
    "name": "string",
    "last_push": "2021-01-22T08:09:48.391Z",
    "pulls": 0,
    "public": true,
    "created_at": "2021-01-22T08:09:48.391Z"
  },
  "tag": {
    "name": "string",
    "author": "string",
    "last_updated": "2021-01-22T08:09:48.391Z",
    "size": 0,
    "created_at": "2021-01-22T08:09:48.391Z",
    "last_scan": "2021-01-22T08:09:48.391Z",
    "scan_status": "string",
    "vulnerabilities": 0,
    "fixes": 0
  },
  "vulnerabilities": [
    {
      "package": "string",
      "name": "string",
      "namespace": "string",
      "description": "string",
      "link": "string",
      "severity": "string",
      "fixed_by": "string"
    }
  ]
}`
			_, _ = fmt.Fprint(writer, resp)
		})
	resp, err := client.ContainerRegistry.GetTag(ctx, "ji84wqtzr77ogo6b", "tag", "tag")
	require.NoError(t, err)
	assert.Len(t, resp.Vulnerabilities, 1)
	assert.Equal(t, "string", resp.Vulnerabilities[0].Package)
}

func TestGenToken(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.RegistryURL(tokenPath), func(writer http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		resp := `{"token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6IkFTNzU6UEhPRTpGNU9NOllQTDU6VDRMRDpRRVNGOjZNTkw6MkNISzpYU1lIOjVWN0E6T0hVNTpVS0dOIn0.eyJpc3MiOiJjci1obi0xLnZjY2xvdWQudm46NTAwMyIsInN1YiI6InN2dHQudHVuZ2RzQHZjY2xvdWQudm4iLCJhdWQiOiJjci1obi0xLnZjY2xvdWQudm46NTAwMiIsImlhdCI6MTYyNTQ4ODQ0NiwibmJmIjoxNjI1NDg4NDQxLCJleHAiOjE2MjU0OTIwNDksImp0aSI6Ijk2NlgzOHZEbGgyU0FxdmFrWDlQWllUSGNnajQxN1gwZFlOVDRMRmtUeiIsImFjY2VzcyI6W3sicHJvamVjdF9pZCI6ImJjOGQyNzkwZmM5YTQ2OTQ5ODE4Yjk0MmMwYTgyNGRlIiwibmFtZSI6ImJjOGQyNzkwZmM5YTQ2OTQ5ODE4Yjk0MmMwYTgyNGRlL3Rlc3RfY3RsMiIsImFjdGlvbnMiOlsicHVsbCIsInB1c2giXX0seyJwcm9qZWN0X2lkIjoiYmM4ZDI3OTBmYzlhNDY5NDk4MThiOTQyYzBhODI0ZGUiLCJuYW1lIjoiYmM4ZDI3OTBmYzlhNDY5NDk4MThiOTQyYzBhODI0ZGUvamgiLCJhY3Rpb25zIjpbInB1bGwiLCJwdXNoIl19XX0.BTCK9vaRC3EAa8wPrvD1Gtrej7BzUyv3mluYT7Un-NeSUzXz7Gr7IRITyGZWH17Zt93VRsEED84yzbM1ageXsz0_nsGc4-RSESUAu980Et6z8FUj40auho4bTKuSp-bzQKFJ_90if02JYOIaB59RYpi93iiwXViUdnG3YjkJ6M5LzvnVGGXifLbGPl-m2r9bwiLfazKzpWHhVRekEfCe3OVVqTYJSYmvHFFi7Y7zVMnY_LjcVt4KHzV8vxpHdYTuX56wYrfYsXvbhzh8fa7R5RkzHyykXjTOv47sSaY5nQekiS5Q_tk7aSdjhGt1SXnkWfztpzImYHRHWFaCpUnQBQ", "expires_in": 3603, "issued_at": "2021-07-05T12:34:06+00:00"} `
		_, _ = fmt.Fprint(writer, resp)
	})
	payload := GenerateTokenPayload{
		ExpiresIn: 6969,
		Scopes: []Scope{
			{
				Action:     []string{"pull", "push"},
				Repository: "test",
			},
		},
	}
	resp, err := client.ContainerRegistry.GenerateToken(ctx, &payload)
	require.NoError(t, err)
	assert.Equal(t, resp.ExpiresIn, 3603)
}
