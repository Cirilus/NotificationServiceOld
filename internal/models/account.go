package models

type Account struct {
	UUID     string  `json:"uuid,omitempty"`
	Telegram *string `json:"telegram,omitempty"`
	Email    *string `json:"email,omitempty"`
}
