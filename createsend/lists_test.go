package createsend

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

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

func TestListSubscribers(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/lists/12CD/active.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"Results": [{"EmailAddress": "alice@example.com", "Name": "alice"}]}`)
	})

	subs, err := client.ListSubscribers("12CD", ActiveSubscribers, nil)
	if err != nil {
		t.Errorf("ListSubscribers returned error: %v", err)
	}

	want := []*Subscriber{{EmailAddress: "alice@example.com", Name: "alice"}}
	if !reflect.DeepEqual(subs, want) {
		t.Errorf("ListSubscribers returned %+v, want %+v", subs, want)
	}
}
