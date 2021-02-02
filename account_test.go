package gobizfly

import (
	"fmt"
	"github.com/bizflycloud/gobizfly/testlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
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
