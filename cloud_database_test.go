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

func Test_CloudDatabaseInstance_List(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.DatabaseURL(cloudDatabaseInstancesResourcePath), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
	"_meta": {
		"results_count": 1
	},
	"instances": [
		{
			"id": "9c727335-4b53-44c3-866b-60ae502b0a3f",
			"name": "instance-test",
			"created_at": "2021-07-16T17:00:18.000000Z",
			"description": "Instance",
			"public_access": true,
			"datastore": {
				"type": "MariaDB",
				"name": "10.4",
				"id": "ee988cc3-bb30-4aaf-9837-e90a34f60d37"
			},
			"autoscaling": {
				"enable": true,
				"volume": {
					"threshold": 80,
					"limited": 2000
				},
				"receivers": [
					{
						"id": "60f1591222ac8ff9e3b90b0e",
						"name": "callback_9c727335-4b53-44c3-866b-60ae502b0a3f",
						"action": "auto_scaling_volume"
					}
				],
				"alarms": [
					{
						"id": "60f15911ed653f049fb83f31",
						"name": "autoscaling_volume_9c727335-4b53-44c3-866b-60ae502b0a3f",
						"receiver_id": "60f1591222ac8ff9e3b90b0e"
					}
				]
			},
			"nodes": [
				{
					"id": "173d0dc3-46ac-40e9-bd01-bce21db0fd28",
					"name": "instance-test-primary-qxxsJMGf",
					"availability_zone": "HN1",
					"role": "primary",
					"dns": {
						"private": "priv-instance-test-primary-yrXqiqaQ-9c72733.dbaas-staging.bizflycloud.vn",
						"public": "pub-instance-test-primary-yrXqiqaQ-9c72733.dbaas-staging.bizflycloud.vn"
					},
					"addresses": {
						"private": [
							{
								"ip_address": "10.20.1.201:3306",
								"network_name": "priv_vctest_devcs_tri19@vccloud.vn"
							}
						],
						"public": [
							{
								"ip_address": "45.124.94.58:3306",
								"network_name": "public@network"
							}
						]
					}
				},
				{
					"id": "48cb294a-cc27-4e1a-b5fb-a2cd57bf5fcd",
					"name": "instance-test-secondary-qxxsJMGf",
					"availability_zone": "HN1",
					"role": "secondary",
					"dns": {
						"private": null,
						"public": null
					},
					"addresses": {
						"private": null,
						"public": null
					}
				},
				{
					"id": "f65c5f3f-8ab3-4394-a160-09303e9e6fd6",
					"name": "instance-4-replica-dXuquJDX",
					"availability_zone": "HN1",
					"role": "replica",
					"dns": {
						"private": "priv-instance-4-replica-UZxg0DJ1-9c72733.dbaas-staging.bizflycloud.vn",
						"public": "pub-instance-4-replica-UZxg0DJ1-9c72733.dbaas-staging.bizflycloud.vn"
					},
					"addresses": {
						"private": [
							{
								"ip_address": "10.20.1.112:3306",
								"network_name": "priv_vctest_devcs_tri19@vccloud.vn"
							}
						],
						"public": [
							{
								"ip_address": "45.124.94.185:3306",
								"network_name": "public@network"
							}
						]
					}
				},
				{
					"id": "8dbaefc6-0ef5-46b1-b6e4-1ce449da0d14",
					"name": "instance-4-replica-seeTl4gB",
					"availability_zone": "HN1",
					"role": "replica",
					"dns": {
						"private": "priv-instance-4-replica-B7nCFXf6-9c72733.dbaas-staging.bizflycloud.vn",
						"public": "pub-instance-4-replica-B7nCFXf6-9c72733.dbaas-staging.bizflycloud.vn"
					},
					"addresses": {
						"private": [
							{
								"ip_address": "10.20.1.202:3306",
								"network_name": "priv_vctest_devcs_tri19@vccloud.vn"
							}
						],
						"public": [
							{
								"ip_address": "45.124.94.244:3306",
								"network_name": "public@network"
							}
						]
					}
				}
			]
		}
	]
}`
		_, _ = fmt.Fprint(w, resp)
	})

	DBInstances, err := client.CloudDatabase.Instances().List(ctx, &CloudDatabaseListOption{})
	require.NoError(t, err)
	DBInstance := DBInstances[0]
	assert.Equal(t, "9c727335-4b53-44c3-866b-60ae502b0a3f", DBInstance.ID)
	assert.Equal(t, "instance-test", DBInstance.Name)
	assert.Equal(t, "Instance", DBInstance.Description)
	assert.Equal(t, true, DBInstance.PublicAccess)
	assert.Equal(t, "ee988cc3-bb30-4aaf-9837-e90a34f60d37", DBInstance.Datastore.ID)
	assert.Equal(t, "173d0dc3-46ac-40e9-bd01-bce21db0fd28", DBInstance.Nodes[0].ID)
	assert.Equal(t, "48cb294a-cc27-4e1a-b5fb-a2cd57bf5fcd", DBInstance.Nodes[1].ID)
	assert.Equal(t, "f65c5f3f-8ab3-4394-a160-09303e9e6fd6", DBInstance.Nodes[2].ID)
}

func Test_CloudDatabaseInstance_ListNodes(t *testing.T) {
	setup()
	defer teardown()
	var ins *cloudDatabaseInstances
	mux.HandleFunc(testlib.DatabaseURL(ins.resourceTypePath("9c727335-4b53-44c3-866b-60ae502b0a3f", nodes)), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
	"id": "9c727335-4b53-44c3-866b-60ae502b0a3f",
	"name": "instance-test",
	"created_at": "2021-07-16T17:00:18.000000Z",
	"description": "Instance",
	"public_access": true,
	"datastore": {
		"type": "MariaDB",
		"name": "10.4",
		"id": "ee988cc3-bb30-4aaf-9837-e90a34f60d37"
	},
	"autoscaling": {
		"enable": true,
		"volume": {
			"threshold": 80,
			"limited": 2000
		},
		"receivers": [
			{
				"id": "60f1591222ac8ff9e3b90b0e",
				"name": "callback_9c727335-4b53-44c3-866b-60ae502b0a3f",
				"action": "auto_scaling_volume"
			}
		],
		"alarms": [
			{
				"id": "60f15911ed653f049fb83f31",
				"name": "autoscaling_volume_9c727335-4b53-44c3-866b-60ae502b0a3f",
				"receiver_id": "60f1591222ac8ff9e3b90b0e"
			}
		]
	},
	"nodes": [
		{
			"id": "173d0dc3-46ac-40e9-bd01-bce21db0fd28",
			"name": "instance-test-primary-qxxsJMGf",
			"availability_zone": "HN1",
			"role": "primary",
			"dns": {
				"private": "priv-instance-test-primary-yrXqiqaQ-9c72733.dbaas-staging.bizflycloud.vn",
				"public": "pub-instance-test-primary-yrXqiqaQ-9c72733.dbaas-staging.bizflycloud.vn"
			},
			"addresses": {
				"private": [
					{
						"ip_address": "10.20.1.201:3306",
						"network_name": "priv_vctest_devcs_tri19@vccloud.vn"
					}
				],
				"public": [
					{
						"ip_address": "45.124.94.58:3306",
						"network_name": "public@network"
					}
				]
			}
		},
		{
			"id": "48cb294a-cc27-4e1a-b5fb-a2cd57bf5fcd",
			"name": "instance-test-secondary-qxxsJMGf",
			"availability_zone": "HN1",
			"role": "secondary",
			"dns": {
				"private": null,
				"public": null
			},
			"addresses": {
				"private": null,
				"public": null
			}
		},
		{
			"id": "f65c5f3f-8ab3-4394-a160-09303e9e6fd6",
			"name": "instance-4-replica-dXuquJDX",
			"availability_zone": "HN1",
			"role": "replica",
			"dns": {
				"private": "priv-instance-4-replica-UZxg0DJ1-9c72733.dbaas-staging.bizflycloud.vn",
				"public": "pub-instance-4-replica-UZxg0DJ1-9c72733.dbaas-staging.bizflycloud.vn"
			},
			"addresses": {
				"private": [
					{
						"ip_address": "10.20.1.112:3306",
						"network_name": "priv_vctest_devcs_tri19@vccloud.vn"
					}
				],
				"public": [
					{
						"ip_address": "45.124.94.185:3306",
						"network_name": "public@network"
					}
				]
			}
		},
		{
			"id": "8dbaefc6-0ef5-46b1-b6e4-1ce449da0d14",
			"name": "instance-4-replica-seeTl4gB",
			"availability_zone": "HN1",
			"role": "replica",
			"dns": {
				"private": "priv-instance-4-replica-B7nCFXf6-9c72733.dbaas-staging.bizflycloud.vn",
				"public": "pub-instance-4-replica-B7nCFXf6-9c72733.dbaas-staging.bizflycloud.vn"
			},
			"addresses": {
				"private": [
					{
						"ip_address": "10.20.1.202:3306",
						"network_name": "priv_vctest_devcs_tri19@vccloud.vn"
					}
				],
				"public": [
					{
						"ip_address": "45.124.94.244:3306",
						"network_name": "public@network"
					}
				]
			}
		}
	]
}`
		_, _ = fmt.Fprint(w, resp)
	})

	DBNodes, err := client.CloudDatabase.Instances().ListNodes(ctx, "9c727335-4b53-44c3-866b-60ae502b0a3f", &CloudDatabaseListOption{})
	require.NoError(t, err)
	DBNode := DBNodes[0]
	assert.Equal(t, "173d0dc3-46ac-40e9-bd01-bce21db0fd28", DBNode.ID)
	assert.Equal(t, "instance-test-primary-qxxsJMGf", DBNode.Name)
	assert.Equal(t, "HN1", DBNode.AvailabilityZone)
	assert.Equal(t, "primary", DBNode.Role)
}

