package integration_test

import (
	accounthttp "Notifications/internal/account/delivery/http"
	accountrepo "Notifications/internal/account/repository/postgres"
	accountusecase "Notifications/internal/account/usecase"
	"Notifications/internal/config"
	"Notifications/internal/models"
	"Notifications/pkg/client/postgresql"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type AccountSuit struct {
	suite.Suite
	router *gin.Engine
}

func (as *AccountSuit) SetupSuite() {
	cfg := config.GetConfig(config.Test)

	logrus.Info("Connecting to db")
	db, err := postgresql.New(fmt.Sprintf("postgresql://%s:%s@%s:%s?sslmode=disable",
		cfg.PostgreSQL.Username, cfg.PostgreSQL.Password,
		cfg.PostgreSQL.Host, cfg.PostgreSQL.Port))
	logrus.Infof("postgresql://%s:%s@%s:%s?sslmode=disable",
		cfg.PostgreSQL.Username, cfg.PostgreSQL.Password,
		cfg.PostgreSQL.Host, cfg.PostgreSQL.Port)
	if err != nil {
		logrus.Fatalf("Problem with connection to db, err= %s", err)
	}

	accountRepo := accountrepo.NewAccountRepository(db)
	accountUC := accountusecase.NewAccountUseCase(accountRepo)

	router := gin.Default()
	gin.SetMode(gin.TestMode)
	as.router = router
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	api := router.Group("/api")
	accountApi := api.Group("/account")
	accounthttp.RegisterHTTPEndpoints(accountApi, accountUC)
}

func (as *AccountSuit) TestCreationAccount() {
	mockAccountID := uuid.New()
	mockTelegram := "test"
	mockEmail := "test"
	mockAccount := models.Account{
		Id:       &mockAccountID,
		Telegram: &mockTelegram,
		Email:    &mockEmail,
	}

	err := as.CreateAccount(mockAccount)
	as.Require().NoError(err)

	accountById, err := as.GetAccountById(mockAccountID)
	as.Require().NoError(err)

	as.Require().Equal(mockAccount, *accountById)
}

func (as *AccountSuit) GetAllAccounts() ([]models.Account, error) {
	router := as.router

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", fmt.Sprintf("/api/account"), nil)

	router.ServeHTTP(w, req)
	as.Require().Equal(http.StatusOK, w.Code)

	var resp []models.Account

	err := json.NewDecoder(w.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (as *AccountSuit) GetAccountById(userId uuid.UUID) (*models.Account, error) {
	router := as.router

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", fmt.Sprintf("/api/account/%s", userId), nil)

	router.ServeHTTP(w, req)
	as.Require().Equal(http.StatusOK, w.Code)

	resp := new(models.Account)

	err := json.NewDecoder(w.Body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

//func (as *AccountSuit) UserAccount(userId uuid.UUID) (*models.Account, error) {
//	router := as.app.GetRouter()
//
//	w := httptest.NewRecorder()
//	req := httptest.NewRequest("GET", fmt.Sprintf("/api/account/%s", userId), bytes.NewBuffer(jsonBody))
//
//	router.ServeHTTP(w, req)
//	as.Require().Equal(http.StatusOK, w.Code)
//
//	resp := new(models.Account)
//
//	err := json.NewDecoder(w.Body).Decode(resp)
//	if err != nil {
//		return nil, err
//	}
//
//	return resp, nil
//}

func (as *AccountSuit) CreateAccount(user models.Account) error {
	router := as.router

	jsonBody, err := json.Marshal(user)
	if err != nil {
		return err
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/account/", bytes.NewBuffer(jsonBody))

	router.ServeHTTP(w, req)

	logrus.Infof("%v", w)

	as.Require().Equal(http.StatusCreated, w.Code)

	return nil
}

func (as *AccountSuit) DeleteAccount(userId uuid.UUID) error {
	router := as.router

	w := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", fmt.Sprintf("/api/account/%s", userId), nil)
	router.ServeHTTP(w, req)

	as.Require().Equal(http.StatusOK, w.Code)

	return nil
}

func (as *AccountSuit) UpdateAccount(userId uuid.UUID, user models.Account) (*models.Account, error) {
	router := as.router

	jsonBody, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest("PUT", fmt.Sprintf("/api/account/%s", userId), bytes.NewBuffer(jsonBody))

	router.ServeHTTP(w, req)
	as.Require().Equal(http.StatusOK, w.Code)

	resp := new(models.Account)

	err = json.NewDecoder(w.Body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func TestSuit(t *testing.T) {
	suite.Run(t, new(AccountSuit))
}
