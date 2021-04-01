package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jerakia/go-jerakia"

	th "github.com/jerakia/go-jerakia/testhelper"
	fake "github.com/jerakia/go-jerakia/testhelper/client"

	"github.com/stretchr/testify/assert"
)

// LookupBasicResponse provides a GET response of a lookup.
const LookupBasicResponse = `
{
  "found": true,
  "payload": {
    "argentina": "buenos aires",
    "france": "paris",
    "spain": "malaga"
  },
  "status": "ok"
}
`

// LookupBasicResult is the expected result of a basic lookup.
var LookupBasicResult = jerakia.LookupResult{
	Status: "ok",
	Found:  true,
	Payload: map[string]interface{}{
		"argentina": "buenos aires",
		"france":    "paris",
		"spain":     "malaga",
	},
}

// HandleLookupBasic tests a basic lookup.
func HandleLookupBasic(t *testing.T) {
	th.Mux.HandleFunc("/lookup/cities", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET")
		assert.Equal(t, r.Header.Get("X-Authentication"), fake.Token)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, LookupBasicResponse)
	})
}

// LookupSingleBoolResponse provides a GET response of a single bool lookup.
const LookupSingleBoolResponse = `
{
  "found": true,
  "payload": true,
  "status": "ok"
}
`

// LookupSingleBoolResult is the expected result of a single bool lookup.
var LookupSingleBoolResult = jerakia.LookupResult{
	Status:  "ok",
	Found:   true,
	Payload: true,
}

// HandleLookupSingleBool tests a single bool lookup.
func HandleLookupSingleBool(t *testing.T) {
	th.Mux.HandleFunc("/lookup/booltrue", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET")
		assert.Equal(t, r.Header.Get("X-Authentication"), fake.Token)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, LookupSingleBoolResponse)
	})
}

// LookupMetadataResponse is the expected response of a metadata lookup.
const LookupMetadataResponse = `
{
    "found": true,
    "payload": [
      "bob",
      "lucy",
      "david"
    ],
    "status": "ok"
}
`

// LookupMetadataResult is the expected result of a metadata lookup.
var LookupMetadataResult = jerakia.LookupResult{
	Status: "ok",
	Found:  true,
	Payload: []interface{}{
		"bob", "lucy", "david",
	},
}

// HandleLookupMetadata tests a metadata lookup.
func HandleLookupMetadata(t *testing.T) {
	th.Mux.HandleFunc("/lookup/users", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET")
		assert.Equal(t, r.Header.Get("X-Authentication"), fake.Token)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, LookupMetadataResponse)
	})
}

// LookupKeylessResponse is the expected response of a keyless lookup.
const LookupKeylessResponse = `
{
    "found": true,
    "payload": {
			"foo": "bar",
			"hello": "world"
    },
    "status": "ok"
}
`

// LookupKeylessResult is the expected result of a keyless lookup.
var LookupKeylessResult = jerakia.LookupResult{
	Status: "ok",
	Found:  true,
	Payload: map[string]interface{}{
		"foo": "bar",
		"hello": "world",
	},
}

// HandleLookupKeyless tests a keyless lookup.
func HandleLookupKeyless(t *testing.T) {
	th.Mux.HandleFunc("/lookup", func(w http.ResponseWriter, r *http.Request) {

		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, fake.Token, r.Header.Get("X-Authentication"))
		assert.Equal(t, []string{"keyless"}, r.URL.Query()["namespace"])

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, LookupKeylessResponse)
	})
}

