package server

import (
	_ "Notifications/docs"
	"Notifications/internal/account"
	accounthttp "Notifications/internal/account/delivery/http"
	accountrepo "Notifications/internal/account/repository/postgres"
	accountusecase "Notifications/internal/account/usecase"
	"Notifications/internal/config"
	"Notifications/internal/notification"
	notificationhttp "Notifications/internal/notification/delivery/http"
	notificationrepo "Notifications/internal/notification/repository/postgres"
	notificationusecase "Notifications/internal/notification/usecase"
	"Notifications/pkg/client/postgresql"
	"Notifications/pkg/health_check"
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	httpServer *http.Server

	notificationUC notification.UseCase

	accountUC account.UseCase
}

func NewApp(cfg *config.Config) *App {
	logrus.Info("Connecting to db")
	db, err := postgresql.New(fmt.Sprintf("postgresql://%s:%s@%s:%s?sslmode=disable",
		cfg.PostgreSQL.Username, cfg.PostgreSQL.Password,
		cfg.PostgreSQL.Host, cfg.PostgreSQL.Port))
	if err != nil {
		logrus.Fatalf("Problem with connection to db, err= %s", err)
	}

	notificationRepo := notificationrepo.NewRepository(db)
	accountRepo := accountrepo.NewAccountRepository(db)

	return &App{
		httpServer:     nil,
		notificationUC: notificationusecase.NewUseCase(notificationRepo),
		accountUC:      accountusecase.NewAccountUseCase(accountRepo),
	}
}

// TODO the notificators(telegram, email)

// Run @title Golang Init
// @version 1.0
// @description This is init
// @host localhost:8000
// @BasePath /
// @schemes http
func (a *App) Run(port string) error {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	health_check.RegisterHTTPEEndpoints(router)

	api := router.Group("/api")

	notificationApi := api.Group("/notification")
	notificationhttp.RegisterHTTPEndpoints(notificationApi, a.notificationUC)

	accountApi := api.Group("/account")
	accounthttp.RegisterHTTPEndpoints(accountApi, a.accountUC)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	logrus.Infof("The server was run on port %s", port)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit
	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
