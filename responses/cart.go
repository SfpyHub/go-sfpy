package responses

type Cart struct {
	Token string `json:"token,omitempty"`

	RequestID string `json:"request_id,omitempty"`

	// Indicates the e-commerce platform or
	// personal website from where the customer
	// and order originates from. For example
	// Shopify, Wordpress, Magento, Custom.
	//
	// The value provided is trimmed and lowered
	Source string `json:"source,omitempty"`

	// This is the external order id or cart id to
	// be associated with this payment. Is useful
	// to reconcile payments with an external source
	// or CMS
	OrderReference string `json:"order_reference,omitempty"`

	// This is the URL to redirect the customer to after
	// they have completed the payment. This can be a
	// success page, order confirmation page or thank you
	// page.
	CompleteURL string `json:"complete_url,omitempty"`

	// This is the URL to redirect the customer to after
	// they cancel the payment or leave it abandoned. This
	// is usually the cart url so they can go back to your
	// site and complete the payment using a different
	// payment method
	CancelURL string `json:"cancel_url,omitempty"`
}
