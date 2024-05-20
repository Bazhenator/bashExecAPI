package v1

import (
	"encoding/json"
	errorlib "github.com/Bazhenator/bashExecAPI/internal/error"
	"github.com/Bazhenator/bashExecAPI/internal/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type CommandHandler struct {
	service *service.Services
}

func NewCommandHandler(service *service.Services) *CommandHandler {
	return &CommandHandler{
		service: service,
	}
}

// CreateCommand creates and executes new command
// @Summary      Create and execute new command
// @Description  Create and execute new command
// @Tags         Commands
// @Accept       json
// @Produce      json
// @Param        command  body  Creation  true  "Command to execute"
// @Success      200  {object}  CreationResponse
// @Failure      500  {object}  Error
// @Router       /commands/create [post]
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

	result, id, err := h.service.CreateCommand(r.Context(), command.Command)
	if err != nil {
		jsonError := errorlib.GetJSONError("Failed to create command", err)
		w.WriteHeader(jsonError.Error.Code)
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id, "result": result})
}

// ListCommands gets list of available commands
// @Summary      Get list of available commands
// @Description  Get list of available commands
// @Tags         Commands
// @Accept       json
// @Produce      json
// @Success      200  {array}  Commands
// @Failure      500  {object}  Error
// @Router       /commands/list [get]
func (h *CommandHandler) ListCommands(w http.ResponseWriter, r *http.Request) {
	commands, err := h.service.ListCommands(r.Context())
	if err != nil {
		jsonError := errorlib.GetJSONError("Failed to get commands", err)
		w.WriteHeader(jsonError.Error.Code)
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	json.NewEncoder(w).Encode(commands)
}

// GetCommand gets command with given id
// @Summary      Get command with given id
// @Description  Get command with given id
// @Tags         Commands
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Command's identifier"
// @Success      200  {object}  Command
// @Failure      500  {object}  Error
// @Router       /commands/{id} [get]
func (h *CommandHandler) GetCommand(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idInt, err := strconv.Atoi(id)
	command, err := h.service.GetCommand(r.Context(), idInt)
	if err != nil {
		jsonError := errorlib.GetJSONError("Failed to get command", err)
		w.WriteHeader(jsonError.Error.Code)
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	json.NewEncoder(w).Encode(command)
}

// RunCommand executes command with given id
// @Summary      Execute command with given id
// @Description  Execute command with given id
// @Tags         Commands
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Command's identifier"
// @Success      200  {object}  RunResponse
// @Failure      500  {object}  Error
// @Router       /commands/run/{id} [post]
func (h *CommandHandler) RunCommand(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	idInt, err := strconv.Atoi(id)
	command, err := h.service.RunCommand(r.Context(), idInt)
	if err != nil {
		jsonError := errorlib.GetJSONError("Failed to run command", err)
		w.WriteHeader(jsonError.Error.Code)
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	json.NewEncoder(w).Encode(command)
}

func (h *CommandHandler) SetRouter(router *mux.Router) {
	router.HandleFunc("/commands/list", h.ListCommands).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/commands/{id}", h.GetCommand).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/commands/create", h.CreateCommand).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/commands/run/{id}", h.RunCommand).Methods(http.MethodPost, http.MethodOptions)
}
