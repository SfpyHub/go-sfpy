package requests

type Notification struct {
	Kind         string `json:"kind"`
	Subscription string `json:"subscription"`
}
