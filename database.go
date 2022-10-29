package discogs

import (
	"context"
	"fmt"
	"net/http"
)

type SearchResult struct {
	ID             uint       `json:"id"`
	Type           string     `json:"type"`
	UserData       *UserData  `json:"user_data"`
	MasterID       uint       `json:"master_id"`
	MasterURL      string     `json:"master_url"`
	URI            string     `json:"uri"`
	Title          string     `json:"title"`
	Thumb          string     `json:"thumb"`
	CoverImage     string     `json:"cover_image"`
	ResourceURL    string     `json:"resource_url"`
	Country        string     `json:"country"`
	Year           string     `json:"year"`
	Format         []string   `json:"format"`
	Label          []string   `json:"label"`
	Genre          []string   `json:"genre"`
	Barcode        []string   `json:"barcode"`
	Catno          string     `json:"catno"`
	Community      *Community `json:"community"`
	Style          []string   `json:"style"`
	FormatQuantity uint       `json:"format_quantity"`
	Formats        []*Format  `json:"formats"`
}
type SearchRequest struct {
	Query        *string `json:"query, omitempty"`
	Type         *string `json:"type, omitempty"`
	Title        *string `json:"title, omitempty"`
	ReleaseTitle *string `json:"release_title, omitempty"`
	Credit       *string `json:"credit,omitempty"`
	Artist       *string `json:"artist,omitempty"`
	Anv          *string `json:"anv,omitempty"`
	Label        *string `json:"label,omitempty"`
	Genre        *string `json:"genre,omitempty"`
	Style        *string `json:"style,omitempty"`
	Country      *string `json:"country,omitempty"`
	Year         *string `json:"year,omitempty"`
	Format       *string `json:"format,omitempty"`
	Catno        *string `json:"catno,omitempty"`
	Barcode      *string `json:"barcode,omitempty"`
	Track        *string `json:"track,omitempty"`
	Submitter    *string `json:"submitter,omitempty"`
	Contributor  *string `json:"contributor,omitempty"`
}

type UserData struct {
	InWantlist   bool `json:"in_wantlist"`
	InCollection bool `json:"in_collection"`
}

type SearchResultList struct {
	Pagination *Pagination     `json:"pagination"`
	Results    []*SearchResult `json:"results"`
}

// GetSearchResults is a function for searching the discogs database
func (c *Client) GetSearchResults(ctx context.Context, searchRequest SearchRequest) (*SearchResultList, error) {
	searchRequestString := buildSearchString(searchRequest)
	searchRequestString = fmt.Sprintf("%s/%s", c.baseURL, searchRequestString)

	req, err := http.NewRequest("GET", searchRequestString, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := SearchResultList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func buildSearchString(searchRequest SearchRequest) string {
	searchString := "/database/search?"
	if searchRequest.Query != nil {
		searchString += fmt.Sprintf("q=%v&", searchRequest.Query)
	}
	if searchRequest.Type != nil {
		searchString += fmt.Sprintf("type=%v&", searchRequest.Type)
	}
	if searchRequest.Title != nil {
		searchString += fmt.Sprintf("title=%v&", searchRequest.Title)
	}
	if searchRequest.ReleaseTitle != nil {
		searchString += fmt.Sprintf("release_title=%v&", searchRequest.ReleaseTitle)
	}
	if searchRequest.Credit != nil {
		searchString += fmt.Sprintf("credit=%v&", searchRequest.Credit)
	}
	if searchRequest.Artist != nil {
		searchString += fmt.Sprintf("artist=%v&", searchRequest.Artist)
	}
	if searchRequest.Anv != nil {
		searchString += fmt.Sprintf("anv=%v&", searchRequest.Anv)
	}
	if searchRequest.Label != nil {
		searchString += fmt.Sprintf("label=%v&", searchRequest.Label)
	}
	if searchRequest.Genre != nil {
		searchString += fmt.Sprintf("genre=%v&", searchRequest.Genre)
	}
	if searchRequest.Style != nil {
		searchString += fmt.Sprintf("style=%v&", searchRequest.Style)
	}
	if searchRequest.Country != nil {
		searchString += fmt.Sprintf("country=%v&", searchRequest.Country)
	}
	if searchRequest.Year != nil {
		searchString += fmt.Sprintf("year=%v&", searchRequest.Year)
	}
	if searchRequest.Catno != nil {
		searchString += fmt.Sprintf("catno=%v&", searchRequest.Catno)
	}
	if searchRequest.Barcode != nil {
		searchString += fmt.Sprintf("barcode=%v&", searchRequest.Barcode)
	}
	if searchRequest.Track != nil {
		searchString += fmt.Sprintf("track=%v&", searchRequest.Track)
	}
	if searchRequest.Submitter != nil {
		searchString += fmt.Sprintf("track=%v&", searchRequest.Submitter)
	}
	if searchRequest.Contributor != nil {
		searchString += fmt.Sprintf("contributor=%v&", searchRequest.Contributor)
	}
	return searchString
}
