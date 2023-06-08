package models

type Account struct {
	Id       string  `json:"id" binding:"required"`
	Telegram *string `json:"telegram,omitempty"`
	Email    *string `json:"email,omitempty"`
}
