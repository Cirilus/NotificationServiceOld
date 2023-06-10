package http

import (
	"Notifications/internal/account"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc account.UseCase) {
	h := NewHandler(uc)

	router.GET("/", h.AllAccounts)
	router.POST("/", h.CreateAccount)

	router.GET("/:id", h.AccountById)
	router.DELETE("/:id", h.DeleteAccount)
	router.PUT("/:id", h.UpdateAccount)
}
