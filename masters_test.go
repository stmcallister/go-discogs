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

	if rel, err := client.GetMaster(ctx, 36390); err != nil {
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
