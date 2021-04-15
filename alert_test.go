// This file is part of gobizfly

package gobizfly

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/bizflycloud/gobizfly/testlib"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_alarms_List(t *testing.T) {
	setup()
	defer teardown()
	var a alarms
	mux.HandleFunc(testlib.CloudWatcherURL(a.resourcePath()), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `{
		    "_items": [{
		        "_updated": "Mon, 13 Jul 2020 10:16:44 GMT",
		        "comparison": {
		            "measurement": "iops",
		            "compare_type": ">=",
		            "value": 5000,
		            "range_time": 300
		        },
		        "enable": true,
		        "user_id": "fake-user-id",
		        "name": "iops300",
		        "receivers": [{
		            "receiver_id": "5ee9841f83019d000e8fb2c4",
		            "autoscale_cluster_name": "group-01",
		            "email_address": "example@domain.com",
		            "name": "telegram",
		            "slack_channel_name": "channel_name",
		            "sms_number": "fake-sms-number",
		            "telegram_chat_id": "fake-telegram-chat-id",
		            "webhook_url": "fake-webhook-url"
		        }],
		        "creator": "user",
		        "_deleted": false,
		        "alert_interval": 300,
		        "project_id": "fake-project-id",
		        "cluster_id": "a2018319-4c80-4a52-bfb1-6045b2b01536",
		        "cluster_name": "group-01",
		        "_links": {
		            "self": {
		                "href": "alarms/5f0c348c4c569c000d0549f0",
		                "title": "alarm"
		            }
		        },
		        "volumes": [{
		            "name": "bncvbnb-node-jtCyyabF_rootdisk",
		            "id": "a0a75b48-1175-491d-89d9-a5a3b631ad5b"
		        }],
		        "instances": [{
		            "name": "bncvbnb-node-jtCyyabF",
		            "id": "a0a75b48-1175-491d-89d9-491d117589d9"
		        }],
		        "load_balancers": [{
		            "load_balancer_id": "5c001a15-3867-42eb-9e0a-40ea8455688d",
		            "load_balancer_name": "fake-loadbalancer-name",
		            "target_id": "ec0e9614-9c89-4785-a56e-4b3dc4c4eb79",
		            "target_type": "frontend"
		        }],
		        "_created": "Mon, 13 Jul 2020 10:16:44 GMT",
		        "_id": "5f0c348c4c569c000d0549f0",
		        "hostname": "1.1.1.1",
		        "http_expected_code": 200,
		        "http_url": "https://google.com",
		        "http_headers": [{
		        	"key": "X-Auth-Token",
		        	"value": "fake-token"
		        }],
		        "resource_type": "volume"
		    }],
		    "_links": {
		        "self": {
		            "href": "alarms?max_results=10&sort=-_created",
		            "title": "alarms"
		        },
		        "parent": {
		            "href": "/",
		            "title": "home"
		        }
		    },
		    "_meta": {
		        "max_results": 10,
		        "total": 1,
		        "page": 1
		    }
		}`
		_, _ = fmt.Fprint(w, resp)
	})
	alarms, err := client.CloudWatcher.Alarms().List(ctx, nil)
	require.NoError(t, err)
	alarm := alarms[0]
	// receivers
	assert.Equal(t, "5ee9841f83019d000e8fb2c4", alarm.Receivers[0].ReceiverID)
	assert.Equal(t, "fake-telegram-chat-id", alarm.Receivers[0].TelegramChatID)
	assert.Equal(t, "group-01", alarm.Receivers[0].AutoscaleClusterName)
	assert.Equal(t, "example@domain.com", alarm.Receivers[0].EmailAddress)
	assert.Equal(t, "channel_name", alarm.Receivers[0].SlackChannelName)
	assert.Equal(t, "fake-webhook-url", alarm.Receivers[0].WebhookURL)
	assert.Equal(t, "fake-sms-number", alarm.Receivers[0].SMSNumber)

	// Volumes
	assert.Equal(t, "bncvbnb-node-jtCyyabF_rootdisk", alarm.Volumes[0].Name)
	assert.Equal(t, "a0a75b48-1175-491d-89d9-a5a3b631ad5b", alarm.Volumes[0].ID)

	// Instances
	assert.Equal(t, "bncvbnb-node-jtCyyabF", alarm.Instances[0].Name)
	assert.Equal(t, "a0a75b48-1175-491d-89d9-491d117589d9", alarm.Instances[0].ID)

	// LoadBalancers
	assert.Equal(t, "5c001a15-3867-42eb-9e0a-40ea8455688d", alarm.LoadBalancers[0].LoadBalancerID)
	assert.Equal(t, "fake-loadbalancer-name", alarm.LoadBalancers[0].LoadBalancerName)
	assert.Equal(t, "ec0e9614-9c89-4785-a56e-4b3dc4c4eb79", alarm.LoadBalancers[0].TargetID)
	assert.Contains(t, []string{"frontend", "backend"}, alarm.LoadBalancers[0].TargetType)

	// http
	assert.Equal(t, "https://google.com", alarm.HTTPURL)
	assert.Equal(t, 200, alarm.HTTPExpectedCode)
	assert.Equal(t, "X-Auth-Token", alarm.HTTPHeaders[0].Key)
	assert.Equal(t, "fake-token", alarm.HTTPHeaders[0].Value)

	// host alive
	assert.Equal(t, "1.1.1.1", alarm.Hostname)

	// alarms
	assert.Equal(t, "5f0c348c4c569c000d0549f0", alarm.ID)
	assert.Equal(t, 300, alarm.AlertInterval)
	assert.Contains(t, []string{
		"instance",
		"load_balancer",
		"simple_storage",
		"volume",
		"http",
		"host",
		"autoscale_group",
	}, alarm.ResourceType)
	assert.Contains(t, []string{
		">=",
		">",
		"<=",
		"<",
	}, alarm.Comparison.CompareType)

	assert.Equal(t, &Comparison{
		CompareType: ">=",
		Measurement: "iops",
		RangeTime:   300,
		Value:       5000,
	}, alarm.Comparison)
}

