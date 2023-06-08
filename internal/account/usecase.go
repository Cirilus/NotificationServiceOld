package account

import (
	"Notifications/internal/models"
	"context"
)

type UseCase interface {
	AccountById(ctx context.Context, id string) (*models.Account, error)
	AllAccounts(ctx context.Context) ([]models.Account, error)

	CreateAccount(ctx context.Context, account models.Account) error
}
