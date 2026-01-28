package discogs

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	BaseURL = "https://api.discogs.com"
)

// Client .
type Client struct {
	apiKey           string
	consumerKey      string
	consumerSecret   string
	oauthToken       string
	oauthTokenSecret string

	baseURL    string
	userAgent  string
	HTTPClient *http.Client
}

// NewClient .
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Minute,
		},
		baseURL: BaseURL,
	}
}

// NewOAuthClient creates a client that signs requests with OAuth 1.0a.
// Pass empty accessToken/accessSecret for the initial request token step.
func NewOAuthClient(consumerKey, consumerSecret, accessToken, accessSecret string) *Client {
	return &Client{
		consumerKey:      consumerKey,
		consumerSecret:   consumerSecret,
		oauthToken:       accessToken,
		oauthTokenSecret: accessSecret,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Minute,
		},
		baseURL: BaseURL,
	}
}

// WithUserAgent sets a Discogs-compliant user agent (required by Discogs policy).
func (c *Client) WithUserAgent(ua string) *Client {
	c.userAgent = ua
	return c
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}

	// Auth: prefer OAuth if configured; else fall back to personal token.
	if c.consumerKey != "" && c.consumerSecret != "" {
		req.Header.Set("Authorization", oauth1AuthorizationHeader(req.Method, req.URL.String(), c.consumerKey, c.consumerSecret, c.oauthToken, c.oauthTokenSecret, nil))
	} else {
		req.Header.Set("Authorization", fmt.Sprintf("Discogs token=%s", c.apiKey))
	}

	// debugging
	//dump, err := httputil.DumpRequestOut(req, true)
	//if err != nil {
	//	fmt.Printf(err.Error())
	//}
	//fmt.Printf("%s\n\n", dump)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	// Try to unmarshall into errorResponse
	if res.StatusCode != http.StatusOK {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	// Unmarshall and populate v
	fullResponse := successResponse{
		Code: res.StatusCode,
		Data: v,
	}

	// debugging
	//if resDump, err := httputil.DumpResponse(res, true); err != nil {
	//	fmt.Printf(err.Error())
	//} else {
	//	fmt.Printf("%s\n\n", resDump)
	//}

	if err = json.NewDecoder(res.Body).Decode(&fullResponse.Data); err != nil {
		return err
	}
	return nil
}