func Test_CloudDatabaseInstance_ListBackups(t *testing.T) {
	setup()
	defer teardown()
	var ins *cloudDatabaseInstances
	mux.HandleFunc(testlib.DatabaseURL(ins.resourceTypePath("ef317a9f-c705-4ea2-8fcb-c18450effb6c", backups)), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "_meta": {
        "page": 1,
        "results_per_page": 5,
        "filter": {
            "name": "",
            "start_time": "2021-01-01 00:00:00",
            "end_time": "2021-02-04 10:19:00"
        },
        "results_count": 2
    },
    "backups": [
        {
            "id": "7393cf38-5eb6-45e7-9ea9-f52dd3497cab",
            "name": "backup_name",
            "description": "MariaDB1 (120f585f-5aec-400b-a715-16c0b3c6a9d1)",
            "node_id": "120f585f-5aec-400b-a715-16c0b3c6a9d1",
            "created": "2021-02-04T10:18:47",
            "updated": "2021-02-04T10:18:52",
            "size": 0.21,
            "status": "COMPLETED",
            "parent_id": null,
            "project_id": "b746214aafba4923a5fbf9478ea64474",
            "datastore": {
                "type": "MariaDB",
                "version": "10.4",
                "version_id": "ee988cc3-bb30-4aaf-9837-e90a34f60d37"
            },
            "type": "manual",
            "meta_data": {
                "role": "primary"
            }
        },
        {
            "id": "75460fd5-3ce8-4f4a-a680-73a0d4914105",
            "name": "backup_name_2",
            "description": "MariaDB1 (120f585f-5aec-400b-a715-16c0b3c6a9d1)",
            "node_id": "120f585f-5aec-400b-a715-16c0b3c6a9d1",
            "created": "2021-02-04T10:18:34",
            "updated": "2021-02-04T10:18:38",
            "size": 0.21,
            "status": "COMPLETED",
            "parent_id": null,
            "project_id": "b746214aafba4923a5fbf9478ea64474",
            "datastore": {
                "type": "MariaDB",
                "version": "10.4",
                "version_id": "ee988cc3-bb30-4aaf-9837-e90a34f60d37"
            },
            "type": "manual",
            "meta_data": {
                "role": "replica"
            }
        }
    ]
}`
		_, _ = fmt.Fprint(w, resp)
	})

	DBBackups, err := client.CloudDatabase.Instances().ListBackups(ctx, "ef317a9f-c705-4ea2-8fcb-c18450effb6c", &CloudDatabaseListOption{})
	require.NoError(t, err)
	DBBackup := DBBackups[0]
	assert.Equal(t, "7393cf38-5eb6-45e7-9ea9-f52dd3497cab", DBBackup.ID)
	assert.Equal(t, "backup_name", DBBackup.Name)
	assert.Equal(t, "120f585f-5aec-400b-a715-16c0b3c6a9d1", DBBackup.NodeID)
	assert.Equal(t, "b746214aafba4923a5fbf9478ea64474", DBBackup.ProjectID)
}

func Test_CloudDatabaseInstance_ListSchedules(t *testing.T) {
	setup()
	defer teardown()
	var ins *cloudDatabaseInstances
	mux.HandleFunc(testlib.DatabaseURL(ins.resourceTypePath("a24509d4-cec1-4055-a476-676009140a3d", schedules)), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "_meta": {
        "page": 1,
        "results_per_page": 10,
        "filter": {
            "name": ""
        },
        "results_count": 2
    },
    "schedules": [
        {
            "id": "6ed3d8bd-c358-48f6-8890-20ed4e1fbfce",
            "name": "schedule_name",
            "pattern": "0 4,5,6 * * 1,2",
            "instance_id": "a24509d4-cec1-4055-a476-676009140a3d",
            "next_execution_time": "16:00 06-03-2021",
            "limit_backup": 3,
            "node_id": "2b138c69-be63-49f5-8b66-0180b7590af3",
            "node_name": "test4-member-ryv5RPgV"
        },
        {
            "id": "6ed3d8bd-c358-48f6-8890-20ed4e1fbfce",
            "name": "schedule_name_2",
            "pattern": "5 4 15,16 * *",
            "instance_id": "a24509d4-cec1-4055-a476-676009140a3d",
            "next_execution_time": "16:00 06-05-2021",
            "limit_backup": 2,
            "node_id": "2b138c69-be63-49f5-8b66-0180b7590af3",
            "node_name": "instance_name-member-ryv5RPgV"
        }
    ]
}`
		_, _ = fmt.Fprint(w, resp)
	})

	DBSchedules, err := client.CloudDatabase.Instances().ListBackupSchedules(ctx, "a24509d4-cec1-4055-a476-676009140a3d", &CloudDatabaseListOption{})
	require.NoError(t, err)
	DBSchedule := DBSchedules[0]
	assert.Equal(t, "6ed3d8bd-c358-48f6-8890-20ed4e1fbfce", DBSchedule.ID)
	assert.Equal(t, "schedule_name", DBSchedule.Name)
	assert.Equal(t, "2b138c69-be63-49f5-8b66-0180b7590af3", DBSchedule.NodeID)
	assert.Equal(t, "a24509d4-cec1-4055-a476-676009140a3d", DBSchedule.InstanceID)
}

