package jerakia

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// DefaultUserAgent is the default User-Agent string set in the request handler.
const DefaultUserAgent = "go-jerakia/1.0.0"

// ClientConfig represents options used for creating a Jerakia client.
type ClientConfig struct {
	// URL is the URL to the Jerakia server.
	URL string

	// Token is the authentication token.
	Token string

	// UserAgent is a custom User-Agent.
	UserAgent UserAgent
}

// UserAgent represents a User-Agent header.
type UserAgent struct {
	prepend []string
}

// Prepend prepends a user-defined string to the default User-Agent string. Users
// may pass in one or more strings to prepend.
func (ua *UserAgent) Prepend(s ...string) {
	ua.prepend = append(s, ua.prepend...)
}

// Join concatenates all the user-defined User-Agend strings with the default
// User-Agent string.
func (ua *UserAgent) Join() string {
	uaSlice := append(ua.prepend, DefaultUserAgent)
	return strings.Join(uaSlice, " ")
}

// RequestOpts represents options used on a per-request basis.
type RequestOpts struct {
	// JSONResponse, if provided, will be populated with the contents of the
	// response body parsed as JSON.
	JSONResponse interface{}
}

// Client represents a Jerakia REST client.
type Client struct {
	httpClient http.Client
	config     ClientConfig
}

// Request performs an HTTP request.
func (client *Client) Request(method, url string, opts *RequestOpts) (*http.Response, error) {
	var body io.Reader

	// Construct the http.Request.
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// Set the content-type.
	req.Header.Set("Accept", "application/json")

	// Set the user-agent.
	req.Header.Set("User-Agent", client.config.UserAgent.Join())

	// Set the header token.
	req.Header.Set("X-Authentication", client.config.Token)

	// Set connection parameter to close the connection immediately when we've
	// got the response
	req.Close = true

	// Issue the request.
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Check for a bad response.
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)

		return nil, fmt.Errorf("Received %d: %s", resp.StatusCode, body)
	}

	// Parse the response body as JSON, if requested to do so.
	if opts.JSONResponse != nil {
		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(opts.JSONResponse); err != nil {
			return nil, err
		}
	}

	return resp, err
}

// Get calls `Request` with the "GET" HTTP verb.
func (client *Client) Get(url string, JSONResponse interface{}, opts *RequestOpts) (*http.Response, error) {
	if opts == nil {
		opts = new(RequestOpts)
	}

	if JSONResponse != nil {
		opts.JSONResponse = JSONResponse
	}

	return client.Request("GET", url, opts)
}

// NewClient will create and return a Client.
func NewClient(httpClient *http.Client, c ClientConfig) Client {
	return Client{
		httpClient: *httpClient,
		config:     c,
	}
}
