package discogs

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestDiscogs_MastersGet(t *testing.T) {
	// todo: make real unit tests using MUX
	key := os.Getenv("DISCOGS_API_KEY")
	client := NewClient(key)
	ctx := context.Background()

	if rel, err := client.GetMasterRelease(ctx, 36390); err != nil {
		panic(err)
	} else {
		fmt.Printf("\n\nID: %d\n", rel.ID)
		fmt.Printf("Title: %s\n", rel.Title)
		fmt.Printf("Year: %d\n", rel.Year)
		fmt.Printf("ResourceURL: %s\n", rel.ResourceURL)
		for _, artist := range rel.Artists {
			fmt.Printf("Artist Name: %s (%d)\n", artist.Name, artist.ID)
		}
		fmt.Printf("NumForSale: %d \n", rel.NumForSale)
		fmt.Printf("LowestPrice: %v \n", rel.LowestPrice)
		fmt.Printf("MainRelease: %d \n", rel.MainRelease)
		fmt.Printf("MainReleaseURL: %s \n", rel.MainReleaseURL)
		for _, v := range rel.Videos {
			fmt.Printf("Video: %s \n", v.Title)
		}
		for _, g := range rel.Genres {
			fmt.Printf("Genre: %s \n", g)
		}
		for _, s := range rel.Styles {
			fmt.Printf("Style: %s \n", s)
		}
		for _, t := range rel.Tracklist {
			fmt.Printf("Track: %s \n", t.Title)
		}
		for _, img := range rel.Images {
			fmt.Printf("Image Type: %s\n", img.Type)
			fmt.Printf("Image ResourceURL: %s\n", img.ResourceURL)

		}
		fmt.Printf("URI: %s\n\n", rel.URI)
	}
}

func TestDiscogs_MastersGetVersions(t *testing.T) {
	// todo: make real unit tests using MUX
	key := os.Getenv("DISCOGS_API_KEY")
	client := NewClient(key)
	ctx := context.Background()
	id := 8883
	sort := "artist"
	order := "asc"
	page := 1
	per := 100

	if rel, err := client.GetMasterReleaseVersions(ctx, id, page, per, sort, order); err != nil {
		panic(err)
	} else {
		fmt.Printf("\n Pagination Per Page: %d ", rel.Pagination.PerPage)
		fmt.Printf("\n Pagination Items: %d ", rel.Pagination.Items)
		fmt.Printf("\n Pagination Page: %d ", rel.Pagination.Page)
		fmt.Printf("\n Pagination URLs: %v ", rel.Pagination.URLs)
		fmt.Printf("\n Pagination Pages: %d \n", rel.Pagination.Pages)

		printVersionItems(rel)
	}
}
func TestDiscogs_MastersGetVersionsAllPages(t *testing.T) {
	// todo: make real unit tests using MUX
	key := os.Getenv("DISCOGS_API_KEY")
	client := NewClient(key)
	ctx := context.Background()
	id := 8883
	sort := "artist"
	order := "asc"

	if rel, err := client.GetMasterReleaseVersionsAllPages(ctx, id, sort, order); err != nil {
		panic(err)
	} else {
		printVersionItems(rel)
	}
}

func printVersionItems(vers *VersionList) {
	for i, rel := range vers.Versions {
		fmt.Printf("\n\nID: %d\n", rel.ID)
		fmt.Printf("Title: %s\n", rel.Title)
		fmt.Printf("ResourceURL: %s\n", rel.ResourceURL)
		fmt.Printf("Label: %s\n", rel.Label)
		fmt.Printf("Format: %s\n", rel.Format)
		fmt.Printf("Country: %s\n", rel.Country)
		fmt.Printf("Released: %s\n", rel.Released)
		fmt.Printf("Catno: %s\n", rel.Catno)
		for _, mf := range rel.MajorFormats {
			fmt.Printf("Major Format: %v \n", mf)
		}
		fmt.Printf("---- Stats ----\n")
		fmt.Printf("User: InCollection: %d\n", rel.Stats.User.InCollection)
		fmt.Printf("User: InWantlist: %d\n", rel.Stats.User.InWantlist)
		fmt.Printf("Community: InCollection: %d\n", rel.Stats.Community.InCollection)
		fmt.Printf("Community: InWantlist: %d\n", rel.Stats.Community.InWantlist)
		fmt.Printf("# %d\n", i)
	}
}
