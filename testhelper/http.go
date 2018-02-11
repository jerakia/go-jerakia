package testhelper

import (
	"net/http"
	"net/http/httptest"
)

var (
	// Mux is a multiplexer that can be used to register handlers.
	Mux *http.ServeMux

	// Server is an in-memory HTTP server for testing.
	Server *httptest.Server
)

// SetupHTTP prepares the Mux and Server.
func SetupHTTP() {
	Mux = http.NewServeMux()
	Server = httptest.NewServer(Mux)
}

// TeardownHTTP releases HTTP-related resources.
func TeardownHTTP() {
	Server.Close()
}

// URL returns a fake URL that will actually target the Mux.
func URL() string {
	return Server.URL + "/"
}
