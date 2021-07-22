package responses

type WrappedRequest struct {
	Merchant *Merchant `json:"merchant"`
	Request  *Request  `json:"request"`
}

type WrappedRequestSlice []*WrappedRequest
