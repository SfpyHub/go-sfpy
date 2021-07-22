package requests

type Fetch struct {
	Token     string `form:"token"`
	Request   string `form:"request"`
	Payment   string `form:"payment"`
	Kind      string `form:"kind"`
	Chain     uint   `form:"chain"`
	SortBy    string `form:"sort_by"`
	SortOrder string `form:"sort_order"`
	Limit     int    `form:"limit"`
	Offset    int    `form:"offset"`
}
