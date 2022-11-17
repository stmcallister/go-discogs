package discogs

import (
	"context"
	"net/http"
	"testing"
)

func TestDiscogs_ListingGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/marketplace/listings/2218783261", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = w.Write([]byte(`{"id": 2218783261, "resource_url": "https://api.discogs.com/marketplace/listings/2218783261", "uri": "https://www.discogs.com/sell/item/2218783261", "status": "For Sale", "condition": "Near Mint (NM or M-)", "sleeve_condition": "Near Mint (NM or M-)", "comments": "B&amp;W paper inner sleeve with photos. Lightly corner bumped. 1-25300-A SR3 SRC 1-2 1-25300-B SR3 SRC 1-1SMI-2", "ships_from": "United States", "posted": "2022-11-01T13:58:56-07:00", "allow_offers": true, "audio": false, "price": {"value": 35.0, "currency": "USD"}, "original_price": {"curr_abbr": "USD", "curr_id": 1, "formatted": "$35.00", "value": 35.0}, "shipping_price": {"value": 13.0, "currency": "USD"}, "original_shipping_price": {"curr_abbr": "USD", "curr_id": 1, "formatted": "$13.00", "value": 13}, "seller": {"id": 13979822, "username": "PrinceofPiney", "avatar_url": "https://secure.gravatar.com/avatar/66a2a988a85b6d59560920cdbe80f15d?s=500&r=pg&d=mm", "stats": {"rating": "100.0", "stars": 5.0, "total": 3}, "min_order_total": 5.0, "html_url": "https://www.discogs.com/user/PrinceofPiney", "uid": 13979822, "url": "https://api.discogs.com/users/PrinceofPiney", "payment": "PayPal Commerce", "shipping": "Visually graded. Audio grade upon request.\r\nCredit/Debit Card only\r\nCustomer is responsible for shipping**************\r\nDOMESTIC POSTAGE (Within the U.S & Puerto Rico)\r\nLP & 12: - U.S.- Priority (2-3 Days to receive) $13.00 .. each additional $2.50\r\nLP & 12: - U.S.- Media $7.25 (1-2 LPs) each additional $.75\r\n78 RPM SHELLAC RECORDS U.S. - Media $2.15 .. each additional $1.00\r\n78 RPM SHELLAC RECORDS U.S. - First Class $13.00 .. up to 2\r\n45's - First Class - U.S. $4.75 (Up to 3)... each additional $.50\r\n45's - Media Mail - U.S. $4.40 (Up To 3) ... each additional $.25\r\n-----------------------------------------------------------------------------------------------------------------------------\r\nINTERNATIONAL POSTAGE RATES (*** PLEASE NOTE - SELLER NOT RESPONSIBLE FOR ANY CUSTOMS FEES INCURRED ON INTERNATIONAL ORDERS *****)\r\nLP & 12\" - Mexico $20 minimum but to be determined by postal weight per the USPS ... 1-2 LP's ... any additional will be based on weight and calculated on request\r\nLP & 12\" - Canada $20 minimum but to be determined by postal weight per the USPS (single LP) .. any additional will be based on weight and calculated on request\r\nLP & 12\" - Europe, South America & International (Other Than Canada/Mexico) $27 minimum but to be determined by postal weight per the USPS ... each additional determined by weight\r\nLP & 12\" - Australia $27 minimum but to be determined by postal weight per the USPS (1-2) LP's\r\n78 RPM SHELLAC RECORDS - Internationally - $27 minimum but to be determined by postal weight per the USPS each additional determined by weight\r\n78 RPM SHELLAC RECORDS - Canada - $23 minimum but to be determined by postal weight per the USPS.. each additional determined by weight\r\n45's - Europe, Australia, South America & International (Other Than Canada/Mexico) $20.00 minimum but to be determined by postal weight per the USPS... (1-3 units)\r\n45's - Canada $15Credit/Debit Card only\r\nCustomer is responsible for shipping**************\r\nDOMESTIC POSTAGE (Within the U.S & Puerto Rico)\r\nLP & 12: - U.S.- Priority (2-3 Days to receive) $13.00 .. each additional $2.50\r\nLP & 12: - U.S.- Media $7.25 (1-2 LPs) each additional $.75\r\n78 RPM SHELLAC RECORDS U.S. - Media $2.15 .. each additional $1.00\r\n78 RPM SHELLAC RECORDS U.S. - First Class $13.00 .. up to 2\r\n45's - First Class - U.S. $4.75 (Up to 3)... each additional $.50\r\n45's - Media Mail - U.S. $4.40 (Up To 3) ... each additional $.25\r\n-----------------------------------------------------------------------------------------------------------------------------\r\nINTERNATIONAL POSTAGE RATES (*** PLEASE NOTE - SELLER NOT RESPONSIBLE FOR ANY CUSTOMS FEES INCURRED ON INTERNATIONAL ORDERS *****)\r\nLP & 12\" - Mexico $20 minimum but to be determined by postal weight per the USPS ... 1-2 LP's ... any additional will be based on weight and calculated on request\r\nLP & 12\" - Canada $20 minimum but to be determined by postal weight per the USPS (single LP) .. any additional will be based on weight and calculated on request\r\nLP & 12\" - Europe, South America & International (Other Than Canada/Mexico) $27 minimum but to be determined by postal weight per the USPS ... each additional determined by weight\r\nLP & 12\" - Australia $27 minimum but to be determined by postal weight per the USPS (1-2) LP's\r\n78 RPM SHELLAC RECORDS - Internationally - $27 minimum but to be determined by postal weight per the USPS each additional determined by weight\r\n78 RPM SHELLAC RECORDS - Canada - $23 minimum but to be determined by postal weight per the USPS.. each additional determined by weight\r\n45's - Europe, Australia, South America & International (Other Than Canada/Mexico) $20.00 minimum but to be determined by postal weight per the USPS... (1-3 units)\r\n45's - Canada $13 minimum but to be determined by postal weight per the USPS ... (1-3 units) minimum but to be determined by postal weight per the USPS ... (1-3 units)\r\n", "resource_url": "https://api.discogs.com/users/PrinceofPiney"}, "release": {"thumbnail": "https://i.discogs.com/0WcPbhM-S6OgvSQtJcA4x8A9NLImz2RCo_KYaSC3ZDM/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xNDg5NjEyNzY1/LTU2MjMuanBlZw.jpeg", "description": "a-ha - Hunting High And Low (LP, Album, SRC)", "images": [{"type": "primary", "uri": "https://i.discogs.com/UgFI5HydKshg4U10q1oI0oX5-iJpNfbUFE6sLImcNX8/rs:fit/g:sm/q:90/h:596/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xNDg5NjEyNzY1/LTU2MjMuanBlZw.jpeg", "resource_url": "https://i.discogs.com/UgFI5HydKshg4U10q1oI0oX5-iJpNfbUFE6sLImcNX8/rs:fit/g:sm/q:90/h:596/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xNDg5NjEyNzY1/LTU2MjMuanBlZw.jpeg", "uri150": "https://i.discogs.com/0WcPbhM-S6OgvSQtJcA4x8A9NLImz2RCo_KYaSC3ZDM/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xNDg5NjEyNzY1/LTU2MjMuanBlZw.jpeg", "width": 600, "height": 596}], "artist": "a-ha", "format": "LP, Album, SRC", "resource_url": "https://api.discogs.com/releases/547689", "title": "Hunting High And Low", "year": 1985, "id": 547689, "catalog_number": "1-25300, 9 25300-1", "stats": {"community": {"in_wantlist": 427, "in_collection": 3110}, "user": {"in_wantlist": 1, "in_collection": 0}}}, "in_cart": false}`))
	})
	client := defaultTestClient(server.URL, "foo")

	ctx := context.Background()
	listingID := 2218783261
	currency := ""
	listing, err := client.GetListing(ctx, listingID, currency)
	if err != nil {
		t.Fatal(err)
	}

	want := &Listing{
		ID:              listingID,
		ResourceURL:     "https://api.discogs.com/marketplace/listings/2218783261",
		URI:             "https://www.discogs.com/sell/item/2218783261",
		Status:          "For Sale",
		Condition:       "Near Mint (NM or M-)",
		SleeveCondition: "Near Mint (NM or M-)",
		Comments:        "B&amp;W paper inner sleeve with photos. Lightly corner bumped. 1-25300-A SR3 SRC 1-2 1-25300-B SR3 SRC 1-1SMI-2",
		ShipsFrom:       "United States",
		Posted:          "2022-11-01T13:58:56-07:00",
		AllowOffers:     true,
		Audio:           false,
		Price: Price{
			Value:    35.00,
			Currency: "USD",
		},
		OriginalPrice: OriginalPrice{
			CurrAbbr:  "USD",
			CurrID:    1,
			Formatted: "$35.00",
			Value:     35.00,
		},
		ShippingPrice: Price{
			Value:    13.00,
			Currency: "USD",
		},
		OriginalShippingPrice: OriginalPrice{
			CurrAbbr:  "USD",
			CurrID:    1,
			Formatted: "$13.00",
			Value:     13.00,
		},
		Seller: Seller{
			ID:        13979822,
			Username:  "PrinceofPiney",
			AvatarURL: "https://secure.gravatar.com/avatar/66a2a988a85b6d59560920cdbe80f15d?s=500&r=pg&d=mm",
			Stats: SellerStats{
				Rating: "100.0",
				Stars:  5.0,
				Total:  3,
			},
			MinOrderTotal: 5.0,
			HtmlURL:       "https://www.discogs.com/user/PrinceofPiney",
			UID:           13979822,
			URL:           "https://api.discogs.com/users/PrinceofPiney",
			Payment:       "PayPal Commerce",
			Shipping:      "Visually graded. Audio grade upon request.\r\nCredit/Debit Card only\r\nCustomer is responsible for shipping**************\r\nDOMESTIC POSTAGE (Within the U.S & Puerto Rico)\r\nLP & 12: - U.S.- Priority (2-3 Days to receive) $13.00 .. each additional $2.50\r\nLP & 12: - U.S.- Media $7.25 (1-2 LPs) each additional $.75\r\n78 RPM SHELLAC RECORDS U.S. - Media $2.15 .. each additional $1.00\r\n78 RPM SHELLAC RECORDS U.S. - First Class $13.00 .. up to 2\r\n45's - First Class - U.S. $4.75 (Up to 3)... each additional $.50\r\n45's - Media Mail - U.S. $4.40 (Up To 3) ... each additional $.25\r\n-----------------------------------------------------------------------------------------------------------------------------\r\nINTERNATIONAL POSTAGE RATES (*** PLEASE NOTE - SELLER NOT RESPONSIBLE FOR ANY CUSTOMS FEES INCURRED ON INTERNATIONAL ORDERS *****)\r\nLP & 12\" - Mexico $20 minimum but to be determined by postal weight per the USPS ... 1-2 LP's ... any additional will be based on weight and calculated on request\r\nLP & 12\" - Canada $20 minimum but to be determined by postal weight per the USPS (single LP) .. any additional will be based on weight and calculated on request\r\nLP & 12\" - Europe, South America & International (Other Than Canada/Mexico) $27 minimum but to be determined by postal weight per the USPS ... each additional determined by weight\r\nLP & 12\" - Australia $27 minimum but to be determined by postal weight per the USPS (1-2) LP's\r\n78 RPM SHELLAC RECORDS - Internationally - $27 minimum but to be determined by postal weight per the USPS each additional determined by weight\r\n78 RPM SHELLAC RECORDS - Canada - $23 minimum but to be determined by postal weight per the USPS.. each additional determined by weight\r\n45's - Europe, Australia, South America & International (Other Than Canada/Mexico) $20.00 minimum but to be determined by postal weight per the USPS... (1-3 units)\r\n45's - Canada $15Credit/Debit Card only\r\nCustomer is responsible for shipping**************\r\nDOMESTIC POSTAGE (Within the U.S & Puerto Rico)\r\nLP & 12: - U.S.- Priority (2-3 Days to receive) $13.00 .. each additional $2.50\r\nLP & 12: - U.S.- Media $7.25 (1-2 LPs) each additional $.75\r\n78 RPM SHELLAC RECORDS U.S. - Media $2.15 .. each additional $1.00\r\n78 RPM SHELLAC RECORDS U.S. - First Class $13.00 .. up to 2\r\n45's - First Class - U.S. $4.75 (Up to 3)... each additional $.50\r\n45's - Media Mail - U.S. $4.40 (Up To 3) ... each additional $.25\r\n-----------------------------------------------------------------------------------------------------------------------------\r\nINTERNATIONAL POSTAGE RATES (*** PLEASE NOTE - SELLER NOT RESPONSIBLE FOR ANY CUSTOMS FEES INCURRED ON INTERNATIONAL ORDERS *****)\r\nLP & 12\" - Mexico $20 minimum but to be determined by postal weight per the USPS ... 1-2 LP's ... any additional will be based on weight and calculated on request\r\nLP & 12\" - Canada $20 minimum but to be determined by postal weight per the USPS (single LP) .. any additional will be based on weight and calculated on request\r\nLP & 12\" - Europe, South America & International (Other Than Canada/Mexico) $27 minimum but to be determined by postal weight per the USPS ... each additional determined by weight\r\nLP & 12\" - Australia $27 minimum but to be determined by postal weight per the USPS (1-2) LP's\r\n78 RPM SHELLAC RECORDS - Internationally - $27 minimum but to be determined by postal weight per the USPS each additional determined by weight\r\n78 RPM SHELLAC RECORDS - Canada - $23 minimum but to be determined by postal weight per the USPS.. each additional determined by weight\r\n45's - Europe, Australia, South America & International (Other Than Canada/Mexico) $20.00 minimum but to be determined by postal weight per the USPS... (1-3 units)\r\n45's - Canada $13 minimum but to be determined by postal weight per the USPS ... (1-3 units) minimum but to be determined by postal weight per the USPS ... (1-3 units)\r\n",
			ResourceURL:   "https://api.discogs.com/users/PrinceofPiney",
		},
		Release: MarketplaceRelease{
			Thumbnail:   "https://i.discogs.com/0WcPbhM-S6OgvSQtJcA4x8A9NLImz2RCo_KYaSC3ZDM/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xNDg5NjEyNzY1/LTU2MjMuanBlZw.jpeg",
			Description: "a-ha - Hunting High And Low (LP, Album, SRC)",
			Artist:      "a-ha",
			Format:      "LP, Album, SRC",
			Images: []*Image{
				{
					Type:        "primary",
					URI:         "https://i.discogs.com/UgFI5HydKshg4U10q1oI0oX5-iJpNfbUFE6sLImcNX8/rs:fit/g:sm/q:90/h:596/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xNDg5NjEyNzY1/LTU2MjMuanBlZw.jpeg",
					ResourceURL: "https://i.discogs.com/UgFI5HydKshg4U10q1oI0oX5-iJpNfbUFE6sLImcNX8/rs:fit/g:sm/q:90/h:596/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xNDg5NjEyNzY1/LTU2MjMuanBlZw.jpeg",
					URI150:      "https://i.discogs.com/0WcPbhM-S6OgvSQtJcA4x8A9NLImz2RCo_KYaSC3ZDM/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xNDg5NjEyNzY1/LTU2MjMuanBlZw.jpeg",
					Width:       600,
					Height:      596,
				},
			},
		},
		InCart: false,
	}

	testEqual(t, want, listing)

}
