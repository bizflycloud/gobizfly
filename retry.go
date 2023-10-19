package gobizfly

import (
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"time"
)

const RetryCount = 3

type retryableTransport struct {
	transport http.RoundTripper
}

func backoff(retries int) time.Duration {
	return time.Duration(math.Pow(2, float64(retries))) * time.Second
}

func shouldRetry(err error, resp *http.Response) bool {
	if err != nil {
		return true
	}

	if resp.StatusCode == http.StatusBadGateway ||
		resp.StatusCode == http.StatusServiceUnavailable ||
		resp.StatusCode == http.StatusGatewayTimeout {
		return true
	}

	return false
}

func drainBody(resp *http.Response) {
	if resp.Body != nil {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}
}
