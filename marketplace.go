package discogs

import (
	"context"
	"fmt"
	"net/http"
)

/* TODO: Test putting an item in the marketplace to confirm the following note in the docs.
   If the authorized user is the listing owner the listing will include the weight,
   xformat_quantity, external_id, location, and quantity keys.
*/
type Listing struct {
	ID                    int                `json:"id"`
	ResourceURL           string             `json:"resource_url"`
	URI                   string             `json:"uri"`
	Status                string             `json:"status"`
	Condition             string             `json:"condition"`
	SleeveCondition       string             `json:"sleeve_condition"`
	Comments              string             `json:"comments"`
	ShipsFrom             string             `json:"ships_from"`
	Posted                string             `json:"posted"`
	AllowOffers           bool               `json:"allow_offers"`
	Audio                 bool               `json:"audio"`
	Price                 Price              `json:"price"`
	OriginalPrice         OriginalPrice      `json:"original_price"`
	ShippingPrice         Price              `json:"shipping_price"`
	OriginalShippingPrice OriginalPrice      `json:"original_shipping_price"`
	Seller                Seller             `json:"seller"`
	Release               MarketplaceRelease `json:"release"`
	InCart                bool               `json:"in_cart"`
}

type Price struct {
	Value    float32 `json:"value"`
	Currency string  `json:"currency"`
}

type OriginalPrice struct {
	CurrAbbr  string  `json:"curr_abbr"`
	CurrID    int     `json:"curr_id"`
	Formatted string  `json:"formatted"`
	Value     float32 `json:"value"`
}

type Seller struct {
	ID            int         `json:"id"`
	Username      string      `json:"username"`
	AvatarURL     string      `json:"avatar_url"`
	Stats         SellerStats `json:"stats"`
	MinOrderTotal float32     `json:"min_order_total"`
	HtmlURL       string      `json:"html_url"`
	UID           int         `json:"uid"`
	URL           string      `json:"url"`
	Payment       string      `json:"payment"`
	Shipping      string      `json:"shipping"`
	ResourceURL   string      `json:"resource_url"`
}

type SellerStats struct {
	Rating string  `json:"rating"`
	Stars  float32 `json:"stars"`
	Total  int     `json:"total"`
}

type MarketplaceRelease struct {
	Thumbnail   string   `json:"thumbnail"`
	Description string   `json:"description"`
	Images      []*Image `json:"images"`
	Artist      string   `json:"artist"`
	Format      string   `json:"format"`
}

// GetListing is a function for getting a single listing
func (c *Client) GetListing(ctx context.Context, ID int, currency string) (*Listing, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/marketplace/listings/%d?%s", c.baseURL, ID, currency), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	listing := Listing{}
	if err := c.sendRequest(req, &listing); err != nil {
		return nil, err
	}

	return &listing, nil
}

