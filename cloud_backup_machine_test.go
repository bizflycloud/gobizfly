package gobizfly

import (
	"fmt"
	"github.com/bizflycloud/gobizfly/testlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"strings"
	"testing"
)

func TestListTenantMachine(t *testing.T) {
	setup()
	defer teardown()
	var cb *cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(cb.machinesPath()),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			resp := `[

  {

    "agent_version": "0.01",

    "created_at": "Wed, 10 Jun 2020 15:14:47 GMT",

    "host_name ": "mysql-server",

    "id": "743d916d-1362-4328-a1d4-c53ccd1bd625",

    "ip_address": "10.0.0.1",

    "name": "mysql",

    "os_version": "windows",

    "machine_storage_size": null,

    "tenant_id": "98dbba57-60c7-4c22-8087-8313997ea776",

    "updated_at": "Wed, 10 Jun 2020 15:14:47 GMT"

  },
  {
    "agent_version": null,
    "created_at": "Wed, 10 Jun 2020 16:30:56 GMT",
    "host_name ": null,
    "id": "3e213039-41c7-4c10-9dc0-1a8dc0bf4823",
    "ip_address": null,
    "name": null,
    "os_version": null,
    "machine_storage_size": null,
    "tenant_id": "98dbba57-60c7-4c22-8087-8313997ea776",
    "updated_at": "Wed, 10 Jun 2020 16:30:56 GMT"
  }

]`
			fmt.Fprint(writer, resp)
		})
	machines, err := client.CloudBackup.ListTenantMachines(ctx, &ListMachineParams{})
	require.NoError(t, err)
	assert.Equal(t, 2, len(machines))
	assert.Equal(t, "98dbba57-60c7-4c22-8087-8313997ea776", machines[1].TenantId)
}

func TestCreateMachine(t *testing.T) {
	setup()
	defer teardown()
	var cb *cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(cb.machinesPath()),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			resp := `{

  "access_key": "UBZ6YHKFHWKCADINEGPQ",

  "agent_version": "0.1",

  "created_at": "Thu, 11 Jun 2020 09:17:51 GMT",

  "host_name ": "host name",

  "id": "8d8ead54-0d81-45e2-bd03-0b94cde42cc0",

  "ip_address": "10.10.10.25",

  "name": "machine name",

  "os_version": "windows 10",

  "secret_key": "19d2b170ae550eb2697ff0b2f0e9aaedcab564080eee8e9c3e335bd749bea621",

  "os_machine_id": null,

  "encryption": false,

  "tenant_id ": "98dbba57-60c7-4c22-8087-8313997ea776",

  "updated_at": "Thu, 11 Jun 2020 09:17:51 GMT",

  "file_content": {

    "access_key": "RZVA6JIAF7JNH3KENRFE",

    "api_url": "config",

    "broker_url": "config",

    "machine_id": "96eb548c-42b3-4fe9-a873-8658769dc293",

    "secret_key": "b57b476c1ef401cfb90e598580efea9c33b37a2ee6f765e9bddf5aa395461387"

  }

}`
			fmt.Fprint(writer, resp)
		})
	machine, err := client.CloudBackup.CreateMachine(ctx, &CreateMachinePayload{
		Name:         "machine_name",
		HostName:     "host_name",
		IpAddress:    "1.2.3.4",
		OsVersion:    "Windows 10",
		AgentVersion: "0.1",
	})
	require.NoError(t, err)
	assert.Equal(t, "RZVA6JIAF7JNH3KENRFE", machine.FileContent.AccessKey)
}

func TestGetMachine(t *testing.T) {
	setup()
	defer teardown()
	var cb cloudBackupService

	mux.HandleFunc(testlib.CloudBackupURL(cb.itemMachinePath("123")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			resp := `{

      "agent_version": null,

      "created_at": "2021-12-27T08:37:35.061792+00:00",

      "deleted": false,

      "encryption": false,

      "host_name": null,

      "id": "e3634e25-c4fc-44c7-8892-e230939f04cb",

      "ip_address": null,

      "machine_storage_size": null,

      "name": "machine-23",

      "operation_status": true,

      "os_version": null,

      "status": "OFFLINE",

      "tenant_id": "2c21d5f1-7237-4d01-bc13-036d47753099",

      "updated_at": "2021-12-27T08:37:35.061792+00:00"

    },}`
			fmt.Fprint(writer, resp)
		})
	machine, err := client.CloudBackup.GetMachine(ctx, "123")
	require.NoError(t, err)
	assert.Equal(t, "OFFLINE", machine.Status)
}

