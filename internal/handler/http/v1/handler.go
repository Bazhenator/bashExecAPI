package v1

import (
	"github.com/Bazhenator/bashExecAPI/internal/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	commandHandler *CommandHandler
	dbHandler      *DataBaseHandler
}

func NewHandler(service *service.Services) *Handler {
	return &Handler{
		commandHandler: NewCommandHandler(service),
		dbHandler:      NewDataBaseHandler(service),
	}
}

func (handler *Handler) SetRouter(router *mux.Router) {
	sub := router.PathPrefix("/v1").Subrouter()
	handler.commandHandler.SetRouter(sub)
	handler.dbHandler.SetRouter(sub)
}
