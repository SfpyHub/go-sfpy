package sfpy

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
