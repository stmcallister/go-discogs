package discogs

import (
	"context"
	"fmt"
	"net/http"
)

// Artist represents an artist in discogs
type Artist struct {
	Name           string   `json:"name"`
	RealName       string   `json:"real_name"`
	ID             int      `json:"id"`
	ResourceURL    string   `json:"resource_url"`
	URI            string   `json:"uri"`
	ReleasesURL    string   `json:"release_url"`
	Images         []*Image `json:"images"`
	Profile        string   `json:"profile"`
	URLs           []string `json:"urls"`
	NameVariations []string `json:"namevariations"`
	Aliases        []*Alias `json:"aliases"`
	Groups         []*Group `json:"groups"`
	DataQuality    string   `json:"data_quality"`
}

// Alias represents an alias for an artist
type Alias struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ResourceURL string `json:"resource_url"`
}

// Group represents a group for an artist
type Group struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ResourceURL string `json:"resource_url"`
	Active      bool   `json:"active"`
}

// GetArtist is a function for getting a single artist
func (c *Client) GetArtist(ctx context.Context, ID int) (*Artist, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/artists/%d", c.baseURL, ID), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := Artist{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
