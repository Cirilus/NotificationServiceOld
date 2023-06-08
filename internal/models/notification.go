package models

import "time"

type Notification struct {
	Id        string    `json:"id"`
	Title     string    `json:"title" binding:"required"`
	Body      string    `json:"body" binding:"required"`
	Telegram  *string   `json:"telegram,omitempty"`
	Email     *string   `json:"email,omitempty"`
	Execution time.Time `json:"execution" binding:"required"`
	AssignTo  []string  `json:"assign_to,omitempty"`
}

type UpdatedNotification struct {
	Id        string    `json:"id,omitempty"`
	Title     *string   `json:"title,omitempty"`
	Body      *string   `json:"body,omitempty"`
	Telegram  *string   `json:"telegram,omitempty"`
	Email     *string   `json:"email,omitempty"`
	Execution time.Time `json:"execution,omitempty"`
	AssignTo  []string  `json:"assign_to,omitempty"`
}
