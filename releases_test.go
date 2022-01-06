package discogs

import (
	"context"
	"net/http"
	"testing"
)

func TestDiscogs_ReleasesGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/releases/249504", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"id": 249504, "status": "Accepted", "year": 1987, "resource_url": "https://api.discogs.com/releases/249504", "uri": "https://www.discogs.com/release/249504-Rick-Astley-Never-Gonna-Give-You-Up", "artists": [{"name": "Rick Astley", "anv": "", "join": "", "role": "", "tracks": "", "id": 72872, "resource_url": "https://api.discogs.com/artists/72872"}], "artists_sort": "Rick Astley", "labels": [{"name": "RCA", "catno": "PB 41447", "entity_type": "1", "entity_type_name": "Label", "id": 895, "resource_url": "https://api.discogs.com/labels/895", "thumbnail_url": "https://img.discogs.com/ja3Z62KhgTLO1LYf6dmRR7F3p3g=/fit-in/600x600/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/L-895-1406902706-6748.jpeg.jpg"}], "series": [], "companies": [{"name": "BMG Records (UK) Ltd.", "catno": "", "entity_type": "13", "entity_type_name": "Phonographic Copyright (p)", "id": 82835, "resource_url": "https://api.discogs.com/labels/82835", "thumbnail_url": "https://img.discogs.com/e92WmIQWWLTFdl5tS_xvnCNlhh0=/fit-in/173x80/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/L-82835-1233170664.gif.jpg"}ogs-i}], "formats": [{"name": "Vinyl", "qty": "1", "descriptions": ["7\"", "45 RPM", "Single", "Stereo"]}], "data_quality": "Needs Vote", "community": {"have": 2372, "want": 408, "rating": {"count": 163, "average": 3.65}, "submitter": {"username": "memory", "resource_url": "https://api.discogs.com/users/memory"}, "contributors": [{"username": "memory", "resource_url": "https://api.discogs.com/users/memory"}], "data_quality": "Needs Vote", "status": "Accepted"}, "format_quantity": 1, "date_added": "2004-04-30T08:10:05-07:00", "date_changed": "2021-09-02T07:40:01-07:00", "num_for_sale": 70, "lowest_price": 1.01, "master_id": 96559, "master_url": "https://api.discogs.com/masters/96559", "title": "Never Gonna Give You Up", "country": "UK", "released": "1987-07-00", "notes": "UK Release has a black label with the text \"Manufactured In England\" printed on it.\n\nSleeve:\n\u2117 1987 \u2022 BMG Records (UK) Ltd. \u00a9 1987 \u2022 BMG Records (UK) Ltd.\nDistributed in the UK by BMG Records \u2022  Distribu\u00e9 en Europe par BMG/Ariola \u2022 Vertrieb en Europa d\u00fcrch BMG/Ariola.\n\nCenter labels:\n\u2117 1987 Pete Waterman Ltd.\nOriginal Sound Recording made by PWL.\nBMG Records (UK) Ltd. are the exclusive licensees for the world.\n\nDurations do not appear on the release.", "released_formatted": "Jul 1987", "identifiers": [{"type": "Barcode", "value": "5012394144777"}], "videos": [{"uri": "https://www.youtube.com/watch?v=dQw4w9WgXcQ", "title": "Rick Astley - Never Gonna Give You Up (Official Music Video)", "description": "The official video for \u201cNever Gonna Give You Up\u201d by Rick Astley\n \n\u201cNever Gonna Give You Up\u201d was a global smash on its release in July 1987, topping the charts in 25 countries including Rick\u2019s native UK and the US Billboard Hot 100.  It also won ", "duration": 213, "embed": true}], "genres": ["Electronic", "Pop"], "styles": ["Euro-Disco"], "tracklist": [{"position": "A", "type_": "track", "title": "Never Gonna Give You Up", "duration": "3:32"}, {"position": "B", "type_": "track", "title": "Never Gonna Give You Up (Instrumental)", "duration": "3:30"}], "extraartists": [{"name": "Me Company", "anv": "Me Co", "join": "", "role": "Design", "tracks": "", "id": 547352, "resource_url": "https://api.discogs.com/artists/547352", "thumbnail_url": "https://img.discogs.com/6QuxK3_zmbj1Zk-n7fuQRlV1obE=/174x49/smart/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/A-547352-1361313237-5460.png.jpg"}], "images": [{"type": "primary", "uri": "https://img.discogs.com/2e3s-w-JF4jdIeUTkIXjTsJY9Kw=/fit-in/600x600/filters:strip_icc():format(jpeg):mode_rgb():quality(90)/discogs-images/R-249504-1334592212.jpeg.jpg", "resource_url": "https://img.discogs.com/2e3s-w-JF4jdIeUTkIXjTsJY9Kw=/fit-in/600x600/filters:strip_icc():format(jpeg):mode_rgb():quality(90)/discogs-images/R-249504-1334592212.jpeg.jpg", "uri150": "https://img.discogs.com/B6Jhg03KWlxWnhEyFkeVF4J2wLk=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/R-249504-1334592212.jpeg.jpg", "width": 600, "height": 600}], "thumb": "https://img.discogs.com/B6Jhg03KWlxWnhEyFkeVF4J2wLk=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/R-249504-1334592212.jpeg.jpg", "estimated_weight": 60, "blocked_from_sale": false}`))
	})
	client := defaultTestClient(server.URL, "foo")

	ctx := context.Background()
	releaseID := 249504
	currency := ""
	rat, err := client.GetRelease(ctx, releaseID, currency)
	if err != nil {
		t.Fatal(err)
	}
	want := &Release{
		ID:          releaseID,
		Status:      "accepted",
		Year:        1987,
		ResourceURL: "https://api.discogs.com/releases/249504",
		URI:         "https://www.discogs.com/release/249504-Rick-Astley-Never-Gonna-Give-You-Up",
		Artists: []*ReleaseArtist{
			{
				Name:        "Rick Astley",
				ID:          72872,
				ResourceURL: "https://api.discogs.com/artists/72872",
			},
		},
		ArtistsSort: "Rick Astley",
		Labels: []*Entity{
			{
				Name:           "RCA",
				Catno:          "PB 41447",
				EntityType:     "1",
				EntityTypeName: "Label",
				ID:             895,
				ResourceURL:    "https://api.discogs.com/labels/895",
			},
		},
		Companies: []*Entity{
			{
				Name:           "BMG Records (UK) Ltd.",
				EntityType:     "13",
				EntityTypeName: "Phonographic Copyright (p)",
				ID:             82835,
				ResourceURL:    "https://api.discogs.com/labels/82835",
			},
		},
		Formats: []*Format{
			{
				Name: "Vinyl",
				Qty:  "1",
				Descriptions: []string{
					"7\"",
					"45 RPM",
					"Single",
					"Stereo",
				},
			},
		},
		DataQuality: "Needs Vote",
		Community: Community{
			Have: 2372,
			Want: 408,
			Rating: Rating{
				Count:   163,
				Average: 3.65,
			},
			Submitter: CommunityPerson{
				Username:    "memory",
				ResourceURL: "https://api.discogs.com/users/memory",
			},
			Contributors: []*CommunityPerson{
				{
					Username:    "memory",
					ResourceURL: "https://api.discogs.com/users/memory",
				},
			},
			DataQuality: "Needs Vote",
			Status:      "Accepted",
		},
		FormatQuantity:    1,
		DateAdded:         "2004-04-30T08:10:05-07:00",
		DateChanged:       "2021-09-02T07:40:01-07:00",
		NumForSale:        70,
		LowestPrice:       1.01,
		MasterID:          96559,
		MasterURL:         "https://api.discogs.com/masters/96559",
		Title:             "Never Gonna Give You Up",
		Country:           "UK",
		Released:          "1987-07-00",
		Notes:             "UK Release has a black label with the text \"Manufactured In England\" printed on it.\n\nSleeve:\n℗ 1987 • BMG Records (UK) Ltd. © 1987 • BMG Records (UK) Ltd.\nDistributed in the UK by BMG Records •  Distribué en Europe par BMG/Ariola • Vertrieb en Europa dürch BMG/Ariola.\n\nCenter labels:\n℗ 1987 Pete Waterman Ltd.\nOriginal Sound Recording made by PWL.\nBMG Records (UK) Ltd. are the exclusive licensees for the world.\n\nDurations do not appear on the release.",
		ReleasedFormatted: "Jul 1987",
		Identifiers: []*Identifier{
			{
				Type:  "Barcode",
				Value: "5012394144777",
			},
		},
		Videos: []*Video{
			{
				URI:         "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
				Title:       "Rick Astley - Never Gonna Give You Up (Official Music Video)",
				Description: "The official video for “Never Gonna Give You Up” by Rick Astley\n \n“Never Gonna Give You Up” was a global smash on its release in July 1987, topping the charts in 25 countries including Rick’s native UK and the US Billboard Hot 100.  It also won ",
				Duration:    213,
				Embed:       true,
			},
		},
		Genres: []string{
			"Electronic",
			"Pop",
		},
		Styles: []string{
			"Euro-Disco",
		},
		Tracklist: []*Track{
			{
				Position: "A",
				Type:     "track",
				Title:    "Never Gonna Give You Up",
				Duration: "3:32",
			},
			{
				Position: "B",
				Type:     "track",
				Title:    "Never Gonna Give You Up (Instrumental)",
				Duration: "3:30",
			},
		},
		ExtraArtists: []*ReleaseArtist{
			{
				Name:        "Me Company",
				Anv:         "Me Co",
				Join:        "",
				Role:        "Design",
				Tracks:      "",
				ID:          547352,
				ResourceURL: "https://api.discogs.com/artists/547352",
			},
		},
		Images: []*Image{
			{
				Type:        "primary",
				URI:         "https://img.discogs.com/2e3s-w-JF4jdIeUTkIXjTsJY9Kw=/fit-in/600x600/filters:strip_icc():format(jpeg):mode_rgb():quality(90)/discogs-images/R-249504-1334592212.jpeg.jpg",
				ResourceURL: "https://img.discogs.com/2e3s-w-JF4jdIeUTkIXjTsJY9Kw=/fit-in/600x600/filters:strip_icc():format(jpeg):mode_rgb():quality(90)/discogs-images/R-249504-1334592212.jpeg.jpg",
				URI150:      "https://img.discogs.com/B6Jhg03KWlxWnhEyFkeVF4J2wLk=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/R-249504-1334592212.jpeg.jpg",
				Width:       600,
				Height:      600,
			},
		},
		Thumb:           "https://img.discogs.com/B6Jhg03KWlxWnhEyFkeVF4J2wLk=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/R-249504-1334592212.jpeg.jpg",
		EstimatedWeight: 60,
	}

	testEqual(t, want, rat)

}
func TestDiscogs_ReleasesGetUserRating(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/releases/1/rating/testuser", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"release_id": 1,"rating": 3, "username": "testuser"}`))
	})
	client := defaultTestClient(server.URL, "foo")

	ctx := context.Background()
	releaseID := 1
	username := "testuser"

	rat, err := client.GetReleaseRatingByUser(ctx, releaseID, username)
	if err != nil {
		t.Fatal(err)
	}
	want := &ReleaseUserRating{
		ReleaseID: releaseID,
		Rating:    3,
		Username:  username,
	}

	testEqual(t, want, rat)
}
func TestDiscogs_ReleasesGetCommunityRating(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/releases/1/rating", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"release_id": 1,"rating": {"count": 165,"average": 3.64}}`))
	})
	client := defaultTestClient(server.URL, "foo")

	ctx := context.Background()
	releaseID := 1

	rat, err := client.GetReleaseCommunityRating(ctx, releaseID)
	if err != nil {
		t.Fatal(err)
	}
	want := &ReleaseRating{
		ReleaseID: releaseID,
		Rating: &Rating{
			Count:   165,
			Average: 3.64,
		},
	}

	testEqual(t, want, rat)
}
func TestDiscogs_ReleasesGetStats(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/releases/1/stats", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"is_offensive":false,"num_have":3,"num_want": 5}`))
	})
	client := defaultTestClient(server.URL, "foo")

	ctx := context.Background()
	releaseID := 1

	stat, err := client.GetReleaseStats(ctx, releaseID)
	if err != nil {
		t.Fatal(err)
	}
	want := &ReleaseStats{
		IsOffensive: false,
		NumHave:     3,
		NumWant:     5,
	}

	testEqual(t, want, stat)
}
func TestDiscogs_ReleasesUpdateUserRating(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/releases/1/rating/testuser", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		_, _ = w.Write([]byte(`{"release_id":1,"rating":3,"username": "testuser"}`))
	})
	client := defaultTestClient(server.URL, "foo")

	ctx := context.Background()
	releaseID := 1
	username := "testuser"
	rating := 3
	rat, err := client.UpdateReleaseRating(ctx, releaseID, rating, username)
	if err != nil {
		t.Fatal(err)
	}
	want := &ReleaseUserRating{
		ReleaseID: releaseID,
		Username:  username,
		Rating:    rating,
	}

	testEqual(t, want, rat)
}
