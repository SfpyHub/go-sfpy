package responses

import "time"

type Refund struct {
	Token        string    `json:"refund_id"`
	RequestID    string    `json:"request_id"`
	PaymentID    string    `json:"payment_id"`
	ChainID      uint      `json:"chain_id"`
	IsETH        bool      `json:"is_eth"`
	TxnHash      string    `json:"txn_hash"`
	From         string    `json:"from"`
	To           string    `json:"to"`
	Amount       string    `json:"amount"`
	TokenAddress string    `json:"token_address"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type RefundSlice []*Refund
