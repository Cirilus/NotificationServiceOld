package account

import (
	"Notifications/internal/models"
	"context"
)

type Repository interface {
	GetAccountById(ctx context.Context, id string) (*models.Account, error)
	GetAllAccounts(ctx context.Context) ([]models.Account, error)

	CreateAccount(ctx context.Context, account models.Account) error
}
