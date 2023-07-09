package models

import "github.com/google/uuid"

type Account struct {
	Id       *uuid.UUID `json:"id,omitempty"`
	Telegram *string    `json:"telegram,omitempty"`
	Email    *string    `json:"email,omitempty"`
}
