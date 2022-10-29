package discogs

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Release represents a release object from discogs
type Release struct {
	ID                int              `json:"id"`
	Status            string           `json:"status"`
	Title             string           `json:"title"`
	Year              int              `json:"year"`
	ResourceURL       string           `json:"resource_url"`
	URI               string           `json:"uri"`
	Artists           []*ReleaseArtist `json:"artists"`
	ArtistsSort       string           `json:"artists_sort"`
	Series            []interface{}    `json:"series"`
	Labels            []*Entity        `json:"labels"`
	Companies         []*Entity        `json:"companies"`
	Formats           []*Format        `json:"formats"`
	DataQuality       string           `json:"data_quality"`
	Community         Community        `json:"community"`
	FormatQuantity    int              `json:"format_quantity"`
	DateAdded         string           `json:"date_added"`
	DateChanged       string           `json:"date_changed"`
	NumForSale        int              `json:"num_for_sale"`
	LowestPrice       float32          `json:"lowest_price"`
	MasterID          int              `json:"master_id"`
	MasterURL         string           `json:"master_url"`
	Country           string           `json:"country"`
	Released          string           `json:"released"`
	Notes             string           `json:"notes"`
	ReleasedFormatted string           `json:"released_formatted"`
	Identifiers       []*Identifier    `json:"identifiers"`
	Videos            []*Video         `json:"videos"`
	Genres            []string         `json:"genres"`
	Styles            []string         `json:"styles"`
	Tracklist         []*Track         `json:"tracklist"`
	ExtraArtists      []*ReleaseArtist `json:"extraartists"`
	Images            []*Image         `json:"images"`
	Thumb             string           `json:"thumb"`
	EstimatedWeight   int              `json:"estimated_weight"`
}

// Image represents image
type Image struct {
	Type        string `json:"type"`
	URI         string `json:"uri"`
	ResourceURL string `json:"resource_url"`
	URI150      string `json:"uri150"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
}

// Track represents track objects
type Track struct {
	Title    string `json:"title"`
	Position string `json:"position"`
	Type     string `json:"type_"`
	Duration string `json:"duration"`
}

// Video represents video objects
type Video struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URI         string `json:"uri"`
	Duration    int    `json:"duration"`
	Embed       bool   `json:"embed"`
}

// Identifier respresents identifier objects
type Identifier struct {
	Type        string `json:"type"`
	Value       string `json:"value"`
	Description string `json:"description, omitempty"`
}

// Community represents community object inside release object
type Community struct {
	Have         int                `json:"have"`
	Want         int                `json:"want"`
	Rating       Rating             `json:"rating"`
	Submitter    CommunityPerson    `json:"submitter"`
	Contributors []*CommunityPerson `json:"contributors"`
	DataQuality  string             `json:"data_quality"`
	Status       string             `json:"status"`
}

// Rating represents the rating of a release
type Rating struct {
	Count   int     `json:"count"`
	Average float32 `json:"average"`
}

// CommunityPerson represents generic community persons in release object
type CommunityPerson struct {
	Username    string `json:"username"`
	ResourceURL string `json:"resource_url"`
}

// Format represents format object in discogs
type Format struct {
	Name         string   `json:"name"`
	Qty          string   `json:"qty"`
	Text         string   `json:"text"`
	Descriptions []string `json:"descriptions"`
}

// Entity represents a generic entity object from discogs. represents labels and companies
type Entity struct {
	Name           string `json:"name"`
	Catno          string `json:"catno"`
	EntityType     string `json:"entity_type"`
	EntityTypeName string `json:"entity_type_name"`
	ID             int    `json:"id"`
	ResourceURL    string `json:"resource_url"`
}

// ReleaseArtist represents an artist object from discogs
type ReleaseArtist struct {
	Name        string `json:"name"`
	Anv         string `json:"anv"`
	Join        string `json:"join"`
	Role        string `json:"role"`
	Tracks      string `json:"tracks"`
	ID          int    `json:"id"`
	ResourceURL string `json:"resource_url"`
}

// ReleaseUserRating is a rating a user gives for a release
type ReleaseUserRating struct {
	ReleaseID int    `json:"release_id"`
	Username  string `json:"username"`
	Rating    int    `json:"rating"`
}

// ReleaseRating is a community rating for a release
type ReleaseRating struct {
	ReleaseID int     `json:"release_id"`
	Rating    *Rating `json:"rating"`
}

// ReleaseRatingRequest is a struct for the request object required to set a rating on a release
type ReleaseRatingRequest struct {
	Rating int `json:"rating"`
}

// ReleaseStats represents the stats for a release
type ReleaseStats struct {
	IsOffensive bool `json:"is_offensive"`
	NumHave     int  `json:"num_have"`
	NumWant     int  `json:"num_want"`
}

// GetRelease is a function for getting a single release
func (c *Client) GetRelease(ctx context.Context, ID int, currency string) (*Release, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/releases/%d?%s", c.baseURL, ID, currency), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := Release{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetReleaseRatingByUser is a function for getting a single release
func (c *Client) GetReleaseRatingByUser(ctx context.Context, ID int, username string) (*ReleaseUserRating, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/releases/%d/rating/%s", c.baseURL, ID, username), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	rat := ReleaseUserRating{}
	if err := c.sendRequest(req, &rat); err != nil {
		return nil, err
	}

	return &rat, nil
}

// GetReleaseCommunityRating is a function for getting a community rating of a single release
func (c *Client) GetReleaseCommunityRating(ctx context.Context, ID int) (*ReleaseRating, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/releases/%d/rating", c.baseURL, ID), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	rat := ReleaseRating{}
	if err := c.sendRequest(req, &rat); err != nil {
		return nil, err
	}

	return &rat, nil
}

// GetReleaseStats is a function for getting stats for a single release
func (c *Client) GetReleaseStats(ctx context.Context, ID int) (*ReleaseStats, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/releases/%d/stats", c.baseURL, ID), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	rat := ReleaseStats{}
	if err := c.sendRequest(req, &rat); err != nil {
		return nil, err
	}

	return &rat, nil
}

// UpdateReleaseUserRating is a function that updates a release's rating for a given user
func (c *Client) UpdateReleaseUserRating(ctx context.Context, ID, rating int, username string) (*ReleaseUserRating, error) {
	ratingReq := ReleaseRatingRequest{
		Rating: rating,
	}
	ratingJSON, err := json.Marshal(ratingReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/releases/%d/rating/%s", c.baseURL, ID, username), bytes.NewBuffer(ratingJSON))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	rat := ReleaseUserRating{}
	if err := c.sendRequest(req, &rat); err != nil {
		return nil, err
	}

	return &rat, nil
}

// DeleteReleaseUserRating is a function that updates a release's rating for a given user
func (c *Client) DeleteReleaseUserRating(ctx context.Context, ID int, username string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/releases/%d/rating/%s", c.baseURL, ID, username), nil)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)
	rat := ReleaseUserRating{}

	err = c.sendRequest(req, &rat)
	return err
}
