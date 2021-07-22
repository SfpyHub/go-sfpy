package sfpy

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/sfpyhub/go-sfpy/errors"
	"github.com/sfpyhub/go-sfpy/responses"
)

type endpoints struct {
	fetchapikey        endpoint.Endpoint
	updateapikey       endpoint.Endpoint
	addmerchant        endpoint.Endpoint
	updatemerchant     endpoint.Endpoint
	fetchmerchant      endpoint.Endpoint
	fetchevents        endpoint.Endpoint
	enablewebhooks     endpoint.Endpoint
	existswebhooks     endpoint.Endpoint
	fetchsecret        endpoint.Endpoint
	updatesecret       endpoint.Endpoint
	addendpoint        endpoint.Endpoint
	updateendpoint     endpoint.Endpoint
	deleteendpoint     endpoint.Endpoint
	fetchendpoints     endpoint.Endpoint
	subscribe          endpoint.Endpoint
	fetchnotifications endpoint.Endpoint
	testnotification   endpoint.Endpoint
	addorderlink       endpoint.Endpoint
	addorder           endpoint.Endpoint
	fetchorders        endpoint.Endpoint
	fetchpayments      endpoint.Endpoint
	fetchrefunds       endpoint.Endpoint
	addemailpswd       endpoint.Endpoint
}

func newEndpoints(apikey, url string) *endpoints {

	return &endpoints{
		fetchapikey:        makeFetchApiKey(apikey, url),
		updateapikey:       makeUpdateApiKey(apikey, url),
		addmerchant:        makeAddMerchant(apikey, url),
		updatemerchant:     makeUpdateMerchant(apikey, url),
		fetchmerchant:      makeFetchMerchant(apikey, url),
		fetchevents:        makeFetchEvents(apikey, url),
		enablewebhooks:     makeEnableWebhooks(apikey, url),
		existswebhooks:     makeExistsWebhooks(apikey, url),
		fetchsecret:        makeFetchSecret(apikey, url),
		updatesecret:       makeRotateSecret(apikey, url),
		addendpoint:        makeAddEndpoint(apikey, url),
		updateendpoint:     makeUpdateEndpoint(apikey, url),
		deleteendpoint:     makeDeleteEndpoint(apikey, url),
		fetchendpoints:     makeFetchEndpoints(apikey, url),
		subscribe:          makeSubscribe(apikey, url),
		fetchnotifications: makeFetchNotifications(apikey, url),
		testnotification:   makeTestNotification(apikey, url),
		addorderlink:       makeAddOrderLink(apikey, url),
		addorder:           makeAddOrder(apikey, url),
		fetchorders:        makeFetchOrders(apikey, url),
		fetchpayments:      makeFetchPayments(apikey, url),
		fetchrefunds:       makeFetchRefunds(apikey, url),
		addemailpswd:       makeAddEmailPassword(apikey, url),
	}
}

func (e *endpoints) FetchApiKey(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := e.fetchapikey(ctx, request)
	if err != nil {
		return nil, err
	}

	return e.parseResponse(res)
}

func (e *endpoints) UpdateApiKey(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := e.updateapikey(ctx, request)
	if err != nil {
		return nil, err
	}

	return e.parseResponse(res)
}

func (e *endpoints) AddMerchant(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := e.addmerchant(ctx, request)
	if err != nil {
		return nil, err
	}

	return e.parseResponse(res)
}

func (e *endpoints) UpdateMerchant(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := e.updatemerchant(ctx, request)
	if err != nil {
		return nil, err
	}

	return e.parseResponse(res)
}

func (e *endpoints) FetchMerchant(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := e.fetchmerchant(ctx, request)
	if err != nil {
		return nil, err
	}

	return e.parseResponse(res)
}

func (e *endpoints) FetchEvents(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := e.fetchevents(ctx, request)
	if err != nil {
		return nil, err
	}

	return e.parseResponse(res)
}

func (e *endpoints) WebhooksExist(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := e.existswebhooks(ctx, request)
	if err != nil {
		return nil, err
	}

	return e.parseResponse(res)
}

func (e *endpoints) EnableWebhooks(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := e.enablewebhooks(ctx, request)
	if err != nil {
		return nil, err
	}

	return e.parseResponse(res)
}

func (e *endpoints) FetchSecret(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := e.fetchsecret(ctx, request)
	if err != nil {
		return nil, err
	}

	return e.parseResponse(res)
}

func (e *endpoints) UpdateSecret(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := e.updatesecret(ctx, request)
	if err != nil {
		return nil, err
	}

	return e.parseResponse(res)
}

func (e *endpoints) DeleteEndpoint(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := e.deleteendpoint(ctx, request)
	if err != nil {
		return nil, err
	}

	return e.parseResponse(res)
}

func (e *endpoints) UpdateEndpoint(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := e.updateendpoint(ctx, request)
	if err != nil {
		return nil, err
	}

	return e.parseResponse(res)
}

func (e *endpoints) AddEndpoint(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := e.addendpoint(ctx, request)
	if err != nil {
		return nil, err
	}

	return e.parseResponse(res)
}

func (e *endpoints) FetchEndpoints(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := e.fetchendpoints(ctx, request)
	if err != nil {
		return nil, err
	}

	return e.parseResponse(res)
}

func (e *endpoints) Subscribe(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := e.subscribe(ctx, request)
	if err != nil {
		return nil, err
	}

	return e.parseResponse(res)
}

func (e *endpoints) TestNotification(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := e.testnotification(ctx, request)
	if err != nil {
		return nil, err
	}

	return e.parseResponse(res)
}

func (e *endpoints) FetchNotifications(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := e.fetchnotifications(ctx, request)
	if err != nil {
		return nil, err
	}

	return e.parseResponse(res)
}

func (e *endpoints) AddOrderLink(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := e.addorderlink(ctx, request)
	if err != nil {
		return nil, err
	}

	return e.parseResponse(res)
}

func (e *endpoints) AddOrder(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := e.addorder(ctx, request)
	if err != nil {
		return nil, err
	}

	return e.parseResponse(res)
}

func (e *endpoints) FetchOrders(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := e.fetchorders(ctx, request)
	if err != nil {
		return nil, err
	}

	return e.parseResponse(res)
}

func (e *endpoints) FetchPayments(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := e.fetchpayments(ctx, request)
	if err != nil {
		return nil, err
	}

	return e.parseResponse(res)
}

func (e *endpoints) FetchRefunds(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := e.fetchrefunds(ctx, request)
	if err != nil {
		return nil, err
	}

	return e.parseResponse(res)
}

func (e *endpoints) AddEmailPassword(ctx context.Context, request interface{}) (interface{}, error) {
	res, err := e.addemailpswd(ctx, request)
	if err != nil {
		return nil, err
	}

	return e.parseResponse(res)
}

func (e *endpoints) parseResponse(res interface{}) (interface{}, error) {
	response, ok := res.(responses.Response)
	if !ok {
		return nil, &errors.Error{
			Category: errors.APIERROR,
			Message:  "Unexpected response received from service",
		}
	}

	if response.DoesError() {
		return nil, response.GetErrors()
	}

	return response, nil
}