func TestPatchMachine(t *testing.T) {
	setup()
	defer teardown()

	var cb *cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(cb.itemMachinePath("123")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPatch, r.Method)
			resp := `{

  "agent_version": "version: dev, commit: , built: ",

  "broker_url": "mqtt://localhost:1883",

  "created_at": "2021-11-01T04:21:22.368979+00:00",

  "deleted": false,

  "encryption": false,

  "host_name": "vinhvn",

  "id": "7b2e1783-9066-4bf6-bb09-84c5d2d3165f",

  "ip_address": "192.168.18.13",

  "name": "machine123",

  "operation_status": true,

  "os_version": "Arch Linux\n",

  "status": "OFFLINE",

  "tenant_id": "2c21d5f1-7237-4d01-bc13-036d47753099",

  "updated_at": "2022-01-12T08:57:47.961537+00:00"

}`
			fmt.Fprint(writer, resp)
		})
	machine, err := client.CloudBackup.PatchMachine(ctx, "123", &PatchMachinePayload{
		Name:         "Nginx Server",
		HostName:     "nginx-1",
		IpAddress:    "10.3.35.12",
		OsVersion:    "linux",
		AgentVersion: "0.1",
	})
	require.NoError(t, err)
	assert.Equal(t, "vinhvn", machine.HostName)
}

func TestDeleteMachine(t *testing.T) {
	setup()
	defer teardown()
	var cb cloudBackupService

	mux.HandleFunc(testlib.CloudBackupURL(cb.itemMachinePath("123")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
		})
	directoryIds := []string{"123", "456"}
	err := client.CloudBackup.DeleteMachine(ctx, "123", &DeleteMachinePayload{
		Keep:         true,
		DirectoryIds: directoryIds,
	})
	require.NoError(t, err)
}

func TestActionMachine(t *testing.T) {
	setup()
	defer teardown()

	var cb cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(strings.Join([]string{cb.itemMachinePath("123"), "action"}, "/")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
		})
	err := client.CloudBackup.ActionMachine(ctx, "123", &ActionMachinePayload{
		Action: "enable",
	})
	require.NoError(t, err)
}

func TestResetMachineSecretKey(t *testing.T) {
	setup()
	defer teardown()
	var cb cloudBackupService
	mux.HandleFunc(testlib.CloudBackupURL(strings.Join([]string{cb.itemMachinePath("123"), "reset-secret-key"}, "/")),
		func(writer http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			resp := `{

  "access_key": "AVI4CEFZWLBPJXD63YBK",

  "agent_version": "version: dev, commit: , built: ",

  "created_at": "2021-11-01T04:21:22.368979+00:00",

  "deleted": false,

  "encryption": false,

  "file_content": {

    "access_key": "AVI4CEFZWLBPJXD63YBK",

    "api_url": "https://localhost:5000",

    "machine_id": "7b2e1783-9066-4bf6-bb09-84c5d2d3165f",

    "secret_key": "7e2d2c27739ad33567a638f7971969d1bfac49b5ec4790df37eb643e74bdcff5"

  },

  "host_name": "vinhvn",

  "id": "7b2e1783-9066-4bf6-bb09-84c5d2d3165f",

  "ip_address": "192.168.18.13",

  "name": "machine123",

  "operation_status": true,

  "os_version": "Arch Linux\n",

  "secret_key": "7e2d2c27739ad33567a638f7971969d1bfac49b5ec4790df37eb643e74bdcff5",

  "status": "OFFLINE",

  "tenant_id": "2c21d5f1-7237-4d01-bc13-036d47753099",

  "updated_at": "2022-01-12T09:24:18.129449+00:00"

}`
			fmt.Fprint(writer, resp)
		})
	machine, err := client.CloudBackup.ResetMachineSecretKey(ctx, "123")
	require.NoError(t, err)
	assert.Equal(t, "7e2d2c27739ad33567a638f7971969d1bfac49b5ec4790df37eb643e74bdcff5", machine.SecretKey)
}
