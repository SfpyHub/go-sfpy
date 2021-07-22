package sfpy

import (
	"context"

	"github.com/sfpyhub/go-sfpy/responses"
)

type SFPY interface {
	FetchApiKey(ctx context.Context, request interface{}) (*responses.Response, error)
	UpdateApiKey(ctx context.Context, request interface{}) (*responses.Response, error)
	AddMerchant(ctx context.Context, request interface{}) (*responses.Response, error)
	UpdateMerchant(ctx context.Context, request interface{}) (*responses.Response, error)
	FetchMerchant(ctx context.Context, request interface{}) (*responses.Response, error)
	FetchEvents(ctx context.Context, request interface{}) (*responses.Response, error)
	WebhooksExist(ctx context.Context, request interface{}) (*responses.Response, error)
	EnableWebhooks(ctx context.Context, request interface{}) (*responses.Response, error)
	FetchSecret(ctx context.Context, request interface{}) (*responses.Response, error)
	UpdateSecret(ctx context.Context, request interface{}) (*responses.Response, error)
	DeleteEndpoint(ctx context.Context, request interface{}) (*responses.Response, error)
	UpdateEndpoint(ctx context.Context, request interface{}) (*responses.Response, error)
	AddEndpoint(ctx context.Context, request interface{}) (*responses.Response, error)
	FetchEndpoints(ctx context.Context, request interface{}) (*responses.Response, error)
	Subscribe(ctx context.Context, request interface{}) (*responses.Response, error)
	TestNotification(ctx context.Context, request interface{}) (*responses.Response, error)
	FetchNotifications(ctx context.Context, request interface{}) (*responses.Response, error)
	AddOrderLink(ctx context.Context, request interface{}) (*responses.Response, error)
	AddOrder(ctx context.Context, request interface{}) (*responses.Response, error)
	FetchOrders(ctx context.Context, request interface{}) (*responses.Response, error)
	FetchPayments(ctx context.Context, request interface{}) (*responses.Response, error)
	FetchRefunds(ctx context.Context, request interface{}) (*responses.Response, error)
	AddEmailPassword(ctx context.Context, request interface{}) (*responses.Response, error)
}
