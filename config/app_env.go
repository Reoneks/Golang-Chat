package config

import (
	"github.com/caarlos0/env"
)

type AppEnv struct {
	App_Env string `env:"APP_ENV" envDefault:"development"`
}

func (c *ConfigImpl) AppEnvironment() string {
	if c.app_env != "" {
		return c.app_env
	}

	app_env := &AppEnv{}
	if err := env.Parse(app_env); err != nil {
		panic(err)
	}

	c.app_env = app_env.App_Env

	return c.app_env
}
