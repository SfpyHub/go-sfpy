package requests

type Order struct {
	Token          string         `json:"-" form:"request"`
	Address        string         `json:"address"`
	Reference      string         `json:"reference"`
	ChainID        uint           `json:"chain_id"`
	Cart           *Cart          `json:"cart"`
	PurchaseTotals *PurchaseTotal `json:"purchase_total"`
}
