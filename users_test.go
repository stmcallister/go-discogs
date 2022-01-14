package discogs

import (
	"context"
	"net/http"
	"testing"
)

func TestDiscogs_UsersGetCollectionFolders(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/testuser/collection/folders", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"folders": [{"id": 0, "name": "All", "count": 253, "resource_url": "https://api.discogs.com/users/stmcallister/collection/folders/0"}, {"id": 2368502, "name": "*\\WorkTunes/*", "count": 0, "resource_url": "https://api.discogs.com/users/stmcallister/collection/folders/2368502"}, {"id": 2368501, "name": "lp", "count": 220, "resource_url": "https://api.discogs.com/users/stmcallister/collection/folders/2368501"}, {"id": 2298906, "name": "Seven inch", "count": 29, "resource_url": "https://api.discogs.com/users/stmcallister/collection/folders/2298906"}, {"id": 1, "name": "Uncategorized", "count": 4, "resource_url": "https://api.discogs.com/users/stmcallister/collection/folders/1"}]}`))
	})
	client := defaultTestClient(server.URL, "foo")

	ctx := context.Background()
	username := "testuser"
	folders, err := client.GetUserCollectionFolders(ctx, username)

	if err != nil {
		t.Fatal(err)
	}
	want := &FolderList{
		Folders: []*Folder{
			{
				ID:          0,
				Name:        "All",
				Count:       253,
				ResourceURL: "https://api.discogs.com/users/stmcallister/collection/folders/0",
			},
			{
				ID:          2368502,
				Name:        "*\\WorkTunes/*",
				Count:       0,
				ResourceURL: "https://api.discogs.com/users/stmcallister/collection/folders/2368502",
			},
			{
				ID:          2368501,
				Name:        "lp",
				Count:       220,
				ResourceURL: "https://api.discogs.com/users/stmcallister/collection/folders/2368501",
			},
			{
				ID:          2298906,
				Name:        "Seven inch",
				Count:       29,
				ResourceURL: "https://api.discogs.com/users/stmcallister/collection/folders/2298906",
			},
			{
				ID:          1,
				Name:        "Uncategorized",
				Count:       4,
				ResourceURL: "https://api.discogs.com/users/stmcallister/collection/folders/1",
			},
		},
	}

	testEqual(t, want, folders)
}

func TestDiscogs_UsersGetCollectionFolder(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/testuser/collection/folders/0", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"id": 0, "name": "All", "count": 253, "resource_url": "https://api.discogs.com/users/stmcallister/collection/folders/0"}`))
	})
	client := defaultTestClient(server.URL, "foo")

	ctx := context.Background()
	username := "testuser"
	ID := 0

	folder, err := client.GetUserCollectionFolder(ctx, username, ID)

	if err != nil {
		t.Fatal(err)
	}
	want := &Folder{
		ID:          0,
		Name:        "All",
		Count:       253,
		ResourceURL: "https://api.discogs.com/users/stmcallister/collection/folders/0",
	}

	testEqual(t, want, folder)
}

