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

// MasterVersion is the version object inside of the master object
type MasterVersion struct {
	Status       string        `json:"status"`
	Stats        *VersionStats `json:"stats"`
	Thumb        string        `json:"thumb"`
	Format       string        `json:"format"`
	Country      string        `json:"country"`
	Title        string        `json:"title"`
	Label        string        `json:"label"`
	Released     string        `json:"released"`
	MajorFormats []*string     `json:"major_formats"`
	Catno        string        `json:"catno"`
	ResourceURL  string        `json:"resource_url"`
	ID           int           `json:"id"`
}

// VersionStats holds the user and community stats for a master version
type VersionStats struct {
	User      *Stats `json:"user"`
	Community *Stats `json:"community"`
}

// Stats is the generic object that defines stats for a master version
type Stats struct {
	InCollection int `json:"in_collection"`
	InWantlist   int `json:"in_wantlist"`
}

// VersionList is the list response for versions
type VersionList struct {
	Pagination *Pagination      `json:"pagination"`
	Versions   []*MasterVersion `json:"versions"`
}

// GetMasterRelease is a function for getting a single release
func (c *Client) GetMasterRelease(ctx context.Context, ID int) (*Master, error) {
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

// GetMasterReleaseVersions is a function for getting a single release
func (c *Client) GetMasterReleaseVersions(ctx context.Context, ID, page, per int, sort, sortOrder string) (*VersionList, error) {
	// note: only added parameters for pagination, didn't want to initially confuse things by adding all possible parameters, when the others
	// didn't feel as widely used. please post issue or PR to project if you need the other parameters

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/masters/%d/versions?page=%d&per_page=%d&sort=%s&sort_order=%s", c.baseURL, ID, page, per, sort, sortOrder), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := VersionList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetMasterReleaseVersionsAllPages is a function for getting a single release
func (c *Client) GetMasterReleaseVersionsAllPages(ctx context.Context, ID int, sort, sortOrder string) (*VersionList, error) {
	page := 1
	per := 100
	vers := new(VersionList)
	vers.Versions = make([]*MasterVersion, 0)

	if temp, err := c.GetMasterReleaseVersions(ctx, ID, page, per, sort, sortOrder); err != nil {
		return nil, err
	} else {
		vers.Versions = append(vers.Versions, temp.Versions...)
		for temp.Pagination.Pages > 1 && temp.Pagination.Page < temp.Pagination.Pages {
			// increase page
			nextPage := temp.Pagination.Page + 1
			if temp, err = c.GetMasterReleaseVersions(ctx, ID, nextPage, per, sort, sortOrder); err != nil {
				return nil, err
			}
			vers.Versions = append(vers.Versions, temp.Versions...)
		}

	}
	return vers, nil
}
