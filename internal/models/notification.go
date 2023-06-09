package models

import (
	"github.com/google/uuid"
	"time"
)

type Notification struct {
	Id        uuid.UUID   `json:"id"`
	Title     string      `json:"title" binding:"required"`
	Body      string      `json:"body" binding:"required"`
	Telegram  *string     `json:"telegram,omitempty"`
	Email     *string     `json:"email,omitempty"`
	Execution time.Time   `json:"execution" binding:"required"`
	AssignTo  []uuid.UUID `json:"assign_to,omitempty"`
}

type UpdatedNotification struct {
	Id        uuid.UUID `json:"id,omitempty"`
	Title     *string   `json:"title,omitempty"`
	Body      *string   `json:"body,omitempty"`
	Telegram  *string   `json:"telegram,omitempty"`
	Email     *string   `json:"email,omitempty"`
	Execution time.Time `json:"execution,omitempty"`
}
