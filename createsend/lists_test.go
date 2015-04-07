package createsend

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

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

func TestListCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/lists/AAE3.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, "\"12CD\"")
	})

	id, err := client.ListCreate("AAE3", &ListCreateOptions{Title: "Test", UnsubscribeSetting: AllClientLists})
	if err != nil {
		t.Errorf("ListCreate returned error: %v", err)
	}

	if id != "12CD" {
		t.Errorf("Return value incorrect")
	}
}

func TestListDelete(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/lists/12CD.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	err := client.ListDelete("12CD")
	if err != nil {
		t.Errorf("ListDelete returned error: %v", err)
	}
}
