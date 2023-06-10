package account

import (
	"Notifications/internal/models"
	"context"
	"github.com/google/uuid"
)

// TODO Add the delete, update endpoints
type Repository interface {
	GetAccountById(ctx context.Context, id string) (*models.Account, error)
	GetAllAccounts(ctx context.Context) ([]models.Account, error)

	DeleteAccount(ctx context.Context, id uuid.UUID) error

	UpdateAccount(ctx context.Context, id uuid.UUID, account models.UpdateAccount) (*models.Account, error)

	CreateAccount(ctx context.Context, account models.Account) error
}
