package jerakia

import (
	"fmt"
	"net/url"
)

const LookupURL = "lookup"

// LookupOpts represents options for a lookup.
type LookupOpts struct {
	// Namespace is the namespace to use for the request.
	// Nested namespaces should be delimited with /.
	Namespace string

	// Policy optionally override the policy used for the request.
	Policy string

	// LookupType optionally overrides the type of lookup (first, cascade).
	LookupType string

	// Merge optionally override the merge strategy to use (array,
	// deep_hash, hash).
	Merge string

	// Scope optionally provides an alternative scope handler to use
	// for the request (eg: puppetdb).
	Scope string

	// ScopeOptions are sent as parameters for the Scope.
	ScopeOptions map[string]string

	// Metadata specifies metadata for the request.
	Metadata map[string]string
}

// ToLookupQuery converts LookupOpts to a query string.
func (opts LookupOpts) ToLookupQuery() (string, error) {
	params := url.Values{}

	if opts.Namespace == "" {
		return "", fmt.Errorf("Namespace is required")
	}

	params.Add("namespace", opts.Namespace)

	if opts.Policy != "" {
		params.Add("policy", opts.Policy)
	}

	if opts.LookupType != "" {
		params.Add("lookup_type", opts.LookupType)
	}

	if opts.Scope != "" {
		params.Add("scope", opts.Scope)
	}

	for k, v := range opts.ScopeOptions {
		s := fmt.Sprintf("scope_%s", k)
		params.Add(s, v)
	}

	for k, v := range opts.Metadata {
		m := fmt.Sprintf("metadata_%s", k)
		params.Add(m, v)
	}

	u := url.URL{
		RawQuery: params.Encode(),
	}

	return u.String(), nil
}

// LookupResult represents a lookup result.
type LookupResult struct {
	// Status is the result of the request.
	Status string `json:"status"`

	// Found is if a value was found or not.
	Found bool `json:"found"`

	// Payload is the data returned from the lookup.
	Payload interface{} `json:"payload"`

	// Message provides details of the error if status is "failed".
	Message string `json:"message"`
}

// Lookup performs a lookup.
func Lookup(client *Client, key string, opts *LookupOpts) (*LookupResult, error) {
	var r LookupResult
	url := client.config.URL + "/" + LookupURL + "/" + key

	if opts != nil {
		query, err := opts.ToLookupQuery()
		if err != nil {
			return nil, err
		}

		url += query
	}

	_, err := client.Get(url, &r, nil)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
