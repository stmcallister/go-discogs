package discogs

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestDiscogs_ReleasesGet(t *testing.T) {
	// todo: make real unit tests using MUX
	key := os.Getenv("DISCOGS_API_KEY")
	client := NewClient(key)
	ctx := context.Background()
	cur := ""

	if rel, err := client.GetRelease(ctx, 249504, cur); err != nil {
		panic(err)
	} else {
		fmt.Printf("\n\nID: %d\n", rel.ID)
		fmt.Printf("Title: %s\n", rel.Title)
		fmt.Printf("Year: %d\n", rel.Year)
		fmt.Printf("Status: %s\n", rel.Status)
		for _, artist := range rel.Artists {
			fmt.Printf("Artist Name: %s (%d)\n", artist.Name, artist.ID)
		}
		for _, label := range rel.Labels {
			fmt.Printf("Label Name: %s (%d)\n", label.Name, label.ID)
			fmt.Printf("Label Catno: %s \n", label.Catno)
		}
		for _, co := range rel.Companies {
			fmt.Printf("Company Name: %s (%d)\n", co.Name, co.ID)
		}
		for _, f := range rel.Formats {
			fmt.Printf("Format: %s \n", f.Name)
		}
		fmt.Printf("Data Quality: %s \n", rel.DataQuality)
		fmt.Println("--- Community ---")
		fmt.Printf("Have: %d \n", rel.Community.Have)
		fmt.Printf("Want: %d \n", rel.Community.Want)
		fmt.Printf("Rating Count: %d \n", rel.Community.Rating.Count)
		fmt.Printf("Rating Avg: %v \n", rel.Community.Rating.Average)
		fmt.Printf("Submitter: %s \n", rel.Community.Submitter.Username)
		for _, c := range rel.Community.Contributors {
			fmt.Printf("Contributor: %s \n", c.Username)
		}
		fmt.Printf("Community Data Quality: %s \n", rel.Community.DataQuality)
		fmt.Printf("Community Status: %s \n", rel.Community.Status)
		fmt.Printf("FormatQuantity: %d \n", rel.FormatQuantity)
		fmt.Printf("DateAdded: %s \n", rel.DateAdded)
		fmt.Printf("DateChanged: %s \n", rel.DateChanged)
		fmt.Printf("NumForSale: %d \n", rel.NumForSale)
		fmt.Printf("LowestPrice: %v \n", rel.LowestPrice)
		fmt.Printf("MasterID: %d \n", rel.MasterID)
		fmt.Printf("MasterURL: %s \n", rel.MasterURL)
		fmt.Printf("Country: %s \n", rel.Country)
		fmt.Printf("Released: %s \n", rel.Released)
		fmt.Printf("ReleasedFormatted: %s \n", rel.ReleasedFormatted)
		fmt.Printf("Notes: %s \n", rel.Notes)
		for _, i := range rel.Identifiers {
			fmt.Printf("Identifier: %s \n", i.Value)
		}
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
		for _, ea := range rel.ExtraArtists {
			fmt.Printf("Extra Artist: %s \n", ea.Name)
		}
		for _, img := range rel.Images {
			fmt.Printf("Image: %s\n", img.Type)
		}
		fmt.Printf("Thumb: %s\n", rel.Thumb)
		fmt.Printf("Estimated Weight: %d\n", rel.EstimatedWeight)
		fmt.Printf("URI: %s\n\n", rel.URI)
	}
}
func TestDiscogs_ReleasesGetUserRating(t *testing.T) {
	// todo: make real unit tests using MUX
	key := os.Getenv("DISCOGS_API_KEY")
	client := NewClient(key)
	ctx := context.Background()
	username := "bartman"

	if rel, err := client.GetReleaseRatingByUser(ctx, 249504, username); err != nil {
		panic(err)
	} else {
		fmt.Printf("\n\nRelease ID: %d", rel.ReleaseID)
		fmt.Printf("\nUsername: %s", rel.Username)
		fmt.Printf("\nRating: %d\n", rel.Rating)
	}
}
func TestDiscogs_ReleasesGetCommunityRating(t *testing.T) {
	// todo: make real unit tests using MUX
	key := os.Getenv("DISCOGS_API_KEY")
	client := NewClient(key)
	ctx := context.Background()

	if rel, err := client.GetReleaseCommunityRating(ctx, 249504); err != nil {
		panic(err)
	} else {
		fmt.Printf("\n\nRelease ID: %d", rel.ReleaseID)
		fmt.Printf("\nRating Count: %d", rel.Rating.Count)
		fmt.Printf("\nRating Avg: %v\n", rel.Rating.Average)
	}
}
func TestDiscogs_ReleasesStats(t *testing.T) {
	// todo: make real unit tests using MUX
	key := os.Getenv("DISCOGS_API_KEY")
	client := NewClient(key)
	ctx := context.Background()

	if rel, err := client.GetReleaseStats(ctx, 249504); err != nil {
		panic(err)
	} else {
		fmt.Printf("\n Is Offensive: %v", rel.IsOffensive)
		fmt.Printf("\n Num Have: %d", rel.NumHave)
		fmt.Printf("\n Num Want: %d\n", rel.NumWant)
	}
}
