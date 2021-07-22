package responses

// Represents an amount of money.

// Money fields can be signed or unsigned. Fields that do not
// explicitly define whether they are signed or unsigned
// are considered unsigned and can only hold positive amounts.
// For signed fields, the sign of the value indicates the purpose
// of the money transfer. See Working with
// Monetary Amounts for more information.
type PriceMoney struct {

	// The amount of money, in the smallest denomination of the currency
	// indicated by currency. For example, when currency is USD,
	// amount is in cents. Monetary amounts can be positive or negative.
	// See the specific field description to determine the meaning
	// of the sign in a particular case.
	Amount int64 `json:"amount"`

	// The type of currency, in ISO 4217 format.
	// For example, the currency code for US dollars is USD.
	Currency string `json:"currency"`
}

type PriceMoneySlice []*PriceMoney
