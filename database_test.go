package discogs

import (
	"context"
	"net/http"
	"testing"
)

func TestDiscogs_Search(t *testing.T) {
	setup()
	defer teardown()
	// mux is mocking a realistic response the discogs api would give. this is not calling the live api
	mux.HandleFunc("/database/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"pagination":{"page":1,"pages":28,"per_page":50,"items":1381,"urls":{"last":"https://api.discogs.com/database/search?q=jimmy+eat+world&page=28&per_page=50","next":"https://api.discogs.com/database/search?q=jimmy+eat+world&page=28&per_page=50"}},"results":[{"id":82079,"type":"artist","user_data":{"in_wantlist":false,"in_collection":true},"master_id":null,"master_url":null,"uri":"/artist/82079-Jimmy-Eat-World","title":"Jimmy Eat World","thumb":"https://i.discogs.com/nf8mu5qMqLjvdqLfdXH93xGVihp7hEXMYTEQZ-a3mDY/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWltYWdlcy9BLTgy/MDc5LTE2MTM0OTQz/NzItODIzMC5qcGVn.jpeg","cover_image":"https://i.discogs.com/I_vVLO6a2K2atNydPFDxaDt5WU1ccpN00UZ3QFFG8Qw/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWltYWdlcy9BLTgy/MDc5LTE2MTM0OTQz/NzItODIzMC5qcGVn.jpeg","resource_url":"https://api.discogs.com/artists/82079"},{"country":"US","year":"1999","format":["CD","EP","Reissue"],"label":["Fueled By Ramen","An Industry For Outer Space","Box Canyon"],"type":"master","genre":["Rock"],"style":["Alternative Rock","Acoustic","Indie Rock"],"id":127909,"barcode":["6 45131 20202 4","37962AM-02 FBR-020 011901-06","37962AM-01 FBR-020 001009-05","ascap"],"user_data":{"in_wantlist":false,"in_collection":false},"master_id":127909,"master_url":"https://api.discogs.com/masters/127909","uri":"/master/127909-Jimmy-Eat-World-Jimmy-Eat-World","catno":"FBR-020","title":"Jimmy Eat World - Jimmy Eat World","thumb":"https://i.discogs.com/M6OBYHgG2eKsaC36bn3lkw4fWDgmF-c7CDSEjoXceCg/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWltYWdlcy9SLTE0/MjM2NzUtMTUzMDI4/MzA0OC01OTc2Lmpw/ZWc.jpeg","cover_image":"https://i.discogs.com/m3JDdWLjDKiH_NwhzZST7HIckUr1nQz8OCROmqBkbsQ/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWltYWdlcy9SLTE0/MjM2NzUtMTUzMDI4/MzA0OC01OTc2Lmpw/ZWc.jpeg","resource_url":"https://api.discogs.com/masters/127909","community":{"want":782,"have":828}},{"country":"US","year":"1996","format":["Vinyl","7\"","Single"],"label":["Abridged Records","Mind's Eye Digital","Big House Studio"],"type":"release","genre":["Rock"],"style":["Emo","Indie Rock"],"id":1594296,"barcode":["ABR05 JIMMY EAT WORLD","ABR05 BLUEPRINT"],"user_data":{"in_wantlist":false,"in_collection":false},"master_id":218107,"master_url":"https://api.discogs.com/masters/218107","uri":"/Jimmy-Eat-World-Blueprint-Jimmy-Eat-World-Blueprint/release/1594296","catno":"ABR05","title":"Jimmy Eat World / Blueprint (4) - Jimmy Eat World / Blueprint","thumb":"https://i.discogs.com/Qaw4_LYr2GQA5K_fUqs4DkKaQORr9Wa0TiYiiGCCFOE/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWltYWdlcy9SLTE1/OTQyOTYtMTMwNjk0/Njg5NC5qcGVn.jpeg","cover_image":"https://i.discogs.com/EIwXl57t1rfcNnv1tCvMM6rLl2Yi7iPf5Eiu8Eok9xE/rs:fit/g:sm/q:90/h:590/w:600/czM6Ly9kaXNjb2dz/LWltYWdlcy9SLTE1/OTQyOTYtMTMwNjk0/Njg5NC5qcGVn.jpeg","resource_url":"https://api.discogs.com/releases/1594296","community":{"want":166,"have":567},"format_quantity":1,"formats":[{"name":"Vinyl","qty":"1","text":"Truck Cover","descriptions":["7\"","Single"]}]}]}`))
	})
	client := defaultTestClient(server.URL, "foo")

	ctx := context.Background()
	searchQuery := "jimmy+eat+world"
	request := SearchRequest{
		Query: &searchQuery,
	}

	url := "https://api.discogs.com/database/search?q=jimmy+eat+world&page=28&per_page=50"
	results, err := client.GetSearchResults(ctx, request)
	if err != nil {
		t.Fatal(err)
	}
	want := &SearchResultList{
		Pagination: &Pagination{
			Page:    1,
			Pages:   28,
			PerPage: 50,
			Items:   1381,
			URLs: PaginationURLs{
				Last: &url,
				Next: &url,
			},
		},
		Results: []*SearchResult{
			{
				ID:   82079,
				Type: "artist",
				UserData: &UserData{
					InWantlist:   false,
					InCollection: true,
				},
				// "master_id": null,
				// "master_url": null,
				URI:         "/artist/82079-Jimmy-Eat-World",
				Title:       "Jimmy Eat World",
				Thumb:       "https://i.discogs.com/nf8mu5qMqLjvdqLfdXH93xGVihp7hEXMYTEQZ-a3mDY/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWltYWdlcy9BLTgy/MDc5LTE2MTM0OTQz/NzItODIzMC5qcGVn.jpeg",
				CoverImage:  "https://i.discogs.com/I_vVLO6a2K2atNydPFDxaDt5WU1ccpN00UZ3QFFG8Qw/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWltYWdlcy9BLTgy/MDc5LTE2MTM0OTQz/NzItODIzMC5qcGVn.jpeg",
				ResourceURL: "https://api.discogs.com/artists/82079",
			},
			{
				Country: "US",
				Year:    "1999",
				Format: []string{
					"CD",
					"EP",
					"Reissue",
				},
				Label: []string{
					"Fueled By Ramen",
					"An Industry For Outer Space",
					"Box Canyon",
				},
				Type: "master",
				Genre: []string{
					"Rock",
				},
				Style: []string{
					"Alternative Rock",
					"Acoustic",
					"Indie Rock",
				},
				ID: 127909,
				Barcode: []string{
					"6 45131 20202 4",
					"37962AM-02 FBR-020 011901-06",
					"37962AM-01 FBR-020 001009-05",
					"ascap",
				},
				UserData: &UserData{
					InWantlist:   false,
					InCollection: false,
				},
				MasterID:    127909,
				MasterURL:   "https://api.discogs.com/masters/127909",
				URI:         "/master/127909-Jimmy-Eat-World-Jimmy-Eat-World",
				Catno:       "FBR-020",
				Title:       "Jimmy Eat World - Jimmy Eat World",
				Thumb:       "https://i.discogs.com/M6OBYHgG2eKsaC36bn3lkw4fWDgmF-c7CDSEjoXceCg/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWltYWdlcy9SLTE0/MjM2NzUtMTUzMDI4/MzA0OC01OTc2Lmpw/ZWc.jpeg",
				CoverImage:  "https://i.discogs.com/m3JDdWLjDKiH_NwhzZST7HIckUr1nQz8OCROmqBkbsQ/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWltYWdlcy9SLTE0/MjM2NzUtMTUzMDI4/MzA0OC01OTc2Lmpw/ZWc.jpeg",
				ResourceURL: "https://api.discogs.com/masters/127909",
				Community: &Community{
					Want: 782,
					Have: 828,
				},
			},
			{
				Country: "US",
				Year:    "1996",
				Format: []string{
					"Vinyl",
					"7\"",
					"Single",
				},
				Label: []string{
					"Abridged Records",
					"Mind's Eye Digital",
					"Big House Studio",
				},
				Type: "release",
				Genre: []string{
					"Rock",
				},
				Style: []string{
					"Emo",
					"Indie Rock",
				},
				ID: 1594296,
				Barcode: []string{
					"ABR05 JIMMY EAT WORLD",
					"ABR05 BLUEPRINT",
				},
				UserData: &UserData{
					InWantlist:   false,
					InCollection: false,
				},
				MasterID:    218107,
				MasterURL:   "https://api.discogs.com/masters/218107",
				URI:         "/Jimmy-Eat-World-Blueprint-Jimmy-Eat-World-Blueprint/release/1594296",
				Catno:       "ABR05",
				Title:       "Jimmy Eat World / Blueprint (4) - Jimmy Eat World / Blueprint",
				Thumb:       "https://i.discogs.com/Qaw4_LYr2GQA5K_fUqs4DkKaQORr9Wa0TiYiiGCCFOE/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWltYWdlcy9SLTE1/OTQyOTYtMTMwNjk0/Njg5NC5qcGVn.jpeg",
				CoverImage:  "https://i.discogs.com/EIwXl57t1rfcNnv1tCvMM6rLl2Yi7iPf5Eiu8Eok9xE/rs:fit/g:sm/q:90/h:590/w:600/czM6Ly9kaXNjb2dz/LWltYWdlcy9SLTE1/OTQyOTYtMTMwNjk0/Njg5NC5qcGVn.jpeg",
				ResourceURL: "https://api.discogs.com/releases/1594296",
				Community: &Community{
					Want: 166,
					Have: 567,
				},
				FormatQuantity: 1,
				Formats: []*Format{
					{
						Name: "Vinyl",
						Qty:  "1",
						Text: "Truck Cover",
						Descriptions: []string{
							"7\"",
							"Single",
						},
					},
				},
			},
		},
	}

	testEqual(t, want, results)
}
