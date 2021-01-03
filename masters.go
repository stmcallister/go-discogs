package discogs

import (
	"context"
	"fmt"
	"net/http"
)

// Master represents a master release object from discogs
type Master struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	Year           int       `json:"year"`
	ResourceURL    string    `json:"resource_url"`
	URI            string    `json:"uri"`
	Artists        []*Artist `json:"artists"`
	Styles         []string  `json:"styles"`
	Videos         []*Video  `json:"videos"`
	Genres         []string  `json:"genres"`
	MainRelease    int       `json:"main_release"`
	MainReleaseURL string    `json:"main_release_url"`
	VersionsURL    string    `json:"versions_url"`
	Images         []*Image  `json:"images"`
	Tracklist      []*Track  `json:"tracklist"`
	DataQuality    string    `json:"data_quality"`
	NumForSale     int       `json:"num_for_sale"`
	LowestPrice    float32   `json:"lowest_price"`
}

// GetMaster is a function for getting a single release
func (c *Client) GetMaster(ctx context.Context, ID int) (*Master, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/masters/%d", c.baseURL, ID), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := Master{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
