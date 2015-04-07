package createsend

import (
	"errors"
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

// ListDelete deletes a given list
//
// See https://www.campaignmonitor.com/api/lists/#deleting_a_list for more
// information.
func (c *APIClient) ListDelete(listID string) error {
	u := fmt.Sprintf("lists/%s.json", listID)

	req, err := c.NewRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	err = c.Do(req, nil)
	return err
}

type UnsubcribeSetting string

const (
	AllClientLists UnsubcribeSetting = "AllClientLists"
	OnlyThisList                     = "OnlyThisList"
)

// ListCreateOptions represents the parameters needed
// to create a new list.
//
// See https://www.campaignmonitor.com/api/lists/#creating_a_list for
// more information.
type ListCreateOptions struct {
	Title                   string            `json:"Title"`
	UnsubscribePage         string            `json:"UnsubscribePage"`
	UnsubscribeSetting      UnsubcribeSetting `json:"UnsubscribeSetting"`
	ConfirmedOptin          bool              `json:"ConfirmedOptin"`
	ConfirmationSuccessPage string            `json:"ConfirmationSuccessPage"`
}

// ListCreate creates a new list
//
// See https://www.campaignmonitor.com/api/lists/#creating_a_list for more
// information.
func (c *APIClient) ListCreate(clientID string, opt *ListCreateOptions) (string, error) {
	if opt.UnsubscribeSetting == "" {
		return "", errors.New("Unsubscribesetting not set")
	}

	u := fmt.Sprintf("lists/%s.json", clientID)

	req, err := c.NewRequest("POST", u, opt)
	if err != nil {
		return "", err
	}

	var v interface{}
	err = c.Do(req, &v)
	if err != nil {
		return "", err
	}

	s, ok := v.(string)
	if !ok {
		return "", errors.New("Returned value is not a string")
	}

	return s, nil
}