func Test_receivers_List(t *testing.T) {
	setup()
	defer teardown()
	var r receivers
	mux.HandleFunc(testlib.CloudWatcherURL(r.resourcePath()), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `{
		    "_items": [{
		        "_updated": "Thu, 02 Jul 2020 08:38:53 GMT",
		        "project_id": "fake-project-id",
		        "autoscale": {
		            "cluster_name": "group-01",
		            "cluster_id": "a2018319-4c80-4a52-bfb1-6045b2b01536",
		            "action_type": "CLUSTER SCALE OUT",
		            "action_id": "e169f975-a8e0-43d6-a9c9-06140929ab3c"
		        },
		        "name": "fake-receiver-name",
		        "webhook_url": "fake-webhook-url",
		        "verified_webhook_url": true,
		        "verified_telegram_chat_id": true,
		        "slack": {
		            "channel_name": "#focus",
		            "webhook_url": "fake-webhook-url"
		        },
		        "telegram_chat_id": "fake-telegram-chat-id",
		        "verified_email_address": true,
		        "email_address": "example@domain.com",
		        "sms_number": "fake-sms-number",
		        "verified_sms_number": true,
		        "creator": "autoscale",
		        "_deleted": false,
		        "_links": {
		            "self": {
		                "href": "receivers/5efd9d1f773561000d45b583",
		                "title": "receiver"
		            }
		        },
		        "_created": "Thu, 02 Jul 2020 08:38:53 GMT",
		        "user_id": "fake-user-id",
		        "_id": "5efd9d1f773561000d45b583"

		    }],
		    "_links": {
		        "self": {
		            "href": "receivers?max_results=10&sort=-_created",
		            "title": "receivers"
		        },
		        "parent": {
		            "href": "/",
		            "title": "home"
		        }
		    },
		    "_meta": {
		        "max_results": 10,
		        "total": 7,
		        "page": 1
		    }
		}`
		_, _ = fmt.Fprint(w, resp)
	})
	receivers, err := client.CloudWatcher.Receivers().List(ctx, nil)
	require.NoError(t, err)
	receiver := receivers[0]
	// receiver
	assert.Equal(t, "fake-receiver-name", receiver.Name)
	assert.Equal(t, "fake-webhook-url", receiver.WebhookURL)
	assert.Equal(t, "fake-telegram-chat-id", receiver.TelegramChatID)
	assert.Equal(t, "fake-sms-number", receiver.SMSNumber)
	assert.Equal(t, "example@domain.com", receiver.EmailAddress)
	assert.Equal(t, &Slack{
		SlackChannelName: "#focus",
		WebhookURL:       "fake-webhook-url",
	}, receiver.Slack)
	assert.Equal(t, &AutoScalingWebhook{
		ClusterName: "group-01",
		ClusterID:   "a2018319-4c80-4a52-bfb1-6045b2b01536",
		ActionID:    "e169f975-a8e0-43d6-a9c9-06140929ab3c",
		ActionType:  "CLUSTER SCALE OUT",
	}, receiver.AutoScale)
}

