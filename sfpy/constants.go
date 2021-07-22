package sfpy

const API_BASE_URL = "https://api.sfpy.co"
const APP_BASE_URL = "https://app.sfpy.co"

const (
	FETCH_API_KEY  = "/v1/apikey/fetch"
	UPDATE_API_KEY = "/v1/apikey/update"

	ADD_MERCHANT    = "/v1/merchant/add"
	UPDATE_MERCHANT = "/v1/merchant/update"
	FETCH_MERCHANT  = "/v1/merchant/fetch"

	FETCH_EVENTS = "/v1/notify/events"

	ENABLE_WEBHOOKS = "/v1/notify/merchant"
	EXISTS_WEBHOOKS = "/v1/notify/exists"

	FETCH_SECRET  = "/v1/notify/secret"
	ROTATE_SECRET = "/v1/notify/secret"

	ADD_ENDPOINT    = "/v1/notify/subscription"
	UPDATE_ENDPOINT = "/v1/notify/subscription"
	FETCH_ENDPOINTS = "/v1/notify/subscription"
	DELETE_ENDPOINT = "/v1/notify/subscription"

	SUBSCRIBE = "/v1/notify/subscribe"

	FETCH_NOTIFICATIONS = "/v1/notify/notification"

	TEST_NOTIFICATION = "/v1/notify/test"

	ADD_ORDER_LINK = "/v1/order/link"
	ADD_ORDER      = "/v1/order/"
	FETCH_ORDERS   = "/v1/order/"

	FETCH_PAYMENTS = "/v1/payment"
	FETCH_REFUNDS  = "/v1/refund"

	ADD_EMAIL_PASSWORD = "/v1/security/set"
)
