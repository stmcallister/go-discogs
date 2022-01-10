package discogs

import (
	"context"
	"net/http"
	"testing"
)

func TestDiscogs_MastersGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/masters/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"id": 1, "main_release": 742394, "most_recent_release": 527468, "resource_url": "https://api.discogs.com/masters/96554", "uri": "https://www.discogs.com/master/96554-Rick-Astley-It-Would-Take-A-Strong-Strong-Man", "versions_url": "https://api.discogs.com/masters/96554/versions", "main_release_url": "https://api.discogs.com/releases/742394", "most_recent_release_url": "https://api.discogs.com/releases/527468", "num_for_sale": 339, "lowest_price": 0.49, "images": [{"type": "primary", "uri": "https://img.discogs.com/_EDIfDDmRQEV6iBVWJqvTyCtTM8=/fit-in/600x585/filters:strip_icc():format(jpeg):mode_rgb():quality(90)/discogs-images/R-742394-1360019823-8548.jpeg.jpg", "resource_url": "https://img.discogs.com/_EDIfDDmRQEV6iBVWJqvTyCtTM8=/fit-in/600x585/filters:strip_icc():format(jpeg):mode_rgb():quality(90)/discogs-images/R-742394-1360019823-8548.jpeg.jpg", "uri150": "https://img.discogs.com/u4jXdo4lfLDLzM6oXGH101NsfVg=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/R-742394-1360019823-8548.jpeg.jpg", "width": 600, "height": 585}], "genres": ["Electronic", "Pop"], "styles": ["Synth-pop"], "year": 1987, "tracklist": [{"position": "A", "type_": "track", "title": "It Would Take A Strong Strong Man", "extraartists": [{"name": "Stock, Aitken & Waterman", "anv": "Stock/Aitken/Waterman", "join": "", "role": "Producer, Written-By", "tracks": "", "id": 20942, "resource_url": "https://api.discogs.com/artists/20942", "thumbnail_url": "https://img.discogs.com/QHmrphRHNlEa4OgdiEwDS3JJzI0=/483x396/smart/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/A-20942-1415554004-8862.jpeg.jpg"}], "duration": "3:39"}, {"position": "B", "type_": "track", "title": "You Move Me", "extraartists": [{"name": "Daize Washbourn", "anv": "", "join": "", "role": "Producer", "tracks": "", "id": 111338, "resource_url": "https://api.discogs.com/artists/111338", "thumbnail_url": ""}, {"name": "Rick Astley", "anv": "R. Astley", "join": "", "role": "Written-By", "tracks": "", "id": 72872, "resource_url": "https://api.discogs.com/artists/72872", "thumbnail_url": "https://img.discogs.com/6n0Nd2VXHLWiRyjg7BxP3kddfO4=/512x512/smart/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/A-72872-1458958828-9832.jpeg.jpg"}], "duration": "3:40"}], "artists": [{"name": "Rick Astley", "anv": "", "join": "", "role": "", "tracks": "", "id": 72872, "resource_url": "https://api.discogs.com/artists/72872", "thumbnail_url": "https://img.discogs.com/6n0Nd2VXHLWiRyjg7BxP3kddfO4=/512x512/smart/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/A-72872-1458958828-9832.jpeg.jpg"}], "title": "It Would Take A Strong Strong Man", "data_quality": "Correct", "videos": [{"uri": "https://www.youtube.com/watch?v=hxoAMnMBiM0", "title": "It Would Take A Strong Strong Man (Matt's Jazzy Guitar Mix) - Rick Astley", "description": "From the maxi single It Would Take A Strong Strong Man (1987)", "duration": 470, "embed": true}]}`))
	})
	client := defaultTestClient(server.URL, "foo")

	ctx := context.Background()
	ID := 1
	rat, err := client.GetMasterRelease(ctx, ID)
	if err != nil {
		t.Fatal(err)
	}
	want := &Master{
		ID:          ID,
		Year:        1987,
		MainRelease: 742394,
		VersionsURL: "https://api.discogs.com/masters/96554/versions",
		MainReleaseURL: "https://api.discogs.com/releases/742394",
		ResourceURL: "https://api.discogs.com/masters/96554",
		URI:         "https://www.discogs.com/master/96554-Rick-Astley-It-Would-Take-A-Strong-Strong-Man",
		Artists: []*Artist{
			{
				Name:        "Rick Astley",
				ID:          72872,
				ResourceURL: "https://api.discogs.com/artists/72872",
			},
		},
		DataQuality: "Correct",
		NumForSale:        339,
		LowestPrice:       0.49,
		Title:             "It Would Take A Strong Strong Man",
		Videos: []*Video{
			{
				URI:         "https://www.youtube.com/watch?v=hxoAMnMBiM0",
				Title:       "It Would Take A Strong Strong Man (Matt's Jazzy Guitar Mix) - Rick Astley",
				Description: "From the maxi single It Would Take A Strong Strong Man (1987)",
				Duration:    470,
				Embed:       true,
			},
		},
		Genres: []string{
			"Electronic",
			"Pop",
		},
		Styles: []string{
			"Synth-pop",
		},
		Tracklist: []*Track{
			{
				Position: "A",
				Type:     "track",
				Title:    "It Would Take A Strong Strong Man",
				Duration: "3:39",
			},
			{
				Position: "B",
				Type:     "track",
				Title:    "You Move Me",
				Duration: "3:40",
			},
		},
		Images: []*Image{
			{
				Type:        "primary",
				URI:         "https://img.discogs.com/_EDIfDDmRQEV6iBVWJqvTyCtTM8=/fit-in/600x585/filters:strip_icc():format(jpeg):mode_rgb():quality(90)/discogs-images/R-742394-1360019823-8548.jpeg.jpg",
				ResourceURL: "https://img.discogs.com/_EDIfDDmRQEV6iBVWJqvTyCtTM8=/fit-in/600x585/filters:strip_icc():format(jpeg):mode_rgb():quality(90)/discogs-images/R-742394-1360019823-8548.jpeg.jpg",
				URI150:      "https://img.discogs.com/u4jXdo4lfLDLzM6oXGH101NsfVg=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/R-742394-1360019823-8548.jpeg.jpg",
				Width:       600,
				Height:      585,
			},
		},
	}

	testEqual(t, want, rat)

}

func TestDiscogs_MastersGetVersions(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/masters/1/versions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"pagination":{"page":1,"pages":1,"per_page":50,"items":11,"urls":{}},"versions":[{"id":14204896,"label":"RCA","country":"US","title":"It Would Take A Strong Strong Man","major_formats":["Cassette"],"format":"Single","catno":"8663-4-RS","released":"1987","status":"Accepted","resource_url":"https://api.discogs.com/releases/14204896","thumb":"https://img.discogs.com/A8R2ELWGEoE0HYO2jlFvAH5UvII=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/R-14204896-1569820009-2315.jpeg.jpg","stats":{"community":{"in_wantlist":7,"in_collection":6},"user":{"in_wantlist":0,"in_collection":0}}}]}`))
	})
	client := defaultTestClient(server.URL, "foo")

	ctx := context.Background()
	ID := 1
	sort := "artist"
	order := "asc"
	page := 1
	per := 100
	rat, err := client.GetMasterReleaseVersions(ctx, ID, page, per, sort, order)
	if err != nil {
		t.Fatal(err)
	}
	format := "Cassette"

	want := &VersionList{
		Pagination: &Pagination {
			Page: 1,
			Pages: 1,
			PerPage: 50,
			Items: 11,
		},
		Versions: []*MasterVersion{
			{
				ID:      14204896,
				Label:   "RCA",
				Country: "US",
				Title:   "It Would Take A Strong Strong Man",
				MajorFormats: []*string{
					&format,
				},
				Format:      "Single",
				Catno:       "8663-4-RS",
				Released:    "1987",
				Status:      "Accepted",
				ResourceURL: "https://api.discogs.com/releases/14204896",
				Thumb:       "https://img.discogs.com/A8R2ELWGEoE0HYO2jlFvAH5UvII=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/R-14204896-1569820009-2315.jpeg.jpg",
				Stats: &VersionStats{
					Community: &Stats{
						InWantlist:   7,
						InCollection: 6,
					},
					User: &Stats{
						InWantlist:   0,
						InCollection: 0,
					},
				},
			},
		},
	}

	testEqual(t, want, rat)
}
func TestDiscogs_MastersGetVersionsAllPages(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/masters/1/versions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"pagination":{"page":1,"pages":1,"per_page":50,"items":11,"urls":{}},"versions":[{"id":14204896,"label":"RCA","country":"US","title":"It Would Take A Strong Strong Man","major_formats":["Cassette"],"format":"Single","catno":"8663-4-RS","released":"1987","status":"Accepted","resource_url":"https://api.discogs.com/releases/14204896","thumb":"https://img.discogs.com/A8R2ELWGEoE0HYO2jlFvAH5UvII=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/R-14204896-1569820009-2315.jpeg.jpg","stats":{"community":{"in_wantlist":7,"in_collection":6},"user":{"in_wantlist":0,"in_collection":0}}}]}`))
	})
	client := defaultTestClient(server.URL, "foo")

	ctx := context.Background()
	ID := 1
	sort := "artist"
	order := "asc"

	rat, err := client.GetMasterReleaseVersionsAllPages(ctx, ID, sort, order)
	if err != nil {
		t.Fatal(err)
	}
	format := "Cassette"

	want := &VersionList{
		Versions: []*MasterVersion{
			{
				ID:      14204896,
				Label:   "RCA",
				Country: "US",
				Title:   "It Would Take A Strong Strong Man",
				MajorFormats: []*string{
					&format,
				},
				Format:      "Single",
				Catno:       "8663-4-RS",
				Released:    "1987",
				Status:      "Accepted",
				ResourceURL: "https://api.discogs.com/releases/14204896",
				Thumb:       "https://img.discogs.com/A8R2ELWGEoE0HYO2jlFvAH5UvII=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/R-14204896-1569820009-2315.jpeg.jpg",
				Stats: &VersionStats{
					Community: &Stats{
						InWantlist:   7,
						InCollection: 6,
					},
					User: &Stats{
						InWantlist:   0,
						InCollection: 0,
					},
				},
			},
		},
	}

	testEqual(t, want, rat)
}