func Test_CloudDatabaseInstance_Get(t *testing.T) {
	setup()
	defer teardown()
	var dbc cloudDatabaseInstances
	mux.HandleFunc(testlib.DatabaseURL(dbc.resourcePath("9c727335-4b53-44c3-866b-60ae502b0a3f")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
	"id": "9c727335-4b53-44c3-866b-60ae502b0a3f",
	"name": "instance-test",
	"created_at": "2021-07-16T17:00:18.000000Z",
	"description": "Instance",
	"public_access": true,
	"datastore": {
		"type": "MariaDB",
		"name": "10.4",
		"id": "ee988cc3-bb30-4aaf-9837-e90a34f60d37"
	},
	"autoscaling": {
		"enable": true,
		"volume": {
			"threshold": 80,
			"limited": 2000
		},
		"receivers": [
			{
				"id": "60f1591222ac8ff9e3b90b0e",
				"name": "callback_9c727335-4b53-44c3-866b-60ae502b0a3f",
				"action": "auto_scaling_volume"
			}
		],
		"alarms": [
			{
				"id": "60f15911ed653f049fb83f31",
				"name": "autoscaling_volume_9c727335-4b53-44c3-866b-60ae502b0a3f",
				"receiver_id": "60f1591222ac8ff9e3b90b0e"
			}
		]
	},
	"nodes": [
		{
			"id": "173d0dc3-46ac-40e9-bd01-bce21db0fd28",
			"name": "instance-test-primary-qxxsJMGf",
			"availability_zone": "HN1",
			"role": "primary",
			"dns": {
				"private": "priv-instance-test-primary-yrXqiqaQ-9c72733.dbaas-staging.bizflycloud.vn",
				"public": "pub-instance-test-primary-yrXqiqaQ-9c72733.dbaas-staging.bizflycloud.vn"
			},
			"addresses": {
				"private": [
					{
						"ip_address": "10.20.1.201:3306",
						"network_name": "priv_vctest_devcs_tri19@vccloud.vn"
					}
				],
				"public": [
					{
						"ip_address": "45.124.94.58:3306",
						"network_name": "public@network"
					}
				]
			}
		},
		{
			"id": "48cb294a-cc27-4e1a-b5fb-a2cd57bf5fcd",
			"name": "instance-test-secondary-qxxsJMGf",
			"availability_zone": "HN1",
			"role": "secondary",
			"dns": {
				"private": null,
				"public": null
			},
			"addresses": {
				"private": null,
				"public": null
			}
		},
		{
			"id": "f65c5f3f-8ab3-4394-a160-09303e9e6fd6",
			"name": "instance-4-replica-dXuquJDX",
			"availability_zone": "HN1",
			"role": "replica",
			"dns": {
				"private": "priv-instance-4-replica-UZxg0DJ1-9c72733.dbaas-staging.bizflycloud.vn",
				"public": "pub-instance-4-replica-UZxg0DJ1-9c72733.dbaas-staging.bizflycloud.vn"
			},
			"addresses": {
				"private": [
					{
						"ip_address": "10.20.1.112:3306",
						"network_name": "priv_vctest_devcs_tri19@vccloud.vn"
					}
				],
				"public": [
					{
						"ip_address": "45.124.94.185:3306",
						"network_name": "public@network"
					}
				]
			}
		},
		{
			"id": "8dbaefc6-0ef5-46b1-b6e4-1ce449da0d14",
			"name": "instance-4-replica-seeTl4gB",
			"availability_zone": "HN1",
			"role": "replica",
			"dns": {
				"private": "priv-instance-4-replica-B7nCFXf6-9c72733.dbaas-staging.bizflycloud.vn",
				"public": "pub-instance-4-replica-B7nCFXf6-9c72733.dbaas-staging.bizflycloud.vn"
			},
			"addresses": {
				"private": [
					{
						"ip_address": "10.20.1.202:3306",
						"network_name": "priv_vctest_devcs_tri19@vccloud.vn"
					}
				],
				"public": [
					{
						"ip_address": "45.124.94.244:3306",
						"network_name": "public@network"
					}
				]
			}
		}
	]
}`
		_, _ = fmt.Fprint(w, resp)
	})

	DBInstance, err := client.CloudDatabase.Instances().Get(ctx, "9c727335-4b53-44c3-866b-60ae502b0a3f")
	require.NoError(t, err)
	assert.Equal(t, "9c727335-4b53-44c3-866b-60ae502b0a3f", DBInstance.ID)
	assert.Equal(t, "instance-test", DBInstance.Name)
	assert.Equal(t, "Instance", DBInstance.Description)
	assert.Equal(t, true, DBInstance.PublicAccess)
	assert.Equal(t, "ee988cc3-bb30-4aaf-9837-e90a34f60d37", DBInstance.Datastore.ID)
	assert.Equal(t, "173d0dc3-46ac-40e9-bd01-bce21db0fd28", DBInstance.Nodes[0].ID)
	assert.Equal(t, "48cb294a-cc27-4e1a-b5fb-a2cd57bf5fcd", DBInstance.Nodes[1].ID)
	assert.Equal(t, "f65c5f3f-8ab3-4394-a160-09303e9e6fd6", DBInstance.Nodes[2].ID)
}

func Test_CloudDatabaseInstance_Create(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.DatabaseURL(cloudDatabaseInstancesResourcePath), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var instance *CloudDatabaseInstanceCreate
		require.NoError(t, json.NewDecoder(r.Body).Decode(&instance))
		assert.Equal(t, "instance_test", instance.Name)
		assert.Equal(t, "1c_2g", instance.FlavorName)
		assert.Equal(t, "HN1", instance.AvailabilityZone)
		assert.Equal(t, "ee988cc3-bb30-4aaf-9837-e90a34f60d37", instance.Datastore.VersionID)
		resp := `{"Result": "ok"}`
		_, _ = fmt.Fprint(w, resp)
	})
	dbr := &CloudDatabaseInstanceCreate{
		Name:       "instance_test",
		FlavorName: "1c_2g",
		VolumeSize: 10,
		Datastore: CloudDatabaseDatastore{
			Type:      "MariaDB",
			VersionID: "ee988cc3-bb30-4aaf-9837-e90a34f60d37",
		},
		EnableFailover:   false,
		Networks:         []CloudDatabaseNetworks{{"489a1cb8-92f4-4393-9df5-3866c025bcac"}},
		PublicAccess:     false,
		AvailabilityZone: "HN1",
		AutoScaling: &CloudDatabaseAutoScaling{
			Enable: true,
			Volume: CloudDatabaseAutoScalingVolume{
				Threshold: 80,
				Limited:   2000,
			},
		},
	}
	_, err := client.CloudDatabase.Instances().Create(ctx, dbr)
	require.NoError(t, err)
}

func Test_CloudDatabaseInstance_Action(t *testing.T) {
	setup()
	defer teardown()
	var dbc cloudDatabaseInstances
	mux.HandleFunc(testlib.DatabaseURL(dbc.resourceActionPath("9c727335-4b53-44c3-866b-60ae502b0a3f")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		resp := `{"Result": "ok"}`
		_, _ = fmt.Fprint(w, resp)
	})

	_, err := client.CloudDatabase.Instances().Action(ctx, "9c727335-4b53-44c3-866b-60ae502b0a3f", &CloudDatabaseAction{
		Action:  "resize_volume",
		NewSize: 20,
	})
	require.NoError(t, err)
}

func Test_CloudDatabaseInstance_ResizeFlavor(t *testing.T) {
	setup()
	defer teardown()
	var dbc cloudDatabaseInstances
	mux.HandleFunc(testlib.DatabaseURL(dbc.resourceActionPath("9c727335-4b53-44c3-866b-60ae502b0a3f")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		resp := `{"Result": "ok"}`
		_, _ = fmt.Fprint(w, resp)
	})

	_, err := client.CloudDatabase.Instances().ResizeFlavor(ctx, "9c727335-4b53-44c3-866b-60ae502b0a3f", CloudDatabaseDatastore{}, "instanceType", "1c_2g")
	require.NoError(t, err)
}

func Test_CloudDatabaseInstance_ResizeVolume(t *testing.T) {
	setup()
	defer teardown()
	var dbc cloudDatabaseInstances
	mux.HandleFunc(testlib.DatabaseURL(dbc.resourceActionPath("9c727335-4b53-44c3-866b-60ae502b0a3f")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		resp := `{"Result": "ok"}`
		_, _ = fmt.Fprint(w, resp)
	})

	_, err := client.CloudDatabase.Instances().ResizeVolume(ctx, "9c727335-4b53-44c3-866b-60ae502b0a3f", CloudDatabaseDatastore{}, "instanceType", 20)
	require.NoError(t, err)
}

func Test_CloudDatabaseInstance_Delete(t *testing.T) {
	setup()
	defer teardown()
	var dbc cloudDatabaseInstances
	mux.HandleFunc(testlib.DatabaseURL(dbc.resourcePath("9c727335-4b53-44c3-866b-60ae502b0a3f")), func(w http.ResponseWriter, r *http.Request) {

		resp := `{"Result": "ok"}`
		_, _ = fmt.Fprint(w, resp)
	})

	_, err := client.CloudDatabase.Instances().Delete(ctx, "9c727335-4b53-44c3-866b-60ae502b0a3f", &CloudDatabaseDelete{})
	require.NoError(t, err)
}

func Test_CloudDatabaseNode_List(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.DatabaseURL(cloudDatabaseNodesResourcePath), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "_meta": {
        "results_count": 2
    },
    "nodes": [
        {
            "id": "173d0dc3-46ac-40e9-bd01-bce21db0fd28",
            "name": "instance-1-primary-qxxsJMGf",
            "status": "ACTIVE",
            "flavor": "1c_1g",
            "datastore": {
                "type": "MariaDB",
                "name": "10.4",
                "id": "ee988cc3-bb30-4aaf-9837-e90a34f60d37"
            },
            "region_name": "HaNoi",
            "availability_zone": "HN1",
            "volume": {
                "size": 30,
                "used": 0.23
            },
            "created_at": "2021-07-16T09:59:33.000000Z",
            "dns": {
                "private": "priv-instance-4-primary-yrXqiqaQ-9c72733.dbaas-staging.bizflycloud.vn",
                "public": "pub-instance-4-primary-yrXqiqaQ-9c72733.dbaas-staging.bizflycloud.vn"
            },
            "addresses": {
                "private": [
                    {
                        "ip_address": "10.20.1.201:3306",
                        "network_name": "priv_vctest_devcs_tri19@vccloud.vn"
                    }
                ],
                "public": [
                    {
                        "ip_address": "45.124.94.58:3306",
                        "network_name": "public@network"
                    }
                ]
            },
            "meta_data": {
                "failover": true
            },
            "role": "primary",
            "instance_id": "9c727335-4b53-44c3-866b-60ae502b0a3f",
            "description": "Primary Node",
            "replicas": [
                {
                    "id": "48cb294a-cc27-4e1a-b5fb-a2cd57bf5fcd",
                    "name": "instance-1-secondary-qxxsJMGf",
                    "availability_zone": "HN1",
                    "role": "secondary",
                    "dns": {
                        "private": null,
                        "public": null
                    },
                    "addresses": {
                        "private": null,
                        "public": null
                    }
                },
                {
                    "id": "8dbaefc6-0ef5-46b1-b6e4-1ce449da0d14",
                    "name": "instance-1-replica-seeTl4gB",
                    "availability_zone": "HN1",
                    "role": "replica",
                    "dns": {
                        "private": "priv-instance-4-replica-B7nCFXf6-9c72733.dbaas-staging.bizflycloud.vn",
                        "public": "pub-instance-4-replica-B7nCFXf6-9c72733.dbaas-staging.bizflycloud.vn"
                    },
                    "addresses": {
                        "private": [
                            {
                                "ip_address": "10.20.1.202:3306",
                                "network_name": "priv_vctest_devcs_tri19@vccloud.vn"
                            }
                        ],
                        "public": [
                            {
                                "ip_address": "45.124.94.244:3306",
                                "network_name": "public@network"
                            }
                        ]
                    }
                },
                {
                    "id": "f65c5f3f-8ab3-4394-a160-09303e9e6fd6",
                    "name": "instance-1-replica-dXuquJDX",
                    "availability_zone": "HN1",
                    "role": "replica",
                    "dns": {
                        "private": "priv-instance-4-replica-UZxg0DJ1-9c72733.dbaas-staging.bizflycloud.vn",
                        "public": "pub-instance-4-replica-UZxg0DJ1-9c72733.dbaas-staging.bizflycloud.vn"
                    },
                    "addresses": {
                        "private": [
                            {
                                "ip_address": "10.20.1.112:3306",
                                "network_name": "priv_vctest_devcs_tri19@vccloud.vn"
                            }
                        ],
                        "public": [
                            {
                                "ip_address": "45.124.94.185:3306",
                                "network_name": "public@network"
                            }
                        ]
                    }
                }
            ]
        },
        {
            "id": "67f7a342-72ea-4320-aaff-b86f19355674",
            "name": "instance-2-primary-addTfkgC",
            "status": "ACTIVE",
            "flavor": "1c_1g",
            "datastore": null,
            "region_name": "HaNoi",
            "availability_zone": "HN1",
            "volume": {
                "size": 10,
                "used": 0.22
            },
            "created_at": "2021-07-16T09:43:16.000000Z",
            "dns": {
                "private": "priv-instance-2-primary-jBM60XMK-6808107.dbaas-staging.bizflycloud.vn",
                "public": "pub-instance-2-primary-jBM60XMK-6808107.dbaas-staging.bizflycloud.vn"
            },
            "addresses": {
                "private": [
                    {
                        "ip_address": "10.20.1.40:3306",
                        "network_name": "priv_vctest_devcs_tri19@vccloud.vn"
                    }
                ],
                "public": [
                    {
                        "ip_address": "45.124.94.64:3306",
                        "network_name": "public@network"
                    }
                ]
            },
            "meta_data": {
                "failover": true
            },
            "role": "primary",
            "instance_id": "68081071-fde7-4738-942a-c690a7b46582",
            "description": "Primary Node",
            "replicas": [
                {
                    "id": "6eb7d3b3-1b83-4db7-9d53-a242b2dc50e2",
                    "name": "instance-2-secondary-addTfkgC",
                    "availability_zone": "HN1",
                    "role": "secondary",
                    "dns": {
                        "private": null,
                        "public": null
                    },
                    "addresses": {
                        "private": null,
                        "public": null
                    }
                }
            ]
        }
	]
}`
		_, _ = fmt.Fprint(w, resp)
	})

	DBNodes, err := client.CloudDatabase.Nodes().List(ctx, &CloudDatabaseListOption{})
	require.NoError(t, err)
	DBNode := DBNodes[0]
	assert.Equal(t, "173d0dc3-46ac-40e9-bd01-bce21db0fd28", DBNode.ID)
	assert.Equal(t, "instance-1-primary-qxxsJMGf", DBNode.Name)
	assert.Equal(t, "2021-07-16T09:59:33.000000Z", DBNode.CreatedAt)
	assert.Equal(t, "primary", DBNode.Role)
	assert.Equal(t, "9c727335-4b53-44c3-866b-60ae502b0a3f", DBNode.InstanceID)
	assert.Equal(t, "48cb294a-cc27-4e1a-b5fb-a2cd57bf5fcd", DBNode.Replicas[0].ID)
	assert.Equal(t, "8dbaefc6-0ef5-46b1-b6e4-1ce449da0d14", DBNode.Replicas[1].ID)
}

