package createsend

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAddSubscriber(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/subscribers/12CD.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `"alice@example.com"`)
	})

	sub := NewSubscriber{
		EmailAddress: "alice@example.com",
		Name:         "Alice",
	}
	err := client.AddSubscriber("12CD", sub)
	if err != nil {
		t.Errorf("AddSubscriber returned error: %v", err)
	}
}
