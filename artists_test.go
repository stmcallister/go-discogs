package discogs

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestDiscogs_ArtistsGet(t *testing.T) {
	// todo: make real unit tests using MUX
	key := os.Getenv("DISCOGS_API_KEY")
	client := NewClient(key)
	ctx := context.Background()

	if a, err := client.GetArtist(ctx, 35301); err != nil {
		panic(err)
	} else {
		fmt.Printf("\n\nID: %d\n", a.ID)
		fmt.Printf("Name: %s\n", a.Name)
		fmt.Printf("Real Name: %s\n", a.RealName)
		fmt.Printf("ResourceURL: %s\n", a.ResourceURL)
		fmt.Printf("Profile: %s \n", a.Profile)
		fmt.Printf("DataQuality: %s \n", a.DataQuality)
		for _, v := range a.URLs {
			fmt.Printf("URL: %s \n", v)
		}
		for _, g := range a.NameVariations {
			fmt.Printf("Genre: %s \n", g)
		}
		for _, img := range a.Images {
			fmt.Printf("Image Type: %s\n", img.Type)
			fmt.Printf("Image ResourceURL: %s\n", img.ResourceURL)

		}
		fmt.Printf("URI: %s\n\n", a.URI)
	}
}

func TestDiscogs_ArtistsReleases(t *testing.T) {
	// todo: make real unit tests using MUX
	key := os.Getenv("DISCOGS_API_KEY")
	client := NewClient(key)
	ctx := context.Background()
	aid := 35301
	sort := "title"
	order := "asc"
	page := 1
	per := 100

	if res, err := client.GetArtistReleases(ctx, sort, order, aid, page, per); err != nil {
		panic(err)
	} else {
		fmt.Printf("\n Pagination Per Page: %d ", res.Pagination.PerPage)
		fmt.Printf("\n Pagination Items: %d ", res.Pagination.Items)
		fmt.Printf("\n Pagination Page: %d ", res.Pagination.Page)
		fmt.Printf("\n Pagination URLs: %v ", res.Pagination.URLs)
		fmt.Printf("\n Pagination Pages: %d \n", res.Pagination.Pages)

		for i, rel := range res.Releases {
			fmt.Printf("\n\nID: %d\n", rel.ID)
			fmt.Printf("Title: %s\n", rel.Title)
			fmt.Printf("Year: %d\n", rel.Year)
			fmt.Printf("Artist: %s \n", rel.Artist)
			fmt.Printf("MainRelease: %d \n", rel.MainRelease)
			fmt.Printf("ResourceURL: %s \n", rel.ResourceURL)
			fmt.Printf("Role: %s \n", rel.Role)
			fmt.Printf("Thumb: %s \n", rel.Thumb)
			fmt.Printf("Type: %s \n", rel.Type)
			fmt.Printf("# %d\n", i)
		}
	}
}
func TestDiscogs_ArtistsAllReleases(t *testing.T) {
	// todo: make real unit tests using MUX
	key := os.Getenv("DISCOGS_API_KEY")
	client := NewClient(key)
	ctx := context.Background()
	aid := 35301
	sort := "title"
	order := "asc"

	if res, err := client.GetAllArtistReleases(ctx, sort, order, aid); err != nil {
		panic(err)
	} else {
		for i, rel := range res.Releases {
			fmt.Printf("\n\nID: %d\n", rel.ID)
			fmt.Printf("Title: %s\n", rel.Title)
			fmt.Printf("Year: %d\n", rel.Year)
			fmt.Printf("Artist: %s \n", rel.Artist)
			fmt.Printf("MainRelease: %d \n", rel.MainRelease)
			fmt.Printf("ResourceURL: %s \n", rel.ResourceURL)
			fmt.Printf("Role: %s \n", rel.Role)
			fmt.Printf("Thumb: %s \n", rel.Thumb)
			fmt.Printf("Type: %s \n", rel.Type)
			fmt.Printf("# %d\n", i)
		}
	}
}
