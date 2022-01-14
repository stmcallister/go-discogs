package discogs

import (
	"context"
	"net/http"
	"testing"
)

func TestDiscogs_ArtistsGet(t *testing.T) {
	setup()
	defer teardown()
	// mux is mocking a realistic response the discogs api would give. this is not calling the live api
	mux.HandleFunc("/artists/72872", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"name": "Rick Astley", "id": 72872, "resource_url": "https://api.discogs.com/artists/72872", "uri": "https://www.discogs.com/artist/72872-Rick-Astley", "releases_url": "https://api.discogs.com/artists/72872/releases", "images": [{"type": "primary", "uri": "https://img.discogs.com/NZwVCXCCHh9QvbYrfF-XmCG6B3E=/512x512/smart/filters:strip_icc():format(jpeg):mode_rgb():quality(90)/discogs-images/A-72872-1458958828-9832.jpeg.jpg", "resource_url": "https://img.discogs.com/NZwVCXCCHh9QvbYrfF-XmCG6B3E=/512x512/smart/filters:strip_icc():format(jpeg):mode_rgb():quality(90)/discogs-images/A-72872-1458958828-9832.jpeg.jpg", "uri150": "https://img.discogs.com/DwQWSekjqVKl6REDDGUPoe_Ezvw=/150x150/smart/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/A-72872-1458958828-9832.jpeg.jpg", "width": 512, "height": 512}], "realname": "Richard Paul Astley", "profile": "Rick Astley (b. February 6, 1966 in Warrington, United Kingdom) was one of figures to rise up in the dance/pop-driven late '80s. He started as a studio-tea-boy in 1985 at the [l279756] and released his first own single in 1987, [i][m=96559][/i].\r\n\r\nIn 2007, Rick Astley became the subject of a viral internet meme in which millions of internet users were tricked into watching Rick Astley's video [i][url=https://youtu.be/dQw4w9WgXcQ]Never Gonna Give You Up[/url][/i] by posting it under the name of other popular video titles. The practice is now known as [url=https://knowyourmeme.com/memes/rickroll]rickrolling[/url].\r\n", "urls": ["http://www.rickastley.co.uk", "http://www.facebook.com/RickAstley", "http://www.instagram.com/officialrickastley", "http://myspace.com/rickastley", "http://twitter.com/rickastley", "http://en.wikipedia.org/wiki/Rick_Astley", "http://www.youtube.com/channel/UCuAXFkgsw1L7xaCfnd5JJOw"], "namevariations": ["Astley", "Astley R.", "R Astley", "R.", "R. Asley", "R. Astley", "R.Astley", "Richard Astley", "Rick", "Rick Ashley", "Rick Asley", "Rick Astle", "Rick Astly", "Rick x", "Rock Astley", "\u0420\u0438\u043a \u0410\u0441\u0442\u043b\u0438", "\u30ea\u30c3\u30af", "\u30ea\u30c3\u30af\u30fb\u30a2\u30b9\u30c8\u30ea\u30fc"], "aliases": [{"id": 1141583, "name": "Dick Spatsley", "resource_url": "https://api.discogs.com/artists/1141583", "thumbnail_url": ""}], "groups": [{"id": 146979, "name": "Band Aid II", "resource_url": "https://api.discogs.com/artists/146979", "active": true, "thumbnail_url": "https://img.discogs.com/waymB_APueCMjJ52bt1zQnr53ys=/600x454/smart/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/A-146979-1429962400-9967.jpeg.jpg"}, {"id": 420265, "name": "Ferry Aid", "resource_url": "https://api.discogs.com/artists/420265", "active": true, "thumbnail_url": "https://img.discogs.com/lmGPv_nHsYvGi4bKtyXGuhTCtCY=/600x413/smart/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/A-420265-1140467463.jpeg.jpg"}, {"id": 6576225, "name": "NHS Voices", "resource_url": "https://api.discogs.com/artists/6576225", "active": true, "thumbnail_url": "https://img.discogs.com/8ZDJ4sz0MqGg0oigzYBQxMA6Sp0=/600x441/smart/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/A-6576225-1587395423-1890.png.jpg"}], "data_quality": "Needs Vote"}`))
	})
	client := defaultTestClient(server.URL, "foo")

	ctx := context.Background()
	artistID := 72872

	art, err := client.GetArtist(ctx, artistID)
	if err != nil {
		t.Fatal(err)
	}
	want := &Artist{
		Name:        "Rick Astley",
		ID:          72872,
		ResourceURL: "https://api.discogs.com/artists/72872",
		URI:         "https://www.discogs.com/artist/72872-Rick-Astley",
		ReleasesURL: "https://api.discogs.com/artists/72872/releases",
		Images: []*Image{
			{
				Type:        "primary",
				URI:         "https://img.discogs.com/NZwVCXCCHh9QvbYrfF-XmCG6B3E=/512x512/smart/filters:strip_icc():format(jpeg):mode_rgb():quality(90)/discogs-images/A-72872-1458958828-9832.jpeg.jpg",
				ResourceURL: "https://img.discogs.com/NZwVCXCCHh9QvbYrfF-XmCG6B3E=/512x512/smart/filters:strip_icc():format(jpeg):mode_rgb():quality(90)/discogs-images/A-72872-1458958828-9832.jpeg.jpg",
				URI150:      "https://img.discogs.com/DwQWSekjqVKl6REDDGUPoe_Ezvw=/150x150/smart/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/A-72872-1458958828-9832.jpeg.jpg",
				Width:       512,
				Height:      512,
			},
		},
		RealName: "Richard Paul Astley",
		Profile:  "Rick Astley (b. February 6, 1966 in Warrington, United Kingdom) was one of figures to rise up in the dance/pop-driven late '80s. He started as a studio-tea-boy in 1985 at the [l279756] and released his first own single in 1987, [i][m=96559][/i].\r\n\r\nIn 2007, Rick Astley became the subject of a viral internet meme in which millions of internet users were tricked into watching Rick Astley's video [i][url=https://youtu.be/dQw4w9WgXcQ]Never Gonna Give You Up[/url][/i] by posting it under the name of other popular video titles. The practice is now known as [url=https://knowyourmeme.com/memes/rickroll]rickrolling[/url].\r\n",
		URLs: []string{
			"http://www.rickastley.co.uk",
			"http://www.facebook.com/RickAstley",
			"http://www.instagram.com/officialrickastley",
			"http://myspace.com/rickastley",
			"http://twitter.com/rickastley",
			"http://en.wikipedia.org/wiki/Rick_Astley",
			"http://www.youtube.com/channel/UCuAXFkgsw1L7xaCfnd5JJOw",
		},
		NameVariations: []string{
			"Astley",
			"Astley R.",
			"R Astley",
			"R.",
			"R. Asley",
			"R. Astley",
			"R.Astley",
			"Richard Astley",
			"Rick",
			"Rick Ashley",
			"Rick Asley",
			"Rick Astle",
			"Rick Astly",
			"Rick x",
			"Rock Astley",
			"Рик Астли",
			"リック",
			"リック・アストリー",
		},
		Aliases: []*Alias{
			{
				ID:          1141583,
				Name:        "Dick Spatsley",
				ResourceURL: "https://api.discogs.com/artists/1141583",
			},
		},
		Groups: []*Group{
			{
				ID:          146979,
				Name:        "Band Aid II",
				ResourceURL: "https://api.discogs.com/artists/146979",
				Active:      true,
			},
			{
				ID:          420265,
				Name:        "Ferry Aid",
				ResourceURL: "https://api.discogs.com/artists/420265",
				Active:      true,
			},
			{
				ID:          6576225,
				Name:        "NHS Voices",
				ResourceURL: "https://api.discogs.com/artists/6576225",
				Active:      true,
			},
		},
		DataQuality: "Needs Vote",
	}

	testEqual(t, want, art)
}

func TestDiscogs_ArtistsReleasesGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/artists/1/releases", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"pagination": {"page": 1, "pages": 1, "per_page": 1, "items": 1, "urls": {"last": "https://api.discogs.com/artists/72872/releases?page=45&per_page=50", "next": "https://api.discogs.com/artists/72872/releases?page=2&per_page=50"}}, "releases": [{"id": 3686862, "type": "release", "format": "7\", Jukebox", "label": "Jive, RCA", "title": "I Surrender (To The Spirit Of The Night) / Never Gonna Give You Up", "resource_url": "https://api.discogs.com/releases/3686862", "role": "Main", "artist": "Samantha Fox / Rick Astley", "year": 1987, "thumb": "https://img.discogs.com/wAh-Kh2b50fxIY5Y6yp-l3v8ao0=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/R-3686862-1384333956-6567.jpeg.jpg"}]}`))
	})
	client := defaultTestClient(server.URL, "foo")

	ctx := context.Background()
	aid := 1
	sort := "title"
	order := "asc"
	page := 1
	per := 1

	rels, err := client.GetArtistReleases(ctx, sort, order, aid, page, per)
	if err != nil {
		t.Fatal(err)
	}
	last := "https://api.discogs.com/artists/72872/releases?page=45&per_page=50"
	next := "https://api.discogs.com/artists/72872/releases?page=2&per_page=50"
	want := &ArtistReleaseList{
		Pagination: &Pagination{
			Page:    1,
			Pages:   1,
			PerPage: 1,
			Items:   1,
			URLs: PaginationURLs{
				Next: &next,
				Last: &last,
			},
		},
		Releases: []*ArtistRelease{
			{
				ID:          3686862,
				Type:        "release",
				Title:       "I Surrender (To The Spirit Of The Night) / Never Gonna Give You Up",
				ResourceURL: "https://api.discogs.com/releases/3686862",
				Format:      "7\", Jukebox",
				Role:        "Main",
				Label:       "Jive, RCA",
				Artist:      "Samantha Fox / Rick Astley",
				Year:        1987,
				Thumb:       "https://img.discogs.com/wAh-Kh2b50fxIY5Y6yp-l3v8ao0=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/R-3686862-1384333956-6567.jpeg.jpg",
			},
		},
	}
	testEqual(t, want, rels)
}

