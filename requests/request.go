package requests

type Request struct {
	MerchantService *MerchantService `json:"merchant_service,omitempty" form:"-"`
	OrderService    *OrderService    `json:"order_service,omitempty" form:"-"`
	PaymentService  *PaymentService  `json:"payment_service,omitempty" form:"-"`
	RefundService   *RefundService   `json:"refund_service,omitempty" form:"-"`
	SecurityService *SecurityService `json:"security_service,omitempty" form:"-"`
	FetchService    *FetchService    `json:"-"`

	SubscriptionService *SubscriptionService `json:"subscription_service,omitempty" form:"-"`
	NotificationService *NotificationService `json:"notification_service,omitempty" form:"-"`
}
