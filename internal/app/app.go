package app

import (
	provider "bashExecAPI/internal/db"
	handler "bashExecAPI/internal/handler/http"
	"bashExecAPI/internal/repository"
	"bashExecAPI/internal/server"
	"bashExecAPI/internal/service"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

type App struct {
	server   *server.Server
	provider *provider.Provider
}

func NewApp(config *Config, notify chan error) (*App, error) {
	provider, err := provider.NewPsqlProvider(&config.DbConfig)
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize db with error: %w", err)
	}

	repos := repository.NewCommandRepository(provider)

	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	c := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedMethods:   []string{http.MethodOptions, http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodPatch},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Set-Cookie", "User-Agent", "Origin"},
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	})

	router := mux.NewRouter()
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	handlers.SetRouter(router)

	server := server.NewServer(&config.ServerConfig, c.Handler(router), notify)

	return &App{
		server:   server,
		provider: provider,
	}, nil
}

func (app *App) Start() {
	app.server.Start()
}

func (app *App) Stop() error {
	serverErr := app.server.Stop()
	providerErr := app.provider.Close()
	if serverErr != nil || providerErr != nil {
		return fmt.Errorf("Provider error: %w. Server error: %w", providerErr, serverErr)
	}
	return nil
}
