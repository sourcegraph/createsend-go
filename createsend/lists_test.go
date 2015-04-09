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

func TestListCustomFields(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/lists/12CD/customfields.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[
			{
				"FieldName": "website",
				"Key": "[website]",
				"DataType": "Text",
				"FieldOptions": [],
				"VisibleInPreferenceCenter": true
			},
			{
				"FieldName": "age",
				"Key": "[age]",
				"DataType": "Number",
				"FieldOptions": [],
				"VisibleInPreferenceCenter": true
			},
			{
				"FieldName": "subscription date",
				"Key": "[subscriptiondate]",
				"DataType": "Date",
				"FieldOptions": [],
				"VisibleInPreferenceCenter": false
			},
			{
				"FieldName": "newsletterformat",
				"Key": "[newsletterformat]",
				"DataType": "MultiSelectOne",
				"FieldOptions": [
					"HTML",
					"Text"
				],
				"VisibleInPreferenceCenter": true
			}
		]`)
	})

	cfs, err := client.ListCustomFields("12CD")
	if err != nil {
		t.Errorf("ListCustomfields returned error: %v", err)
	}

	if len(cfs) != 4 {
		t.Errorf("Expected 4 customfields but got %d", len(cfs))
	}

	if cfs[0].DataType != "Text" {
		t.Errorf("Wrong datatype: %v", cfs[0].DataType)
	}

	if len(cfs[3].FieldOptions) != 2 {
		t.Errorf("Expected 2 fieldoptions: %d", len(cfs[3].FieldOptions))
	}

}

func TestListCreateCustomField(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/lists/12CD/customfields.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `"AE12"`)
	})

	id, err := client.ListCreateCustomField("12CD", &CustomFieldCreate{FieldName: "test", DataType: Text, VisibleInPreferenceCenter: false})
	if err != nil {
		t.Errorf("ListCreateCustomField return error: %v", err)
	}

	if id != "AE12" {
		t.Errorf("Id returned is wrong: %v", id)
	}
}

func TestListDeleteCustomField(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/lists/12CD/customfields/[test].json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusOK)
	})

	err := client.ListDeleteCustomField("12CD", "[test]")
	if err != nil {
		t.Errorf("ListDeleteCustomField return error: %v", err)
	}
}

func TestListDeleteCustomFieldFail(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/lists/12CD/customfields/test.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"Code":253}`)
	})

	err := client.ListDeleteCustomField("12CD", "test")
	if err == nil {
		t.Errorf("ListDeleteCustomField did not return error")
	}
}

func TestListSegments(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/lists/12CD/segments.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[
				  {
					"ListID": "12CD",
					"SegmentID": "b1b1b1b1b1b1b1b1b1b1b1b1b1b1b1b1",
					"Title": "Segment One"
				  },
				  {
					"ListID": "12CD",
					"SegmentID": "c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1",
					"Title": "Segment Two"
				  }
				]`)
	})

	segments, err := client.ListSegments("12CD")
	if err != nil {
		t.Errorf("ListSegments returned an error: %v", err)
	}

	if len(segments) != 2 {
		t.Errorf("Expected 2 ListsSegment's but got %d", len(segments))
	}

	if segments[0].ListID != "12CD" {
		t.Errorf("Expected first listid to equal \"12CD\" but got \"%s\"", segments[0].ListID)
	}
}
