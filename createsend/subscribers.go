package createsend

import (
	"fmt"
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
