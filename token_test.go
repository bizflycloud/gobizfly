package gobizfly

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToken(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(tokenPath, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		resp := `
{
    "token": "xxx",
    "expires_at": "2019-11-22T15:39:54.000000Z"
}
`
		fmt.Fprint(w, resp)
	})

	tok, err := client.Token.Create(ctx, &TokenCreateRequest{})
	require.NoError(t, err)
	require.Equal(t, "xxx", tok.Token)
}
