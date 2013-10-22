package createsend

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListClients(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/clients.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"ClientID": "12ab", "Name": "Alice"}]`)
	})

	clients, err := client.ListClients()
	if err != nil {
		t.Errorf("ListClients returned error: %v", err)
	}

	want := []Client{{ClientID: "12ab", Name: "Alice"}}
	if !reflect.DeepEqual(clients, want) {
		t.Errorf("ListClients returned %+v, want %+v", clients, want)
	}
}
