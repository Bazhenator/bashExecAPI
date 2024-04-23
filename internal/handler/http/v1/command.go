package v1

import (
	"SQLbash/internal/service"
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

func (h *CommandHandler) ListCommands(w http.ResponseWriter, r *http.Request) {
	commands, err := h.service.GetCommands(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(commands)
}

func (h *CommandHandler) RunCommand(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	output, err := h.service.RunCommand(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"output": output})
}

func (h *CommandHandler) AddCommand(w http.ResponseWriter, r *http.Request) {
	var command struct {
		Command string `json:"command"`
	}
	err := json.NewDecoder(r.Body).Decode(&command)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	id, err := h.service.AddCommand(r.Context(), command.Command)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

func (h *CommandHandler) SetRouter(router *mux.Router) {
	router.HandleFunc("/commands", h.ListCommands).Methods(http.MethodGet)
	router.HandleFunc("/commands/{id}", h.RunCommand).Methods(http.MethodGet)
	router.HandleFunc("/commands", h.AddCommand).Methods(http.MethodPost)
}