func TestDiscogs_UsersGetUserCollectionItemsByRelease(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/testuser/collection/releases/0", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"pagination":{"page":1,"pages":1,"per_page":50,"items":1},"releases":[{"id":892155,"instance_id":498658685,"date_added":"2020-08-16T21:43:39-07:00","rating":0,"basic_information":{"master_id":36390,"master_url":"https://api.discogs.com/masters/36390","title":"Mary Star Of The Sea","year":2003,"formats":[{"name":"Vinyl","qty":"2","descriptions":["LP","Album"]}],"labels":[{"name":"Martha's Music","catno":"9362-48436-1","entity_type":"1","entity_type_name":"Label","id":36917,"resource_url":"https://api.discogs.com/labels/36917"},{"name":"Reprise Records","catno":"9362-48436-1","entity_type":"1","entity_type_name":"Label","id":157,"resource_url":"https://api.discogs.com/labels/157"}],"artists":[{"name":"Zwan","id":277868,"resource_url":"https://api.discogs.com/artists/277868"}],"genres":["Rock"],"styles":["Alternative Rock"]},"folder_id":2368501}]}`))
	})
	client := defaultTestClient(server.URL, "foo")

	ctx := context.Background()
	username := "testuser"
	ID := 0

	rels, err := client.GetUserCollectionItemsByRelease(ctx, username, ID)
	if err != nil {
		t.Fatal(err)
	}

	want := &ReleaseList{
		Pagination: &Pagination{
			Page:    1,
			Pages:   1,
			PerPage: 50,
			Items:   1,
		},
		Releases: []*CollectionRelease{
			{
				ID:         892155,
				InstanceID: 498658685,
				DateAdded:  "2020-08-16T21:43:39-07:00",
				Rating:     0,
				BasicInformation: &BasicInformation{
					MasterID:  36390,
					MasterURL: "https://api.discogs.com/masters/36390",
					Title:     "Mary Star Of The Sea",
					Year:      2003,
					Formats: []*Format{
						{
							Name: "Vinyl",
							Qty:  "2",
							Descriptions: []string{
								"LP",
								"Album",
							},
						},
					},
					Labels: []*Entity{
						{
							Name:           "Martha's Music",
							Catno:          "9362-48436-1",
							EntityType:     "1",
							EntityTypeName: "Label",
							ID:             36917,
							ResourceURL:    "https://api.discogs.com/labels/36917",
						},
						{
							Name:           "Reprise Records",
							Catno:          "9362-48436-1",
							EntityType:     "1",
							EntityTypeName: "Label",
							ID:             157,
							ResourceURL:    "https://api.discogs.com/labels/157",
						},
					},
					Artists: []*Artist{
						{
							Name:        "Zwan",
							ID:          277868,
							ResourceURL: "https://api.discogs.com/artists/277868",
						},
					},
					Genres: []string{
						"Rock",
					},
					Styles: []string{
						"Alternative Rock",
					},
				},
				FolderID: 2368501,
			},
		},
	}

	testEqual(t, want, rels)
}

func TestDiscogs_UsersGetUserCollectionItemsByFolder(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/testuser/collection/folders/1/releases", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"pagination":{"page":1,"pages":1,"per_page":50,"items":1},"releases":[{"id":892155,"instance_id":498658685,"date_added":"2020-08-16T21:43:39-07:00","rating":0,"basic_information":{"master_id":36390,"master_url":"https://api.discogs.com/masters/36390","title":"Mary Star Of The Sea","year":2003,"formats":[{"name":"Vinyl","qty":"2","descriptions":["LP","Album"]}],"labels":[{"name":"Martha's Music","catno":"9362-48436-1","entity_type":"1","entity_type_name":"Label","id":36917,"resource_url":"https://api.discogs.com/labels/36917"},{"name":"Reprise Records","catno":"9362-48436-1","entity_type":"1","entity_type_name":"Label","id":157,"resource_url":"https://api.discogs.com/labels/157"}],"artists":[{"name":"Zwan","id":277868,"resource_url":"https://api.discogs.com/artists/277868"}],"genres":["Rock"],"styles":["Alternative Rock"]},"folder_id":2368501}]}`))
	})
	client := defaultTestClient(server.URL, "foo")

	ctx := context.Background()
	username := "testuser"
	folderID := 1
	sort := "artist"
	page := 1
	per := 100

	rels, err := client.GetUserCollectionItemsByFolder(ctx, username, sort, folderID, page, per)
	if err != nil {
		t.Fatal(err)
	}

	want := &ReleaseList{
		Pagination: &Pagination{
			Page:    1,
			Pages:   1,
			PerPage: 50,
			Items:   1,
		},
		Releases: []*CollectionRelease{
			{
				ID:         892155,
				InstanceID: 498658685,
				DateAdded:  "2020-08-16T21:43:39-07:00",
				Rating:     0,
				BasicInformation: &BasicInformation{
					MasterID:  36390,
					MasterURL: "https://api.discogs.com/masters/36390",
					Title:     "Mary Star Of The Sea",
					Year:      2003,
					Formats: []*Format{
						{
							Name: "Vinyl",
							Qty:  "2",
							Descriptions: []string{
								"LP",
								"Album",
							},
						},
					},
					Labels: []*Entity{
						{
							Name:           "Martha's Music",
							Catno:          "9362-48436-1",
							EntityType:     "1",
							EntityTypeName: "Label",
							ID:             36917,
							ResourceURL:    "https://api.discogs.com/labels/36917",
						},
						{
							Name:           "Reprise Records",
							Catno:          "9362-48436-1",
							EntityType:     "1",
							EntityTypeName: "Label",
							ID:             157,
							ResourceURL:    "https://api.discogs.com/labels/157",
						},
					},
					Artists: []*Artist{
						{
							Name:        "Zwan",
							ID:          277868,
							ResourceURL: "https://api.discogs.com/artists/277868",
						},
					},
					Genres: []string{
						"Rock",
					},
					Styles: []string{
						"Alternative Rock",
					},
				},
				FolderID: 2368501,
			},
		},
	}

	testEqual(t, want, rels)
}

