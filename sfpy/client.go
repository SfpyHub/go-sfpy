package sfpy

import (
	"crypto/sha512"
	"encoding/json"

	"github.com/sfpyhub/go-sfpy/errors"
	"github.com/sfpyhub/go-sfpy/responses"
)

type Client struct {
	sharedSecret string
	endpoints    *endpoints
}

func NewClient(apikey, sharedsecret string) *Client {
	return &Client{
		sharedSecret: sharedsecret,
		endpoints:    newEndpoints(apikey, API_BASE_URL),
	}
}

func (c *Client) ValidateSignature(signature string, event *responses.Event) error {
	payload, err := json.Marshal(event)
	if err != nil {
		return &errors.Error{
			Category: errors.DECODINGERROR,
			Message:  err.Error(),
		}
	}

	messageMAC, err := messageMAC(signature)
	if err != nil {
		return err
	}
	if !checkMAC(payload, messageMAC, []byte(c.sharedSecret), sha512.New) {
		return &errors.Error{
			Category: errors.AUTHENTICATIONERROR,
			Message:  "payload signature check failed",
		}
	}
	return nil
}
