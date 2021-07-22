package responses

type Metadata struct {
	Limit  int   `json:"limit,omitempty"`
	Offset int   `json:"offset"`
	Count  int64 `json:"count,omitempty"`
}
