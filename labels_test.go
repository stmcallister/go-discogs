package discogs

import (
	"context"
	"net/http"
	"testing"
)

func TestDiscogs_LabelsGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/labels/99512", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"id":99512,"name":"Temporary Residence Limited","resource_url":"https://api.discogs.com/labels/9952","uri":"https://www.discogs.com/label/9952-Temporary-Residence-Limited","releases_url":"https://api.discogs.com/labels/9952/releases","images":[{"type":"primary","uri":"https://i.discogs.com/3WhxhnP1uMwu9Lo7iX-UkuKgFsMBVJVyiS_YIZgcQ_A/rs:fit/g:sm/q:90/h:636/w:600/czM6Ly9kaXNjb2dz/LWltYWdlcy9MLTk5/NTItMTYwMDg0Njg1/OC0zNTM2LmpwZWc.jpeg","resource_url":"https://i.discogs.com/3WhxhnP1uMwu9Lo7iX-UkuKgFsMBVJVyiS_YIZgcQ_A/rs:fit/g:sm/q:90/h:636/w:600/czM6Ly9kaXNjb2dz/LWltYWdlcy9MLTk5/NTItMTYwMDg0Njg1/OC0zNTM2LmpwZWc.jpeg","uri150":"https://i.discogs.com/NLzhVuBlqcurOzYMEr89GPz0PeuevwvHkodFzgyFDi0/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWltYWdlcy9MLTk5/NTItMTYwMDg0Njg1/OC0zNTM2LmpwZWc.jpeg","width":600,"height":636}],"contact_info":"Temporary Residence Limited\r\nPO Box 60097\r\nBrooklyn\r\nNY 11206\r\nUSA\r\n\r\ntrl@temporaryresidence.com\r\ninfo@temporaryresidence.com\r\n\r\n(1997)\r\nThe Temporary Residence Limited\r\nPost Office Box 22910\r\nBaltimore, Maryland 21203-4910\r\nUSA","profile":"Independent label based in Brooklyn, New York. Established 1996.\r\n[b]For the company and holder of copyrights, use [l203831][/b]","data_quality":"Correct","urls":["http://temporaryresidence.com","http://www.facebook.com/temporaryresidence"],"sublabels":[{"id":648273,"name":"Sounds Of The Geographically Challenged","resource_url":"https://api.discogs.com/labels/648273"}]}`))
	})
	client := defaultTestClient(server.URL, "foo")

	ctx := context.Background()
	ID := 99512
	rat, err := client.GetLabel(ctx, ID)
	if err != nil {
		t.Fatal(err)
	}
	want := &Label{
			ID: 99512,
			Name: "Temporary Residence Limited",
			ResourceURL: "https://api.discogs.com/labels/9952",
			URI: "https://www.discogs.com/label/9952-Temporary-Residence-Limited",
			ReleasesURL: "https://api.discogs.com/labels/9952/releases",
			Images: []*Image{
				{
				Type: "primary",
				URI: "https://i.discogs.com/3WhxhnP1uMwu9Lo7iX-UkuKgFsMBVJVyiS_YIZgcQ_A/rs:fit/g:sm/q:90/h:636/w:600/czM6Ly9kaXNjb2dz/LWltYWdlcy9MLTk5/NTItMTYwMDg0Njg1/OC0zNTM2LmpwZWc.jpeg",
				ResourceURL: "https://i.discogs.com/3WhxhnP1uMwu9Lo7iX-UkuKgFsMBVJVyiS_YIZgcQ_A/rs:fit/g:sm/q:90/h:636/w:600/czM6Ly9kaXNjb2dz/LWltYWdlcy9MLTk5/NTItMTYwMDg0Njg1/OC0zNTM2LmpwZWc.jpeg",
				URI150: "https://i.discogs.com/NLzhVuBlqcurOzYMEr89GPz0PeuevwvHkodFzgyFDi0/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWltYWdlcy9MLTk5/NTItMTYwMDg0Njg1/OC0zNTM2LmpwZWc.jpeg",
				Width: 600,
				Height: 636,
				},
			},
			ContactInfo: "Temporary Residence Limited\r\nPO Box 60097\r\nBrooklyn\r\nNY 11206\r\nUSA\r\n\r\ntrl@temporaryresidence.com\r\ninfo@temporaryresidence.com\r\n\r\n(1997)\r\nThe Temporary Residence Limited\r\nPost Office Box 22910\r\nBaltimore, Maryland 21203-4910\r\nUSA",
			Profile: "Independent label based in Brooklyn, New York. Established 1996.\r\n[b]For the company and holder of copyrights, use [l203831][/b]",
			DataQuality: "Correct",
			URLs: []string{
				"http://temporaryresidence.com",
				"http://www.facebook.com/temporaryresidence",
			},
			SubLabels: []*SubLabel{
				{
					ID : 648273,
					Name: "Sounds Of The Geographically Challenged",
					ResourceURL: "https://api.discogs.com/labels/648273",
				},
			},
		}

	testEqual(t, want, rat)

}

