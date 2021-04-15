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
		var payload *createRepositoryPayload
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, false, payload.Public)
	})
	err := client.ContainerRegistry.Create(ctx, &createRepositoryPayload{
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
			var payload *editRepositoryPayload
			require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
			assert.Equal(t, false, payload.Public)
		})
	err := client.ContainerRegistry.EditRepo(ctx, "ji84wqtzr77ogo6b", &editRepositoryPayload{
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
