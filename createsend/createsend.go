package createsend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "createsend-go/" + libraryVersion
	defaultBaseURL = "https://api.createsend.com/api/v3.1/"
)

// A APIClient manages communication with the Campaign Monitor API.
type APIClient struct {
	// client is the HTTP client used to communicate with the API.
	client *http.Client

	// BaseURL for API requests. Defaults to the public Campaign Monitor V3 API.
	BaseURL *url.URL

	// UserAgent used when communicating with the Campaign Monitor API.
	UserAgent string

	// Log is used to log debugging messages, if set.
	Log *log.Logger
}

// NewAPIClient returns a new Campaign Monitor API client. If a nil httpClient
// is provided, http.DefaultClient will be used. To use API methods which
// require authentication, provide an http.Client that will perform the
// authentication for you (such as that provided by the goauth2 library).
func NewAPIClient(httpClient *http.Client) *APIClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &APIClient{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}
	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the APIClient.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *APIClient) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.UserAgent)
	return req, nil
}

type CreatesendError struct {
	Code       int
	Message    string
	ResultData interface{}
}

func (e *CreatesendError) Error() string {
	return fmt.Sprintf("%s (createsend error %d)", e.Message, e.Code)
}

// Do sends an API request and returns the API response. The API response is
// decoded and stored in the value pointed to by v, or returned as an error if
// an API error has occurred.
func (c *APIClient) Do(req *http.Request, v interface{}) error {
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		var e CreatesendError
		err = json.NewDecoder(resp.Body).Decode(&e)
		if err != nil {
			return err
		}
		return &e
	} else if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		if c.Log != nil {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Printf("ReadAll failed: %s")
			}
			c.Log.Printf("http response %d body:\n%s", resp.StatusCode, body)
		}
		return fmt.Errorf("http response status code %d", resp.StatusCode)
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	return err
}
