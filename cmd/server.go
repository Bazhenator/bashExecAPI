package main

import (
	"fmt"
	"github.com/Bazhenator/bashExecAPI/internal/app"
	"github.com/Bazhenator/bashExecAPI/internal/config"
	_ "github.com/Bazhenator/bashExecAPI/internal/domain"
	"github.com/Bazhenator/bashExecAPI/internal/logger"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

const (
	configFile string = "../configs/config.yaml"
)

// @title			BashExecAPI
// @version         0.0.1
// @description     This web application provides API for executing bash-scripts.
//
// @host      localhost:8080
// @BasePath  /api/v1

func main() {
	conf, err := config.ReadConfigFromYAML[app.Config](configFile)
	if err != nil {
		panic(fmt.Errorf("Reading config from '%s' failed: %w", configFile, err))
	}
	err = config.ValidateConfig(conf)
	if err != nil {
		panic(fmt.Errorf("'%s' parsing failed: %w", configFile, err))
	}

	logs, err := os.OpenFile(conf.LoggerConfig.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModeAppend)
	if err != nil {
		panic(fmt.Errorf("Unable to create '%s' file. Error: %w", conf.LoggerConfig.LogFile, err))
	}
	defer func() {
		logs.Close()
		if err := os.Remove(conf.LoggerConfig.LogFile); err != nil {
			log.Println("Error removing warn file:", err)
		}
	}()

	errLogs, err := os.OpenFile(conf.LoggerConfig.WarnFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModeAppend)
	if err != nil {
		panic(fmt.Errorf("Unable to create '%s' file. Error: %w", conf.LoggerConfig.WarnFile, err))
	}
	defer func() {
		errLogs.Close()
		if err := os.Remove(conf.LoggerConfig.WarnFile); err != nil {
			log.Println("Error removing warn file:", err)
		}
	}()

	logger.SetupLogrus(logs, errLogs)
	log.Info("Starting...")

	notify := make(chan error, 1)
	defer close(notify)

	app, err := app.NewApp(conf, notify)
	if err != nil {
		log.Fatal(err)
	}

	app.Start()
	defer app.Stop()

	interupt := make(chan os.Signal, 1)
	defer close(interupt)

	signal.Notify(interupt, os.Interrupt, syscall.SIGTERM)

	select {
	case serr := <-notify:
		log.Error(fmt.Errorf("Notified with app error: %w", serr))
	case signl := <-interupt:
		log.Info("Cought signal while App running: " + signl.String())
	}

	log.Info("Shutting down...")
}
