package requests

type Refund struct {
	Payment      string `json:"payment"`
	ChainId      uint   `json:"chain_id"`
	IsETH        *bool  `json:"is_eth"`
	TxnHash      string `json:"txn_hash"`
	From         string `json:"from"`
	To           string `json:"to"`
	Amount       string `json:"amount"`
	TokenAddress string `json:"token_address"`
}