func Test_CloudDatabaseNode_ListBackups(t *testing.T) {
	setup()
	defer teardown()
	var no *cloudDatabaseNodes
	mux.HandleFunc(testlib.DatabaseURL(no.resourceTypePath("ef317a9f-c705-4ea2-8fcb-c18450effb6c", backups)), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "_meta": {
        "page": 1,
        "results_per_page": 5,
        "filter": {
            "name": "",
            "start_time": "2021-01-01 00:00:00",
            "end_time": "2021-02-04 10:19:00"
        },
        "results_count": 2
    },
    "backups": [
        {
            "id": "7393cf38-5eb6-45e7-9ea9-f52dd3497cab",
            "name": "backup_name",
            "description": "MariaDB1 (120f585f-5aec-400b-a715-16c0b3c6a9d1)",
            "node_id": "120f585f-5aec-400b-a715-16c0b3c6a9d1",
            "created": "2021-02-04T10:18:47",
            "updated": "2021-02-04T10:18:52",
            "size": 0.21,
            "status": "COMPLETED",
            "parent_id": null,
            "project_id": "b746214aafba4923a5fbf9478ea64474",
            "datastore": {
                "type": "MariaDB",
                "version": "10.4",
                "version_id": "ee988cc3-bb30-4aaf-9837-e90a34f60d37"
            },
            "type": "manual",
            "meta_data": {
                "role": "primary"
            }
        },
        {
            "id": "75460fd5-3ce8-4f4a-a680-73a0d4914105",
            "name": "backup_name_2",
            "description": "MariaDB1 (120f585f-5aec-400b-a715-16c0b3c6a9d1)",
            "node_id": "120f585f-5aec-400b-a715-16c0b3c6a9d1",
            "created": "2021-02-04T10:18:34",
            "updated": "2021-02-04T10:18:38",
            "size": 0.21,
            "status": "COMPLETED",
            "parent_id": null,
            "project_id": "b746214aafba4923a5fbf9478ea64474",
            "datastore": {
                "type": "MariaDB",
                "version": "10.4",
                "version_id": "ee988cc3-bb30-4aaf-9837-e90a34f60d37"
            },
            "type": "manual",
            "meta_data": {
                "role": "replica"
            }
        }
    ]
}`
		_, _ = fmt.Fprint(w, resp)
	})

	DBBackups, err := client.CloudDatabase.Nodes().ListBackups(ctx, "ef317a9f-c705-4ea2-8fcb-c18450effb6c", &CloudDatabaseListOption{})
	require.NoError(t, err)
	DBBackup := DBBackups[0]
	assert.Equal(t, "7393cf38-5eb6-45e7-9ea9-f52dd3497cab", DBBackup.ID)
	assert.Equal(t, "backup_name", DBBackup.Name)
	assert.Equal(t, "120f585f-5aec-400b-a715-16c0b3c6a9d1", DBBackup.NodeID)
	assert.Equal(t, "b746214aafba4923a5fbf9478ea64474", DBBackup.ProjectID)
}

func Test_CloudDatabaseNode_ListSchedules(t *testing.T) {
	setup()
	defer teardown()
	var no *cloudDatabaseNodes
	mux.HandleFunc(testlib.DatabaseURL(no.resourceTypePath("a24509d4-cec1-4055-a476-676009140a3d", schedules)), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "_meta": {
        "page": 1,
        "results_per_page": 10,
        "filter": {
            "name": ""
        },
        "results_count": 2
    },
    "schedules": [
        {
            "id": "6ed3d8bd-c358-48f6-8890-20ed4e1fbfce",
            "name": "schedule_name",
            "pattern": "0 4,5,6 * * 1,2",
            "instance_id": "a24509d4-cec1-4055-a476-676009140a3d",
            "next_execution_time": "16:00 06-03-2021",
            "limit_backup": 3,
            "node_id": "2b138c69-be63-49f5-8b66-0180b7590af3",
            "node_name": "test4-member-ryv5RPgV"
        },
        {
            "id": "6ed3d8bd-c358-48f6-8890-20ed4e1fbfce",
            "name": "schedule_name_2",
            "pattern": "5 4 15,16 * *",
            "instance_id": "a24509d4-cec1-4055-a476-676009140a3d",
            "next_execution_time": "16:00 06-05-2021",
            "limit_backup": 2,
            "node_id": "2b138c69-be63-49f5-8b66-0180b7590af3",
            "node_name": "instance_name-member-ryv5RPgV"
        }
    ]
}`
		_, _ = fmt.Fprint(w, resp)
	})

	DBSchedules, err := client.CloudDatabase.Nodes().ListBackupSchedules(ctx, "a24509d4-cec1-4055-a476-676009140a3d", &CloudDatabaseListOption{})
	require.NoError(t, err)
	DBSchedule := DBSchedules[0]
	assert.Equal(t, "6ed3d8bd-c358-48f6-8890-20ed4e1fbfce", DBSchedule.ID)
	assert.Equal(t, "schedule_name", DBSchedule.Name)
	assert.Equal(t, "2b138c69-be63-49f5-8b66-0180b7590af3", DBSchedule.NodeID)
	assert.Equal(t, "a24509d4-cec1-4055-a476-676009140a3d", DBSchedule.InstanceID)
}

