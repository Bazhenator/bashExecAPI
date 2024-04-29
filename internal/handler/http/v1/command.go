package v1

import (
	errorlib "bashExecAPI/internal/error"
	"bashExecAPI/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type CommandHandler struct {
	service *service.Service
}

func NewCommandHandler(service *service.Service) *CommandHandler {
	return &CommandHandler{
		service: service,
	}
}

func (h *CommandHandler) GetCommands(w http.ResponseWriter, r *http.Request) {
	commands, err := h.service.GetCommands(r.Context())
	if err != nil {
		jsonError := errorlib.GetJSONError("Failed to get commands", err)
		w.WriteHeader(jsonError.Error.Code)
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	json.NewEncoder(w).Encode(commands)
}

func (h *CommandHandler) GetCommand(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	command, err := h.service.GetCommand(r.Context(), id)
	if err != nil {
		jsonError := errorlib.GetJSONError("Failed to get command", err)
		w.WriteHeader(jsonError.Error.Code)
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	json.NewEncoder(w).Encode(command)
}

func (h *CommandHandler) CreateCommand(w http.ResponseWriter, r *http.Request) {
	var command struct {
		Command string `json:"command"`
	}
	err := json.NewDecoder(r.Body).Decode(&command)
	if err != nil {
		jsonError := errorlib.GetJSONError("Invalid request body", err)
		w.WriteHeader(jsonError.Error.Code)
		json.NewEncoder(w).Encode(jsonError)
		return
	}

	id, err := h.service.CreateCommand(r.Context(), command.Command)
	if err != nil {
		jsonError := errorlib.GetJSONError("Failed to create command", err)
		w.WriteHeader(jsonError.Error.Code)
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func (h *CommandHandler) RunCommand(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	command, err := h.service.RunCommand(r.Context(), id)
	if err != nil {
		jsonError := errorlib.GetJSONError("Failed to run command", err)
		w.WriteHeader(jsonError.Error.Code)
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	json.NewEncoder(w).Encode(command)
}

func (h *CommandHandler) StopCommand(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.service.StopCommand(r.Context(), id)
	if err != nil {
		jsonError := errorlib.GetJSONError("Failed to stop command", err)
		w.WriteHeader(jsonError.Error.Code)
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *CommandHandler) SetRouter(router *mux.Router) {
	router.HandleFunc("/commands/list", h.GetCommands).Methods(http.MethodGet)
	router.HandleFunc("/commands/{id}", h.GetCommand).Methods(http.MethodGet)
	router.HandleFunc("/commands/create", h.CreateCommand).Methods(http.MethodPost)
	router.HandleFunc("/commands/run/{id}", h.RunCommand).Methods(http.MethodPost)
	router.HandleFunc("/commands/stop/{id}", h.StopCommand).Methods(http.MethodPost)
}
