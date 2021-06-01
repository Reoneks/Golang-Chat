package config

import (
	"github.com/caarlos0/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	DSN string `env:"DSN"`
}

func (c *ConfigImpl) DBClient() *gorm.DB {
	if c.dbClient != nil {
		return c.dbClient
	}

	c.Lock()
	defer c.Unlock()

	dbConfig := &DBConfig{}
	if err := env.Parse(dbConfig); err != nil {
		panic(err)
	}

	client, err := gorm.Open(postgres.Open(dbConfig.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	c.dbClient = client

	return client
}