func Test_histories_List(t *testing.T) {
	setup()
	defer teardown()
	var h histories
	mux.HandleFunc(testlib.CloudWatcherURL(h.resourcePath()), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `{
		    "_items": [{
		        "_updated": "Thu, 01 Jan 1970 00:00:00 GMT",
		        "project_id": "fake-project-id",
		        "resource": "https://google.com",
		        "alarm": {
		            "_updated": "Sun, 19 Jul 2020 11:03:27 GMT",
		            "enable": true,
		            "user_id": "fake-user-id",
		            "name": "google",
		            "receivers": [{
		                "receiver_id": "5ea2a5bf248e68000dc480ae",
		                "methods": ["telegram"]
		            }],
		            "http_headers": [{
		                "key": "asadasd",
		                "value": "asdasda"
		            }],
		            "_deleted": false,
		            "alert_interval": 0,
		            "project_id": "fake-project-id",
		            "_created": "Sun, 19 Jul 2020 11:03:27 GMT",
		            "http_url": "https://google.com",
		            "_id": "5f14287f016b80000cd92a4d",
		            "creator": "user",
		            "resource_type": "http",
		            "http_expected_code": 200
		        },
		        "_deleted": false,
		        "alarm_id": "5f14287f016b80000cd92a4d",
		        "state": "DOWN",
		        "_links": {
		            "self": {
		                "href": "histories/5f1428a411084e00236c6ec4",
		                "title": "history"
		            }
		        },
		        "measurement": null,
		        "user_id": "fake-user-id",
		        "_id": "5f1428a411084e00236c6ec4",
		        "_created": "Sun, 19 Jul 2020 11:04:04 GMT"
		    }],
		    "_links": {
		        "self": {
		            "href": "histories?max_results=10&sort=-_created",
		            "title": "histories"
		        },
		        "last": {
		            "href": "histories?max_results=10&sort=-_created&page=7795",
		            "title": "last page"
		        },
		        "parent": {
		            "href": "/",
		            "title": "home"
		        },
		        "next": {
		            "href": "histories?max_results=10&sort=-_created&page=2",
		            "title": "next page"
		        }
		    },
		    "_meta": {
		        "max_results": 10,
		        "total": 77942,
		        "page": 1
		    }
		}`
		_, _ = fmt.Fprint(w, resp)
	})
	histories, err := client.CloudWatcher.Histories().List(ctx, nil)
	require.NoError(t, err)
	history := histories[0]
	assert.Equal(t, "DOWN", history.State)
	assert.Equal(t, "", history.Measurement)
	assert.Equal(t, "https://google.com", history.Resource)
	assert.Equal(t, history.AlarmID, history.Alarm.ID)
	assert.Equal(t, "5f14287f016b80000cd92a4d", history.AlarmID)
}
