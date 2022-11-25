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

// TODO: GetUserInventory is a function for getting inventory of a user
// TODO: EditListing is a function for editing a single listing
// TODO: DeleteListing is a function for deleting a single listing
// TODO: CreateListing is a function for creating a single listing

// TODO: GetOrder
// TODO: EditOrder
// TODO: ListOrders
// TODO: ListMessages
// TODO: CreateMessage
// TODO: GetFee
// TODO: GetFeeWithCurrency
// TODO: GetPriceSuggestions
// TODO: GetStatsForRelease
// TODO: ExportInventory
