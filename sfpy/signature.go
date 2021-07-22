package sfpy

import (
	"crypto/hmac"
	"encoding/hex"
	"fmt"
	"hash"

	"github.com/sfpyhub/go-sfpy/errors"
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
