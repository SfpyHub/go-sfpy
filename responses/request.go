package responses

import (
	"time"

	"api/types"
)

type Request struct {
	Token         string         `json:"token"`
	ClientID      string         `json:"merchant"`
	Address       string         `json:"address"`
	Reference     string         `json:"reference"`
	ChainID       types.ChainID  `json:"chain_id"`
	State         State          `json:"state"`
	Cart          *Cart          `json:"cart,omitempty"`
	PurchaseTotal *PurchaseTotal `json:"purchase_total,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

type RequestSlice []*Request
