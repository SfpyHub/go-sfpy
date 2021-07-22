package requests

type Subscription struct {
	Kind           string   `json:"kind"`
	Token          string   `json:"token"`
	Endpoint       string   `json:"endpoint"`
	EventsToAdd    []string `json:"events_to_add"`
	EventsToRemove []string `json:"events_to_remove"`
}
