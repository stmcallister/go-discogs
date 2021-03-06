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

// ArtistReleaseList represents a list of artist releases
type ArtistReleaseList struct {
	Pagination *Pagination      `json:"pagination"`
	Releases   []*ArtistRelease `json:"releases"`
}

// ArtistRelease represents a release object on the artists endpoint
type ArtistRelease struct {
	Artist      string `json:"artist"`
	ID          int    `json:"id"`
	MainRelease int    `json:"main_release"`
	ResourceURL string `json:"resource_url"`
	Role        string `json:"role"`
	Thumb       string `json:"thumb"`
	Title       string `json:"title"`
	Type        string `json:"master"`
	Year        int    `json:"year"`
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

// GetArtistReleases is a function for getting a single artist
func (c *Client) GetArtistReleases(ctx context.Context, sort, order string, ID, page, per int) (*ArtistReleaseList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/artists/%d/releases?sort=%s&order=%s&page=%d&per_page=%d", c.baseURL, ID, sort, order, page, per), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := ArtistReleaseList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetAllArtistReleases continues to call GetArtistReleases until all Items are returned
func (c *Client) GetAllArtistReleases(ctx context.Context, sort, order string, ID int) (*ArtistReleaseList, error) {
	page := 1
	per := 100
	res := new(ArtistReleaseList)
	res.Releases = make([]*ArtistRelease, 0)

	if temp, err := c.GetArtistReleases(ctx, sort, order, ID, page, per); err != nil {
		return nil, err
	} else {
		res.Releases = append(res.Releases, temp.Releases...)

		for temp.Pagination.Pages > 1 && temp.Pagination.Page < temp.Pagination.Pages {
			// increase page
			nextPage := temp.Pagination.Page + 1
			if temp, err = c.GetArtistReleases(ctx, sort, order, ID, nextPage, per); err != nil {
				return nil, err
			}
			res.Releases = append(res.Releases, temp.Releases...)
		}
	}

	return res, nil
}