func Test_CloudDatabaseNode_Get(t *testing.T) {
	setup()
	defer teardown()
	var dbc cloudDatabaseNodes
	mux.HandleFunc(testlib.DatabaseURL(dbc.resourcePath("411e212d-0913-47a5-833b-e7a7332fcb01")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "id": "411e212d-0913-47a5-833b-e7a7332fcb01",
    "name": "instance-8-primary-reWtenvi",
    "status": "ACTIVE",
    "flavor": "1c_1g",
    "datastore": {
        "type": "MariaDB",
        "name": "10.4",
        "id": "ee988cc3-bb30-4aaf-9837-e90a34f60d37"
    },
    "region_name": "HaNoi",
    "availability_zone": "HN1",
    "volume": {
        "size": 10,
        "used": 0.22
    },
    "created_at": "2021-07-19T07:53:14.000000Z",
    "dns": {},
    "addresses": {},
    "meta_data": {
        "failover": true
    },
    "role": "primary",
    "instance_id": "415691db-b74d-4908-81ad-0d40a6e2900e",
    "description": "Primary Node"
}`
		_, _ = fmt.Fprint(w, resp)
	})

	DBNode, err := client.CloudDatabase.Nodes().Get(ctx, "411e212d-0913-47a5-833b-e7a7332fcb01")
	require.NoError(t, err)
	assert.Equal(t, "411e212d-0913-47a5-833b-e7a7332fcb01", DBNode.ID)
	assert.Equal(t, "instance-8-primary-reWtenvi", DBNode.Name)
	assert.Equal(t, "ACTIVE", DBNode.Status)
	assert.Equal(t, "2021-07-19T07:53:14.000000Z", DBNode.CreatedAt)
	assert.Equal(t, "415691db-b74d-4908-81ad-0d40a6e2900e", DBNode.InstanceID)
	assert.Equal(t, "ee988cc3-bb30-4aaf-9837-e90a34f60d37", DBNode.Datastore.ID)
	assert.Equal(t, "HaNoi", DBNode.RegionName)
}

func Test_CloudDatabaseNode_Create(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.DatabaseURL(cloudDatabaseNodesResourcePath), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var payload *CloudDatabaseNodeCreate
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "33a7577b-1955-492f-b02b-165a5f264f39", payload.ReplicaOf)
		assert.Equal(t, "secondary", payload.Role)
		resp := `{"Result": "ok"}`
		_, _ = fmt.Fprint(w, resp)
	})
	dbr := &CloudDatabaseNodeCreate{
		ReplicaOf: "33a7577b-1955-492f-b02b-165a5f264f39",
		Role:      "secondary",
		Secondaries: &CloudDatabaseReplicaNodeCreate{
			Configurations: CloudDatabaseReplicasConfiguration{
				AvailabilityZone: "HN1",
				Region:           "HaNoi",
			},
			Quantity: 1,
		},
	}
	_, err := client.CloudDatabase.Nodes().Create(ctx, dbr)
	require.NoError(t, err)
}

func Test_CloudDatabaseNode_Action(t *testing.T) {
	setup()
	defer teardown()
	var no cloudDatabaseNodes
	mux.HandleFunc(testlib.DatabaseURL(no.resourceActionPath("9c727335-4b53-44c3-866b-60ae502b0a3f")), func(w http.ResponseWriter, r *http.Request) {

		resp := `{"Result": "ok"}`
		_, _ = fmt.Fprint(w, resp)
	})

	_, err := client.CloudDatabase.Nodes().Action(ctx, "9c727335-4b53-44c3-866b-60ae502b0a3f", &CloudDatabaseAction{
		Action:     "resize",
		FlavorName: "1c_2g",
	})
	require.NoError(t, err)
}

func Test_CloudDatabaseNode_ResizeFlavor(t *testing.T) {
	setup()
	defer teardown()
	var no cloudDatabaseNodes
	mux.HandleFunc(testlib.DatabaseURL(no.resourceActionPath("9c727335-4b53-44c3-866b-60ae502b0a3f")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		resp := `{"Result": "ok"}`
		_, _ = fmt.Fprint(w, resp)
	})

	_, err := client.CloudDatabase.Nodes().ResizeFlavor(ctx, "9c727335-4b53-44c3-866b-60ae502b0a3f", "1c_2g")
	require.NoError(t, err)
}

func Test_CloudDatabaseNode_ResizeVolume(t *testing.T) {
	setup()
	defer teardown()
	var no cloudDatabaseNodes
	mux.HandleFunc(testlib.DatabaseURL(no.resourceActionPath("9c727335-4b53-44c3-866b-60ae502b0a3f")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		resp := `{"Result": "ok"}`
		_, _ = fmt.Fprint(w, resp)
	})

	_, err := client.CloudDatabase.Nodes().ResizeVolume(ctx, "9c727335-4b53-44c3-866b-60ae502b0a3f", 20)
	require.NoError(t, err)
}

func Test_CloudDatabaseNode_Restart(t *testing.T) {
	setup()
	defer teardown()
	var no cloudDatabaseNodes
	mux.HandleFunc(testlib.DatabaseURL(no.resourceActionPath("9c727335-4b53-44c3-866b-60ae502b0a3f")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		resp := `{"Result": "ok"}`
		_, _ = fmt.Fprint(w, resp)
	})

	_, err := client.CloudDatabase.Nodes().Restart(ctx, "9c727335-4b53-44c3-866b-60ae502b0a3f")
	require.NoError(t, err)
}

func Test_CloudDatabaseNode_DetachReplica(t *testing.T) {
	setup()
	defer teardown()
	var no cloudDatabaseNodes
	mux.HandleFunc(testlib.DatabaseURL(no.resourceActionPath("9c727335-4b53-44c3-866b-60ae502b0a3f")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		resp := `{"Result": "ok"}`
		_, _ = fmt.Fprint(w, resp)
	})

	_, err := client.CloudDatabase.Nodes().DetachReplica(ctx, "9c727335-4b53-44c3-866b-60ae502b0a3f")
	require.NoError(t, err)
}

func Test_CloudDatabaseNode_EnableRoot(t *testing.T) {
	setup()
	defer teardown()
	var no cloudDatabaseNodes
	mux.HandleFunc(testlib.DatabaseURL(no.resourceActionPath("9c727335-4b53-44c3-866b-60ae502b0a3f")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		resp := `{"Result": "ok"}`
		_, _ = fmt.Fprint(w, resp)
	})

	_, err := client.CloudDatabase.Nodes().EnableRoot(ctx, "9c727335-4b53-44c3-866b-60ae502b0a3f")
	require.NoError(t, err)
}

func Test_CloudDatabaseNode_Delete(t *testing.T) {
	setup()
	defer teardown()
	var dbc cloudDatabaseNodes
	mux.HandleFunc(testlib.DatabaseURL(dbc.resourcePath("9c727335-4b53-44c3-866b-60ae502b0a3f")), func(w http.ResponseWriter, r *http.Request) {

		resp := `{"Result": "ok"}`
		_, _ = fmt.Fprint(w, resp)
	})

	_, err := client.CloudDatabase.Nodes().Delete(ctx, "9c727335-4b53-44c3-866b-60ae502b0a3f", &CloudDatabaseDelete{
		PurgeBackup:     false,
		PurgeAutobackup: false,
	})
	require.NoError(t, err)
}

func Test_CloudDatabaseConfiguration_List(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.DatabaseURL(cloudDatabaseConfigurationsResourcePath), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "_meta": {
        "page": 1,
        "results_per_page": 10,
        "results_count": 3,
        "filter": {}
    },
    "configurations": [
        {
            "name": "test1",
            "id": "27e7ad85-9966-40c5-b2e8-090789144412",
            "created_at": "2021-07-16T10:02:59.000000Z",
            "datastore": {
                "type": "MariaDB",
                "name": "10.4",
                "id": "ee988cc3-bb30-4aaf-9837-e90a34f60d37"
            },
            "node_count": 1
        },
        {
            "name": "test2",
            "id": "51e72cd4-9042-4aed-817a-62c6fc012e18",
            "created_at": "2021-07-19T10:15:04.000000Z",
            "datastore": {
                "type": "MariaDB",
                "name": "10.4",
                "id": "ee988cc3-bb30-4aaf-9837-e90a34f60d37"
            },
            "node_count": 2
        },
        {
            "name": "test3",
            "id": "d062e0c1-4906-43cc-84bc-8d3f2be5961c",
            "created_at": "2021-07-16T11:27:51.000000Z",
            "datastore": {
                "type": "MariaDB",
                "name": "10.4",
                "id": "ee988cc3-bb30-4aaf-9837-e90a34f60d37"
            },
            "node_count": 0
        }
    ]
}`
		_, _ = fmt.Fprint(w, resp)
	})

	DBConfigurations, err := client.CloudDatabase.Configurations().List(ctx, &CloudDatabaseListOption{})
	require.NoError(t, err)
	DBConfiguration := DBConfigurations[0]
	assert.Equal(t, "27e7ad85-9966-40c5-b2e8-090789144412", DBConfiguration.ID)
	assert.Equal(t, "test1", DBConfiguration.Name)
	assert.Equal(t, "2021-07-16T10:02:59.000000Z", DBConfiguration.CreatedAt)
	assert.Equal(t, "ee988cc3-bb30-4aaf-9837-e90a34f60d37", DBConfiguration.Datastore.ID)
}