/*

    "seller": {
        "id": 13979822,
        "username": "PrinceofPiney",
        "avatar_url": "https://secure.gravatar.com/avatar/66a2a988a85b6d59560920cdbe80f15d?s=500&r=pg&d=mm",
        "stats": {
            "rating": "100.0",
            "stars": 5.0,
            "total": 3
        },
        "min_order_total": 5.0,
        "html_url": "https://www.discogs.com/user/PrinceofPiney",
        "uid": 13979822,
        "url": "https://api.discogs.com/users/PrinceofPiney",
        "payment": "PayPal Commerce",
        "shipping": "Visually graded. Audio grade upon request.\r\nCredit/Debit Card only\r\nCustomer is responsible for shipping**************\r\nDOMESTIC POSTAGE (Within the U.S & Puerto Rico)\r\nLP & 12: - U.S.- Priority (2-3 Days to receive) $13.00 .. each additional $2.50\r\nLP & 12: - U.S.- Media $7.25 (1-2 LPs) each additional $.75\r\n78 RPM SHELLAC RECORDS U.S. - Media $2.15 .. each additional $1.00\r\n78 RPM SHELLAC RECORDS U.S. - First Class $13.00 .. up to 2\r\n45's - First Class - U.S. $4.75 (Up to 3)... each additional $.50\r\n45's - Media Mail - U.S. $4.40 (Up To 3) ... each additional $.25\r\n-----------------------------------------------------------------------------------------------------------------------------\r\nINTERNATIONAL POSTAGE RATES (*** PLEASE NOTE - SELLER NOT RESPONSIBLE FOR ANY CUSTOMS FEES INCURRED ON INTERNATIONAL ORDERS *****)\r\nLP & 12\" - Mexico $20 minimum but to be determined by postal weight per the USPS ... 1-2 LP's ... any additional will be based on weight and calculated on request\r\nLP & 12\" - Canada $20 minimum but to be determined by postal weight per the USPS (single LP) .. any additional will be based on weight and calculated on request\r\nLP & 12\" - Europe, South America & International (Other Than Canada/Mexico) $27 minimum but to be determined by postal weight per the USPS ... each additional determined by weight\r\nLP & 12\" - Australia $27 minimum but to be determined by postal weight per the USPS (1-2) LP's\r\n78 RPM SHELLAC RECORDS - Internationally - $27 minimum but to be determined by postal weight per the USPS each additional determined by weight\r\n78 RPM SHELLAC RECORDS - Canada - $23 minimum but to be determined by postal weight per the USPS.. each additional determined by weight\r\n45's - Europe, Australia, South America & International (Other Than Canada/Mexico) $20.00 minimum but to be determined by postal weight per the USPS... (1-3 units)\r\n45's - Canada $15Credit/Debit Card only\r\nCustomer is responsible for shipping**************\r\nDOMESTIC POSTAGE (Within the U.S & Puerto Rico)\r\nLP & 12: - U.S.- Priority (2-3 Days to receive) $13.00 .. each additional $2.50\r\nLP & 12: - U.S.- Media $7.25 (1-2 LPs) each additional $.75\r\n78 RPM SHELLAC RECORDS U.S. - Media $2.15 .. each additional $1.00\r\n78 RPM SHELLAC RECORDS U.S. - First Class $13.00 .. up to 2\r\n45's - First Class - U.S. $4.75 (Up to 3)... each additional $.50\r\n45's - Media Mail - U.S. $4.40 (Up To 3) ... each additional $.25\r\n-----------------------------------------------------------------------------------------------------------------------------\r\nINTERNATIONAL POSTAGE RATES (*** PLEASE NOTE - SELLER NOT RESPONSIBLE FOR ANY CUSTOMS FEES INCURRED ON INTERNATIONAL ORDERS *****)\r\nLP & 12\" - Mexico $20 minimum but to be determined by postal weight per the USPS ... 1-2 LP's ... any additional will be based on weight and calculated on request\r\nLP & 12\" - Canada $20 minimum but to be determined by postal weight per the USPS (single LP) .. any additional will be based on weight and calculated on request\r\nLP & 12\" - Europe, South America & International (Other Than Canada/Mexico) $27 minimum but to be determined by postal weight per the USPS ... each additional determined by weight\r\nLP & 12\" - Australia $27 minimum but to be determined by postal weight per the USPS (1-2) LP's\r\n78 RPM SHELLAC RECORDS - Internationally - $27 minimum but to be determined by postal weight per the USPS each additional determined by weight\r\n78 RPM SHELLAC RECORDS - Canada - $23 minimum but to be determined by postal weight per the USPS.. each additional determined by weight\r\n45's - Europe, Australia, South America & International (Other Than Canada/Mexico) $20.00 minimum but to be determined by postal weight per the USPS... (1-3 units)\r\n45's - Canada $13 minimum but to be determined by postal weight per the USPS ... (1-3 units) minimum but to be determined by postal weight per the USPS ... (1-3 units)\r\n",
        "resource_url": "https://api.discogs.com/users/PrinceofPiney"
    },
    "release": {
        "thumbnail": "https://i.discogs.com/0WcPbhM-S6OgvSQtJcA4x8A9NLImz2RCo_KYaSC3ZDM/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xNDg5NjEyNzY1/LTU2MjMuanBlZw.jpeg",
        "description": "a-ha - Hunting High And Low (LP, Album, SRC)",
        "images": [
            {
                "type": "primary",
                "uri": "https://i.discogs.com/UgFI5HydKshg4U10q1oI0oX5-iJpNfbUFE6sLImcNX8/rs:fit/g:sm/q:90/h:596/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xNDg5NjEyNzY1/LTU2MjMuanBlZw.jpeg",
                "resource_url": "https://i.discogs.com/UgFI5HydKshg4U10q1oI0oX5-iJpNfbUFE6sLImcNX8/rs:fit/g:sm/q:90/h:596/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xNDg5NjEyNzY1/LTU2MjMuanBlZw.jpeg",
                "uri150": "https://i.discogs.com/0WcPbhM-S6OgvSQtJcA4x8A9NLImz2RCo_KYaSC3ZDM/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xNDg5NjEyNzY1/LTU2MjMuanBlZw.jpeg",
                "width": 600,
                "height": 596
            },
            {
                "type": "secondary",
                "uri": "https://i.discogs.com/d_xQH7kb_nhk6aTlRDrXxVMVZnTz3FE3OjT6FBEOdMI/rs:fit/g:sm/q:90/h:604/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xNDg5NjEyNzY2/LTU5ODEuanBlZw.jpeg",
                "resource_url": "https://i.discogs.com/d_xQH7kb_nhk6aTlRDrXxVMVZnTz3FE3OjT6FBEOdMI/rs:fit/g:sm/q:90/h:604/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xNDg5NjEyNzY2/LTU5ODEuanBlZw.jpeg",
                "uri150": "https://i.discogs.com/vCPvLPCyiTMNOo68oulcWwgEBDDhJtwbMxao9MHjP8c/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xNDg5NjEyNzY2/LTU5ODEuanBlZw.jpeg",
                "width": 600,
                "height": 604
            },
            {
                "type": "secondary",
                "uri": "https://i.discogs.com/8clJ2zrDNHU_ZLmNNk8Ou6vnCe-h2mO5Pbo5SjQyH4Y/rs:fit/g:sm/q:90/h:570/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xMTU1ODkzMTY4/LmpwZWc.jpeg",
                "resource_url": "https://i.discogs.com/8clJ2zrDNHU_ZLmNNk8Ou6vnCe-h2mO5Pbo5SjQyH4Y/rs:fit/g:sm/q:90/h:570/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xMTU1ODkzMTY4/LmpwZWc.jpeg",
                "uri150": "https://i.discogs.com/-DXVO3PCMYT8WUuauDrHt17jIqpgjp5-5tiIBqsTnaU/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xMTU1ODkzMTY4/LmpwZWc.jpeg",
                "width": 600,
                "height": 570
            },
            {
                "type": "secondary",
                "uri": "https://i.discogs.com/kq2IHn87n5QxWh4EBZFwwIrzSybxgAp5cz9g0GiIPoE/rs:fit/g:sm/q:90/h:580/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xMTU1ODkzMTkz/LmpwZWc.jpeg",
                "resource_url": "https://i.discogs.com/kq2IHn87n5QxWh4EBZFwwIrzSybxgAp5cz9g0GiIPoE/rs:fit/g:sm/q:90/h:580/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xMTU1ODkzMTkz/LmpwZWc.jpeg",
                "uri150": "https://i.discogs.com/WDJU56eyfdGd-OanwT5efKQxzQ_T75hcOfTwjfhxJYQ/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xMTU1ODkzMTkz/LmpwZWc.jpeg",
                "width": 600,
                "height": 580
            },
            {
                "type": "secondary",
                "uri": "https://i.discogs.com/zjzUnAlAabKsTVYUOocXfrREukbcAqeWozt2HAME8eY/rs:fit/g:sm/q:90/h:587/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xMzA0MTg5ODMx/LmpwZWc.jpeg",
                "resource_url": "https://i.discogs.com/zjzUnAlAabKsTVYUOocXfrREukbcAqeWozt2HAME8eY/rs:fit/g:sm/q:90/h:587/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xMzA0MTg5ODMx/LmpwZWc.jpeg",
                "uri150": "https://i.discogs.com/jOiHWntwv3rwSPuV4GBWTqL0KQKt6_ZYw5Pynzwajk4/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xMzA0MTg5ODMx/LmpwZWc.jpeg",
                "width": 600,
                "height": 587
            },
            {
                "type": "secondary",
                "uri": "https://i.discogs.com/1BLdADq5yI-GQ7fE62C4ae9TNeykBnsg2HpXhFfGIZk/rs:fit/g:sm/q:90/h:585/w:599/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xMzA0MTg5ODM5/LmpwZWc.jpeg",
                "resource_url": "https://i.discogs.com/1BLdADq5yI-GQ7fE62C4ae9TNeykBnsg2HpXhFfGIZk/rs:fit/g:sm/q:90/h:585/w:599/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xMzA0MTg5ODM5/LmpwZWc.jpeg",
                "uri150": "https://i.discogs.com/WHg8eeaQlMl2djlisLTXXdSuqUyakX5GDB67CeFR-UI/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xMzA0MTg5ODM5/LmpwZWc.jpeg",
                "width": 599,
                "height": 585
            },
            {
                "type": "secondary",
                "uri": "https://i.discogs.com/eEQUqur6cGQCl8krTRo2kpiB1lkuY2zZi3FfkBsgAWY/rs:fit/g:sm/q:90/h:196/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xNjYzOTQwMzIy/LTQyODYuanBlZw.jpeg",
                "resource_url": "https://i.discogs.com/eEQUqur6cGQCl8krTRo2kpiB1lkuY2zZi3FfkBsgAWY/rs:fit/g:sm/q:90/h:196/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xNjYzOTQwMzIy/LTQyODYuanBlZw.jpeg",
                "uri150": "https://i.discogs.com/h4hxYBvophO45o_U4j6h8wzlw8hZAFm6AoYaOqxyzRo/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xNjYzOTQwMzIy/LTQyODYuanBlZw.jpeg",
                "width": 600,
                "height": 196
            },
            {
                "type": "secondary",
                "uri": "https://i.discogs.com/zFOB9EHXTRP2DrRtQ_OxU6YUHhkP5VgvDbnmdhZ-z3A/rs:fit/g:sm/q:90/h:507/w:338/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xNjYzOTQwMzAy/LTczNTMuanBlZw.jpeg",
                "resource_url": "https://i.discogs.com/zFOB9EHXTRP2DrRtQ_OxU6YUHhkP5VgvDbnmdhZ-z3A/rs:fit/g:sm/q:90/h:507/w:338/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xNjYzOTQwMzAy/LTczNTMuanBlZw.jpeg",
                "uri150": "https://i.discogs.com/g7Tb72DRKqVDOhX1V8tK6M05I8OxFZgWDfXNbclpirg/rs:fit/g:sm/q:40/h:150/w:150/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTU0NzY4/OS0xNjYzOTQwMzAy/LTczNTMuanBlZw.jpeg",
                "width": 338,
                "height": 507
            }
        ],
        "artist": "a-ha",
        "format": "LP, Album, SRC",
        "resource_url": "https://api.discogs.com/releases/547689",
        "title": "Hunting High And Low",
        "year": 1985,
        "id": 547689,
        "catalog_number": "1-25300, 9 25300-1",
        "stats": {
            "community": {
                "in_wantlist": 426,
                "in_collection": 3067
            },
            "user": {
                "in_wantlist": 1,
                "in_collection": 0
            }
        }
    },
    "in_cart": false
}
*/
