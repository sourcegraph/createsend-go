package createsend

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
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

func TestGetSubscriber(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/subscribers/12CD.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testQuerystring(t, r, "email=alice@example.com")
		fmt.Fprint(w, `{"EmailAddress":"alice@example.com","Name":"alice","Date":"2010-10-25 10:28:00"}`)
	})

	want := Subscriber{
		EmailAddress: "alice@example.com",
		Name:         "alice",
		Date:         time.Date(2010, 10, 25, 10, 28, 0, 0, time.UTC),
		DateStr:      "2010-10-25T10:28:00Z",
	}
	sub, err := client.GetSubscriber("12CD", "alice@example.com")
	if err != nil {
		t.Errorf("GetSubscriber returned error: %v", err)
	}
	if !reflect.DeepEqual(*sub, want) {
		t.Errorf("GetSubscriber returned %+v, want %+v", *sub, want)
	}
}

func TestGetSubscriber_NotInList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/subscribers/12CD.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"Code": 203, "Message": "Subscriber not in list"}`)
	})

	want := &CreatesendError{Code: 203, Message: "Subscriber not in list"}
	sub, err := client.GetSubscriber("12CD", "alice@example.com")
	if !reflect.DeepEqual(err, want) {
		t.Errorf("GetSubscriber returned error %+v, want %+v", err, want)
	}
	if sub != nil {
		t.Errorf("GetSubscriber returned non-nil subscriber %+v", sub)
	}
}

func TestUnsubscribe(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/subscribers/12CD/unsubscribe.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
	})

	err := client.Unsubscribe("12CD", "alice@example.com")
	if err != nil {
		t.Errorf("Unsubscribe returned error: %v", err)
	}
}
