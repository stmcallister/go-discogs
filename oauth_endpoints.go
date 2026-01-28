package discogs

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// RequestToken starts the OAuth 1.0a flow.
func (c *Client) RequestToken(ctx context.Context, callbackURL string) (token string, secret string, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/oauth/request_token", nil)
	if err != nil {
		return "", "", err
	}
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}
	req.Header.Set("Authorization", oauth1AuthorizationHeader(http.MethodPost, req.URL.String(), c.consumerKey, c.consumerSecret, "", "", map[string]string{
		"oauth_callback": callbackURL,
	}))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return "", "", fmt.Errorf("status %d: %s", res.StatusCode, strings.TrimSpace(string(body)))
	}
	vals, err := url.ParseQuery(string(body))
	if err != nil {
		return "", "", err
	}
	token = vals.Get("oauth_token")
	secret = vals.Get("oauth_token_secret")
	if token == "" || secret == "" {
		return "", "", fmt.Errorf("unexpected response: %s", strings.TrimSpace(string(body)))
	}
	return token, secret, nil
}

// AccessToken completes the OAuth 1.0a flow.
func (c *Client) AccessToken(ctx context.Context, requestToken, requestSecret, verifier string) (token string, secret string, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/oauth/access_token", nil)
	if err != nil {
		return "", "", err
	}
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}
	req.Header.Set("Authorization", oauth1AuthorizationHeader(http.MethodPost, req.URL.String(), c.consumerKey, c.consumerSecret, requestToken, requestSecret, map[string]string{
		"oauth_token":    requestToken,
		"oauth_verifier": verifier,
	}))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return "", "", fmt.Errorf("status %d: %s", res.StatusCode, strings.TrimSpace(string(body)))
	}
	vals, err := url.ParseQuery(string(body))
	if err != nil {
		return "", "", err
	}
	token = vals.Get("oauth_token")
	secret = vals.Get("oauth_token_secret")
	if token == "" || secret == "" {
		return "", "", fmt.Errorf("unexpected response: %s", strings.TrimSpace(string(body)))
	}
	return token, secret, nil
}

type Identity struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

// Identity fetches the authenticated user's identity.
func (c *Client) Identity(ctx context.Context) (*Identity, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/oauth/identity", nil)
	if err != nil {
		return nil, err
	}
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", oauth1AuthorizationHeader(http.MethodGet, req.URL.String(), c.consumerKey, c.consumerSecret, c.oauthToken, c.oauthTokenSecret, nil))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		b, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("status %d: %s", res.StatusCode, strings.TrimSpace(string(b)))
	}
	var out Identity
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

