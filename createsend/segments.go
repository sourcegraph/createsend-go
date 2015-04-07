package createsend

import "fmt"

type SegmentCreate struct {
	Title      string            `json:"Title"`
	RuleGroups []RuleGroupCreate `json:"RuleGroups"`
}

type RuleGroupCreate struct {
	Rules []RuleCreate `json:"Rules"`
}

type RuleCreate struct {
	RuleType string `json:"RuleType"`
	Clause   string `json:"Clause"`
}

// Create a new segment on the given list.
//
// See https://www.campaignmonitor.com/api/segments/#creating_a_segment for more
// information.
func (c *APIClient) SegmentCreate(listID string, sgmt *SegmentCreate) (string, error) {
	u := fmt.Sprintf("segments/%s.json", listID)

	req, err := c.NewRequest("POST", u, sgmt)
	if err != nil {
		return "", err
	}

	var r string
	err = c.Do(req, &r)
	if err != nil {
		return "", err
	}

	return r, nil
}

type SegmentDetail struct {
	ActiveSubscribers int               `json:"ActiveSubscribers"`
	RuleGroups        []RuleGroupCreate `json:"RuleGroups"`
	ListID            string            `json:"ListID"`
	SegmentID         string            `json:"SegmentID"`
	Title             string            `json:"Title"`
}

// Fetch segment details for a given segment.
//
// See https://www.campaignmonitor.com/api/segments/#getting_a_segments_details for more
// information.
func (c *APIClient) SegmentDetail(segmentID string) (*SegmentDetail, error) {
	u := fmt.Sprintf("segments/%s.json", segmentID)

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var s SegmentDetail
	err = c.Do(req, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil

}
