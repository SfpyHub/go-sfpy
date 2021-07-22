package responses

import "time"

type Payment struct {
	Token        string    `json:"payment_id"`
	RequestID    string    `json:"request_id"`
	ChainID      uint      `json:"chain_id"`
	State        string    `json:"state"`
	IsETH        bool      `json:"is_eth"`
	TxnHash      string    `json:"txn_hash"`
	From         string    `json:"from"`
	Amount       string    `json:"amount"`
	TokenAddress string    `json:"token_address"`
	Rate         string    `json:"rate"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type PaymentSlice []*Payment
