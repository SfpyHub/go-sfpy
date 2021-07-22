package responses

import "time"

type ApiKeys struct {
	Token         string    `json:"token"`
	ClientID      string    `json:"client_id"`
	PrivateApiKey string    `json:"pvt_api_key"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
