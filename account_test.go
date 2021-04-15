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

func TestRegionList(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.AccountURL(regionsPath), func(writer http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		resp := `{
  "HCM": {
    "active": true, 
    "icon": "https://bizfly-ui.ss-hn-1.vccloud.vn/v1/staging/icons/HCM2.svg", 
    "name": "H\u1ed3 Ch\u00ed Minh", 
    "order": 2, 
    "region_name": "HoChiMinh", 
    "short_name": "HCM", 
    "zones": [
      {
        "active": true, 
        "icon": "https://bizfly-ui.ss-hn-1.vccloud.vn/v1/staging/icons/HCM2.svg", 
        "name": "H\u1ed3 Ch\u00ed Minh 1", 
        "order": 1, 
        "short_name": "HCM1"
      }
    ]
  }, 
  "HN": {
    "active": true, 
    "icon": "https://bizfly-ui.ss-hn-1.vccloud.vn/v1/staging/icons/hn1.png", 
    "name": "H\u00e0 N\u1ed9i", 
    "order": 1, 
    "region_name": "HaNoi", 
    "short_name": "HN", 
    "zones": [
      {
        "active": true, 
        "icon": "https://bizfly-ui.ss-hn-1.vccloud.vn/v1/staging/icons/hn1.png", 
        "name": "H\u00e0 N\u1ed9i 1", 
        "order": 1, 
        "short_name": "HN1"
      }, 
      {
        "active": true, 
        "icon": "https://bizfly-ui.ss-hn-1.vccloud.vn/v1/staging/icons/hn1.png", 
        "name": "H\u00e0 N\u1ed9i 2", 
        "order": 2, 
        "short_name": "HN2"
      }
    ]
  }
}
`
		_, _ = fmt.Fprint(writer, resp)
	})
	regions, err := client.Account.ListRegion(ctx)
	require.NoError(t, err)

	assert.Equal(t, "https://bizfly-ui.ss-hn-1.vccloud.vn/v1/staging/icons/HCM2.svg", regions.HCM.Icon)

}

