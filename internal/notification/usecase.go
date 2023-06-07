package notification

import (
	"Notifications/internal/models"
	"context"
)

type UseCase interface {
	CreateNotification(ctx context.Context, notification models.Notification) error

	DeleteNotification(ctx context.Context, uuid string) error

	UpdateNotifications(ctx context.Context, uuid string, notification models.UpdatedNotification) (*models.UpdatedNotification, error)

	NotificationById(ctx context.Context, uuid string) (*models.Notification, error)
	AllNotifications(ctx context.Context) ([]models.Notification, error)
	NotificationsByUser(ctx context.Context, uuid string) ([]models.Notification, error)
}
