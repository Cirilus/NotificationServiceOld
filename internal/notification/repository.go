package notification

import (
	"Notifications/internal/models"
	"context"
)

type Repository interface {
	CreateNotification(ctx context.Context, notification models.Notification) error

	DeleteNotification(ctx context.Context, uuid string) error

	UpdateNotifications(ctx context.Context, uuid string, notification models.UpdatedNotification) (*models.UpdatedNotification, error)

	GetNotificationById(ctx context.Context, uuid string) (*models.Notification, error)
	GetAllNotifications(ctx context.Context) ([]models.Notification, error)
	GetNotificationsByUser(ctx context.Context, uuid string) ([]models.Notification, error)
}