func TestDiscogs_ArtistsReleasesGetAll(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/artists/1/releases", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"pagination": {"page": 1, "pages": 1, "per_page": 1, "items": 1, "urls": {"last": "https://api.discogs.com/artists/72872/releases?page=45&per_page=50", "next": "https://api.discogs.com/artists/72872/releases?page=2&per_page=50"}}, "releases": [{"id": 3686862, "type": "release", "format": "7\", Jukebox", "label": "Jive, RCA", "title": "I Surrender (To The Spirit Of The Night) / Never Gonna Give You Up", "resource_url": "https://api.discogs.com/releases/3686862", "role": "Main", "artist": "Samantha Fox / Rick Astley", "year": 1987, "thumb": "https://img.discogs.com/wAh-Kh2b50fxIY5Y6yp-l3v8ao0=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/R-3686862-1384333956-6567.jpeg.jpg"}]}`))
	})
	client := defaultTestClient(server.URL, "foo")

	ctx := context.Background()
	aid := 1
	sort := "title"
	order := "asc"

	rels, err := client.GetAllArtistReleases(ctx, sort, order, aid)
	if err != nil {
		t.Fatal(err)
	}
	want := &ArtistReleaseList{
		Releases: []*ArtistRelease{
			{
				ID:          3686862,
				Type:        "release",
				Title:       "I Surrender (To The Spirit Of The Night) / Never Gonna Give You Up",
				ResourceURL: "https://api.discogs.com/releases/3686862",
				Format:      "7\", Jukebox",
				Role:        "Main",
				Label:       "Jive, RCA",
				Artist:      "Samantha Fox / Rick Astley",
				Year:        1987,
				Thumb:       "https://img.discogs.com/wAh-Kh2b50fxIY5Y6yp-l3v8ao0=/fit-in/150x150/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/R-3686862-1384333956-6567.jpeg.jpg",
			},
		},
	}

	testEqual(t, want, rels)
}