func Test_CloudDatabaseConfiguration_Get(t *testing.T) {
	setup()
	defer teardown()
	var dbc cloudDatabaseConfigurations
	mux.HandleFunc(testlib.DatabaseURL(dbc.resourcePath("27e7ad85-9966-40c5-b2e8-090789144412")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "name": "test2",
    "id": "27e7ad85-9966-40c5-b2e8-090789144412",
    "created_at": "2021-07-16T10:02:59.000000Z",
    "datastore": {
        "type": "MariaDB",
        "name": "10.4",
        "id": "ee988cc3-bb30-4aaf-9837-e90a34f60d37"
    },
    "node_count": 1,
    "values": {
        "character_set_server": "armscii8",
        "max_connections": 100
    },
    "nodes": [
        {
            "name": "instance-2-primary-QSN92LKb",
            "id": "e810c3f7-2469-4d9a-8062-5b41512791c4"
        }
    ]
}`
		_, _ = fmt.Fprint(w, resp)
	})

	DBConfiguration, err := client.CloudDatabase.Configurations().Get(ctx, "27e7ad85-9966-40c5-b2e8-090789144412")
	require.NoError(t, err)
	assert.Equal(t, "27e7ad85-9966-40c5-b2e8-090789144412", DBConfiguration.ID)
	assert.Equal(t, "test2", DBConfiguration.Name)
	assert.Equal(t, "2021-07-16T10:02:59.000000Z", DBConfiguration.CreatedAt)
	assert.Equal(t, 1, DBConfiguration.NodeCount)
}

func Test_CloudDatabaseConfiguration_Create(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.DatabaseURL(cloudDatabaseConfigurationsResourcePath), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var payload *CloudDatabaseConfigurationCreate
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "test_cfg", payload.Name)
		assert.Equal(t, "MariaDB", payload.Datastore.Type)
		resp := `{"Result": "ok"}`
		_, _ = fmt.Fprint(w, resp)
	})
	dcr := &CloudDatabaseConfigurationCreate{
		Name: "test_cfg",
		Datastore: CloudDatabaseDatastore{
			Type:      "MariaDB",
			VersionID: "ee988cc3-bb30-4aaf-9837-e90a34f60d37",
		},
		Parameters: map[string]interface{}{
			"character_set_server": "armscii8",
			"max_connections":      200,
		},
	}
	_, err := client.CloudDatabase.Configurations().Create(ctx, dcr)
	require.NoError(t, err)
}

func Test_CloudDatabaseConfiguration_Update(t *testing.T) {
	setup()
	defer teardown()
	var dbc cloudDatabaseConfigurations
	mux.HandleFunc(testlib.DatabaseURL(dbc.resourcePath("3b2de5a3-bcd1-4972-a34c-dc48sdsd")), func(w http.ResponseWriter, r *http.Request) {
		resp := `{"Result": "ok"}`
		_, _ = fmt.Fprint(w, resp)
	})

	_, err := client.CloudDatabase.Configurations().Update(ctx, "3b2de5a3-bcd1-4972-a34c-dc48sdsd", &CloudDatabaseConfigurationUpdate{
		Parameters: map[string]interface{}{
			"character_set_server": "armscii8",
			"max_connections":      100,
		},
	})
	require.NoError(t, err)
}

func Test_CloudDatabaseConfiguration_Action(t *testing.T) {
	setup()
	defer teardown()
	var dbc cloudDatabaseConfigurations
	mux.HandleFunc(testlib.DatabaseURL(dbc.resourceActionPath("9c727335-4b53-44c3-866b-60ae502b0a3f", "3b2de5a3-bcd1-4972-a34c-dc48sdsd")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		resp := `{"Result": "ok"}`
		_, _ = fmt.Fprint(w, resp)
	})

	_, err := client.CloudDatabase.Configurations().Action(ctx, "9c727335-4b53-44c3-866b-60ae502b0a3f", "3b2de5a3-bcd1-4972-a34c-dc48sdsd", &CloudDatabaseAction{
		Action:    "attach",
		ActionAll: true,
	})
	require.NoError(t, err)
}

func Test_CloudDatabaseConfiguration_Attach(t *testing.T) {
	setup()
	defer teardown()
	var dbc cloudDatabaseConfigurations
	mux.HandleFunc(testlib.DatabaseURL(dbc.resourceActionPath("9c727335-4b53-44c3-866b-60ae502b0a3f", "3b2de5a3-bcd1-4972-a34c-dc48sdsd")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		resp := `{"Result": "ok"}`
		_, _ = fmt.Fprint(w, resp)
	})

	_, err := client.CloudDatabase.Configurations().Attach(ctx, "9c727335-4b53-44c3-866b-60ae502b0a3f", "3b2de5a3-bcd1-4972-a34c-dc48sdsd", true)
	require.NoError(t, err)
}

func Test_CloudDatabaseConfiguration_Detach(t *testing.T) {
	setup()
	defer teardown()
	var dbc cloudDatabaseConfigurations
	mux.HandleFunc(testlib.DatabaseURL(dbc.resourceActionPath("9c727335-4b53-44c3-866b-60ae502b0a3f", "3b2de5a3-bcd1-4972-a34c-dc48sdsd")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		resp := `{"Result": "ok"}`
		_, _ = fmt.Fprint(w, resp)
	})

	_, err := client.CloudDatabase.Configurations().Detach(ctx, "9c727335-4b53-44c3-866b-60ae502b0a3f", "3b2de5a3-bcd1-4972-a34c-dc48sdsd", false)
	require.NoError(t, err)
}

func Test_CloudDatabaseConfiguration_Delete(t *testing.T) {
	setup()
	defer teardown()
	var dbc cloudDatabaseConfigurations
	mux.HandleFunc(testlib.DatabaseURL(dbc.resourcePath("3b2de5a3-bcd1-4972-a34c-dc48sdsd")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		resp := `{"Result": "ok"}`
		_, _ = fmt.Fprint(w, resp)
	})

	_, err := client.CloudDatabase.Configurations().Delete(ctx, "3b2de5a3-bcd1-4972-a34c-dc48sdsd")
	require.NoError(t, err)
}

func Test_CloudDatabaseBackup_List(t *testing.T) {
	setup()
	defer teardown()
	var bk cloudDatabaseBackups
	mux.HandleFunc(testlib.DatabaseURL(bk.resourceCreatePath("instances", "ef317a9f-c705-4ea2-8fcb-c18450effb6c")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "_meta": {
        "page": 1,
        "results_per_page": 5,
        "filter": {
            "name": "",
            "start_time": "2021-01-01 00:00:00",
            "end_time": "2021-02-04 10:19:00"
        },
        "results_count": 2
    },
    "backups": [
        {
            "id": "7393cf38-5eb6-45e7-9ea9-f52dd3497cab",
            "name": "backup_name",
            "description": "MariaDB1 (120f585f-5aec-400b-a715-16c0b3c6a9d1)",
            "node_id": "120f585f-5aec-400b-a715-16c0b3c6a9d1",
            "created": "2021-02-04T10:18:47",
            "updated": "2021-02-04T10:18:52",
            "size": 0.21,
            "status": "COMPLETED",
            "parent_id": null,
            "project_id": "b746214aafba4923a5fbf9478ea64474",
            "datastore": {
                "type": "MariaDB",
                "version": "10.4",
                "version_id": "ee988cc3-bb30-4aaf-9837-e90a34f60d37"
            },
            "type": "manual",
            "meta_data": {
                "role": "primary"
            }
        },
        {
            "id": "75460fd5-3ce8-4f4a-a680-73a0d4914105",
            "name": "backup_name_2",
            "description": "MariaDB1 (120f585f-5aec-400b-a715-16c0b3c6a9d1)",
            "node_id": "120f585f-5aec-400b-a715-16c0b3c6a9d1",
            "created": "2021-02-04T10:18:34",
            "updated": "2021-02-04T10:18:38",
            "size": 0.21,
            "status": "COMPLETED",
            "parent_id": null,
            "project_id": "b746214aafba4923a5fbf9478ea64474",
            "datastore": {
                "type": "MariaDB",
                "version": "10.4",
                "version_id": "ee988cc3-bb30-4aaf-9837-e90a34f60d37"
            },
            "type": "manual",
            "meta_data": {
                "role": "replica"
            }
        }
    ]
}`
		_, _ = fmt.Fprint(w, resp)
	})

	DBBackups, err := client.CloudDatabase.Backups().List(ctx, &CloudDatabaseBackupResource{
		ResourceType: "instances",
		ResourceID:   "ef317a9f-c705-4ea2-8fcb-c18450effb6c",
	}, &CloudDatabaseListOption{})
	require.NoError(t, err)
	DBBackup := DBBackups[0]
	assert.Equal(t, "7393cf38-5eb6-45e7-9ea9-f52dd3497cab", DBBackup.ID)
	assert.Equal(t, "backup_name", DBBackup.Name)
	assert.Equal(t, "120f585f-5aec-400b-a715-16c0b3c6a9d1", DBBackup.NodeID)
	assert.Equal(t, "b746214aafba4923a5fbf9478ea64474", DBBackup.ProjectID)
}

