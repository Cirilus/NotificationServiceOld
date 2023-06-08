package http

import (
	"Notifications/internal/account"
	"Notifications/internal/models"
	"github.com/gin-gonic/gin"
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

// @Summary Create the Account
// @Tags account
// @Accept  json
// @Produce json
// @Param notification body models.Account true "Account object that needs to be created"
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
