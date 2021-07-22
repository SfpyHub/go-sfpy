package responses

import (
	"time"
)

// Merchant internal definition
type Merchant struct {
	ClientID         string    `json:"client_id"`
	RegisteredName   string    `json:"registered_name"`
	WebsiteURL       string    `json:"website_url"`
	TwitterURL       string    `json:"twitter_url"`
	InstagramURL     string    `json:"instagram_url"`
	ProfileImgURL    string    `json:"profile_img_url"`
	BackgroundImgURL string    `json:"background_img_url"`
	ApiKeys          *ApiKeys  `json:"-"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// MerchantSlice is a convenient collection
type MerchantSlice []*Merchant
