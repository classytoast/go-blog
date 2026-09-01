package core_logger

import (
	"fmt"
	"os"
)

type Config struct {
	Level  string
	Folder string
}

func NewConfig() (Config, error) {
	level, ok := os.LookupEnv("LOGGER_LEVEL")
	if !ok {
		level = "DEBUG"
	}

	folder, ok := os.LookupEnv("LOGGER_FOLDER")
	if !ok {
		return Config{}, fmt.Errorf("environment variable LOGGER_FOLDER is not set")
	}

	return Config{
		Level:  level,
		Folder: folder,
	}, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get Logger config: %w", err)
		panic(err)
	}

	return config
}