func TestUserGet(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc(testlib.AccountURL(usersInfoPath), func(writer http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		resp := `{
    "service": null,
    "url_type": "service",
    "origin": null,
    "client_type": null,
    "billing_balance": 108768098,
    "balances": {
        "cloud_server": 0.0,
        "load_balancer": 0.0,
        "mail_inbox": 0.0,
        "promotion_general": 0.0,
        "cloud_storage": 0.0,
        "call_center": 0.0,
        "pentest": 0.0,
        "container_registry": 0.0,
        "media_fullstack": 0.0,
        "primary": 108768098,
        "vpn": 0.0,
        "security": 0.0,
        "sysadmin": 0.0,
        "drive": 0.0,
        "ddos": 0.0,
        "auto_scaling": 0.0,
        "cdn": 0.0,
        "cloud_watcher": 0.0,
        "kubernetes_engine": 0.0,
        "simple_storage": 0.0
    },
    "payment_method": "pre-paid",
    "billing_acc_id": "292e0b89-aa07-47fc-8573-7ebe81bbc342",
    "debit": false,
    "email": "vctest_devcs_tung491@vccloud.vn",
    "phone": "",
    "fullname": "",
    "verified_email": true,
    "verified_phone": true,
    "login_alert": true,
    "verified_payment": true,
    "last_region": "HN",
    "last_project": "vctest_devcs_tung491@vccloud.vn",
    "type": "Owner",
    "otp": false,
    "services": [
        {
            "canonical_name": "mail_inbox",
            "name": "Mail Inbox",
            "service_url": "https://mail.vccloud.vn",
            "description": "D\u1ecbch v\u1ee5 m\u1eb7c \u0111\u1ecbnh c\u1ee7a VCE. M\u1ed7i ng\u01b0\u1eddi d\u00f9ng tr\u00ean h\u1ec7 th\u1ed1ng s\u1ebd \u0111\u01b0\u1ee3c cung c\u1ea5p 1 \u0111\u1ecba ch\u1ec9 email khi kh\u1edfi t\u1ea1o.",
            "enabled": true,
            "code": "",
            "region": "",
            "icon": ""
        },
        {
            "canonical_name": "admin",
            "name": "Admin",
            "service_url": "https://staging-admin.bizflycloud.vn",
            "description": "Dich vu admin",
            "enabled": true,
            "code": "",
            "region": "",
            "icon": ""
        },
        {
            "canonical_name": "cloud_server",
            "name": "Cloud Server",
            "service_url": "https://hn-staging.manage.bizflycloud.vn/iaas-cloud/api",
            "description": "",
            "enabled": true,
            "code": "CS",
            "region": "HN",
            "icon": "https://hn-staging.manage.bizflycloud.vn/iaas-cloud/api"
        },
        {
            "canonical_name": "cloud_server",
            "name": "Cloud Server",
            "service_url": "https://hcm-staging.manage.bizflycloud.vn/iaas-cloud/api",
            "description": "",
            "enabled": true,
            "code": "CS",
            "region": "HCM",
            "icon": "https://hcm-staging.manage.bizflycloud.vn/iaas-cloud/api"
        },
        {
            "canonical_name": "load_balancer",
            "name": "Load Balancer",
            "service_url": "https://hn-staging.manage.bizflycloud.vn/api/loadbalancers",
            "description": "",
            "enabled": true,
            "code": "LB",
            "region": "HN",
            "icon": "https://hn-staging.manage.bizflycloud.vn/api/loadbalancers"
        },
        {
            "canonical_name": "load_balancer",
            "name": "Load Balancer",
            "service_url": "https://hcm-staging.manage.bizflycloud.vn/api/loadbalancers",
            "description": "",
            "enabled": true,
            "code": "LB",
            "region": "HCM",
            "icon": "https://hcm-staging.manage.bizflycloud.vn/api/loadbalancers"
        },
        {
            "canonical_name": "vpn_site_to_site",
            "name": "VPN Site to Site",
            "service_url": "https://hn-staging.manage.bizflycloud.vn/api/vpnaas",
            "description": "",
            "enabled": true,
            "code": "VPN",
            "region": "HN",
            "icon": "https://hn-staging.manage.bizflycloud.vn/api/vpnaas"
        },
        {
            "canonical_name": "vpn_site_to_site",
            "name": "VPN Site to Site",
            "service_url": "https://hcm-staging.manage.bizflycloud.vn/api/vpnaas",
            "description": "",
            "enabled": true,
            "code": "VPN",
            "region": "HCM",
            "icon": "https://hcm-staging.manage.bizflycloud.vn/api/vpnaas"
        },
        {
            "canonical_name": "dns",
            "name": "DNS",
            "service_url": "https://staging.bizflycloud.vn/api/dns",
            "description": "",
            "enabled": true,
            "code": "DNS",
            "region": "HN",
            "icon": "https://staging.bizflycloud.vn/api/dns"
        },
        {
            "canonical_name": "auto_scaling",
            "name": "Auto Scaling",
            "service_url": "https://hn-staging.manage.bizflycloud.vn/api/auto-scaling",
            "description": "",
            "enabled": true,
            "code": "BAS",
            "region": "HN",
            "icon": "https://hn-staging.manage.bizflycloud.vn/api/auto-scaling"
        },
        {
            "canonical_name": "auto_scaling",
            "name": "Auto Scaling",
            "service_url": "https://hcm-staging.manage.bizflycloud.vn/api/auto-scaling",
            "description": "",
            "enabled": true,
            "code": "BAS",
            "region": "HCM",
            "icon": "https://hcm-staging.manage.bizflycloud.vn/api/auto-scaling"
        }
    ],
    "whitelist": [],
    "expires": "2021-02-01T12:08:33.000000Z",
    "tenant_id": "ebbed256d9414b0598719c42dc17e837",
    "tenant_name": "vctest_devcs_tung491@vccloud.vn",
    "ks_user_id": "7156c45b82cb4fabba997a90b032c0de",
    "iam": {
        "expire": "2021-02-01T12:08:33.000000Z",
        "tenant_id": "ebbed256d9414b0598719c42dc17e837",
        "tenant_name": "vctest_devcs_tung491@vccloud.vn",
        "verified_phone": true,
        "verified_email": true,
        "verified_payment": true
    },
    "domains": [],
    "payment_type": "pre-paid",
    "dob": "",
    "_gender": "male",
    "trial": {
        "started_at": "2020-10-02T12:14:37.752000Z",
        "expired_at": "2020-10-05T12:14:37.752000Z",
        "active": false,
        "enable": false,
        "service_level": 2
    },
    "has_expired_invoice": false,
    "negative_balance": false,
    "promotion": []
}
`
		_, _ = fmt.Fprint(writer, resp)
	})
	user, err := client.Account.GetUserInfo(ctx)
	require.NoError(t, err)
	assert.Equal(t, false, user.NegativeBalance)
}
