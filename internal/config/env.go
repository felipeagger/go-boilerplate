package config

import (
	env "github.com/Netflix/go-env"
	"log"
)

var configs Environment

const (
	ServiceName = "go-boilerplate-auth"
)

type Environment struct {
	Env           string `env:"ENVIRONMENT,default=dev"`
	Port          string `env:"PORT,default=8000"`
	Debug         bool   `env:"DEBUG,default=false"`
	CachePort     string `env:"CACHE_PORT,default=6379"`
	CacheHost     string `env:"CACHE_HOST,default=localhost"`
	CachePassword string `env:"CACHE_PASSWORD"`
	TokenSecret   string `env:"TOKEN_SECRET,required=true"`
}

func init() {
	var environment Environment
	_, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatal(err)
	}

	configs = environment
}

// GetEnv return class with all env vars
func GetEnv() Environment {
	return configs
}
