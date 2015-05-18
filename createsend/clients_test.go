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
		fmt.Fprint(w, `[{"ListID": "34cd", "ListName": "mylist", "SubscriberState": "Active"}]`)
	})

	lists, err := client.ListsForEmail("12ab", "alice@example.com")
	if err != nil {
		t.Errorf("ListsForEmail returned error: %v", err)
	}

	want := []*ListForEmail{{ListID: "34cd", ListName: "mylist", SubscriberState: "Active"}}
	if !reflect.DeepEqual(lists, want) {
		t.Errorf("ListsForEmail returned %+v, want %+v", lists, want)
	}
}

func TestCampaigns(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/clients/12ab/campaigns.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `
				[
					{
						"FromName": "My Name",
						"FromEmail": "myemail@example.com",
						"ReplyTo": "myemail@example.com",
						"WebVersionURL": "http://createsend.com/t/r-765E86829575EE2C/",
						"WebVersionTextURL": "http://createsend.com/t/r-765E86829575EE2C/t",
						"CampaignID": "fc0ce7105baeaf97f47c99be31d02a91",
						"Subject": "Campaign One",
						"Name": "Campaign One",
						"SentDate": "2010-10-12 12:58:00",
						"TotalRecipients": 2245
					},
					{
						"FromName": "My Name",
						"FromEmail": "myemail@example.com",
						"ReplyTo": "myemail@example.com",
						"WebVersionURL": "http://createsend.com/t/r-DD543566A87C9B8B/",
						"WebVersionTextURL": "http://createsend.com/t/r-DD543566A87C9B8B/t",
						"CampaignID": "072472b88c853ae5dedaeaf549a8d607",
						"Subject": "Campaign Two",
						"Name": "Campaign Two",
						"SentDate": "2010-10-06 16:20:00",
						"TotalRecipients": 11222
					}
				]`)
	})

	campaigns, err := client.Campaigns("12ab")
	if err != nil {
		t.Errorf("Campaigns returned error: %v", err)
	}

	want := []*Campaign{
		{
			FromName:          "My Name",
			FromEmail:         "myemail@example.com",
			ReplyTo:           "myemail@example.com",
			WebVersionURL:     "http://createsend.com/t/r-765E86829575EE2C/",
			WebVersionTextURL: "http://createsend.com/t/r-765E86829575EE2C/t",
			CampaignID:        "fc0ce7105baeaf97f47c99be31d02a91",
			Subject:           "Campaign One",
			Name:              "Campaign One",
			SentDate:          "2010-10-12 12:58:00",
			TotalRecipients:   2245,
		},
		{
			FromName:          "My Name",
			FromEmail:         "myemail@example.com",
			ReplyTo:           "myemail@example.com",
			WebVersionURL:     "http://createsend.com/t/r-DD543566A87C9B8B/",
			WebVersionTextURL: "http://createsend.com/t/r-DD543566A87C9B8B/t",
			CampaignID:        "072472b88c853ae5dedaeaf549a8d607",
			Subject:           "Campaign Two",
			Name:              "Campaign Two",
			SentDate:          "2010-10-06 16:20:00",
			TotalRecipients:   11222,
		},
	}

	if !reflect.DeepEqual(campaigns, want) {
		t.Errorf("Campaigns return %+v, want %+v", campaigns, want)
	}
}
