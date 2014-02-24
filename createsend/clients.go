package createsend

import "fmt"

// A Client represents a client of a Campaign Monitor account.
//
// See http://www.campaignmonitor.com/api/account/#getting_your_clients for more
// information.
type Client struct {
	ClientID string
	Name     string
}

// ListClients lists the clients associated with the authenticated account.
//
// See http://www.campaignmonitor.com/api/account/#getting_your_clients for more
// information.
func (c *APIClient) ListClients() ([]Client, error) {
	u := "clients.json"

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	clients := new([]Client)
	err = c.Do(req, clients)
	if err != nil {
		return nil, err
	}

	return *clients, err
}

// ListLists returns all of the subscriber lists that belong to a client.
//
// See http://www.campaignmonitor.com/api/clients/#subscriber_lists for more
// information.
func (c *APIClient) ListLists(clientID string) ([]*List, error) {
	u := fmt.Sprintf("clients/%s/lists.json", clientID)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var lists []*List
	err = c.Do(req, &lists)
	return lists, err
}

// ListForEmail represents a subscriber list *and* a specific email address's
// subscription to that list. The schema differs from that of List.
//
// See http://www.campaignmonitor.com/api/clients/#lists_for_email for more
// information.
type ListForEmail struct {
	ListID                 string
	ListName               string
	SubscriberState        string
	DateSubscriberAddedStr string `json:"DateSubscriberAdded"`
}

func (e *ListForEmail) IsSubscribed() bool {
	return e.SubscriberState == "Active"
}

func (e *ListForEmail) IsUnsubscribed() bool {
	return e.SubscriberState == "Unsubscribed"
}

// ListsForEmail returns all of the client's subscriber lists to which the email
// address is subscribed.
//
// See http://www.campaignmonitor.com/api/clients/#lists_for_email for more
// information.
func (c *APIClient) ListsForEmail(clientID string, email string) ([]*ListForEmail, error) {
	u := fmt.Sprintf("clients/%s/listsforemail.json?email=%s", clientID, email)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var lists []*ListForEmail
	err = c.Do(req, &lists)
	if err != nil {
		return nil, err
	}

	return lists, err
}
