package discogs

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

// {"id": 0, "name": "All", "count": 240, "resource_url": "https://api.discogs.com/users/stmcallister/collection/folders/0"}
// Folder is a folder object inside a user collection
type Folder struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Count       int    `json:"count"`
	ResourceURL string `json:"resource_url"`
}

// FolderList is the list response to grabbing a list of folders
type FolderList struct {
	Folders []*Folder `jsont:"folders"`
}

// ReleaseList is the list response for releases
type ReleaseList struct {
	Pagination *Pagination          `json:"pagination"`
	Releases   []*CollectionRelease `json:"releases"`
}

// CollectionRelease is slightly different than a regular Release object
type CollectionRelease struct {
	ID               int               `json:"id"`
	InstanceID       int               `json:"instance_id"`
	Rating           int               `json:"rating"`
	BasicInformation *BasicInformation `json:"basic_information"`
	FolderID         int               `json:"folder_id"`
	DateAdded        string            `json:"date_added"`
}

// BasicInformation is a stripped down version of the release object used in the CollectionRelease obj
type BasicInformation struct {
	Title     string    `json:"title"`
	Year      int       `json:"year"`
	Artists   []*Artist `json:"artists"`
	Labels    []*Entity `json:"labels"`
	Formats   []*Format `json:"formats"`
	Genres    []string  `json:"genres"`
	Styles    []string  `json:"styles"`
	MasterID  int       `json:"master_id"`
	MasterURL string    `json:"master_url"`
}

// Pagination is the struct that represents pagination objects in list objects (this will probably go somwhere else eventually)
type Pagination struct {
	PerPage int         `json:"per_page"`
	Items   int         `json:"items"`
	Page    int         `json:"page"`
	URLs    interface{} `json:"urls"`
	Pages   int         `json:"pages"`
}

// GetUserCollectionFolders is a function for getting a single release
func (c *Client) GetUserCollectionFolders(ctx context.Context, username string) (*FolderList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users/%s/collection/folders", c.baseURL, username), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := FolderList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetUserCollectionFolder is a function for getting a single collection folder
func (c *Client) GetUserCollectionFolder(ctx context.Context, username string, folderID int) (*Folder, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users/%s/collection/folders/%d", c.baseURL, username, folderID), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := Folder{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetUserCollectionItemsByRelease is a function for getting a single release
func (c *Client) GetUserCollectionItemsByRelease(ctx context.Context, username string, releaseID int) (*ReleaseList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users/%s/collection/releases/%d", c.baseURL, username, releaseID), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := ReleaseList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetUserCollectionItemsByFolder is a function for getting a single release
func (c *Client) GetUserCollectionItemsByFolder(ctx context.Context, username, sort string, folderID, page, per int) (*ReleaseList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users/%s/collection/folders/%d/releases?page=%d&sort=%s&per_page=%d", c.baseURL, username, folderID, page, sort, per), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := ReleaseList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetUserCollectionAllItemsByFolder continues to call GetUserCollectionItemsByFolder until all Items are returned
func (c *Client) GetUserCollectionAllItemsByFolder(ctx context.Context, username, sort string, folderID int) (*ReleaseList, error) {
	page := 1
	per := 100
	res := new(ReleaseList)
	res.Releases = make([]*CollectionRelease, 0)
	key := os.Getenv("DISCOGS_API_KEY")
	client := NewClient(key)

	if temp, err := client.GetUserCollectionItemsByFolder(ctx, username, sort, folderID, page, per); err != nil {
		return nil, err
	} else {
		res.Releases = append(res.Releases, temp.Releases...)

		for temp.Pagination.Pages > 1 && temp.Pagination.Page < temp.Pagination.Pages {
			// increase page
			fmt.Printf("> > > Moody Page: %d\n", temp.Pagination.Page)
			nextPage := temp.Pagination.Page + 1
			if temp, err = client.GetUserCollectionItemsByFolder(ctx, username, sort, folderID, nextPage, per); err != nil {
				return nil, err
			}
			res.Releases = append(res.Releases, temp.Releases...)
		}
	}

	return res, nil
}
