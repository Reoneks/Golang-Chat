package config

import (
	"github.com/caarlos0/env"
	"github.com/go-chi/jwtauth"
)

type JWT struct {
	Secret    string `env:"JWT_SECRET" envDefault:"9caf06bb4436cdbfa20af9121a626bc1093c4f54b31c0fa937957856135345b6"`
	Algorithm string `env:"JWT_ALGORITHM" envDefault:"HS256"`
}

func (jwt *JWT) GetJWTEntry() *jwtauth.JWTAuth {
	return jwtauth.New(jwt.Algorithm, []byte(jwt.Secret), nil)
}

func (c *ConfigImpl) JWT() *jwtauth.JWTAuth {
	if c.jwt != nil {
		return c.jwt
	}

	jwt := &JWT{}
	if err := env.Parse(jwt); err != nil {
		panic(err)
	}

	c.jwt = jwt.GetJWTEntry()

	return c.jwt
}