func Test_CloudDatabaseBackup_Get(t *testing.T) {
	setup()
	defer teardown()
	var dbc cloudDatabaseBackups
	mux.HandleFunc(testlib.DatabaseURL(dbc.resourcePath("62a3c4ae-0f53-4a79-8a3b-8d49fff7f0c3")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "id": "62a3c4ae-0f53-4a79-8a3b-8d49fff7f0c3",
    "name": "backup_name_2",
    "description": "instance_name-replica-member-QlvtfDBA (296dee42-3b44-4c3b-872f-5d7bfd28ec64)",
    "node_id": "296dee42-3b44-4c3b-872f-5d7bfd28ec64",
    "created": "2021-04-13T07:47:23.000000Z",
    "updated": "2021-04-13T07:47:34",
    "size": 0.21,
    "status": "COMPLETED",
    "parent_id": null,
    "project_id": "4e7bfc86923349b6bb3272945f46f5ae",
    "datastore": {
        "type": "MariaDB",
        "version": "10.4",
        "version_id": "ee988cc3-bb30-4aaf-9837-e90a34f60d37"
    },
    "type": "manual",
    "meta_data": {
        "role": "replica"
    }
}`
		_, _ = fmt.Fprint(w, resp)
	})

	DBBackup, err := client.CloudDatabase.Backups().Get(ctx, "62a3c4ae-0f53-4a79-8a3b-8d49fff7f0c3")
	require.NoError(t, err)
	assert.Equal(t, "62a3c4ae-0f53-4a79-8a3b-8d49fff7f0c3", DBBackup.ID)
	assert.Equal(t, "backup_name_2", DBBackup.Name)
	assert.Equal(t, "296dee42-3b44-4c3b-872f-5d7bfd28ec64", DBBackup.NodeID)
	assert.Equal(t, "4e7bfc86923349b6bb3272945f46f5ae", DBBackup.ProjectID)
}

func Test_CloudDatabaseBackup_Create(t *testing.T) {
	setup()
	defer teardown()
	var bk cloudDatabaseBackups
	mux.HandleFunc(testlib.DatabaseURL(bk.resourceCreatePath("instances", "e5454cb9-b23c-431f-84a7-12e3e0dec691")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var payload *CloudDatabaseBackupCreate
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "backup_name", payload.Name)
		assert.Equal(t, "edf80901-1168-4f7d-9c47-5b44ad62befa", payload.ParentID)
		assert.Equal(t, "e5454cb9-b23c-431f-84a7-12e3e0dec691", payload.NodeID)
		resp := `{"Result": "ok"}`
		_, _ = fmt.Fprint(w, resp)
	})
	bkr := &CloudDatabaseBackupCreate{
		Name:     "backup_name",
		ParentID: "edf80901-1168-4f7d-9c47-5b44ad62befa",
		NodeID:   "e5454cb9-b23c-431f-84a7-12e3e0dec691",
	}
	_, err := client.CloudDatabase.Backups().Create(ctx, "instances", "e5454cb9-b23c-431f-84a7-12e3e0dec691", bkr)
	require.NoError(t, err)
}

func Test_CloudDatabaseBackup_Delete(t *testing.T) {
	setup()
	defer teardown()
	var dbc cloudDatabaseBackups
	mux.HandleFunc(testlib.DatabaseURL(dbc.resourcePath("e5454cb9-b23c-431f-84a7-12e3e0dec691")), func(w http.ResponseWriter, r *http.Request) {

		resp := `{"Result": "ok"}`
		_, _ = fmt.Fprint(w, resp)
	})

	_, err := client.CloudDatabase.Backups().Delete(ctx, "e5454cb9-b23c-431f-84a7-12e3e0dec691")
	require.NoError(t, err)
}

func Test_CloudDatabaseBackupSchedule_List(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.DatabaseURL(cloudDatabaseBackupSchedulesResourcePath), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "_meta": {
        "page": 1,
        "results_per_page": 10,
        "filter": {
            "name": ""
        },
        "results_count": 2
    },
    "schedules": [
        {
            "id": "6ed3d8bd-c358-48f6-8890-20ed4e1fbfce",
            "name": "schedule_name",
            "pattern": "0 4,5,6 * * 1,2",
            "instance_id": "a24509d4-cec1-4055-a476-676009140a3d",
            "next_execution_time": "16:00 06-03-2021",
            "limit_backup": 3,
            "node_id": "2b138c69-be63-49f5-8b66-0180b7590af3",
            "node_name": "test4-member-ryv5RPgV"
        },
        {
            "id": "6ed3d8bd-c358-48f6-8890-20ed4e1fbfce",
            "name": "schedule_name_2",
            "pattern": "5 4 15,16 * *",
            "instance_id": "a24509d4-cec1-4055-a476-676009140a3d",
            "next_execution_time": "16:00 06-05-2021",
            "limit_backup": 2,
            "node_id": "2b138c69-be63-49f5-8b66-0180b7590af3",
            "node_name": "instance_name-member-ryv5RPgV"
        }
    ]
}`
		_, _ = fmt.Fprint(w, resp)
	})

	DBSchedules, err := client.CloudDatabase.BackupSchedules().List(ctx, &CloudDatabaseBackupScheduleListResourceOption{
		ResourceType: "instances",
		ResourceID:   "ef317a9f-c705-4ea2-8fcb-c18450effb6c",
		All:          true,
	}, &CloudDatabaseListOption{})
	require.NoError(t, err)
	DBSchedule := DBSchedules[0]
	assert.Equal(t, "6ed3d8bd-c358-48f6-8890-20ed4e1fbfce", DBSchedule.ID)
	assert.Equal(t, "schedule_name", DBSchedule.Name)
	assert.Equal(t, "2b138c69-be63-49f5-8b66-0180b7590af3", DBSchedule.NodeID)
	assert.Equal(t, "a24509d4-cec1-4055-a476-676009140a3d", DBSchedule.InstanceID)
}

func Test_CloudDatabaseBackupSchedule_ListBackups(t *testing.T) {
	setup()
	defer teardown()
	var sc *cloudDatabaseBackupSchedules
	mux.HandleFunc(testlib.DatabaseURL(sc.resourceBackupPath("ef317a9f-c705-4ea2-8fcb-c18450effb6c")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "_meta": {
        "page": 1,
        "results_per_page": 5,
        "filter": {
            "name": "",
            "start_time": "2021-01-01 00:00:00",
            "end_time": "2021-02-04 10:19:00"
        },
        "results_count": 2
    },
    "backups": [
        {
            "id": "7393cf38-5eb6-45e7-9ea9-f52dd3497cab",
            "name": "backup_name",
            "description": "MariaDB1 (120f585f-5aec-400b-a715-16c0b3c6a9d1)",
            "node_id": "120f585f-5aec-400b-a715-16c0b3c6a9d1",
            "created": "2021-02-04T10:18:47",
            "updated": "2021-02-04T10:18:52",
            "size": 0.21,
            "status": "COMPLETED",
            "parent_id": null,
            "project_id": "b746214aafba4923a5fbf9478ea64474",
            "datastore": {
                "type": "MariaDB",
                "version": "10.4",
                "version_id": "ee988cc3-bb30-4aaf-9837-e90a34f60d37"
            },
            "type": "manual",
            "meta_data": {
                "role": "primary"
            }
        },
        {
            "id": "75460fd5-3ce8-4f4a-a680-73a0d4914105",
            "name": "backup_name_2",
            "description": "MariaDB1 (120f585f-5aec-400b-a715-16c0b3c6a9d1)",
            "node_id": "120f585f-5aec-400b-a715-16c0b3c6a9d1",
            "created": "2021-02-04T10:18:34",
            "updated": "2021-02-04T10:18:38",
            "size": 0.21,
            "status": "COMPLETED",
            "parent_id": null,
            "project_id": "b746214aafba4923a5fbf9478ea64474",
            "datastore": {
                "type": "MariaDB",
                "version": "10.4",
                "version_id": "ee988cc3-bb30-4aaf-9837-e90a34f60d37"
            },
            "type": "manual",
            "meta_data": {
                "role": "replica"
            }
        }
    ]
}`
		_, _ = fmt.Fprint(w, resp)
	})

	DBBackups, err := client.CloudDatabase.BackupSchedules().ListBackups(ctx, "ef317a9f-c705-4ea2-8fcb-c18450effb6c", &CloudDatabaseListOption{})
	require.NoError(t, err)
	DBBackup := DBBackups[0]
	assert.Equal(t, "7393cf38-5eb6-45e7-9ea9-f52dd3497cab", DBBackup.ID)
	assert.Equal(t, "backup_name", DBBackup.Name)
	assert.Equal(t, "120f585f-5aec-400b-a715-16c0b3c6a9d1", DBBackup.NodeID)
	assert.Equal(t, "b746214aafba4923a5fbf9478ea64474", DBBackup.ProjectID)
}

func Test_CloudDatabaseBackupSchedule_Get(t *testing.T) {
	setup()
	defer teardown()
	var sc cloudDatabaseBackupSchedules
	mux.HandleFunc(testlib.DatabaseURL(sc.resourcePath("6ed3d8bd-c358-48f6-8890-20ed4e1fbfce")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "id": "6ed3d8bd-c358-48f6-8890-20ed4e1fbfce",
    "name": "schedule_name_2",
    "pattern": "5 4 15,16 * *",
    "instance_id": "a24509d4-cec1-4055-a476-676009140a3d",
    "next_execution_time": "16:00 06-05-2021",
    "limit_backup": 2,
    "node_id": "2b138c69-be63-49f5-8b66-0180b7590af3",
    "node_name": "instance_name-member-ryv5RPgV"
}`
		_, _ = fmt.Fprint(w, resp)
	})

	DBSchedule, err := client.CloudDatabase.BackupSchedules().Get(ctx, "6ed3d8bd-c358-48f6-8890-20ed4e1fbfce")
	require.NoError(t, err)
	assert.Equal(t, "6ed3d8bd-c358-48f6-8890-20ed4e1fbfce", DBSchedule.ID)
	assert.Equal(t, "schedule_name_2", DBSchedule.Name)
	assert.Equal(t, "2b138c69-be63-49f5-8b66-0180b7590af3", DBSchedule.NodeID)
	assert.Equal(t, "a24509d4-cec1-4055-a476-676009140a3d", DBSchedule.InstanceID)
}

func Test_CloudDatabaseBackupSchedule_Create(t *testing.T) {
	setup()
	defer teardown()
	var sc cloudDatabaseBackupSchedules
	mux.HandleFunc(testlib.DatabaseURL(sc.resourceCreatePath("nodes", "2b138c69-be63-49f5-8b66-0180b7590af3")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var payload *CloudDatabaseBackupScheduleCreate
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, "schedule_name", payload.Name)
		assert.Equal(t, 3, payload.LimitBackup)
		assert.Equal(t, "20,50 * * * *", payload.CronExpression)
		resp := `{"Result": "ok"}`
		_, _ = fmt.Fprint(w, resp)
	})
	bkr := &CloudDatabaseBackupScheduleCreate{
		Name:           "schedule_name",
		LimitBackup:    3,
		CronExpression: "20,50 * * * *",
	}
	_, err := client.CloudDatabase.BackupSchedules().Create(ctx, "2b138c69-be63-49f5-8b66-0180b7590af3", bkr)
	require.NoError(t, err)
}

