package sfpy

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func makeAddEmailPassword(apikey, instance string) endpoint.Endpoint {
	u := formatURL(instance, ADD_EMAIL_PASSWORD)

	return httptransport.NewClient(
		"POST",
		u,
		encodePostRequest,
		decodeResponse,
		httptransport.ClientBefore(
			setApiKey(apikey),
		),
	).Endpoint()
}

func makeFetchRefunds(apikey, instance string) endpoint.Endpoint {
	u := formatURL(instance, FETCH_REFUNDS)

	return httptransport.NewClient(
		"GET",
		u,
		encodeGetRequest,
		decodeResponse,
		httptransport.ClientBefore(
			setApiKey(apikey),
		),
	).Endpoint()
}

func makeFetchPayments(apikey, instance string) endpoint.Endpoint {
	u := formatURL(instance, FETCH_PAYMENTS)

	return httptransport.NewClient(
		"GET",
		u,
		encodeGetRequest,
		decodeResponse,
		httptransport.ClientBefore(
			setApiKey(apikey),
		),
	).Endpoint()
}

func makeFetchOrders(apikey, instance string) endpoint.Endpoint {
	u := formatURL(instance, FETCH_ORDERS)

	return httptransport.NewClient(
		"GET",
		u,
		encodeGetRequest,
		decodeResponse,
		httptransport.ClientBefore(
			setApiKey(apikey),
		),
	).Endpoint()
}

func makeAddOrder(apikey, instance string) endpoint.Endpoint {
	u := formatURL(instance, ADD_ORDER)

	return httptransport.NewClient(
		"POST",
		u,
		encodePostRequest,
		decodeResponse,
		httptransport.ClientBefore(
			setApiKey(apikey),
		),
	).Endpoint()
}

func makeAddOrderLink(apikey, instance string) endpoint.Endpoint {
	u := formatURL(instance, ADD_ORDER_LINK)

	return httptransport.NewClient(
		"POST",
		u,
		encodePostRequest,
		decodeResponse,
		httptransport.ClientBefore(
			setApiKey(apikey),
		),
	).Endpoint()
}

func makeFetchNotifications(apikey, instance string) endpoint.Endpoint {
	u := formatURL(instance, FETCH_NOTIFICATIONS)

	return httptransport.NewClient(
		"GET",
		u,
		encodeGetRequest,
		decodeResponse,
		httptransport.ClientBefore(
			setApiKey(apikey),
		),
	).Endpoint()
}

func makeTestNotification(apikey, instance string) endpoint.Endpoint {
	u := formatURL(instance, TEST_NOTIFICATION)

	return httptransport.NewClient(
		"POST",
		u,
		encodePostRequest,
		decodeResponse,
		httptransport.ClientBefore(
			setApiKey(apikey),
		),
	).Endpoint()
}

func makeSubscribe(apikey, instance string) endpoint.Endpoint {
	u := formatURL(instance, SUBSCRIBE)

	return httptransport.NewClient(
		"POST",
		u,
		encodePostRequest,
		decodeResponse,
		httptransport.ClientBefore(
			setApiKey(apikey),
		),
	).Endpoint()
}

func makeFetchEndpoints(apikey, instance string) endpoint.Endpoint {
	u := formatURL(instance, FETCH_ENDPOINTS)

	return httptransport.NewClient(
		"GET",
		u,
		encodeGetRequest,
		decodeResponse,
		httptransport.ClientBefore(
			setApiKey(apikey),
		),
	).Endpoint()
}

func makeAddEndpoint(apikey, instance string) endpoint.Endpoint {
	u := formatURL(instance, ADD_ENDPOINT)

	return httptransport.NewClient(
		"POST",
		u,
		encodePostRequest,
		decodeResponse,
		httptransport.ClientBefore(
			setApiKey(apikey),
		),
	).Endpoint()
}

func makeUpdateEndpoint(apikey, instance string) endpoint.Endpoint {
	u := formatURL(instance, UPDATE_ENDPOINT)

	return httptransport.NewClient(
		"PUT",
		u,
		encodePostRequest,
		decodeResponse,
		httptransport.ClientBefore(
			setApiKey(apikey),
		),
	).Endpoint()
}

