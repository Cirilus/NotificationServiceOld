package main

import (
	"Notifications/internal/config"
	"Notifications/server"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.GetConfig("")
	app := server.NewApp(cfg)

	err := app.Run("8000")
	if err != nil {
		logrus.Fatalf("The app can't run, err=%s", err)
	}
}
