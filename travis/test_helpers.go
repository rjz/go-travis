package travis

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	tc     *Client
)

func setup(token string) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	tc = &Client{&server.URL, &token, http.DefaultClient}
}

func teardown() {
	server.Close()
}

func expectStringEqual(t *testing.T, desc, expected, actual string) {
	if expected != actual {
		t.Error(fmt.Sprintf("Expected %s to be '%s'; got '%s'", desc, expected, actual))
	}
}

func expectQueryParam(t *testing.T, r *http.Request, key, val string) {
	expectStringEqual(t, fmt.Sprintf("query param '%s'", key), val, r.URL.Query().Get(key))
}

func expectMethod(t *testing.T, r *http.Request, method string) {
	expectStringEqual(t, "method", method, r.Method)
}

func expectHeader(t *testing.T, r *http.Request, key, expected string) {
	expectStringEqual(t, "header "+key, expected, r.Header.Get(key))
}

func expectTravisHeaders(t *testing.T, r *http.Request, token string) {
	expectHeader(t, r, "Accept", "application/vnd.travis-ci.2+json")
	expectHeader(t, r, "User-agent", "go-travis/0.0.1")
	expectHeader(t, r, "Authorization", fmt.Sprintf("token %s", token))
}
