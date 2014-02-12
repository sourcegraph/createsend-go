package createsend

import (
	"fmt"
	"time"
)

// List represents a subscriber list.
//
// See http://www.campaignmonitor.com/api/clients/#subscriber_lists for more
// information.
type List struct {
	ListID string
	Name   string
}

type SubscriberGroup string

const (
	ActiveSubscribers       SubscriberGroup = "active"
	UnconfirmedSubscribers                  = "unconfirmed"
	UnsubscribedSubscribers                 = "unsubscribed"
	BouncedSubscribers                      = "bounced"
	DeletedSubscribers                      = "deleted"
)

// ListSubcribersOptions represents the URL parameters that may be used to
// filter a subscriber list.
//
// See http://www.campaignmonitor.com/api/lists/#unconfirmed_subscribers for
// more information.
type ListSubscribersOptions struct {
	Date           time.Time
	Page           int
	PageSize       int
	OrderField     string
	OrderDirection string
}

// ListSubscribers lists all of the subscribers (in a given group, such as
// ActiveSubscribers, UnconfirmedSubscribers, etc.).
//
// See http://www.campaignmonitor.com/api/lists/#active_subscribers for more
// information.
func (c *APIClient) ListSubscribers(listID string, group SubscriberGroup, opt *ListSubscribersOptions) ([]*Subscriber, error) {
	if opt != nil {
		panic("opt is not yet implemented")
	}

	u := fmt.Sprintf("lists/%s/%s.json", listID, group)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var results struct {
		Results []*Subscriber
	}
	err = c.Do(req, &results)
	return results.Results, err
}
