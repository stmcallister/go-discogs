package discogs

import (
	"context"
	"fmt"
	"net/http"
)

type Label struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	ResourceURL string      `json:"resource_url"`
	URI         string      `json:"uri"`
	ReleasesURL string      `json:"releases_url"`
	Images      []*Image    `json:"images"`
	ContactInfo string      `json:"contact_info"`
	Profile     string      `json:"profile"`
	DataQuality string      `json:"data_quality"`
	URLs        []string    `json:"urls"`
	SubLabels   []*SubLabel `json:"sublabels"`
}

type SubLabel struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ResourceURL string `json:"resource_url"`
}

// GetLabel is a function for getting a single record label
func (c *Client) GetLabel(ctx context.Context, ID int) (*Label, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/labels/%d", c.baseURL, ID), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := Label{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
