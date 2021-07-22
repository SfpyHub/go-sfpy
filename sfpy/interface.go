package sfpy

import "context"

type SFPY interface {
	FetchApiKey(ctx context.Context, request interface{}) (interface{}, error)
	UpdateApiKey(ctx context.Context, request interface{}) (interface{}, error)
	AddMerchant(ctx context.Context, request interface{}) (interface{}, error)
	UpdateMerchant(ctx context.Context, request interface{}) (interface{}, error)
	FetchMerchant(ctx context.Context, request interface{}) (interface{}, error)
	FetchEvents(ctx context.Context, request interface{}) (interface{}, error)
	WebhooksExist(ctx context.Context, request interface{}) (interface{}, error)
	EnableWebhooks(ctx context.Context, request interface{}) (interface{}, error)
	FetchSecret(ctx context.Context, request interface{}) (interface{}, error)
	UpdateSecret(ctx context.Context, request interface{}) (interface{}, error)
	DeleteEndpoint(ctx context.Context, request interface{}) (interface{}, error)
	UpdateEndpoint(ctx context.Context, request interface{}) (interface{}, error)
	AddEndpoint(ctx context.Context, request interface{}) (interface{}, error)
	FetchEndpoints(ctx context.Context, request interface{}) (interface{}, error)
	Subscribe(ctx context.Context, request interface{}) (interface{}, error)
	TestNotification(ctx context.Context, request interface{}) (interface{}, error)
	FetchNotifications(ctx context.Context, request interface{}) (interface{}, error)
	AddOrderLink(ctx context.Context, request interface{}) (interface{}, error)
	AddOrder(ctx context.Context, request interface{}) (interface{}, error)
	FetchOrders(ctx context.Context, request interface{}) (interface{}, error)
	FetchPayments(ctx context.Context, request interface{}) (interface{}, error)
	FetchRefunds(ctx context.Context, request interface{}) (interface{}, error)
	AddEmailPassword(ctx context.Context, request interface{}) (interface{}, error)
}
