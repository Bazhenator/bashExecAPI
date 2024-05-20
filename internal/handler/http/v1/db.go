package v1

import (
	"encoding/json"
	"github.com/Bazhenator/bashExecAPI/internal/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type DataBaseHandler struct {
	service *service.Services
}

func NewDataBaseHandler(service *service.Services) *DataBaseHandler {
	return &DataBaseHandler{
		service: service,
	}
}

// DeleteAllRows deletes all rows in table Commands
// @Summary      Delete all rows in table Commands
// @Description  Delete all rows in table Commands
// @Tags         DataBase
// @Accept       json
// @Produce      json
// @Success      200  {object}  DeleteAllRowsResponse
// @Failure      500  {object}  Error
// @Router       /commands/delete [delete]
func (h *DataBaseHandler) DeleteAllRows(w http.ResponseWriter, r *http.Request) {
	err := h.service.DeleteAllRows(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("all rows have been successfully deleted")
}

// DeleteRow deletes row with given id in table Commands
// @Summary      Delete row with given id in table Commands
// @Description  Delete row with given id in table Commands
// @Tags         DataBase
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Row's identifier"
// @Success      200  {object}  DeleteRowResponse
// @Failure      500  {object}  Error
// @Router       /commands/delete/{id} [delete]
func (h *DataBaseHandler) DeleteRow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.DeleteRow(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("row " + strconv.Itoa(id) + " has been successfully deleted")
}

func (h *DataBaseHandler) SetRouter(router *mux.Router) {
	router.HandleFunc("/commands/delete", h.DeleteAllRows).Methods(http.MethodDelete, http.MethodOptions)
	router.HandleFunc("/commands/delete/{id}", h.DeleteRow).Methods(http.MethodDelete, http.MethodOptions)
}
