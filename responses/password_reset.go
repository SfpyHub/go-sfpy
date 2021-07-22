package responses

import "time"

type PasswordReset struct {
	Token              string    `json:"token"`
	ClientID           string    `json:"client_id"`
	Email              string    `json:"email"`
	ResetToken         string    `json:"-"`
	ResetTokenUsed     bool      `json:"-"`
	ResetTokenCreated  time.Time `json:"-"`
	ResetTokenValidity time.Time `json:"-"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type PasswordResetSlice []*PasswordReset