func TestDiscogs_UsersGetAllUserCollectionItemsByFolder(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/testuser/collection/folders/1/releases", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"pagination":{"page":1,"pages":1,"per_page":50,"items":1},"releases":[{"id":892155,"instance_id":498658685,"date_added":"2020-08-16T21:43:39-07:00","rating":0,"basic_information":{"master_id":36390,"master_url":"https://api.discogs.com/masters/36390","title":"Mary Star Of The Sea","year":2003,"formats":[{"name":"Vinyl","qty":"2","descriptions":["LP","Album"]}],"labels":[{"name":"Martha's Music","catno":"9362-48436-1","entity_type":"1","entity_type_name":"Label","id":36917,"resource_url":"https://api.discogs.com/labels/36917"},{"name":"Reprise Records","catno":"9362-48436-1","entity_type":"1","entity_type_name":"Label","id":157,"resource_url":"https://api.discogs.com/labels/157"}],"artists":[{"name":"Zwan","id":277868,"resource_url":"https://api.discogs.com/artists/277868"}],"genres":["Rock"],"styles":["Alternative Rock"]},"folder_id":2368501}]}`))
	})
	client := defaultTestClient(server.URL, "foo")

	ctx := context.Background()
	username := "testuser"
	folderID := 1
	sort := "artist"

	rels, err := client.GetUserCollectionAllItemsByFolder(ctx, username, sort, folderID)
	if err != nil {
		t.Fatal(err)
	}

	want := &ReleaseList{
		Releases: []*CollectionRelease{
			{
				ID:         892155,
				InstanceID: 498658685,
				DateAdded:  "2020-08-16T21:43:39-07:00",
				Rating:     0,
				BasicInformation: &BasicInformation{
					MasterID:  36390,
					MasterURL: "https://api.discogs.com/masters/36390",
					Title:     "Mary Star Of The Sea",
					Year:      2003,
					Formats: []*Format{
						{
							Name: "Vinyl",
							Qty:  "2",
							Descriptions: []string{
								"LP",
								"Album",
							},
						},
					},
					Labels: []*Entity{
						{
							Name:           "Martha's Music",
							Catno:          "9362-48436-1",
							EntityType:     "1",
							EntityTypeName: "Label",
							ID:             36917,
							ResourceURL:    "https://api.discogs.com/labels/36917",
						},
						{
							Name:           "Reprise Records",
							Catno:          "9362-48436-1",
							EntityType:     "1",
							EntityTypeName: "Label",
							ID:             157,
							ResourceURL:    "https://api.discogs.com/labels/157",
						},
					},
					Artists: []*Artist{
						{
							Name:        "Zwan",
							ID:          277868,
							ResourceURL: "https://api.discogs.com/artists/277868",
						},
					},
					Genres: []string{
						"Rock",
					},
					Styles: []string{
						"Alternative Rock",
					},
				},
				FolderID: 2368501,
			},
		},
	}
	testEqual(t, want, rels)
}
