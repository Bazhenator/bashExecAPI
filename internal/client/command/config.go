package command

import "time"

type Config struct {
	Host               string        `yaml:"host"`
	CommandUrl         string        `yaml:"command-url"`
	CommandsPerRequest int           `yaml:"commands-per-request"`
	ThreadCount        int           `yaml:"thread-count"`
	MaxTimeSleep       time.Duration `yaml:"max-time-sleep"`
	MinTimeSleep       time.Duration `yaml:"min-time-sleep"`
	ReadTimeout        time.Duration `yaml:"read-timeout"`
	WriteTimeout       time.Duration `yaml:"write-timeout"`
}
