package responses

type WrappedMerchant struct {
	ApiKey   *ApiKeys  `json:"api_key"`
	Merchant *Merchant `json:"merchant"`
}
