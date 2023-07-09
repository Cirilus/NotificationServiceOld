package http

import (
	"Notifications/internal/account"
	"Notifications/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	uc account.UseCase
}

func NewHandler(uc account.UseCase) *Handler {
	return &Handler{uc: uc}
}

// @Summary Return all accounts
// @Tags account
// @Accept  json
// @Produce json
// @Success 201 {array} models.Account
// @Failure 500
// @Router /api/account/ [get]
func (h Handler) AllAccounts(c *gin.Context) {
	allAccounts, err := h.uc.AllAccounts(c.Request.Context())
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, allAccounts)
}

// @Summary Return account by id
// @Tags account
// @Accept  json
// @Produce json
// @Param id path string true "Id of the account"
// @Success 201 {object} models.Account
// @Failure 500
// @Router /api/account/{id} [get]
func (h Handler) AccountById(c *gin.Context) {
	id := c.Param("id")
	accountById, err := h.uc.AccountById(c.Request.Context(), id)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, accountById)
}

// @Summary Return account by id
// @Tags account
// @Accept  json
// @Produce json
// @Success 201 {object} models.Account
// @Failure 500
// @Router /api/account/me [get]
func (h Handler) UserAccount(c *gin.Context) {
	userAccount, err := h.uc.UserAccount(c.Request.Context())
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, userAccount)
}

// @Summary Create the Account
// @Tags account
// @Accept  json
// @Produce json
// @Param account body models.Account true "Account object that needs to be created"
// @Success 201
// @Failure 500
// @Router /api/account/ [post]
func (h Handler) CreateAccount(c *gin.Context) {
	inp := new(models.Account)
	err := c.BindJSON(inp)
	if err != nil {
		logrus.Errorf("Bad request = %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = h.uc.CreateAccount(c.Request.Context(), *inp)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
}

// @Summary DeleteAccount
// @Tags account
// @Accept  json
// @Produce json
// @Param id path string true "Id of the account"
// @Success 201
// @Failure 500
// @Router /api/account/{id} [delete]
func (h Handler) DeleteAccount(c *gin.Context) {
	id := c.Param("id")
	accountUUID, err := uuid.Parse(id)
	if err != nil {
		logrus.Errorf("Cannot convert id as uuid = %s", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = h.uc.DeleteAccount(c.Request.Context(), accountUUID)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

// @Summary Update the Account
// @Tags account
// @Accept  json
// @Produce json
// @Param id path string true "Id of the account"
// @Param account body models.Account true "Account object that needs to be updated"
// @Success 201
// @Failure 500
// @Router /api/account/{id} [put]
func (h Handler) UpdateAccount(c *gin.Context) {
	id := c.Param("id")
	accountUUID, err := uuid.Parse(id)
	if err != nil {
		logrus.Errorf("Cannot convert id as uuid = %s", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	inp := new(models.Account)
	err = c.BindJSON(inp)
	if err != nil {
		logrus.Errorf("Bad request = %s", err)
		return
	}

	updateAccount, err := h.uc.UpdateAccount(c.Request.Context(), accountUUID, *inp)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, updateAccount)
}
