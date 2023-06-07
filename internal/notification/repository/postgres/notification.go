package postgres

import (
	"Notifications/internal/models"
	"Notifications/pkg/client/postgresql"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type NotificationRepository struct {
	*postgresql.Postgres
}

func NewRepository(db *postgresql.Postgres) *NotificationRepository {
	return &NotificationRepository{db}
}

func (n NotificationRepository) CreateNotification(ctx context.Context, notification models.Notification) error {
	id := uuid.New()
	sql, args, err := n.Builder.
		Insert("notification").
		Columns("id, title, body, telegram, email, execution, assignTo").
		Values(id, notification.Title, notification.Body, notification.Telegram,
			notification.Email, notification.Execution, notification.AssignTo).
		ToSql()
	if err != nil {
		logrus.Error("Notification - Repository - CreateNotification")
		return err
	}
	_, err = n.Pool.Exec(ctx, sql, args...)
	if err != nil {
		logrus.Error("Notification - Repository - CreateNotification")
		return err
	}
	return nil
}

func (n NotificationRepository) DeleteNotification(ctx context.Context, id string) error {
	sql, args, err := n.Builder.
		Delete("notification").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		logrus.Error("Notification - Repository - DeleteNotification")
		return err
	}
	_, err = n.Pool.Exec(ctx, sql, args...)
	if err != nil {
		logrus.Error("Notification - Repository - DeleteNotification")
		return err
	}
	return nil
}

func (n NotificationRepository) UpdateNotifications(ctx context.Context, id string, notification models.UpdatedNotification) (*models.UpdatedNotification, error) {
	sql := `UPDATE notification SET title=COALESCE($2, title), 
                 body=COALESCE($3, body), telegram =COALESCE($4, telegram),
                 email=COALESCE($5, email), execution=CASE WHEN $6=TO_DATE('0001-01-01T00:00:00Z', 'YYYY-MM-DD"T"hh24:MI:SS"Z"') then execution else $6 END, 
                 assignto=COALESCE($7, assignto)
                 WHERE id=$1 
                 RETURNING id, title, body, telegram, email, execution, AssignTo`
	updatedNotification := new(models.UpdatedNotification)
	err := n.Pool.QueryRow(ctx, sql, id, notification.Title,
		notification.Body, notification.Telegram, notification.Email,
		notification.Execution, notification.AssignTo).Scan(&updatedNotification.Id, &updatedNotification.Title,
		&updatedNotification.Body, &updatedNotification.Telegram, &updatedNotification.Email,
		&updatedNotification.Execution, &updatedNotification.AssignTo)
	if err != nil {
		logrus.Error("Notification - Repository - UpdateNotifications")
		return nil, err
	}
	return updatedNotification, nil
}

func (n NotificationRepository) GetNotificationById(ctx context.Context, id string) (*models.Notification, error) {
	sql, args, err := n.Builder.
		Select("id, title, body, telegram, email, execution, assignTo").
		From("notification").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	notification := new(models.Notification)
	if err != nil {
		logrus.Error("Notification - Repository - GetNotificationById")
		return nil, err
	}
	err = n.Pool.QueryRow(ctx, sql, args...).Scan(&notification.Id, &notification.Title,
		&notification.Body, &notification.Telegram, &notification.Email, &notification.Execution, &notification.AssignTo)
	if err != nil {
		logrus.Error("Notification - Repository - GetNotificationById")
		return nil, err
	}
	return notification, nil

}

func (n NotificationRepository) GetAllNotifications(ctx context.Context) ([]models.Notification, error) {
	sql, args, err := n.Builder.
		Select("id, title, body, telegram, email, execution, assignTo").
		From("notification").
		ToSql()
	notifications := make([]models.Notification, 0)
	if err != nil {
		logrus.Error("Notification - Repository - GetAllNotifications")
		return nil, err
	}
	rows, err := n.Pool.Query(ctx, sql, args...)
	if err != nil {
		logrus.Error("Notification - Repository - GetAllNotifications")
		return nil, err
	}
	for rows.Next() {
		notification := new(models.Notification)
		err = rows.Scan(&notification.Id, &notification.Title,
			&notification.Body, &notification.Telegram, &notification.Email,
			&notification.Execution, &notification.AssignTo)
		if err != nil {
			logrus.Error("Notification - Repository - GetAllNotifications")
			return nil, err
		}
		notifications = append(notifications, *notification)
	}
	return notifications, nil
}

func (n NotificationRepository) GetNotificationsByUser(ctx context.Context, id string) ([]models.Notification, error) {
	sql, args, err := n.Builder.
		Select("id, title, body, telegram, email, execution, assignTo").
		From("notification").
		Where(squirrel.Eq{"assignTo": id}).
		ToSql()
	notifications := make([]models.Notification, 0)
	if err != nil {
		logrus.Error("Notification - Repository - GetAllNotifications")
		return nil, err
	}
	rows, err := n.Pool.Query(ctx, sql, args...)
	if err != nil {
		logrus.Error("Notification - Repository - GetAllNotifications")
		return nil, err
	}
	for rows.Next() {
		notification := new(models.Notification)
		err = rows.Scan(&notification.Id, &notification.Title,
			&notification.Body, &notification.Telegram, &notification.Email,
			&notification.Execution, &notification.AssignTo)
		if err != nil {
			logrus.Error("Notification - Repository - GetAllNotifications")
			return nil, err
		}
		notifications = append(notifications, *notification)
	}
	return notifications, nil
}
