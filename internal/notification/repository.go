package notification

import (
	"Notifications/internal/models"
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	CreateNotification(ctx context.Context, notification models.Notification) error
	AddNotificationToAccount(ctx context.Context, notificationId uuid.UUID, accountId []uuid.UUID) error
	DeleteNotificationFromAccount(ctx context.Context, notificationId uuid.UUID, accountId uuid.UUID) error

	DeleteNotification(ctx context.Context, uuid string) error

	UpdateNotifications(ctx context.Context, uuid string, notification models.UpdatedNotification) (*models.UpdatedNotification, error)

	GetNotificationById(ctx context.Context, uuid string) (*models.Notification, error)
	GetAllNotifications(ctx context.Context) ([]models.Notification, error)
	GetNotificationsByUser(ctx context.Context, uuid string) ([]models.Notification, error)
}
