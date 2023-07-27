package tests

import (
	"Notifications/internal/models"
	"Notifications/internal/notification/mocks"
	"Notifications/internal/notification/usecase"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestCreateNotification(t *testing.T) {
	mockRepository := new(mocks.Repository)
	id := uuid.New()
	mockNotification := models.Notification{
		Id:        id,
		Title:     "test",
		Body:      "test",
		Telegram:  nil,
		Email:     nil,
		Execution: time.Now(),
		AssignTo:  nil,
	}
	t.Run("success", func(t *testing.T) {
		mockRepository.On("CreateNotification", mock.Anything, mockNotification).Return(nil).Once()
		u := usecase.NewUseCase(mockRepository)
		err := u.CreateNotification(context.Background(), mockNotification)
		assert.NoError(t, err)

		mockRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepository.On("CreateNotification", mock.Anything, mockNotification).Return(errors.New("test")).Once()
		u := usecase.NewUseCase(mockRepository)

		err := u.CreateNotification(context.Background(), mockNotification)

		assert.Error(t, err)

		mockRepository.AssertExpectations(t)
	})
}

func TestAddNotificationToAccount(t *testing.T) {
	mockRepository := new(mocks.Repository)
	notificationId := uuid.New()
	accountId := make([]uuid.UUID, 5)
	for i := range accountId {
		accountId[i] = uuid.New()
	}
	t.Run("success", func(t *testing.T) {
		mockRepository.On("AddNotificationToAccount", mock.Anything, notificationId, accountId).Return(nil).Once()
		u := usecase.NewUseCase(mockRepository)
		err := u.AddNotificationToAccount(context.Background(), notificationId, accountId)

		assert.NoError(t, err)

		mockRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepository.On("AddNotificationToAccount", mock.Anything, notificationId, accountId).Return(errors.New("test")).Once()
		u := usecase.NewUseCase(mockRepository)
		err := u.AddNotificationToAccount(context.Background(), notificationId, accountId)

		assert.Error(t, err)
		mockRepository.AssertExpectations(t)
	})
}

func TestDeleteNotificationFromAccount(t *testing.T) {
	mockRepository := new(mocks.Repository)
	notificationId := uuid.New()
	accountId := uuid.New()
	t.Run("success", func(t *testing.T) {
		mockRepository.On("DeleteNotificationFromAccount", mock.Anything, notificationId, accountId).Return(nil).Once()
		u := usecase.NewUseCase(mockRepository)
		err := u.DeleteNotificationFromAccount(context.Background(), notificationId, accountId)

		assert.NoError(t, err)

		mockRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepository.On("DeleteNotificationFromAccount", mock.Anything, notificationId, accountId).Return(errors.New("test")).Once()
		u := usecase.NewUseCase(mockRepository)
		err := u.DeleteNotificationFromAccount(context.Background(), notificationId, accountId)

		assert.Error(t, err)
		mockRepository.AssertExpectations(t)
	})
}

func TestDeleteNotification(t *testing.T) {
	mockRepository := new(mocks.Repository)
	id := uuid.New()
	t.Run("success", func(t *testing.T) {
		mockRepository.On("DeleteNotification", mock.Anything, id.String()).Return(nil).Once()
		u := usecase.NewUseCase(mockRepository)
		err := u.DeleteNotification(context.Background(), id.String())

		assert.NoError(t, err)

		mockRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepository.On("DeleteNotification", mock.Anything, id.String()).Return(errors.New("test")).Once()
		u := usecase.NewUseCase(mockRepository)

		err := u.DeleteNotification(context.Background(), id.String())

		assert.Error(t, err)
		mockRepository.AssertExpectations(t)
	})
}

func TestUpdateNotifications(t *testing.T) {
	mockRepository := new(mocks.Repository)
	id := uuid.New()
	title := "test"
	body := "test"
	t.Run("success", func(t *testing.T) {
		mockUpdateNotification := models.UpdatedNotification{
			Id:        id,
			Title:     &title,
			Body:      &body,
			Telegram:  nil,
			Email:     nil,
			Execution: time.Now(),
		}

		mockRepository.On("UpdateNotifications", mock.Anything, id.String(), mockUpdateNotification).Return(&mockUpdateNotification, nil).Once()
		u := usecase.NewUseCase(mockRepository)
		notification, err := u.UpdateNotifications(context.Background(), id.String(), mockUpdateNotification)

		assert.NoError(t, err)
		assert.NotNil(t, notification)
		assert.Equal(t, mockUpdateNotification, *notification)

		mockRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUpdateNotification := models.UpdatedNotification{
			Id:        id,
			Title:     &title,
			Body:      &body,
			Telegram:  nil,
			Email:     nil,
			Execution: time.Now(),
		}
		mockRepository.On("UpdateNotifications", mock.Anything, id.String(), mockUpdateNotification).Return(nil, errors.New("test")).Once()
		u := usecase.NewUseCase(mockRepository)

		notification, err := u.UpdateNotifications(context.Background(), id.String(), mockUpdateNotification)

		assert.Error(t, err)
		assert.Nil(t, notification)

		mockRepository.AssertExpectations(t)
	})
}
func TestNotificationById(t *testing.T) {
	mockRepository := new(mocks.Repository)
	id := uuid.New()
	t.Run("success", func(t *testing.T) {
		mockNotification := models.Notification{
			Id:        id,
			Title:     "test",
			Body:      "test",
			Telegram:  nil,
			Email:     nil,
			Execution: time.Now(),
			AssignTo:  nil,
		}
		mockRepository.On("GetNotificationById", mock.Anything, id.String()).Return(&mockNotification, nil).Once()
		u := usecase.NewUseCase(mockRepository)
		notification, err := u.NotificationById(context.Background(), id.String())
		assert.NoError(t, err)
		assert.NotNil(t, notification)
		assert.Equal(t, mockNotification, *notification)

		mockRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepository.On("GetNotificationById", mock.Anything, id.String()).Return(nil, errors.New("test")).Once()
		u := usecase.NewUseCase(mockRepository)

		notification, err := u.NotificationById(context.Background(), id.String())

		assert.Error(t, err)
		assert.Nil(t, notification)

		mockRepository.AssertExpectations(t)
	})
}

func TestAllNotifications(t *testing.T) {
	mockRepository := new(mocks.Repository)
	t.Run("success", func(t *testing.T) {
		mockNotifications := make([]models.Notification, 2)
		mockNotifications = append(mockNotifications,
			models.Notification{
				Id:        uuid.New(),
				Title:     "test1",
				Body:      "test2",
				Telegram:  nil,
				Email:     nil,
				Execution: time.Now(),
				AssignTo:  nil,
			},
			models.Notification{
				Id:        uuid.New(),
				Title:     "test1",
				Body:      "test2",
				Telegram:  nil,
				Email:     nil,
				Execution: time.Now(),
				AssignTo:  nil,
			},
		)
		mockRepository.On("GetAllNotifications", mock.Anything).Return(mockNotifications, nil).Once()
		u := usecase.NewUseCase(mockRepository)
		notifications, err := u.AllNotifications(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, notifications)
		assert.Equal(t, mockNotifications, notifications)

		mockRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepository.On("GetAllNotifications", mock.Anything).Return(nil, errors.New("test")).Once()
		u := usecase.NewUseCase(mockRepository)

		notifications, err := u.AllNotifications(context.Background())

		assert.Error(t, err)
		assert.Nil(t, notifications)

		mockRepository.AssertExpectations(t)
	})

}

func TestNotificationsByUser(t *testing.T) {
	mockRepository := new(mocks.Repository)
	userId := uuid.New()
	assignTo := make([]uuid.UUID, 1)
	assignTo = append(assignTo, userId)
	t.Run("success", func(t *testing.T) {
		mockNotifications := make([]models.Notification, 0)
		mockNotifications = append(mockNotifications,
			models.Notification{
				Id:        uuid.New(),
				Title:     "test1",
				Body:      "test2",
				Telegram:  nil,
				Email:     nil,
				Execution: time.Now(),
				AssignTo:  assignTo,
			},
			models.Notification{
				Id:        uuid.New(),
				Title:     "test1",
				Body:      "test2",
				Telegram:  nil,
				Email:     nil,
				Execution: time.Now(),
				AssignTo:  assignTo,
			},
		)
		mockRepository.On("GetNotificationsByUser", mock.Anything, userId.String()).Return(mockNotifications, nil).Once()
		u := usecase.NewUseCase(mockRepository)
		notifications, err := u.NotificationsByUser(context.Background(), userId.String())
		assert.NoError(t, err)
		assert.NotNil(t, notifications)
		for _, notification := range notifications {
			assert.Contains(t, notification.AssignTo, userId)
		}

		mockRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepository.On("GetNotificationsByUser", mock.Anything, userId.String()).Return(nil, errors.New("test")).Once()
		u := usecase.NewUseCase(mockRepository)

		notification, err := u.NotificationsByUser(context.Background(), userId.String())

		assert.Error(t, err)
		assert.Nil(t, notification)

		mockRepository.AssertExpectations(t)
	})
}
