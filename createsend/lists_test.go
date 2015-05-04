package createsend

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestListSubscribers(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/lists/12CD/active.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"Results": [{"EmailAddress": "alice@example.com", "Name": "alice"}],
		"ResultsOrderedBy": "email",
	    "OrderDirection": "asc",
	    "PageNumber": 1,
	    "PageSize": 1000,
	    "RecordsOnThisPage": 1,
	    "TotalNumberOfRecords": 1,
	    "NumberOfPages": 1}]}`)
	})

	subs, err := client.ListSubscribers("12CD", ActiveSubscribers, nil)
	if err != nil {
		t.Errorf("ListSubscribers returned error: %v", err)
	}

	want := &ListSubscribersResponse{Results: []*Subscriber{{EmailAddress: "alice@example.com", Name: "alice"}}, ResultsOrderedBy: "email", OrderDirection: "asc", PageNumber: 1, PageSize: 1000, RecordsOnThisPage: 1, TotalNumberOfRecords: 1, NumberOfPages: 1}
	if !reflect.DeepEqual(subs, want) {
		t.Errorf("ListSubscribers returned %+v, want %+v", subs, want)
	}
}

func TestListSubscribersEmptyOptions(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/lists/12CD/active.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"Results": [{"EmailAddress": "alice@example.com", "Name": "alice"}],
		"ResultsOrderedBy": "email",
	    "OrderDirection": "asc",
	    "PageNumber": 1,
	    "PageSize": 1000,
	    "RecordsOnThisPage": 1,
	    "TotalNumberOfRecords": 1,
	    "NumberOfPages": 1}]}`)
	})

	subs, err := client.ListSubscribers("12CD", ActiveSubscribers, &ListSubscribersOptions{})
	if err != nil {
		t.Errorf("ListSubscribers returned error: %v", err)
	}

	want := &ListSubscribersResponse{Results: []*Subscriber{{EmailAddress: "alice@example.com", Name: "alice"}}, ResultsOrderedBy: "email", OrderDirection: "asc", PageNumber: 1, PageSize: 1000, RecordsOnThisPage: 1, TotalNumberOfRecords: 1, NumberOfPages: 1}
	if !reflect.DeepEqual(subs, want) {
		t.Errorf("ListSubscribers returned %+v, want %+v", subs, want)
	}
}

func TestListSubscribersOptionsDate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/lists/12CD/active.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if r.FormValue("date") != "2001-02-03" {
			t.Errorf("Expected date to equsl 2001-02-03 but was %s", r.FormValue("date"))
		}
		if r.FormValue("page") != "" || r.FormValue("pagesize") != "" || r.FormValue("orderfield") != "" || r.FormValue("orderdirection") != "" {
			t.Errorf("Unexpected extra parameters")
		}
		fmt.Fprint(w, `{"Results": [{"EmailAddress": "alice@example.com", "Name": "alice"}],
		"ResultsOrderedBy": "email",
	    "OrderDirection": "asc",
	    "PageNumber": 1,
	    "PageSize": 1000,
	    "RecordsOnThisPage": 1,
	    "TotalNumberOfRecords": 1,
	    "NumberOfPages": 1}]}`)
	})

	subs, err := client.ListSubscribers("12CD", ActiveSubscribers, &ListSubscribersOptions{Date: time.Date(2001, time.February, 3, 0, 0, 0, 0, time.UTC)})
	if err != nil {
		t.Errorf("ListSubscribers returned error: %v", err)
	}

	want := &ListSubscribersResponse{Results: []*Subscriber{{EmailAddress: "alice@example.com", Name: "alice"}}, ResultsOrderedBy: "email", OrderDirection: "asc", PageNumber: 1, PageSize: 1000, RecordsOnThisPage: 1, TotalNumberOfRecords: 1, NumberOfPages: 1}
	if !reflect.DeepEqual(subs, want) {
		t.Errorf("ListSubscribers returned %+v, want %+v", subs, want)
	}
}

func TestListSubscribersOptionsPage(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/lists/12CD/active.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if r.FormValue("page") != "3" {
			t.Errorf("Expected page to equal 3 but was %s", r.FormValue("page"))
		}
		if r.FormValue("pagesize") != "123" {
			t.Errorf("Expected pagesize to equal 123 but was %s", r.FormValue("pagesize"))
		}
		if r.FormValue("date") != "" || r.FormValue("orderfield") != "" || r.FormValue("orderdirection") != "" {
			t.Errorf("Unexpected extra parameters")
		}
		fmt.Fprint(w, `{"Results": [{"EmailAddress": "alice@example.com", "Name": "alice"}],
		"ResultsOrderedBy": "email",
	    "OrderDirection": "asc",
	    "PageNumber": 3,
	    "PageSize": 1000,
	    "RecordsOnThisPage": 1,
	    "TotalNumberOfRecords": 1,
	    "NumberOfPages": 1}]}`)
	})

	subs, err := client.ListSubscribers("12CD", ActiveSubscribers, &ListSubscribersOptions{Page: 3, PageSize: 123})
	if err != nil {
		t.Errorf("ListSubscribers returned error: %v", err)
	}

	want := &ListSubscribersResponse{Results: []*Subscriber{{EmailAddress: "alice@example.com", Name: "alice"}}, ResultsOrderedBy: "email", OrderDirection: "asc", PageNumber: 3, PageSize: 1000, RecordsOnThisPage: 1, TotalNumberOfRecords: 1, NumberOfPages: 1}
	if !reflect.DeepEqual(subs, want) {
		t.Errorf("ListSubscribers returned %+v, want %+v", subs, want)
	}
}

