package http

import (
	"Notifications/internal/account"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc account.UseCase) {
	h := NewHandler(uc)

	router.GET("/", h.AllAccounts)
	router.GET("/:id", h.AccountById)
	router.POST("/", h.CreateAccount)
}
