package usecase

import (
	"Notifications/internal/account"
	"Notifications/internal/models"
	"context"
	"github.com/sirupsen/logrus"
)

type AccountUseCase struct {
	repo account.Repository
}

func NewAccountUseCase(repo account.Repository) *AccountUseCase {
	return &AccountUseCase{repo: repo}
}

func (a AccountUseCase) AccountById(ctx context.Context, id string) (*models.Account, error) {
	accountById, err := a.repo.GetAccountById(ctx, id)
	if err != nil {
		logrus.Error("Account - UseCase - AccountById")
		return nil, err
	}
	return accountById, nil
}

func (a AccountUseCase) AllAccounts(ctx context.Context) ([]models.Account, error) {
	allAccounts, err := a.repo.GetAllAccounts(ctx)
	if err != nil {
		logrus.Error("Account - UseCase - AllAccounts")
		return nil, err
	}
	return allAccounts, nil
}

func (a AccountUseCase) CreateAccount(ctx context.Context, account models.Account) error {
	err := a.repo.CreateAccount(ctx, account)
	if err != nil {
		logrus.Error("Account - UseCase - CreateAccount")
		return err
	}
	return nil
}
