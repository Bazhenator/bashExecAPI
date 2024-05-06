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

func (h *DataBaseHandler) DeleteAllRows(w http.ResponseWriter, r *http.Request) {
	err := h.service.DeleteAllRows(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("all rows have been successfully deleted")
}

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
	router.HandleFunc("/commands/delete", h.DeleteAllRows).Methods(http.MethodDelete)
	router.HandleFunc("/commands/delete/{id}", h.DeleteRow).Methods(http.MethodDelete)
}
