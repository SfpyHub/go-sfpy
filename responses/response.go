package responses

import "github.com/sfpyhub/go-sfpy/errors"

type Response struct {
	HTTPStatus int           `json:"-"`
	Data       interface{}   `json:"data,omitempty"`
	Metadata   *Metadata     `json:"metadata,omitempty"`
	Error      *errors.Error `json:"error,omitempty"`
}

// GetErrors returns the error
func (r Response) GetErrors() error {
	return r.Error
}

// DoesError returns whether we do not have an error
func (r Response) DoesError() bool {
	return r.Error != nil
}
