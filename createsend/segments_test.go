package createsend

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSegmentCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/segments/12CD.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `"EQS1"`)
	})

	rg := make([]RuleGroupCreate, 1)
	rg[0].Rules = []RuleCreate{RuleCreate{RuleType: "DateSubscribed", Clause: "AFTER 2009-01-01"}}

	sgmt := SegmentCreate{Title: "Test", RuleGroups: rg}
	id, err := client.SegmentCreate("12CD", &sgmt)

	if err != nil {
		t.Errorf("SegmentCreate returned error: %v", err)
	}

	if id != "EQS1" {
		t.Errorf("Incorrect id returned: %v", id)
	}

}

func TestSegmentCreateFail(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/segments/12CD.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"Code" : 275}`)
	})

	rg := make([]RuleGroupCreate, 1)
	rg[0].Rules = []RuleCreate{RuleCreate{RuleType: "DateSubscribed", Clause: "AFTER 2009-01-01"}}

	sgmt := SegmentCreate{Title: "Test", RuleGroups: rg}
	_, err := client.SegmentCreate("12CD", &sgmt)

	if err == nil {
		t.Errorf("SegmentCreate returned no error")
	}
}

func TestSegmentDetail(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/segments/12CD.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
				"ActiveSubscribers": 1,
				"RuleGroups": [
					{
						"Rules": [
							{
								"RuleType": "Subject1_1",
								"Clause": "Clause1_1"
							},
							{
								"RuleType": "Subject1_2",
								"Clause": "Clause1_2"
							}
						]
					},
					{
						"Rules": [
							{
								"RuleType": "Subject2_1",
								"Clause": "Clause2_1"
							}
						]
					}
				],
				"ListID": "a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1",
				"SegmentID": "12CD",
				"Title": "Segment Title"
			}`)
	})

	s, err := client.SegmentDetail("12CD")
	if err != nil {
		t.Errorf("SegmentDetail returned an error")
	}

	if s.ActiveSubscribers != 1 {
		t.Errorf("Expected 1 active subscriber but got: %d", s.ActiveSubscribers)
	}

	if s.SegmentID != "12CD" {
		t.Errorf("Expected SegmentID 12CD but got: %s", s.SegmentID)
	}
}
