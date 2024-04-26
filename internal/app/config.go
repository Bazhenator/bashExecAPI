package app

import (
	"bashExecAPI/internal/client/command"
	provider "bashExecAPI/internal/db"
	"bashExecAPI/internal/logger"
)

type Config struct {
	DbConfig      *provider.DbConfig   `yaml:"db"`
	LoggerConfig  *logger.LoggerConfig `yaml:"logger"`
	CommandConfig *command.Config      `yaml:"command-service"`
}
