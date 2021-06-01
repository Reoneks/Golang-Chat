package config

import (
	"net/url"

	"github.com/caarlos0/env"
)

type URL struct {
	URL string `env:"HTTP_SERVER_ADDRESS"`
}

func (c *ConfigImpl) ServerAddress() *url.URL {
	if c.url != nil {
		return c.url
	}

	serverURL := &URL{}
	if err := env.Parse(serverURL); err != nil {
		panic(err)
	}

	result, err := url.Parse(serverURL.URL)
	if err != nil {
		panic(err)
	}
	c.url = result

	return c.url
}
