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

func TestListLists(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/clients/12ab/lists.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"ListID": "34cd", "Name": "mylist"}]`)
	})

	lists, err := client.ListLists("12ab")
	if err != nil {
		t.Errorf("ListLists returned error: %v", err)
	}

	want := []*List{{ListID: "34cd", Name: "mylist"}}
	if !reflect.DeepEqual(lists, want) {
		t.Errorf("ListLists returned %+v, want %+v", lists, want)
	}
}

func TestListsForEmail(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/clients/12ab/listsforemail.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testQuerystring(t, r, "email=alice@example.com")
		fmt.Fprint(w, `[{"ListID": "34cd", "Name": "mylist"}]`)
	})

	lists, err := client.ListsForEmail("12ab", "alice@example.com")
	if err != nil {
		t.Errorf("ListsForEmail returned error: %v", err)
	}

	want := []*List{{ListID: "34cd", Name: "mylist"}}
	if !reflect.DeepEqual(lists, want) {
		t.Errorf("ListsForEmail returned %+v, want %+v", lists, want)
	}
}
