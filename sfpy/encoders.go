package sfpy

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-playground/form"
	"github.com/sfpyhub/go-sfpy/requests"
)

func encodePostRequest(c context.Context, r *http.Request, request interface{}) error {
	setHeaders(r, request)

	if req, ok := request.(*requests.Request); ok {
		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(req); err != nil {
			return err
		}
		r.Body = ioutil.NopCloser(&buf)
		r.ContentLength = int64(buf.Len())
	}

	return nil
}

func encodeGetRequest(c context.Context, r *http.Request, request interface{}) error {
	setHeaders(r, request)

	if req, ok := request.(*requests.Request); ok {
		if req.FetchService != nil {
			encoder := form.NewEncoder()
			values, err := encoder.Encode(req.FetchService.Fetch)
			if err != nil {
				return err
			}
			r.URL.RawQuery = values.Encode()
		}
	}

	return nil
}

func setHeaders(r *http.Request, request interface{}) {
	r.Header.Set("Content-Type", "application/json; charset=utf-8")

	if headerer, ok := request.(httptransport.Headerer); ok {
		headers := headerer.Headers()
		for k := range headers {
			r.Header.Set(k, headers.Get(k))
		}
	}
}
