package http

import (
	"Notifications/internal/models"
	"Notifications/internal/notification"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	useCase notification.UseCase
}

func NewHandler(useCase notification.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

// @Summary Create the User's notifications.
// @Tags notification
// @Accept  json
// @Produce json
// @Param notification body models.Notification true "Notification object that needs to be created"
// @Success 201
// @Failure 501
// @Router /api/notification/ [post]
func (h Handler) CreateNotification(c *gin.Context) {
	inp := new(models.Notification)
	err := c.BindJSON(inp)
	if err != nil {
		logrus.Errorf("Bad request, err= %s", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = h.useCase.CreateNotification(c.Request.Context(), *inp)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}

// @Summary Return all notifications
// @Tags notification
// @Accept  json
// @Produce json
// @Success 201 {array} models.Notification
// @Failure 501
// @Router /api/notification/ [get]
func (h Handler) GetAllNotifications(c *gin.Context) {
	notifications, err := h.useCase.AllNotifications(c.Request.Context())
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, notifications)
}

// @Summary		Return notification by id
// @Tags 		notification
// @Accept  	json
// @Produce 	json
// @Param		id path string true "the id of the notification"
// @Success 	201 {object} models.Notification
// @Failure 	501
// @Router		/api/notification/{id} [get]
func (h Handler) GetNotificationsById(c *gin.Context) {
	id := c.Param("id")
	notifications, err := h.useCase.NotificationById(c.Request.Context(), id)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, notifications)
}

// @Summary		Return notification by user_id
// @Tags 		notification
// @Accept  	json
// @Produce 	json
// @Param		id path string true "user id of the notifications"
// @Success 	201 {array} models.Notification
// @Failure 	501
// @Router		/api/notification/user/{id} [get]
func (h Handler) GetNotificationsByUser(c *gin.Context) {
	id := c.Param("id")
	notifications, err := h.useCase.NotificationsByUser(c.Request.Context(), id)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, notifications)
}

// @Summary		Update the notification
// @Tags 		notification
// @Accept  	json
// @Produce 	json
// @Param		id path string true "id of the notifications"
// @Param		updatednotification body models.Notification true "The Update notification"
// @Success 	200 {object} models.Notification
// @Failure 	501
// @Router		/api/notification/{id} [put]
func (h Handler) UpdateNotification(c *gin.Context) {
	id := c.Param("id")
	inp := new(models.UpdatedNotification)
	err := c.BindJSON(inp)
	if err != nil {
		logrus.Errorf("Bad request, err= %s", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	updateNotifications, err := h.useCase.UpdateNotifications(c.Request.Context(), id, *inp)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, updateNotifications)
}

// @Summary		Delete the notification
// @Tags 		notification
// @Accept  	json
// @Produce 	json
// @Param		id path string true "delete the notifications"
// @Success 	200
// @Failure 	501
// @Router		/api/notification/{id} [delete]
func (h Handler) DeleteNotification(c *gin.Context) {
	id := c.Param("id")
	err := h.useCase.DeleteNotification(c.Request.Context(), id)
	if err != nil {
		logrus.Error(err)
		return
	}
	c.Status(http.StatusOK)
}
