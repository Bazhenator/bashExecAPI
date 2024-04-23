package v1

import (
	"SQLbash/internal/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	commandHandler *CommandHandler
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		commandHandler: NewCommandHandler(service),
	}
}

func (handler *Handler) SetRouter(router *mux.Router) {
	sub := router.PathPrefix("/v1").Subrouter()
	handler.commandHandler.SetRouter(sub)
}
