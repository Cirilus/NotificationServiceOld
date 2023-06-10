package models

import "github.com/google/uuid"

type Account struct {
	Id       uuid.UUID `json:"id" binding:"required"`
	Telegram *string   `json:"telegram,omitempty"`
	Email    *string   `json:"email,omitempty"`
}

type UpdateAccount struct {
	Id       *uuid.UUID `json:"id"`
	Telegram *string    `json:"telegram,omitempty"`
	Email    *string    `json:"email,omitempty"`
}
