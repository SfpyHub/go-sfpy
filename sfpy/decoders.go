package sfpy

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sfpyhub/go-sfpy/responses"
)

func decodeResponse(c context.Context, r *http.Response) (interface{}, error) {
	var response responses.Response
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	response.HTTPStatus = r.StatusCode
	return response, nil
}
