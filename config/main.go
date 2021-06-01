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
}

type ConfigImpl struct {
	sync.Mutex

	//internal objects
	dbClient *gorm.DB
	jwt      *jwtauth.JWTAuth
	url      *url.URL
	log      *logrus.Entry
}

func NewConfig() Config {
	return &ConfigImpl{
		Mutex: sync.Mutex{},
	}
}