func TestListSubscribersOptionsOrder(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/lists/12CD/active.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if r.FormValue("orderfield") != "of" {
			t.Errorf("Expected orderfield to equal of but was %s", r.FormValue("orderfield"))
		}
		if r.FormValue("orderdirection") != "desc" {
			t.Errorf("Expected orderdirection to equal desc but was %s", r.FormValue("orderdirection"))
		}
		if r.FormValue("date") != "" || r.FormValue("page") != "" || r.FormValue("pagesize") != "" {
			t.Errorf("Unexpected extra parameters")
		}
		fmt.Fprint(w, `{"Results": [{"EmailAddress": "alice@example.com", "Name": "alice"}],
		"ResultsOrderedBy": "of",
	    "OrderDirection": "desc",
	    "PageNumber": 1,
	    "PageSize": 1000,
	    "RecordsOnThisPage": 1,
	    "TotalNumberOfRecords": 1,
	    "NumberOfPages": 1}]}`)
	})

	subs, err := client.ListSubscribers("12CD", ActiveSubscribers, &ListSubscribersOptions{OrderField: "of", OrderDirection: "desc"})
	if err != nil {
		t.Errorf("ListSubscribers returned error: %v", err)
	}

	want := &ListSubscribersResponse{Results: []*Subscriber{{EmailAddress: "alice@example.com", Name: "alice"}}, ResultsOrderedBy: "of", OrderDirection: "desc", PageNumber: 1, PageSize: 1000, RecordsOnThisPage: 1, TotalNumberOfRecords: 1, NumberOfPages: 1}
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

func TestListWebhooks(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/lists/12CD/webhooks.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[
				{
					"WebhookID": "ee1b3864e5ca61618q98su98qsu9q",
					"Events": [
						"Subscribe"
					],
					"Url": "http://example.com/subscribe",
					"Status": "Active",
					"PayloadFormat": "Json"
				},
				{
					"WebhookID": "89d8quw98du9qw8ud9q8wud8u98u8",
					"Events": [
						"Deactivate"
					],
					"Url": "http://example.com/deactivate",
					"Status": "Active",
					"PayloadFormat": "Json"
				}
			]`)
	})

	webhooks, err := client.ListWebhooks("12CD")
	if err != nil {
		t.Errorf("ListWebhooks returned an error: %v", err)
	}

	if len(webhooks) != 2 {
		t.Errorf("Expected 2 webhooks but got %d", len(webhooks))
	}

	if webhooks[0].Url != "http://example.com/subscribe" {
		t.Errorf("Wrong url: %s", webhooks[0].Url)
	}
}

func TestListCreateWebhook(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/lists/12CD/webhooks.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `"QWE123"`)
	})

	id, err := client.ListCreateWebhook("12CD", &WebhookCreate{Events: []string{"Subscribe"}, Url: "http://example.com/subscribe", PayloadFormat: "json"})
	if err != nil {
		t.Errorf("ListCreateWebhook returned an error: %v", err)
	}

	if id != "QWE123" {
		t.Errorf("Wrong ID returned: %s", id)
	}
}

func TestListCreateWebhookFail(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/lists/12CD/webhooks.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"Code" : 602}`)
	})

	_, err := client.ListCreateWebhook("12CD", &WebhookCreate{Events: []string{"Subscribe"}, Url: "http://example.com/subscribe"})
	if err == nil {
		t.Errorf("ListCreateWebhook did not return an error")
	}
}

func TestListTestWebhook(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/lists/12CD/webhooks/QWE123/test.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusOK)
	})

	err := client.ListTestWebhook("12CD", "QWE123")
	if err != nil {
		t.Errorf("ListTestWebhook returned an error: %v", err)
	}
}

func TestListTestWebhookFail(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/lists/12CD/webhooks/QWE123/test.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{
				"ResultData": {
					"FailureStatus": "ProtocolError",
					"FailureResponseMessage": "NotFound",
					"FailureResponseCode": 404,
					"FailureResponse": ""
				},
				"Code": 610,
				"Message": "The webhook request has failed"
			}`)
	})

	err := client.ListTestWebhook("12CD", "QWE123")
	if err == nil {
		t.Errorf("ListTestWebhook did not return an error")
	}
}

func TestListDeleteWebhook(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/lists/12CD/webhooks/QWE123.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusOK)
	})

	err := client.ListDeleteWebhook("12CD", "QWE123")
	if err != nil {
		t.Errorf("ListDeleteWebhook returned an error: %v", err)
	}
}

func TestListActivateWebhook(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/lists/12CD/webhooks/QWE123/activate.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		w.WriteHeader(http.StatusOK)
	})

	err := client.ListActivateWebhook("12CD", "QWE123")
	if err != nil {
		t.Errorf("ListActivateWebhook returned an error: %v", err)
	}
}

func TestListDeactivateWebhook(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/lists/12CD/webhooks/QWE123/deactivate.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		w.WriteHeader(http.StatusOK)
	})

	err := client.ListDeactivateWebhook("12CD", "QWE123")
	if err != nil {
		t.Errorf("ListDeactivateWebhook returned an error: %v", err)
	}
}
