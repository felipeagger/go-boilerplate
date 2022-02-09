package config

import (
	"log"

	env "github.com/Netflix/go-env"
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
	DBHost        string `env:"DB_HOST,default=localhost"`
	DBName        string `env:"DB_NAME,default=auth"`
	DBUser        string `env:"DB_USER,default=root"`
	DBPass        string `env:"DB_PASS,default=toor"`
	TraceHost     string `env:"TRACE_HOST,default=localhost:14268"`
	TokenSecret   string `env:"TOKEN_SECRET,default=mySecretJWT"`
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
