package acceptance

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jtopjian/go-jerakia"
)

// AcceptanceTestChoices contains information required for acceptance tests.
type AcceptanceTestChoices struct {
	// URL is the URL to the Jerakia server.
	URL string

	// Token is a Jerakia auth token.
	Token string
}

// AcceptanceTestChoicesFromEnv populates AcceptanceTestChoices from
// environment variables.
func AcceptanceTestChoicesFromEnv() (*AcceptanceTestChoices, error) {
	url := os.Getenv("JERAKIA_URL")
	if url == "" {
		return nil, fmt.Errorf("JERAKIA_URL is required")
	}

	token := os.Getenv("JERAKIA_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("JERAKIA_TOKEN is required")
	}

	choices := AcceptanceTestChoices{
		URL:   url,
		Token: token,
	}

	return &choices, nil
}

// NewClient returns a new Jerakia client for acceptance testing.
func NewClient() (*jerakia.Client, error) {
	choices, err := AcceptanceTestChoicesFromEnv()
	if err != nil {
		return nil, err
	}

	config := jerakia.ClientConfig{
		URL:   choices.URL,
		Token: choices.Token,
	}

	client := jerakia.NewClient(http.DefaultClient, config)
	return &client, nil
}
