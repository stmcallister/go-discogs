package discogs

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestDiscogs_UsersGetCollectionFolders(t *testing.T) {
	// todo: make real unit tests using MUX
	key := os.Getenv("DISCOGS_API_KEY")
	client := NewClient(key)
	ctx := context.Background()
	username := "stmcallister"

	if res, err := client.GetUserCollectionFolders(ctx, username); err != nil {
		panic(err)
	} else {
		for _, f := range res.Folders {
			fmt.Printf("\nFolder name: %s (%d)", f.Name, f.ID)
			fmt.Printf("\nFolder count: %d \n", f.Count)
			fmt.Println("--")
		}
	}
}

func TestDiscogs_UsersGetCollectionFolder(t *testing.T) {
	// todo: make real unit tests using MUX
	key := os.Getenv("DISCOGS_API_KEY")
	client := NewClient(key)
	ctx := context.Background()
	username := "stmcallister"
	folderID := 2368501

	if res, err := client.GetUserCollectionFolder(ctx, username, folderID); err != nil {
		panic(err)
	} else {
		fmt.Printf("\nFolder name: %s (%d)", res.Name, res.ID)
		fmt.Printf("\nFolder count: %d \n", res.Count)
	}
}

func TestDiscogs_UsersGetUserCollectionItemsByRelease(t *testing.T) {
	// todo: make real unit tests using MUX
	key := os.Getenv("DISCOGS_API_KEY")
	client := NewClient(key)
	ctx := context.Background()
	username := "stmcallister"
	rid := 1964222

	if res, err := client.GetUserCollectionItemsByRelease(ctx, username, rid); err != nil {
		panic(err)
	} else {
		fmt.Printf("\n Pagination Per Page: %d ", res.Pagination.PerPage)
		fmt.Printf("\n Pagination Items: %d ", res.Pagination.Items)
		fmt.Printf("\n Pagination Page: %d ", res.Pagination.Page)
		fmt.Printf("\n Pagination URLs: %v ", res.Pagination.URLs)
		fmt.Printf("\n Pagination Pages: %d \n", res.Pagination.Pages)

		for _, rel := range res.Releases {
			fmt.Printf("\n\nID: %d\n", rel.ID)
			fmt.Printf("Title: %s\n", rel.BasicInformation.Title)
			fmt.Printf("Year: %d\n", rel.BasicInformation.Year)
			for _, artist := range rel.BasicInformation.Artists {
				fmt.Printf("Artist Name: %s (%d)\n", artist.Name, artist.ID)
			}
			for _, label := range rel.BasicInformation.Labels {
				fmt.Printf("Label Name: %s (%d)\n", label.Name, label.ID)
				fmt.Printf("Label Catno: %s \n", label.Catno)
			}
			for _, f := range rel.BasicInformation.Formats {
				fmt.Printf("Format: %s \n", f.Name)
			}
			fmt.Printf("MasterID: %d \n", rel.BasicInformation.MasterID)
			fmt.Printf("MasterURL: %s \n", rel.BasicInformation.MasterURL)
			for _, g := range rel.BasicInformation.Genres {
				fmt.Printf("Genre: %s \n", g)
			}
			for _, s := range rel.BasicInformation.Styles {
				fmt.Printf("Style: %s \n", s)
			}
			fmt.Printf("FolderID: %d\n", rel.FolderID)
		}
	}
}

func TestDiscogs_UsersGetUserCollectionItemsByFolder(t *testing.T) {
	// todo: make real unit tests using MUX
	key := os.Getenv("DISCOGS_API_KEY")
	client := NewClient(key)
	ctx := context.Background()
	username := "stmcallister"
	fid := 2368501

	if res, err := client.GetUserCollectionItemsByFolder(ctx, username, fid); err != nil {
		panic(err)
	} else {
		fmt.Printf("\n Pagination Per Page: %d ", res.Pagination.PerPage)
		fmt.Printf("\n Pagination Items: %d ", res.Pagination.Items)
		fmt.Printf("\n Pagination Page: %d ", res.Pagination.Page)
		fmt.Printf("\n Pagination URLs: %v ", res.Pagination.URLs)
		fmt.Printf("\n Pagination Pages: %d \n", res.Pagination.Pages)

		for _, rel := range res.Releases {
			fmt.Printf("\n\nID: %d\n", rel.ID)
			fmt.Printf("Title: %s\n", rel.BasicInformation.Title)
			fmt.Printf("Year: %d\n", rel.BasicInformation.Year)
			for _, artist := range rel.BasicInformation.Artists {
				fmt.Printf("Artist Name: %s (%d)\n", artist.Name, artist.ID)
			}
			for _, label := range rel.BasicInformation.Labels {
				fmt.Printf("Label Name: %s (%d)\n", label.Name, label.ID)
				fmt.Printf("Label Catno: %s \n", label.Catno)
			}
			for _, f := range rel.BasicInformation.Formats {
				fmt.Printf("Format: %s \n", f.Name)
			}
			fmt.Printf("MasterID: %d \n", rel.BasicInformation.MasterID)
			fmt.Printf("MasterURL: %s \n", rel.BasicInformation.MasterURL)
			for _, g := range rel.BasicInformation.Genres {
				fmt.Printf("Genre: %s \n", g)
			}
			for _, s := range rel.BasicInformation.Styles {
				fmt.Printf("Style: %s \n", s)
			}
			fmt.Printf("FolderID: %d\n", rel.FolderID)
		}
	}
}
