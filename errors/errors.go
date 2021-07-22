package errors

// Category of errors
type Category string

const (
	// APIERROR : An error occurred with the Safepay API itself.
	APIERROR Category = "API_ERROR"

	// VALIDATIONERROR : The request was invalid.
	// Most commonly, a required parameter was missing,
	// or a provided parameter had an invalid value.
	VALIDATIONERROR Category = "VALIDATION_ERROR"

	// BADREQUEST - a general error occurred.
	BADREQUEST Category = "BAD_REQUEST"

	// NOTFOUND ...
	NOTFOUND Category = "NOT_FOUND"

	// AUTHENTICATIONERROR : An authentication error occurred.
	// Most commonly, the request had a missing, malformed, or
	// otherwise invalid Authorization header.
	AUTHENTICATIONERROR Category = "AUTHENTICATION_ERROR"

	// DECODINGERROR: json decoding errors
	DECODINGERROR Category = "DECODING_ERROR"
)

// String <>
func (c Category) String() string {
	return string(c)
}

type Error struct {
	Category Category `json:"category"`
	Message  string   `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}
