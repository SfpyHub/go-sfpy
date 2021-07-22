package requests

type Payment struct {
	Request      string `json:"request"`
	ChainId      uint   `json:"chain_id"`
	IsETH        *bool  `json:"is_eth"`
	TxnHash      string `json:"txn_hash"`
	From         string `json:"from"`
	Amount       string `json:"amount"`
	TokenAddress string `json:"token_address"`
	Rate         string `json:"rate"`
}
