package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"sync"
)

type Configs struct {
	Level Level `yaml:"level" env:"required"`

	Prod Config `yaml:"prod"`

	Test Config `yaml:"test"`

	Dev Config `yaml:"dev"`
}

var instance *Configs
var once sync.Once

func GetConfig(level Level) *Config {
	once.Do(func() {
		logrus.Info("read application config")
		instance = &Configs{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logrus.Info(help)
			logrus.Fatal(err)
		}
	})

	if level == "" {
		level = instance.Level
	}

	switch level {
	case Prod:
		return &instance.Prod
	case Test:
		return &instance.Test
	case Dev:
		return &instance.Dev
	}
	return nil
}
