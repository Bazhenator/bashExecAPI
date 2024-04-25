package server

import "time"

type ServerConfig struct {
	Host             *string        `yaml:"host" validate:"hostname_port"`
	ConnectorHost    *string        `yaml:"connector-host" validate:"required,hostname_port"`
	ResourceTimeout  *time.Duration `yaml:"resource-timeout"`
	AnalyticsTimeout *time.Duration `yaml:"analytics-timeout"`
	ReadTimeout      *time.Duration `yaml:"read-timeout"`
	WriteTimeout     *time.Duration `yaml:"write-timeout"`
}
