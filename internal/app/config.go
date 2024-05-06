package app

import (
	provider "bashExecAPI/internal/db"
	"bashExecAPI/internal/logger"
	"bashExecAPI/internal/server"
	"time"
)

// Server defaults
const (
	defaultReadTimeout      time.Duration = 5 * time.Second
	defaultWriteTimeout     time.Duration = 5 * time.Second
	defaultServerHost       string        = ":8000"
	defaultResourseTimeout  time.Duration = 5 * time.Second
	defaultAnalyticsTimeout time.Duration = 15 * time.Second
)

type Config struct {
	DbConfig     provider.DbConfig   `yaml:"db"`
	LoggerConfig logger.LoggerConfig `yaml:"logger"`
	ServerConfig server.ServerConfig `yaml:"rest"`
}

func (config *Config) PopulateConfig() {
	s := func(s string) *string { return &s }
	d := func(d time.Duration) *time.Duration { return &d }

	if config.ServerConfig.Host == nil {
		config.ServerConfig.Host = s(defaultServerHost)
	}
	if config.ServerConfig.AnalyticsTimeout == nil {
		config.ServerConfig.AnalyticsTimeout = d(defaultAnalyticsTimeout)
	}
	if config.ServerConfig.ResourceTimeout == nil {
		config.ServerConfig.ReadTimeout = d(defaultResourseTimeout)
	}
	if config.ServerConfig.ReadTimeout == nil {
		config.ServerConfig.ReadTimeout = d(defaultReadTimeout)
	}
	if config.ServerConfig.WriteTimeout == nil {
		config.ServerConfig.WriteTimeout = d(defaultWriteTimeout)
	}
}
