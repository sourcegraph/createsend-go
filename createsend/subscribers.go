package createsend

import (
	"fmt"
	"time"
)

// NewSubscriber represents a new subscriber to be added with AddSubscriber.
//
// See http://www.campaignmonitor.com/api/subscribers/#adding_a_subscriber for
// more information.
type NewSubscriber struct {
	EmailAddress                           string
	Name                                   string        `json:",omitempty"`
	CustomFields                           []CustomField `json:",omitempty"`
	Resubscribe                            bool          `json:",omitempty"`
	RestartSubscriptionBasedAutoresponders bool          `json:",omitempty"`
}

// CustomField represents a subscriber custom data field.
//
// See http://www.campaignmonitor.com/api/subscribers/#adding_a_subscriber for
// more information.
type CustomField struct {
	Key   string
	Value interface{}
}

// AddSubscriber adds a subscriber.
//
// See http://www.campaignmonitor.com/api/subscribers/#adding_a_subscriber for
// more information.
func (c *APIClient) AddSubscriber(listID string, sub NewSubscriber) error {
	u := fmt.Sprintf("subscribers/%s.json", listID)

	req, err := c.NewRequest("POST", u, sub)
	if err != nil {
		return err
	}

	return c.Do(req, nil)
}

// UpdateSubscriber updates a subscriber.
//
// See http://www.campaignmonitor.com/api/subscribers/#updating_a_subscriber for
// more information.
func (c *APIClient) UpdateSubscriber(listID string, email string, sub NewSubscriber) error {
	u := fmt.Sprintf("subscribers/%s.json?email=%s", listID, email)

	req, err := c.NewRequest("PUT", u, sub)
	if err != nil {
		return err
	}

	return c.Do(req, nil)
}

// Subscriber represents a subscriber.
//
// See
// http://www.campaignmonitor.com/api/subscribers/#getting_a_subscribers_details
// for more information.
type Subscriber struct {
	EmailAddress   string
	Name           string        `json:",omitempty"`
	Date           time.Time     `json:"-"`
	State          string        `json:",omitempty"`
	CustomFields   []CustomField `json:",omitempty"`
	ReadsEmailWith string        `json:",omitempty"`

	// DateStr holds the createsend API's date format, which is "2010-10-25
	// 10:28:00". This is not the format that encoding/json expects, so we must
	// parse it separately. The parsed date is stored in the Date field, and the
	// RFC3339 date string is overwritten into this field.
	DateStr string `json:"date,omitempty"`
}

// GetSubscriber gets a subscriber's details.
//
// See
// http://www.campaignmonitor.com/api/subscribers/#getting_a_subscribers_details
// for more information.
func (c *APIClient) GetSubscriber(listID string, email string) (*Subscriber, error) {
	u := fmt.Sprintf("subscribers/%s.json?email=%s", listID, email)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var sub Subscriber
	err = c.Do(req, &sub)
	if err != nil {
		return nil, err
	}

	// Parse createsend API date format. (See Subscriber.DateStr field comment.)
	sub.Date, err = time.Parse("2006-01-02 15:04:05", sub.DateStr)
	if err != nil {
		return nil, err
	}
	sub.DateStr = sub.Date.Format(time.RFC3339)

	return &sub, nil
}

// Unsubscribe changes the status of a subscriber from Active to Unsubscribed.
//
// See
// http://www.campaignmonitor.com/api/subscribers/#unsubscribing_a_subscriber
// for more information.
func (c *APIClient) Unsubscribe(listID string, email string) error {
	u := fmt.Sprintf("subscribers/%s/unsubscribe.json", listID)

	req, err := c.NewRequest("POST", u, struct{ EmailAddress string }{email})
	if err != nil {
		return err
	}

	return c.Do(req, nil)
}

// Delete removes the Subscriber from the specified list
//
// See
// https://www.campaignmonitor.com/api/subscribers/#deleting_a_subscriber
// for more information.
func (c *APIClient) DeleteSubscriber(listID string, email string) error {
	u := fmt.Sprintf("subscribers/%s.json?email=%s", listID, email)

	req, err := c.NewRequest("DELETE", u, struct{ EmailAddress string }{email})
	if err != nil {
		return err
	}

	return c.Do(req, nil)
}

// NewSubscriber represents a new subscriber to be added with AddSubscriber.
//
// See http://www.campaignmonitor.com/api/subscribers/#adding_a_subscriber for
// more information.
type ImportSubscriber struct {
	EmailAddress string
	Name         string        `json:",omitempty"`
	CustomFields []CustomField `json:",omitempty"`
}

type ImportSubscribers struct {
	Subscribers                            []ImportSubscriber
	Resubscribe                            bool `json:",omitempty"`
	QueueSubscriptionBasedAutoResponders   bool `json:",omitempty"`
	RestartSubscriptionBasedAutoresponders bool `json:",omitempty"`
}

// Importing many subscribes
//
// See
// https://www.campaignmonitor.com/api/subscribers/#importing_many_subscribers
// for more information.
func (c *APIClient) ImportSubscribers(listID string, importSubscribers ImportSubscribers) (interface{}, error) {
	u := fmt.Sprintf("subscribers/%s/import.json", listID)

	req, err := c.NewRequest("POST", u, importSubscribers)
	if err != nil {
		return nil, err
	}

	var v interface{}
	err = c.Do(req, &v)
	return v, err
}
