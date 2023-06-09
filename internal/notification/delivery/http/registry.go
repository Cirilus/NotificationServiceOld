package http

import (
	"Notifications/internal/notification"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc notification.UseCase) {
	h := NewHandler(uc)

	router.POST("/", h.CreateNotification)
	router.GET("/", h.GetAllNotifications)
	router.GET("/:id", h.GetNotificationsById)
	router.PUT("/:id", h.UpdateNotification)
	router.DELETE("/:id", h.DeleteNotification)

	router.GET("/user/:id", h.GetNotificationsByUser)
	router.POST("/user/:id", h.AddNotificationOnAccount)
	router.DELETE("/user/:id", h.DeleteNotificationFromAccount)
}