func makeDeleteEndpoint(apikey, instance string) endpoint.Endpoint {
	u := formatURL(instance, DELETE_ENDPOINT)

	return httptransport.NewClient(
		"DELETE",
		u,
		encodePostRequest,
		decodeResponse,
		httptransport.ClientBefore(
			setApiKey(apikey),
		),
	).Endpoint()
}

func makeRotateSecret(apikey, instance string) endpoint.Endpoint {
	u := formatURL(instance, ROTATE_SECRET)

	return httptransport.NewClient(
		"PUT",
		u,
		encodePostRequest,
		decodeResponse,
		httptransport.ClientBefore(
			setApiKey(apikey),
		),
	).Endpoint()
}

func makeFetchSecret(apikey, instance string) endpoint.Endpoint {
	u := formatURL(instance, FETCH_SECRET)

	return httptransport.NewClient(
		"GET",
		u,
		encodeGetRequest,
		decodeResponse,
		httptransport.ClientBefore(
			setApiKey(apikey),
		),
	).Endpoint()
}

func makeEnableWebhooks(apikey, instance string) endpoint.Endpoint {
	u := formatURL(instance, ENABLE_WEBHOOKS)

	return httptransport.NewClient(
		"POST",
		u,
		encodePostRequest,
		decodeResponse,
		httptransport.ClientBefore(
			setApiKey(apikey),
		),
	).Endpoint()
}

func makeExistsWebhooks(apikey, instance string) endpoint.Endpoint {
	u := formatURL(instance, EXISTS_WEBHOOKS)

	return httptransport.NewClient(
		"POST",
		u,
		encodePostRequest,
		decodeResponse,
		httptransport.ClientBefore(
			setApiKey(apikey),
		),
	).Endpoint()
}

func makeFetchEvents(apikey, instance string) endpoint.Endpoint {
	u := formatURL(instance, FETCH_EVENTS)

	return httptransport.NewClient(
		"GET",
		u,
		encodeGetRequest,
		decodeResponse,
		httptransport.ClientBefore(
			setApiKey(apikey),
		),
	).Endpoint()
}

func makeFetchMerchant(apikey, instance string) endpoint.Endpoint {
	u := formatURL(instance, FETCH_MERCHANT)

	return httptransport.NewClient(
		"GET",
		u,
		encodePostRequest,
		decodeResponse,
		httptransport.ClientBefore(
			setApiKey(apikey),
		),
	).Endpoint()
}

func makeUpdateMerchant(apikey, instance string) endpoint.Endpoint {
	u := formatURL(instance, UPDATE_MERCHANT)

	return httptransport.NewClient(
		"PUT",
		u,
		encodePostRequest,
		decodeResponse,
		httptransport.ClientBefore(
			setApiKey(apikey),
		),
	).Endpoint()
}

func makeAddMerchant(apikey, instance string) endpoint.Endpoint {
	u := formatURL(instance, ADD_MERCHANT)

	return httptransport.NewClient(
		"POST",
		u,
		encodePostRequest,
		decodeResponse,
		httptransport.ClientBefore(
			setApiKey(apikey),
		),
	).Endpoint()
}

func makeUpdateApiKey(apikey, instance string) endpoint.Endpoint {
	u := formatURL(instance, UPDATE_API_KEY)

	return httptransport.NewClient(
		"PUT",
		u,
		encodePostRequest,
		decodeResponse,
		httptransport.ClientBefore(
			setApiKey(apikey),
		),
	).Endpoint()
}

func makeFetchApiKey(apikey, instance string) endpoint.Endpoint {
	u := formatURL(instance, FETCH_API_KEY)

	return httptransport.NewClient(
		"GET",
		u,
		encodeGetRequest,
		decodeResponse,
		httptransport.ClientBefore(
			setApiKey(apikey),
		),
	).Endpoint()
}

func setApiKey(apikey string) httptransport.RequestFunc {
	return func(ctx context.Context, r *http.Request) context.Context {
		r.Header.Set("X-SFPY-API-KEY", apikey)
		return ctx
	}
}

func formatURL(instance, path string) *url.URL {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		panic(err)
	}
	if u.Path == "" {
		u.Path = path
	}
	return u
}
