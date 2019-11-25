package gobizfly

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestMemberList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(poolPath+"/023f2e34-7806-443b-bfae-16c324569a3d/members", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "members": [
        {
            "backup": false,
            "created_at": "2018-09-18T07:25:04",
            "weight": 1,
            "address": "10.6.169.102",
            "monitor_port": null,
            "subnet_id": "2f866d94-8218-4c9f-8c96-358837e63e6e",
            "protocol_port": 80,
            "project_id": "1e7f10a9850b45b488a3f0417ccb60e0",
            "provisioning_status": "ACTIVE",
            "monitor_address": null,
            "operating_status": "ONLINE",
            "updated_at": "2018-09-18T07:25:21",
            "name": "sapd-lemp-8",
            "admin_state_up": true,
            "tenant_id": "1e7f10a9850b45b488a3f0417ccb60e0",
            "id": "0b9b1602-fb7a-4f9e-ac2e-99f2d4f7b494"
        },
        {
            "backup": false,
            "created_at": "2018-09-18T07:25:22",
            "weight": 1,
            "address": "10.6.169.31",
            "monitor_port": null,
            "subnet_id": "2f866d94-8218-4c9f-8c96-358837e63e6e",
            "protocol_port": 80,
            "project_id": "1e7f10a9850b45b488a3f0417ccb60e0",
            "provisioning_status": "ACTIVE",
            "monitor_address": null,
            "operating_status": "ONLINE",
            "updated_at": "2018-09-18T07:25:27",
            "name": "sapd-lemp-11",
            "admin_state_up": true,
            "tenant_id": "1e7f10a9850b45b488a3f0417ccb60e0",
            "id": "54277bf2-68ea-4ddd-87ee-6bf4c91850a5"
        }
    ],
    "members_links": []
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	members, err := client.Member.List(ctx, "023f2e34-7806-443b-bfae-16c324569a3d", &ListOptions{})
	require.NoError(t, err)
	assert.Len(t, members, 2)
}

func TestMemberGet(t *testing.T) {
	setup()
	defer teardown()

	path := strings.Join([]string{poolPath, "023f2e34-7806-443b-bfae-16c324569a3d", "members", "0b9b1602-fb7a-4f9e-ac2e-99f2d4f7b494"}, "/")
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "backup": false,
    "created_at": "2018-09-18T07:25:04",
    "weight": 1,
    "address": "10.6.169.102",
    "monitor_port": null,
    "subnet_id": "2f866d94-8218-4c9f-8c96-358837e63e6e",
    "protocol_port": 80,
    "project_id": "1e7f10a9850b45b488a3f0417ccb60e0",
    "provisioning_status": "ACTIVE",
    "monitor_address": null,
    "operating_status": "ONLINE",
    "updated_at": "2018-09-18T07:25:21",
    "name": "sapd-lemp-8",
    "admin_state_up": true,
    "tenant_id": "1e7f10a9850b45b488a3f0417ccb60e0",
    "id": "0b9b1602-fb7a-4f9e-ac2e-99f2d4f7b494"
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	member, err := client.Member.Get(ctx, "023f2e34-7806-443b-bfae-16c324569a3d", "0b9b1602-fb7a-4f9e-ac2e-99f2d4f7b494")
	require.NoError(t, err)
	assert.Equal(t, "0b9b1602-fb7a-4f9e-ac2e-99f2d4f7b494", member.ID)
}

func TestMemberUpdate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(memberPath+"/957a1ace-1bd2-449b-8455-820b6e4b63f3", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)

		resp := `
{
    "member": {
        "monitor_port": 8080,
        "project_id": "e3cd678b11784734bc366148aa37580e",
        "name": "web-server-1",
        "weight": 20,
        "backup": false,
        "admin_state_up": true,
        "subnet_id": "bbb35f84-35cc-4b2f-84c2-a6a29bba68aa",
        "created_at": "2017-05-11T17:21:34",
        "provisioning_status": "PENDING_UPDATE",
        "monitor_address": null,
        "updated_at": "2017-05-11T17:21:37",
        "address": "192.0.2.16",
        "protocol_port": 80,
        "id": "957a1ace-1bd2-449b-8455-820b6e4b63f3",
        "operating_status": "NO_MONITOR"
    }
}
`
		_, _ = fmt.Fprint(w, resp)
	})

	name := "MemberUpdated"
	_, err := client.Member.Update(ctx, "957a1ace-1bd2-449b-8455-820b6e4b63f3", &MemberUpdateRequest{
		Name: name,
	})
	require.NoError(t, err)
}

func TestMemberDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(memberPath+"/957a1ace-1bd2-449b-8455-820b6e4b63f3", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		w.WriteHeader(http.StatusNoContent)
	})

	require.NoError(t, client.Member.Delete(ctx, "957a1ace-1bd2-449b-8455-820b6e4b63f3"))
}
