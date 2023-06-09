package usecase

import (
	"Notifications/internal/models"
	"Notifications/internal/notification"
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type NotificationUseCase struct {
	repo notification.Repository
}

func NewUseCase(repo notification.Repository) *NotificationUseCase {
	return &NotificationUseCase{repo: repo}
}

func (n NotificationUseCase) CreateNotification(ctx context.Context, notification models.Notification) error {
	err := n.repo.CreateNotification(ctx, notification)
	if err != nil {
		logrus.Error("Notification - UseCase - CreateNotification")
		return err
	}
	return nil
}

func (n NotificationUseCase) DeleteNotification(ctx context.Context, id string) error {
	err := n.repo.DeleteNotification(ctx, id)
	if err != nil {
		logrus.Error("Notification - UseCase - DeleteNotification")
		return err
	}
	return nil
}

func (n NotificationUseCase) UpdateNotifications(ctx context.Context, id string, notification models.UpdatedNotification) (*models.UpdatedNotification, error) {
	updateNotifications, err := n.repo.UpdateNotifications(ctx, id, notification)
	if err != nil {
		logrus.Error("Notification - UseCase - UpdateNotifications")
		return nil, err
	}
	return updateNotifications, nil
}

func (n NotificationUseCase) NotificationById(ctx context.Context, id string) (*models.Notification, error) {
	notificationById, err := n.repo.GetNotificationById(ctx, id)
	if err != nil {
		logrus.Error("Notification - UseCase - NotificationById")
		return nil, err
	}
	return notificationById, nil
}

func (n NotificationUseCase) AllNotifications(ctx context.Context) ([]models.Notification, error) {
	allNotifications, err := n.repo.GetAllNotifications(ctx)
	if err != nil {
		logrus.Error("Notification - UseCase - AllNotifications")
		return nil, err
	}
	return allNotifications, nil
}

func (n NotificationUseCase) NotificationsByUser(ctx context.Context, id string) ([]models.Notification, error) {
	notificationByUser, err := n.repo.GetNotificationsByUser(ctx, id)
	if err != nil {
		logrus.Error("Notification - UseCase - NotificationById")
		return nil, err
	}
	return notificationByUser, nil
}

func (n NotificationUseCase) AddNotificationToAccount(ctx context.Context, notificationId uuid.UUID, accountId []uuid.UUID) error {
	err := n.repo.AddNotificationToAccount(ctx, notificationId, accountId)
	if err != nil {
		logrus.Error("Notification - UseCase - AddNotificationToAccount")
		return err
	}
	return nil
}
func (n NotificationUseCase) DeleteNotificationFromAccount(ctx context.Context, notificationId uuid.UUID, accountId uuid.UUID) error {
	err := n.repo.DeleteNotificationFromAccount(ctx, notificationId, accountId)
	if err != nil {
		logrus.Error("Notification - UseCase - DeleteNotificationFromAccount")
		return err
	}
	return nil
}
