package createsend

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
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

type ListSubscribersResponse struct {
	Results              []*Subscriber
	ResultsOrderedBy     string
	OrderDirection       string
	PageNumber           int
	PageSize             int
	RecordsOnThisPage    int
	TotalNumberOfRecords int
	NumberOfPages        int
}

// ListSubscribers lists all of the subscribers (in a given group, such as
// ActiveSubscribers, UnconfirmedSubscribers, etc.).
//
// See http://www.campaignmonitor.com/api/lists/#active_subscribers for more
// information.
func (c *APIClient) ListSubscribers(listID string, group SubscriberGroup, opt *ListSubscribersOptions) (*ListSubscribersResponse, error) {
	u := fmt.Sprintf("lists/%s/%s.json", listID, group)

	if opt != nil {
		v := url.Values{}
		if !opt.Date.IsZero() {
			v.Set("date", opt.Date.Format("2006-01-02"))
		}
		if opt.Page > 0 {
			v.Set("page", strconv.Itoa(opt.Page))
		}
		if opt.PageSize > 0 {
			v.Set("pagesize", strconv.Itoa(opt.PageSize))
		}
		if opt.OrderField != "" {
			v.Set("orderfield", opt.OrderField)
		}
		if opt.OrderDirection != "" {
			v.Set("orderdirection", opt.OrderDirection)
		}

		q := v.Encode()
		if q != "" {
			u = fmt.Sprintf("%s?%s", u, q)
		}
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var results ListSubscribersResponse
	err = c.Do(req, &results)
	return &results, err
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

type DataType string

const (
	Text            DataType = "Text"
	Number                   = "Number"
	MultiSelectOne           = "MultiSelectOne"
	MultiSelectMany          = "MultiSelectMany"
	Date                     = "Date"
	Country                  = "Country"
	USState                  = "USState"
)

type CustomFieldDefinition struct {
	FieldName                 string   `json:"FieldName"`
	Key                       string   `json:"Key"`
	DataType                  DataType `json:"DataType"`
	FieldOptions              []string `json:"FieldOptions"`
	VisibleInPreferenceCenter bool     `json:"VisibleInPreferenceCenter"`
}

// ListCustomFields returns a list of CustomFields for the given list.
//
// See https://www.campaignmonitor.com/api/lists/#list_custom_fields for
// more information.
func (c *APIClient) ListCustomFields(listID string) ([]CustomFieldDefinition, error) {
	u := fmt.Sprintf("lists/%s/customfields.json", listID)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var result []CustomFieldDefinition
	err = c.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type CustomFieldCreate struct {
	FieldName                 string   `json:"FieldName"`
	DataType                  DataType `json:"DataType"`
	Options                   []string `json:"Options,omitempty"`
	VisibleInPreferenceCenter bool     `json:"VisibleInPreferenceCenter"`
}

// ListCreateCustomField creates a new CustomField on the given list.
//
// See https://www.campaignmonitor.com/api/lists/#creating_a_custom_field for
// more information.
func (c *APIClient) ListCreateCustomField(listID string, def *CustomFieldCreate) (string, error) {
	u := fmt.Sprintf("lists/%s/customfields.json", listID)

	req, err := c.NewRequest("POST", u, def)
	if err != nil {
		return "", err
	}

	var v interface{}
	err = c.Do(req, &v)
	if err != nil {
		return "", err
	}

	r, ok := v.(string)
	if !ok {
		return "", errors.New("Return is not a string")
	}

	return r, nil
}

// ListDeleteCustomField deletes a CustomField from a given list.
//
// See https://www.campaignmonitor.com/api/lists/#deleting_a_custom_field for
// more information.
func (c *APIClient) ListDeleteCustomField(listID string, cfKey string) error {
	u := fmt.Sprintf("lists/%s/customfields/%s.json", listID, cfKey)

	req, err := c.NewRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	err = c.Do(req, nil)

	return err
}

type ListSegment struct {
	ListID    string `json:"ListID"`
	SegmentID string `json:"SegmentID"`
	Title     string `json:"Title"`
}

// ListSegments list the segments for a given list.
//
// See https://www.campaignmonitor.com/api/lists/#list_segments for
// more information.
func (c *APIClient) ListSegments(listID string) ([]ListSegment, error) {
	u := fmt.Sprintf("lists/%s/segments.json", listID)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var result []ListSegment
	err = c.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type WebhookCreate struct {
	Events        []string `json:"Events"`
	Url           string   `json:"Url"`
	PayloadFormat string   `json:"PayloadFormat"`
}

type Webhook struct {
	WebhookCreate
	WebhookID string `json:"WebhookID"`
	Status    string `json:"Status"`
}

// ListWebhooks lists the webhooks for a given list.
//
// See https://www.campaignmonitor.com/api/lists/#list_webhooks for
// more information.
func (c *APIClient) ListWebhooks(listID string) ([]Webhook, error) {
	u := fmt.Sprintf("lists/%s/webhooks.json", listID)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var result []Webhook
	err = c.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// ListCreateWebhook creates a new webhook for a given list.
//
// See https://www.campaignmonitor.com/api/lists/#list_webhooks for
// more information.
func (c *APIClient) ListCreateWebhook(listID string, webhook *WebhookCreate) (string, error) {
	u := fmt.Sprintf("lists/%s/webhooks.json", listID)

	req, err := c.NewRequest("POST", u, webhook)
	if err != nil {
		return "", err
	}

	var result string
	err = c.Do(req, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

// ListTestWebhook tests a given webhook for a given list.
//
// See https://www.campaignmonitor.com/api/lists/#testing_a_webhook for
// more information.
func (c *APIClient) ListTestWebhook(listID string, webhookID string) error {
	u := fmt.Sprintf("lists/%s/webhooks/%s/test.json", listID, webhookID)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return err
	}

	return c.Do(req, nil)
}

// ListDeleteWebhook deletes a given webhook for a given list.
//
// See https://www.campaignmonitor.com/api/lists/#deleting_a_webhook for
// more information.
func (c *APIClient) ListDeleteWebhook(listID string, webhookID string) error {
	u := fmt.Sprintf("lists/%s/webhooks/%s.json", listID, webhookID)

	req, err := c.NewRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	return c.Do(req, nil)
}

// ListActivateWebhook actives a given webhook for a given list.
//
// See https://www.campaignmonitor.com/api/lists/#activating_a_webhook for
// more information.
func (c *APIClient) ListActivateWebhook(listID string, webhookID string) error {
	u := fmt.Sprintf("lists/%s/webhooks/%s/activate.json", listID, webhookID)

	req, err := c.NewRequest("PUT", u, nil)
	if err != nil {
		return err
	}

	return c.Do(req, nil)
}

// ListDeactivateWebhook deactivates a given webhook for a given list.
//
// See https://www.campaignmonitor.com/api/lists/#deactivating_a_webhook for
// more information.
func (c *APIClient) ListDeactivateWebhook(listID string, webhookID string) error {
	u := fmt.Sprintf("lists/%s/webhooks/%s/deactivate.json", listID, webhookID)

	req, err := c.NewRequest("PUT", u, nil)
	if err != nil {
		return err
	}

	return c.Do(req, nil)
}
