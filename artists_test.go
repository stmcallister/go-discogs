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
