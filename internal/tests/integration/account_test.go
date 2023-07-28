package integration_test

import (
	"Notifications/internal/config"
	"Notifications/server"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AccountSuit struct {
	suite.Suite
}

func TestSuit(t *testing.T) {
	suite.Run(t, &AccountSuit{})
}

func (ts *AccountSuit) SetupSuite() {
	cfg := config.GetConfig(config.Test)
	app := server.NewApp(cfg)

	err := app.Run("8000")
	if err != nil {
		logrus.Fatalf("The app can't run, err=%s", err)
	}
}

func (ts *AccountSuit) AllAccounts() {
}
func (ts *AccountSuit) AccountById() {

}
func (ts *AccountSuit) UserAccount() {

}
func (ts *AccountSuit) CreateAccount() {

}
func (ts *AccountSuit) DeleteAccount() {

}
func (ts *AccountSuit) UpdateAccount() {

}
