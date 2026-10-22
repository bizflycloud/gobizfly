// This file is part of gobizfly

package gobizfly

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/bizflycloud/gobizfly/testlib"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// failingTransport fails the n-th request to a path containing failPath with a network error.
type failingTransport struct {
	failPath string
	failOn   int32
	calls    int32
}

func (f *failingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if strings.Contains(req.URL.Path, f.failPath) && atomic.AddInt32(&f.calls, 1) == f.failOn {
		return nil, errors.New("connection reset by peer")
	}
	return http.DefaultTransport.RoundTrip(req)
}

func handleToken(t *testing.T, token string) {
	mux.HandleFunc(testlib.AuthURL(tokenPath), func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		_, _ = fmt.Fprintf(w, `{"token": %q, "expires_at": "2019-11-22T15:39:54.000000Z"}`, token)
	})
}

func TestDoRetryNetworkErrorReturnsError(t *testing.T) {
	setup()
	defer teardown()
	handleToken(t, "xxx")

	var l cloudLoadBalancerService
	mux.HandleFunc(testlib.LoadBalancerURL(l.resourcePath()), func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	})
	// First call gets the 401, the retry after the token refresh fails at the network level
	client.httpClient = &http.Client{Transport: &failingTransport{failPath: l.resourcePath(), failOn: 2}}
	client.keystoneToken = "yyy"

	assert.NotPanics(t, func() {
		_, err := client.CloudLoadBalancer.List(ctx, &ListOptions{})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "connection reset by peer")
	})
}

func TestDoRetryResendsBody(t *testing.T) {
	setup()
	defer teardown()
	handleToken(t, "xxx")

	var bodies []string
	mux.HandleFunc("/iaas-cloud/api/echo", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		b, _ := io.ReadAll(r.Body)
		bodies = append(bodies, string(b))
		if r.Header.Get("X-Auth-Token") != "xxx" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		_, _ = fmt.Fprint(w, `{}`)
	})
	client.keystoneToken = "yyy"

	req, err := client.NewRequest(ctx, http.MethodPost, serverServiceName, "/echo", map[string]string{"name": "node-a"})
	require.NoError(t, err)
	resp, err := client.Do(ctx, req)
	require.NoError(t, err)
	_ = resp.Body.Close()

	require.Len(t, bodies, 2)
	assert.Equal(t, bodies[0], bodies[1])
	assert.Contains(t, bodies[1], "node-a")
}

func TestDoTokenEndpointUnauthorizedDoesNotRecurse(t *testing.T) {
	setup()
	defer teardown()

	var tokenCalls int32
	mux.HandleFunc(testlib.AuthURL(tokenPath), func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&tokenCalls, 1)
		w.WriteHeader(http.StatusUnauthorized)
	})
	var l cloudLoadBalancerService
	mux.HandleFunc(testlib.LoadBalancerURL(l.resourcePath()), func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	})

	_, err := client.CloudLoadBalancer.List(ctx, &ListOptions{})
	require.Error(t, err)
	assert.Equal(t, int32(1), atomic.LoadInt32(&tokenCalls))
}

func TestNewRequestHonorsContext(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/iaas-cloud/api/slow", func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-r.Context().Done():
		case <-time.After(5 * time.Second):
		}
	})
	cctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	req, err := client.NewRequest(cctx, http.MethodGet, serverServiceName, "/slow", nil)
	require.NoError(t, err)
	start := time.Now()
	_, err = client.Do(cctx, req)
	require.Error(t, err)
	assert.True(t, errors.Is(err, context.DeadlineExceeded))
	assert.True(t, time.Since(start) < 2*time.Second)
}

func TestConcurrentTokenRefresh(t *testing.T) {
	setup()
	defer teardown()
	handleToken(t, "xxx")

	var l cloudLoadBalancerService
	mux.HandleFunc(testlib.LoadBalancerURL(l.resourcePath()), func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Auth-Token") != "xxx" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		_, _ = fmt.Fprint(w, `{"loadbalancers": []}`)
	})
	client.keystoneToken = "yyy"

	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := client.CloudLoadBalancer.List(ctx, &ListOptions{})
			assert.NoError(t, err)
		}()
	}
	wg.Wait()
}
