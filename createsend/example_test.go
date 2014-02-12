package createsend_test

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sourcegraph/createsend-go/createsend"
)

func Example() {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Fprintf(os.Stderr, "You must set your Campaign Monitor API key in the API_KEY environment variable to run example_test.go. (Skipping.)\n")
		return
	}

	authClient := &http.Client{
		Transport: &createsend.APIKeyAuthTransport{APIKey: apiKey},
	}

	c := createsend.NewAPIClient(authClient)
	clients, err := c.ListClients()
	if err != nil {
		fmt.Printf("Error listing clients: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Found %d clients.\n", len(clients))
	for _, client := range clients {
		fmt.Printf(" - %s (ID: [%d-char ID])\n", client.Name, len(client.ClientID))
	}

	// This output will be different for each account.

	// sample output:
	// Found 1 clients.
	//  - Sourcegraph (ID: [32-char ID])
}
