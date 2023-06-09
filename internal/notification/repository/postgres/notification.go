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
	sql := `INSERT INTO notification(id, title, body, execution, telegram, email) 
	VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := n.Pool.Exec(ctx, sql, id, notification.Title, notification.Body,
		notification.Execution, notification.Telegram, notification.Email)
	if err != nil {
		logrus.Error("Notification - Repository - CreateNotification - CreatingNotification")
		return err
	}
	if notification.AssignTo != nil {
		sql = `INSERT INTO account_notification(notification_id, account_id) SELECT $1, id FROM account WHERE id = ANY($2)`
		_, err = n.Pool.Exec(ctx, sql, id, notification.AssignTo)
		if err != nil {
			logrus.Error("Notification - Repository - CreateNotification - CreatingRelations")
			return err
		}
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
                 email=COALESCE($5, email), execution=CASE WHEN $6=TO_DATE('0001-01-01T00:00:00Z', 'YYYY-MM-DD"T"hh24:MI:SS"Z"') then execution else $6 END 
                 WHERE id=$1 
                 RETURNING id, title, body, telegram, email, execution`
	updatedNotification := new(models.UpdatedNotification)
	err := n.Pool.QueryRow(ctx, sql, id, notification.Title,
		notification.Body, notification.Telegram, notification.Email,
		notification.Execution).Scan(&updatedNotification.Id, &updatedNotification.Title,
		&updatedNotification.Body, &updatedNotification.Telegram, &updatedNotification.Email,
		&updatedNotification.Execution)
	if err != nil {
		logrus.Error("Notification - Repository - UpdateNotifications")
		return nil, err
	}
	return updatedNotification, nil
}

func (n NotificationRepository) GetNotificationById(ctx context.Context, id string) (*models.Notification, error) {
	sql := `SELECT notification.id, title, body, telegram, email, execution, array_agg(an.account_id) as account_ids
	FROM notification LEFT JOIN account_notification an on notification.id = an.notification_id
	WHERE notification.id = $1
	GROUP BY notification.id, title, body, telegram, email, execution`

	notification := new(models.Notification)

	err := n.Pool.QueryRow(ctx, sql, id).Scan(&notification.Id, &notification.Title,
		&notification.Body, &notification.Telegram, &notification.Email, &notification.Execution, &notification.AssignTo)
	if err != nil {
		logrus.Error("Notification - Repository - GetNotificationById")
		return nil, err
	}
	if len(notification.AssignTo) == 1 && notification.AssignTo[0].String() == "00000000-0000-0000-0000-000000000000" {
		notification.AssignTo = nil
	}
	return notification, nil

}

func (n NotificationRepository) GetAllNotifications(ctx context.Context) ([]models.Notification, error) {

	sql := `SELECT notification.id, title, body, telegram, email, execution, array_agg(an.account_id) as account_ids
	FROM notification LEFT JOIN account_notification an on notification.id = an.notification_id
	GROUP BY notification.id, title, body, telegram, email, execution;`

	notifications := make([]models.Notification, 0)

	rows, err := n.Pool.Query(ctx, sql)
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
		if len(notification.AssignTo) == 1 && notification.AssignTo[0].String() == "00000000-0000-0000-0000-000000000000" {
			notification.AssignTo = nil
		}
		notifications = append(notifications, *notification)
	}

	return notifications, nil
}

func (n NotificationRepository) GetNotificationsByUser(ctx context.Context, id string) ([]models.Notification, error) {

	sql := `SELECT notification.id, title, body, telegram, email, execution
	FROM notification LEFT JOIN account_notification an on notification.id = an.notification_id
	WHERE account_id = $1`

	notifications := make([]models.Notification, 0)

	rows, err := n.Pool.Query(ctx, sql, id)
	if err != nil {
		logrus.Error("Notification - Repository - GetAllNotifications")
		return nil, err
	}

	for rows.Next() {
		notification := new(models.Notification)
		err = rows.Scan(&notification.Id, &notification.Title,
			&notification.Body, &notification.Telegram, &notification.Email,
			&notification.Execution)
		if err != nil {
			logrus.Error("Notification - Repository - GetAllNotifications")
			return nil, err
		}
		notifications = append(notifications, *notification)
	}

	return notifications, nil
}

func (n NotificationRepository) AddNotificationToAccount(ctx context.Context, notificationId uuid.UUID, accountId []uuid.UUID) error {
	sql := `INSERT INTO account_notification(notification_id, account_id) SELECT $1, id FROM account WHERE id = ANY($2)`
	_, err := n.Pool.Exec(ctx, sql, notificationId, accountId)
	if err != nil {
		logrus.Error("Notification - Repository - AddNotificationToAccount")
		return err
	}
	return nil
}

func (n NotificationRepository) DeleteNotificationFromAccount(ctx context.Context, notificationId uuid.UUID, accountId uuid.UUID) error {
	sql := "DELETE FROM account_notification WHERE notification_id = $1 and account_id = $2"
	_, err := n.Pool.Exec(ctx, sql, notificationId, accountId)
	if err != nil {
		logrus.Error("Notification - Repository - DeleteNotificationFromAccount")
		return err
	}
	return nil
}
