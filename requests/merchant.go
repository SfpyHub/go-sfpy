package requests

// Merchant in the Sfpy ecosystem
type Merchant struct {
	RegisteredName   string `json:"registered_name"`
	WebsiteURL       string `json:"website_url,omitempty"`
	TwitterURL       string `json:"twitter_url,omitempty"`
	InstagramURL     string `json:"instagram_url,omitempty"`
	ProfileImgURL    string `json:"profile_img_url,omitempty"`
	BackgroundImgURL string `json:"background_img_url,omitempty"`
}

// MerchantSlice is a convenient collection
type MerchantSlice []*Merchant
