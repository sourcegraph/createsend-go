package createsend

import (
	"net/http"
)

type APIKeyAuthTransport struct {
	Transport http.RoundTripper
	APIKey    string
}

func (t *APIKeyAuthTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	transport := t.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}

	req.SetBasicAuth(t.APIKey, "x")

	return transport.RoundTrip(req)
}
