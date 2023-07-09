package http

import (
	"Notifications/internal/account"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc account.UseCase) {
	h := NewHandler(uc)

	router.GET("/", h.AllAccounts)
	router.GET("/me", h.UserAccount)
	router.GET("/:id", h.AccountById)

	router.POST("/", h.CreateAccount)

	router.DELETE("/:id", h.DeleteAccount)

	router.PUT("/:id", h.UpdateAccount)
}
