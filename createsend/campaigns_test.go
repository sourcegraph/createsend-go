package createsend

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCampaignRecipients(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/campaigns/13CD/recipients.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
				"Results": [
					{
						"EmailAddress": "example+1@example.com",
						"ListID": "a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1"
					},
					{
						"EmailAddress": "example+2@example.com",
						"ListID": "a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1"
					},
					{
						"EmailAddress": "example+3@example.com",
						"ListID": "a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1"
					},
					{
						"EmailAddress": "example+4@example.com",
						"ListID": "a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1"
					}
				],
				"ResultsOrderedBy": "email",
				"OrderDirection": "asc",
				"PageNumber": 1,
				"PageSize": 1000,
				"RecordsOnThisPage": 4,
				"TotalNumberOfRecords": 4,
				"NumberOfPages": 1
		}`)
	})

	campaigns, err := client.CampaignRecipients("13CD", nil)
	if err != nil {
		t.Errorf("CampaignRecipients returned error: %v", err)
	}

	recs := []*Recipient{
		{
			EmailAddress: "example+1@example.com",
			ListID:       "a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1",
		},
		{
			EmailAddress: "example+2@example.com",
			ListID:       "a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1",
		},
		{
			EmailAddress: "example+3@example.com",
			ListID:       "a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1",
		},
		{
			EmailAddress: "example+4@example.com",
			ListID:       "a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1",
		},
	}
	want := &CampaignRecipients{
		Results:              recs,
		ResultsOrderedBy:     "email",
		OrderDirection:       "asc",
		PageNumber:           1,
		PageSize:             1000,
		RecordsOnThisPage:    4,
		TotalNumberOfRecords: 4,
		NumberOfPages:        1,
	}

	if !reflect.DeepEqual(campaigns, want) {
		t.Errorf("CampaignRecipients returend %+v, want %+v", campaigns, want)
	}
}

func TestCampaignRecipientsOptions(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/campaigns/13CD/recipients.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		if r.FormValue("page") != "2" {
			t.Errorf("Expected page to equal 3 but was %s", r.FormValue("page"))
		}
		if r.FormValue("pagesize") != "300" {
			t.Errorf("Expected pagesize to equal 300 but was %s", r.FormValue("pagesize"))
		}
		if r.FormValue("orderfield") != "email" {
			t.Errorf("Expected orderfield to equal email but was %s", r.FormValue("email"))
		}
		if r.FormValue("orderdirection") != "desc" {
			t.Errorf("Expected orderdirection to equal desc but was %s", r.FormValue("orderdirection"))
		}
		fmt.Fprint(w, `{
				"Results": [
					{
						"EmailAddress": "example+1@example.com",
						"ListID": "a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1"
					},
					{
						"EmailAddress": "example+2@example.com",
						"ListID": "a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1"
					},
					{
						"EmailAddress": "example+3@example.com",
						"ListID": "a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1"
					},
					{
						"EmailAddress": "example+4@example.com",
						"ListID": "a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1"
					}
				],
				"ResultsOrderedBy": "email",
				"OrderDirection": "desc",
				"PageNumber": 2,
				"PageSize": 300,
				"RecordsOnThisPage": 4,
				"TotalNumberOfRecords": 304,
				"NumberOfPages": 2
		}`)
	})

	opts := CampaignRecipientsOptions{
		Page:           2,
		PageSize:       300,
		OrderField:     "email",
		OrderDirection: "desc",
	}
	campaigns, err := client.CampaignRecipients("13CD", &opts)
	if err != nil {
		t.Errorf("CampaignRecipients returned error: %v", err)
	}

	recs := []*Recipient{
		{
			EmailAddress: "example+1@example.com",
			ListID:       "a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1",
		},
		{
			EmailAddress: "example+2@example.com",
			ListID:       "a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1",
		},
		{
			EmailAddress: "example+3@example.com",
			ListID:       "a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1",
		},
		{
			EmailAddress: "example+4@example.com",
			ListID:       "a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1",
		},
	}
	want := &CampaignRecipients{
		Results:              recs,
		ResultsOrderedBy:     "email",
		OrderDirection:       "desc",
		PageNumber:           2,
		PageSize:             300,
		RecordsOnThisPage:    4,
		TotalNumberOfRecords: 304,
		NumberOfPages:        2,
	}

	if !reflect.DeepEqual(campaigns, want) {
		t.Errorf("CampaignRecipients returend %+v, want %+v", campaigns, want)
	}
}
