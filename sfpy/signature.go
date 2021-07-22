package sfpy

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash"

	"github.com/sfpyhub/go-sfpy/errors"
	"github.com/sfpyhub/go-sfpy/responses"
)

// genMAC generates the HMAC signature for a message provided the secret key
// and hashFunc.
func genMAC(message, key []byte, hashFunc func() hash.Hash) []byte {
	mac := hmac.New(hashFunc, key)
	mac.Write(message)
	return mac.Sum(nil)
}

// checkMAC reports whether messageMAC is a valid HMAC tag for message.
func checkMAC(message, messageMAC, key []byte, hashFunc func() hash.Hash) bool {
	expectedMAC := genMAC(message, key, hashFunc)
	return hmac.Equal(messageMAC, expectedMAC)
}

// messageMAC returns the hex-decoded HMAC tag from the signature and its
// corresponding hash function.
func messageMAC(signature string) ([]byte, error) {
	if signature == "" {
		return nil, &errors.Error{
			Category: errors.AUTHENTICATIONERROR,
			Message:  "missing signature",
		}
	}

	buf, err := hex.DecodeString(signature)
	if err != nil {
		return nil, &errors.Error{
			Category: errors.AUTHENTICATIONERROR,
			Message:  fmt.Sprintf("error decoding signature %q: %v", signature, err),
		}
	}
	return buf, nil
}

func ValidateSignature(signature string, event *responses.Event, secretToken []byte) error {
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
	if !checkMAC(payload, messageMAC, secretToken, sha512.New) {
		return &errors.Error{
			Category: errors.AUTHENTICATIONERROR,
			Message:  "payload signature check failed",
		}
	}
	return nil
}