func Test_CloudDatabaseBackupSchedule_Delete(t *testing.T) {
	setup()
	defer teardown()
	var sc cloudDatabaseBackupSchedules
	mux.HandleFunc(testlib.DatabaseURL(sc.resourcePath("2b138c69-be63-49f5-8b66-0180b7590af3")), func(w http.ResponseWriter, r *http.Request) {

		resp := `{"Result": "ok"}`
		_, _ = fmt.Fprint(w, resp)
	})

	_, err := client.CloudDatabase.BackupSchedules().Delete(ctx, "2b138c69-be63-49f5-8b66-0180b7590af3", &CloudDatabaseBackupScheduleDelete{PurgeBackup: true})
	require.NoError(t, err)
}

func Test_CloudDatabaseAutoscaling_Create(t *testing.T) {
	setup()
	defer teardown()
	var au cloudDatabaseAutoScalings
	mux.HandleFunc(testlib.DatabaseURL(au.resourcePath("2b138c69-be63-49f5-8b66-0180b7590af3")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var payload *CloudDatabaseAutoScaling
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, true, payload.Enable)
		assert.Equal(t, 80, payload.Volume.Threshold)
		assert.Equal(t, 2000, payload.Volume.Limited)
		resp := `{"Result": "ok"}`
		_, _ = fmt.Fprint(w, resp)
	})
	auc := &CloudDatabaseAutoScaling{
		Enable: true,
		Volume: CloudDatabaseAutoScalingVolume{
			Threshold: 80,
			Limited:   2000,
		},
	}
	_, err := client.CloudDatabase.AutoScalings().Create(ctx, "2b138c69-be63-49f5-8b66-0180b7590af3", auc)
	require.NoError(t, err)
}

func Test_CloudDatabaseAutoscaling_Update(t *testing.T) {
	setup()
	defer teardown()
	var au cloudDatabaseAutoScalings
	mux.HandleFunc(testlib.DatabaseURL(au.resourcePath("2b138c69-be63-49f5-8b66-0180b7590af3")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
		var payload *CloudDatabaseAutoScaling
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, true, payload.Enable)
		assert.Equal(t, 70, payload.Volume.Threshold)
		assert.Equal(t, 2000, payload.Volume.Limited)
		resp := `{"Result": "ok"}`
		_, _ = fmt.Fprint(w, resp)
	})
	aup := &CloudDatabaseAutoScaling{
		Enable: true,
		Volume: CloudDatabaseAutoScalingVolume{
			Threshold: 70,
			Limited:   2000,
		},
	}
	_, err := client.CloudDatabase.AutoScalings().Update(ctx, "2b138c69-be63-49f5-8b66-0180b7590af3", aup)
	require.NoError(t, err)
}

func Test_CloudDatabaseAutoscaling_Delete(t *testing.T) {
	setup()
	defer teardown()
	var sc cloudDatabaseAutoScalings
	mux.HandleFunc(testlib.DatabaseURL(sc.resourcePath("2b138c69-be63-49f5-8b66-0180b7590af3")), func(w http.ResponseWriter, r *http.Request) {

		resp := `{"Result": "ok"}`
		_, _ = fmt.Fprint(w, resp)
	})

	_, err := client.CloudDatabase.AutoScalings().Delete(ctx, "2b138c69-be63-49f5-8b66-0180b7590af3")
	require.NoError(t, err)
}

func Test_CloudDatabaseTrustedSource_Get(t *testing.T) {
	setup()
	defer teardown()
	var sc cloudDatabaseBackupSchedules
	mux.HandleFunc(testlib.DatabaseURL(sc.resourcePath("6ed3d8bd-c358-48f6-8890-20ed4e1fbfce")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "id": "6ed3d8bd-c358-48f6-8890-20ed4e1fbfce",
    "name": "schedule_name_2",
    "pattern": "5 4 15,16 * *",
    "instance_id": "a24509d4-cec1-4055-a476-676009140a3d",
    "next_execution_time": "16:00 06-05-2021",
    "limit_backup": 2,
    "node_id": "2b138c69-be63-49f5-8b66-0180b7590af3",
    "node_name": "instance_name-member-ryv5RPgV"
}`
		_, _ = fmt.Fprint(w, resp)
	})

	DBSchedule, err := client.CloudDatabase.BackupSchedules().Get(ctx, "6ed3d8bd-c358-48f6-8890-20ed4e1fbfce")
	require.NoError(t, err)
	assert.Equal(t, "6ed3d8bd-c358-48f6-8890-20ed4e1fbfce", DBSchedule.ID)
	assert.Equal(t, "schedule_name_2", DBSchedule.Name)
	assert.Equal(t, "2b138c69-be63-49f5-8b66-0180b7590af3", DBSchedule.NodeID)
	assert.Equal(t, "a24509d4-cec1-4055-a476-676009140a3d", DBSchedule.InstanceID)
}

func Test_CloudDatabaseTrustedSource_Update(t *testing.T) {
	setup()
	defer teardown()
	var fr cloudDatabaseTrustedSources
	mux.HandleFunc(testlib.DatabaseURL(fr.resourcePath("2b138c69-be63-49f5-8b66-0180b7590af3")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		var payload *CloudDatabaseTrustedSources
		require.NoError(t, json.NewDecoder(r.Body).Decode(&payload))
		assert.Equal(t, []string{"1.1.1.1", "2.2.2.2", "192.168.0.0/24"}, payload.TrustedSources)
		resp := `{"Result": "ok"}`
		_, _ = fmt.Fprint(w, resp)
	})
	fru := &CloudDatabaseTrustedSources{
		TrustedSources: []string{"1.1.1.1", "2.2.2.2", "192.168.0.0/24"},
	}
	_, err := client.CloudDatabase.TrustedSources().Update(ctx, "2b138c69-be63-49f5-8b66-0180b7590af3", fru)
	require.NoError(t, err)
}

func Test_CloudDatabaseTask_Get(t *testing.T) {
	setup()
	defer teardown()
	var sc cloudDatabaseTasks
	mux.HandleFunc(testlib.DatabaseURL(sc.resourcePath("6ed3d8bd-c358-48f6-8890-20ed4e1fbfce")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "ready": true,
    "result": {
        "progress": 100,
        "action": "delete_instance_success",
        "message": "Delete successful",
        "data": {}
    }
}`
		_, _ = fmt.Fprint(w, resp)
	})

	DBTask, err := client.CloudDatabase.Tasks().Get(ctx, "6ed3d8bd-c358-48f6-8890-20ed4e1fbfce")
	require.NoError(t, err)
	assert.Equal(t, true, DBTask.Ready)
	assert.Equal(t, 100, DBTask.Result.Progress)
	assert.Equal(t, "delete_instance_success", DBTask.Result.Action)
	assert.Equal(t, "Delete successful", DBTask.Result.Message)
}

func Test_CloudDatabaseEngine_Get(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.DatabaseURL(cloudDatabaseEnginesResourcePath), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "engines": [
        {
            "id": "fbf982a8-ac66-43ee-adfa-68474b11a617",
            "name": "MariaDB",
            "versions": [
                {
                    "id": "ee988cc3-bb30-4aaf-9837-e90a34f60d37",
                    "name": "10.4",
                    "version": "10.4"
                }
            ]
        }
    ]
}`
		_, _ = fmt.Fprint(w, resp)
	})

	Engines, err := client.CloudDatabase.Engines().List(ctx)
	require.NoError(t, err)
	assert.Equal(t, "fbf982a8-ac66-43ee-adfa-68474b11a617", Engines[0].ID)
	assert.Equal(t, "ee988cc3-bb30-4aaf-9837-e90a34f60d37", Engines[0].Versions[0].ID)
	assert.Equal(t, "10.4", Engines[0].Versions[0].Version)
}

func Test_CloudDatabaseEngineParameters_Get(t *testing.T) {
	setup()
	defer teardown()
	var ep *cloudDatabaseEngineParameters
	mux.HandleFunc(testlib.DatabaseURL(ep.resourcePath("MongoDB", "ee988cc3-bb30-4aaf-9837-e90a34f60d37")), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `
{
    "configuration_parameters": [
        {
            "name": "interactive_timeout",
            "datastore_version_id": "ed5e2cd6-bb2b-4f36-a431-8af6b5fd9a9d",
            "restart_required": false,
            "type": "integer",
            "max": 65535,
            "min": 1,
            "default": 28800,
            "description": "Số giây server đợi các hoạt động trên connection trước khi đóng"
        },
        {
            "name": "key_buffer_size",
            "datastore_version_id": "ed5e2cd6-bb2b-4f36-a431-8af6b5fd9a9d",
            "restart_required": false,
            "type": "integer",
            "max": 4294967295,
            "min": 8,
            "default": 209715200,
            "description": "Kích thước của index buffer được giữ trong bộ nhớ, ảnh hưởng đến tốc độ đọc index. Biến này chỉ có tác dụng khi sử dụng MyISAM"
        }
	]
}`
		_, _ = fmt.Fprint(w, resp)
	})

	DBEngineParameters, err := client.CloudDatabase.EngineParameters().Get(ctx, "MongoDB", "ee988cc3-bb30-4aaf-9837-e90a34f60d37")
	require.NoError(t, err)
	assert.Equal(t, 2, len(DBEngineParameters.Parameters))
}
