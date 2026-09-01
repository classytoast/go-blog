package core_http_server

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	Addr            string
	ShutdownTimeout time.Duration
}

func NewConfig() (Config, error) {
	addr, ok := os.LookupEnv("HTTP_ADDR")
	if !ok {
		return Config{}, fmt.Errorf("environment variable HTTP_ADDR is not set")
	}

	timeoutStr, ok := os.LookupEnv("HTTP_SHUTDOWN_TIMEOUT")
	if !ok {
		timeoutStr = "30s"
	}

	shutdownTimeout, err := time.ParseDuration(timeoutStr)
	if err != nil {
		return Config{}, fmt.Errorf(
			"invalid HTTP_SHUTDOWN_TIMEOUT %q: %w",
			timeoutStr,
			err,
		)
	}

	return Config{
		Addr:            addr,
		ShutdownTimeout: shutdownTimeout,
	}, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get server config: %w", err)
		panic(err)
	}

	return config
}
