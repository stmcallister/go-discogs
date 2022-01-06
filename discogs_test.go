package discogs

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the PagerDuty client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	authToken := "foo"
	client = NewClient(authToken)
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()

	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testEqual(t *testing.T, want interface{}, got interface{}) {
	t.Helper()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("values not equal (-want / +got):\n%s", diff)
	}
}

func defaultTestClient(serverURL, apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Minute,
		},
		baseURL: serverURL,
	}
}
