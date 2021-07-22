package responses

type PurchaseTotal struct {
	Token string `json:"token,omitempty"`

	RequestID string `json:"request_id,omitempty"`

	Discount *PriceMoney `json:"discount"`

	SubTotal *PriceMoney `json:"sub_total"`

	TaxTotal *PriceMoney `json:"tax_total"`

	GrandTotal *PriceMoney `json:"grand_total"`
}
