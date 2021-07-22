package responses

import "time"

type Auth struct {
	ClientID  string    `json:"client_id"`
	Token     string    `json:"token"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
