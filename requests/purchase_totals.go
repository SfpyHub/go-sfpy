package requests

type PurchaseTotal struct {
	SubTotal *PriceMoney `json:"sub_total"`
	Discount *PriceMoney `json:"discount"`
	TaxTotal *PriceMoney `json:"tax_total"`
}
