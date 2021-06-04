package config

import (
	"net/url"
	"sync"

	"github.com/go-chi/jwtauth"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Config interface {
	DBClient() *gorm.DB
	JWT() *jwtauth.JWTAuth
	ServerAddress() *url.URL
	Log() *logrus.Entry
	AppEnvironment() string
}

type ConfigImpl struct {
	sync.Mutex

	//internal objects
	dbClient *gorm.DB
	jwt      *jwtauth.JWTAuth
	url      *url.URL
	log      *logrus.Entry
	app_env  string
}

func NewConfig() Config {
	return &ConfigImpl{
		Mutex: sync.Mutex{},
	}
}
