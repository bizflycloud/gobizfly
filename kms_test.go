package gobizfly

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ClientInit() (*Client, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := NewClient(
		WithAPIURL("https://manage.bizflycloud.vn"),
		WithProjectID("cecd854937c8421b81d1d73789acad52"),
		WithRegionName("HN"),
	)
	if err != nil {
		fmt.Printf("Failed to create Bizfly client: %v", err)
		return nil, err
	}

	token, err := client.Token.Init(
		ctx,
		&TokenCreateRequest{
			AuthMethod:    "application_credential",                                                                 // application_credential
			Username:      "",                                                                                       // ""
			Password:      "",                                                                                       // ""
			AppCredID:     "22c2c3b3f43e4e2385dc00c575ab857d",                                                       // => from env
			AppCredSecret: "WznT3PwGFdsZ-WOmDBNsr3ZXGAtJ7GuWdPfPT9evK8RcR8JCw8VmvPt7rvWQOBkzUmoCJcF5IaH2EM3j-KExDw", //  => from env
		})

	if err != nil {
		fmt.Printf("Failed to create token: %v", err)
		return nil, err
	}
	client.SetKeystoneToken(token)

	return client, nil
}

func TestListKMSCertificates1(t *testing.T) {
	c, err := ClientInit()
	if err != nil {
		t.Fatal(err)
	}

	cert, err := c.KMS.Certificate().List(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// require.NotEmpty(t, cert)

	t.Log(cert)
}

func TestListKMSCertificates(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/ssl/certificate_container", func(writer http.ResponseWriter, request *http.Request) {
		require.Equal(t, http.MethodGet, request.Method)
		resp := `
{
  "data": [
	{
	  "id": 1,
	  "name": "test",
	}
  ]
}`

		_, _ = fmt.Fprint(writer, resp)
	})

	certs, err := client.KMS.Certificate().List(ctx)
	require.NoError(t, err)
	assert.Len(t, certs, 1)
	assert.LessOrEqual(t, "1", certs[0].ContainerId)
}
