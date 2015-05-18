package createsend

import (
	"fmt"
	"net/url"
	"strconv"
)

// CampaignRecipientsOptions represents the URL parameters that may be used to
// fetch campaign recipients.
//
// See https://www.campaignmonitor.com/api/campaigns/#campaign_recipients for
// more information.
type CampaignRecipientsOptions struct {
	Page           int
	PageSize       int
	OrderField     string
	OrderDirection string
}

// CampaignRecipients lists all the recipients from a campaign.
// See https://www.campaignmonitor.com/api/campaigns/#campaign_recipients for
// more information.
type CampaignRecipients struct {
	Results              []*Recipient `json:"Results"`
	ResultsOrderedBy     string       `json:"ResultsOrderedBy"`
	OrderDirection       string       `json:"OrderDirection"`
	PageNumber           int          `json:"PageNumber"`
	PageSize             int          `json:"PageSize"`
	RecordsOnThisPage    int          `json:"RecordsOnThisPage"`
	TotalNumberOfRecords int          `json:"TotalNumberOfRecords"`
	NumberOfPages        int          `json:"NumberOfPages"`
}

type Recipient struct {
	EmailAddress string `json:"EmailAddress"`
	ListID       string `json:"ListID"`
}

func (c *APIClient) CampaignRecipients(campaignID string, opt *CampaignRecipientsOptions) (*CampaignRecipients, error) {

	u := fmt.Sprintf("campaigns/%s/recipients.json", campaignID)

	if opt != nil {
		v := url.Values{}
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

	var results CampaignRecipients
	err = c.Do(req, &results)
	return &results, err
}
