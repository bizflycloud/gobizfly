// This file is part of gobizfly

package gobizfly

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/bizflycloud/gobizfly/testlib"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_autoScalingGroup_List(t *testing.T) {
	setup()
	defer teardown()
	var asg autoScalingGroup
	mux.HandleFunc(testlib.AutoScalingURL(asg.resourcePath()), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
        {
          "_meta": {
            "max_results": 1,
            "total": 10,
            "page": 1
          },
          "clusters": [
            {
              "domain": null,
              "profile_id": "2b1a00f1-28cd-4daf-b886-c0f07b401ed4",
              "updated_at": "2020-01-07T09:26:52Z",
              "profile_name": "Centos-8",
              "id": "31e8465b-7275-4055-aeba-e3984453a223",
              "data": {},
              "desired_capacity": 1,
              "node_ids": [
                "7133bd53-507b-4648-b9e5-df648d00e9ec"
              ],
              "config": {},
              "metadata": {
                "deletion_policy": "04e6595b-a8e5-4701-a118-6c4095dd7cd7",
                "webhook_ids": {
                  "scale_out": {
                    "id": "e7ff2be5-50bc-491e-b9bf-69eba6865a2f",
                    "name": "receiver-31e8465b-CLUSTER_SCALE_OUT"
                  },
                  "scale_in": {
                    "id": "ef69815f-d7a7-4107-9f13-cc14d5b7e742",
                    "name": "receiver-31e8465b-CLUSTER_SCALE_IN"
                  }
                },
                "scale_in_receiver": "5e0eaa80045cab0010473a22",
                "scale_out_receiver": "5e0eaa83045cab0010473a23"
              },
              "status": "ACTIVE",
              "min_size": 1,
              "user": "bb08773452954176ad34de108ac517bd",
              "status_reason": "CLUSTER_RESIZE: number of active nodes is equal or above desired_capacity (1).",
              "max_size": 1,
              "name": "mongo",
              "created_at": "2020-01-03T02:44:59Z",
              "project": "fd109a330e1a4ffbb85c336972d5aed4",
              "init_at": "2020-01-03T02:44:07Z",
              "timeout": 3600,
              "dependents": {}
            }
          ]
        }`
		_, _ = fmt.Fprint(w, resp)
	})

	ASGroups, err := client.AutoScaling.AutoScalingGroups().List(ctx, false)
	require.NoError(t, err)
	ASGroup := ASGroups[0]
	assert.Equal(t, "31e8465b-7275-4055-aeba-e3984453a223", ASGroup.ID)
	assert.Equal(t, "2b1a00f1-28cd-4daf-b886-c0f07b401ed4", ASGroup.ProfileID)
	assert.Equal(t, "Centos-8", ASGroup.ProfileName)
	assert.Equal(t, "mongo", ASGroup.Name)
	assert.Equal(t, 1, ASGroup.MaxSize)
	assert.Equal(t, 1, ASGroup.MinSize)
	assert.Equal(t, 1, ASGroup.DesiredCapacity)
	assert.Equal(t, "receiver-31e8465b-CLUSTER_SCALE_IN", ASGroup.Metadata.WebhookIDs.ScaleIn.Name)
	assert.Equal(t, "ef69815f-d7a7-4107-9f13-cc14d5b7e742", ASGroup.Metadata.WebhookIDs.ScaleIn.ID)
	assert.Equal(t, "receiver-31e8465b-CLUSTER_SCALE_OUT", ASGroup.Metadata.WebhookIDs.ScaleOut.Name)
	assert.Equal(t, "e7ff2be5-50bc-491e-b9bf-69eba6865a2f", ASGroup.Metadata.WebhookIDs.ScaleOut.ID)
	assert.Equal(t, "04e6595b-a8e5-4701-a118-6c4095dd7cd7", ASGroup.Metadata.DeletionPolicy)
	assert.Equal(t, []string{"7133bd53-507b-4648-b9e5-df648d00e9ec"}, ASGroup.NodeIDs)

}

func Test_launchConfiguration_List(t *testing.T) {
	setup()
	defer teardown()
	var lc launchConfiguration
	mux.HandleFunc(testlib.AutoScalingURL(lc.resourcePath()), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `{
          "profiles": [
            {
              "name": "LaunchConfiguration",
              "availability_zone": "HN1",
              "datadisks": [
                {
                  "type": "HDD",
                  "delete_on_termination": true,
                  "size": 20
                },
                {
                  "type": "HDD",
                  "delete_on_termination": true,
                  "size": 40
                }
              ],
              "key_name": "ministry",
              "created_at": "2020-01-06T10:03:02Z",
              "rootdisk": {
                "type": "HDD",
                "delete_on_termination": true,
                "size": 20
              },
              "user_data": "#!/bin/bash\\necho \\\"Hello Linux\\\" > /tmp/greeting.txt",
              "os": {
                "os_name": "CentOS 8.0@2019-11-30-f62e620e",
                "type": "snapshot",
                "id": "a135f52c-f93c-466d-a87e-6db66acad4e0",
                "ios_name": "snapshot-10-24-0-26-12-2019"
              },
              "id": "243f772f-5fb1-4fee-bdf1-70413de1e2d0",
              "flavor": "nix.2c_2g",
              "type": "premium",
              "network_plan": "free_datatransfer",
              "networks": [{
                "id": "wan",
                "security_groups": [
                  "7102e422-2c81-4943-86bf-d90b8e7520d4",
                  "46b3f48a-9f3c-43a6-9927-3c790fbb166b"
                ]
              }],
              "security_groups": [
                "7102e422-2c81-4943-86bf-d90b8e7520d4",
                "46b3f48a-9f3c-43a6-9927-3c790fbb166b"
              ],
              "metadata": {
                "category": "premium"
              }
            }
          ],
          "_meta": {
            "max_results": 1,
            "total": 10,
            "page": 1
          }
        }`
		_, _ = fmt.Fprint(w, resp)
	})

	launchConfigs, err := client.AutoScaling.LaunchConfigurations().List(ctx, false)
	require.NoError(t, err)
	launchConfig := launchConfigs[0]

	assert.Equal(t, "243f772f-5fb1-4fee-bdf1-70413de1e2d0", launchConfig.ID)
	assert.Equal(t, "LaunchConfiguration", launchConfig.Name)
	assert.Equal(t, "nix.2c_2g", launchConfig.Flavor)
	assert.Equal(t, "HDD", launchConfig.RootDisk.Type)
	assert.Equal(t, "snapshot", launchConfig.OperatingSystem.CreateFrom)
	assert.Equal(t, "free_datatransfer", launchConfig.NetworkPlan)
}

func Test_webhook_List(t *testing.T) {
	setup()
	defer teardown()
	var wh webhook
	mux.HandleFunc(testlib.AutoScalingURL(wh.resourcePath("da6d6ab3-f765-46d0-ad6d-b5bfa237b920")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `[
          {
            "cluster_name": "group_01",
            "cluster_id": "da6d6ab3-f765-46d0-ad6d-b5bfa237b920",
            "action_type": "CLUSTER SCALE IN",
            "action_id": "ce596947-d88c-4af0-9fcd-55d8ceac42ad"
          },
          {
            "cluster_name": "group_01",
            "cluster_id": "da6d6ab3-f765-46d0-ad6d-b5bfa237b920",
            "action_type": "CLUSTER SCALE OUT",
            "action_id": "5a267cee-e600-4c23-915b-877f74cf9601"
          }
        ]`
		_, _ = fmt.Fprint(w, resp)
	})

	webhooks, err := client.AutoScaling.Webhooks().List(ctx, "da6d6ab3-f765-46d0-ad6d-b5bfa237b920")
	require.NoError(t, err)
	webhook := webhooks[0]

	assert.Equal(t, "ce596947-d88c-4af0-9fcd-55d8ceac42ad", webhook.ActionID)
	assert.Equal(t, "CLUSTER SCALE IN", webhook.ActionType)
	assert.Equal(t, "da6d6ab3-f765-46d0-ad6d-b5bfa237b920", webhook.ClusterID)
	assert.Equal(t, "group_01", webhook.ClusterName)
}

func Test_event_List(t *testing.T) {
	setup()
	defer teardown()
	var e event
	mux.HandleFunc(testlib.AutoScalingURL(e.resourcePath("09ea7069-2767-4d86-b125-0549827e30f7", 1, 10)), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := ` {
          "_meta": {
            "max_results": 12,
            "total": 10,
            "page": 1
          },
          "events": [
            {
              "level": "INFO",
              "timestamp": "2020-07-28T09:49:26+00:00",
              "otype": "NODE",
              "meta_data": {
                "action": {
                  "outputs": {
                    "node_name": "group_01-node-oiPmVAkF"
                  },
                  "created_at": "2020-07-28T09:49:06Z"
                }
              },
              "cluster_id": "da6d6ab3-f765-46d0-ad6d-b5bfa237b920",
              "type": "removed",
              "action": "NODE_DELETE",
              "status_reason": "Node deleted successfully.",
              "id": "121e9210-99bb-4abe-8ac9-2f367e996812"
            }]
        }`
		_, _ = fmt.Fprint(w, resp)
	})
	events, err := client.AutoScaling.Events().List(ctx, "09ea7069-2767-4d86-b125-0549827e30f7", 1, 10)
	require.NoError(t, err)
	event := events[0]

	assert.Equal(t, "da6d6ab3-f765-46d0-ad6d-b5bfa237b920", event.ClusterID)
	assert.Equal(t, "removed", event.ActionType)
	assert.Equal(t, "NODE_DELETE", event.ActionName)
}

func Test_policy_List(t *testing.T) {
	setup()
	defer teardown()
	var p policy
	mux.HandleFunc(testlib.AutoScalingURL(p.resourcePath("09ea7069-2767-4d86-b125-0549827e30f7")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `{
          "scale_out_policy": [
            {
              "range_time": 300,
              "type": "CHANGE_IN_CAPACITY",
              "metric": "ram_used",
              "number": 1,
              "cooldown": 300,
              "threshold": 95,
              "best_effort": true,
              "id": "5f1fd6f0e48a89000e37a1b5"
            }
          ],
          "deletion_policy": {},
          "load_balancer_policy": {},
          "scale_in_policy": [],
          "doing_task": []
        }`
		_, _ = fmt.Fprint(w, resp)
	})
	policies, err := client.AutoScaling.Policies().List(ctx, "09ea7069-2767-4d86-b125-0549827e30f7")
	require.NoError(t, err)
	scaleOutPolicy := policies.ScaleOutPolicies[0]

	policyID := "5f1fd6f0e48a89000e37a1b5"
	policyType := "CHANGE_IN_CAPACITY"
	assert.Equal(t, ScalePolicy{
		RangeTime:  300,
		BestEffort: true,
		CoolDown:   300,
		Threshold:  95,
		ScaleSize:  1,
		MetricType: "ram_used",
		ID:         &policyID,
		Type:       &policyType,
	}, scaleOutPolicy)
}

func Test_node_List(t *testing.T) {
	setup()
	defer teardown()
	var n node
	mux.HandleFunc(testlib.AutoScalingURL(n.resourcePath("09ea7069-2767-4d86-b125-0549827e30f7")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `{
          "nodes": [
            {
              "status": "ACTIVE",
              "name": "group_01-node-AxMHi4cG",
              "profile_id": "98f73f63-2e4d-426a-9e5c-431b0bdec9c8",
              "profile_name": "LaunchConfiguration",
              "physical_id": "a08c872e-46b5-455b-adc0-c79adf929aab",
              "status_reason": "Creation succeeded",
              "id": "17bb9458-f76c-45e8-84dc-7a831ac1ef96",
              "addresses": {
                "EXT_DIRECTNET_3": [
                  {
                    "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:03:57:8b",
                    "version": 4,
                    "addr": "103.56.157.210",
                    "OS-EXT-IPS:type": "fixed"
                  }
                ],
                "priv_botv@vccloud.vn": [
                  {
                    "OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:7e:57:ce",
                    "version": 4,
                    "addr": "10.21.46.50",
                    "OS-EXT-IPS:type": "fixed"
                  }
                ]
              }
            }
          ],
          "_meta": {
            "max_results": 1,
            "total": 10,
            "page": 1
          }
        }`
		_, _ = fmt.Fprint(w, resp)
	})
	nodes, err := client.AutoScaling.Nodes().List(ctx, "09ea7069-2767-4d86-b125-0549827e30f7", true)
	require.NoError(t, err)
	node := nodes[0]

	assert.Equal(t, "ACTIVE", node.Status)
	assert.Equal(t, "group_01-node-AxMHi4cG", node.Name)
	assert.Equal(t, "17bb9458-f76c-45e8-84dc-7a831ac1ef96", node.ID)
	assert.Equal(t, "a08c872e-46b5-455b-adc0-c79adf929aab", node.PhysicalID)
	assert.Equal(t, "LaunchConfiguration", node.ProfileName)
	assert.Equal(t, "98f73f63-2e4d-426a-9e5c-431b0bdec9c8", node.ProfileID)
}

func Test_schedule_List(t *testing.T) {
	setup()
	defer teardown()
	var s schedule
	mux.HandleFunc(testlib.AutoScalingURL(s.resourcePath("09ea7069-2767-4d86-b125-0549827e30f7")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `{
            "cron_triggers": [{
                "sizing": {
                    "_type": "daily",
                    "_from": {
                        "inputs": {
                            "min_size": 1,
                            "desired_capacity": 1,
                            "cluster_id": "da6d6ab3-f765-46d0-ad6d-b5bfa237b920",
                            "max_size": 2
                        },
                        "cron_pattern": "05 22 * * *"
                    },
                    "_to": {
                        "inputs": {
                            "min_size": 1,
                            "desired_capacity": 1,
                            "cluster_id": "da6d6ab3-f765-46d0-ad6d-b5bfa237b920",
                            "max_size": 5
                        },
                        "cron_pattern": "06 22 * * *"
                    }
                },
                "_id": "5f203ee4bf32db000effe9ce",
                "name": "fake-name",
                "created_at": "22:06 28/07/2020",
                "updated_at": null,
                "valid": {
                    "_from": "00:00 30/07/2020",
                    "_to": "23:59 31/07/2020"
                },
                "cluster_id": "da6d6ab3-f765-46d0-ad6d-b5bfa237b920",
                "project_id": "fd109a330e1a4ffbb85c336972d5aed4"
            }],
            "_meta": {
                "max_results": 1,
                "total": 10,
                "page": 1
            }
        }`
		_, _ = fmt.Fprint(w, resp)
	})
	schedules, err := client.AutoScaling.Schedules().List(ctx, "09ea7069-2767-4d86-b125-0549827e30f7")
	require.NoError(t, err)
	cronTrigger := schedules[0]

	assert.Equal(t, "5f203ee4bf32db000effe9ce", cronTrigger.ID)
	assert.Equal(t, "fake-name", cronTrigger.Name)
	assert.Equal(t, "da6d6ab3-f765-46d0-ad6d-b5bfa237b920", cronTrigger.ClusterID)
	assert.Equal(t, "00:00 30/07/2020", cronTrigger.Valid.From)
	_to := "23:59 31/07/2020"
	assert.Equal(t, &_to, cronTrigger.Valid.To)
	assert.Equal(t, "daily", cronTrigger.Sizing.Type)
	assert.Equal(t, "05 22 * * *", cronTrigger.Sizing.From.CronPattern)
	assert.Equal(t, 5, cronTrigger.Sizing.To.Inputs.MaxSize)
}

func Test_common_UsingResource(t *testing.T) {
	setup()
	defer teardown()
	var c common
	mux.HandleFunc(testlib.AutoScalingURL(c.usingResourcePath()), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `{
          "ssh_keys": [
            "ministry"
          ],
          "snapshots": [
            "e54124ee-667d-4b7c-9e7b-9f451bc1883e",
            "dcf46400-0ecd-4971-aaa6-cebce0d7ec5c"
          ]
        }`
		_, _ = fmt.Fprint(w, resp)
	})
	resources, err := client.AutoScaling.Common().AutoScalingUsingResource(ctx)
	require.NoError(t, err)

	assert.Equal(t, []string{"ministry"}, resources.SSHKeys)
	assert.Equal(t, []string{
		"e54124ee-667d-4b7c-9e7b-9f451bc1883e",
		"dcf46400-0ecd-4971-aaa6-cebce0d7ec5c",
	}, resources.Snapshots)
}

func TestIsValidQuotas(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.AutoScalingURL(getQuotasResourcePath()), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		resp := `{
            "message": {
                "valid": true
            }
        }`
		_, _ = fmt.Fprint(w, resp)
	})
	valid, err := isValidQuotas(ctx, client, "", "ProfileID", 2, 9)
	require.NoError(t, err)

	assert.Equal(t, true, valid)

}

func Test_common_GetSuggestion(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		ctx             context.Context
		ProfileID       string
		desiredCapacity int
		maxSize         int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &common{
				client: tt.fields.client,
			}
			got, err := c.AutoScalingGetSuggestion(tt.args.ctx, tt.args.ProfileID, tt.args.desiredCapacity, tt.args.maxSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("common.GetSuggestion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("common.GetSuggestion() = %v, want %v", got, tt.want)
			}
		})
	}
}
