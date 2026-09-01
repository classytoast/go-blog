package core_postgres_pool

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
	Timeout  time.Duration
}

func NewConfig() (Config, error) {
	host, ok := os.LookupEnv("POSTGRES_HOST")
	if !ok {
		return Config{}, fmt.Errorf("environment variable POSTGRES_HOST is not set")
	}

	port, ok := os.LookupEnv("POSTGRES_PORT")
	if !ok {
		return Config{}, fmt.Errorf("environment variable POSTGRES_PORT is not set")
	}

	user, ok := os.LookupEnv("POSTGRES_USER")
	if !ok {
		return Config{}, fmt.Errorf("environment variable POSTGRES_USER is not set")
	}

	password, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		return Config{}, fmt.Errorf("environment variable POSTGRES_PASSWORD is not set")
	}

	database, ok := os.LookupEnv("POSTGRES_DB")
	if !ok {
		return Config{}, fmt.Errorf("environment variable POSTGRES_DB is not set")
	}

	sslMode, ok := os.LookupEnv("POSTGRES_SSLMODE")
	if !ok {
		return Config{}, fmt.Errorf("environment variable POSTGRES_SSLMODE is not set")
	}

	timeoutStr, ok := os.LookupEnv("POSTGRES_TIMEOUT")
	if !ok {
		return Config{}, fmt.Errorf("environment variable POSTGRES_TIMEOUT is not set")
	}

	shutdownTimeout, err := time.ParseDuration(timeoutStr)
	if err != nil {
		return Config{}, fmt.Errorf(
			"invalid POSTGRES_TIMEOUT %q: %w",
			timeoutStr,
			err,
		)
	}

	return Config{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Database: database,
		SSLMode:  sslMode,
		Timeout:  shutdownTimeout,
	}, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get postgres connection pool config: %w", err)
		panic(err)
	}

	return config
}
