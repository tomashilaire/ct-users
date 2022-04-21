package testhttphdl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test/internal/core/domain"
	"test/internal/core/ports"

	"github.com/gorilla/mux"
)

type HTTPHandler struct {
	ts ports.TestService
}

func NewHTTPHandler(ts ports.TestService) *HTTPHandler {
	return &HTTPHandler{ts: ts}
}

func (h *HTTPHandler) GetTest(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	result, err := h.ts.ShowById(id)
	if err != nil {
		WriteAsJson(w, http.StatusUnprocessableEntity, err)
		return
	}
	WriteAsJson(w, http.StatusOK, result)
}

func (h *HTTPHandler) GetAllTests(w http.ResponseWriter, r *http.Request) {
	result, err := h.ts.ShowAll()
	if err != nil {
		WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}

	WriteAsJson(w, http.StatusOK, result)
}

func (h *HTTPHandler) DeleteTest(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := h.ts.Delete(id)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	WriteAsJson(w, http.StatusOK, fmt.Sprintf("Succesfully deleted document of id: %s", id))
}

func (h *HTTPHandler) PostTest(w http.ResponseWriter, r *http.Request) {
	var t domain.Test
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	result, err := h.ts.Create(t.Name, t.Action)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	WriteAsJson(w, http.StatusOK, result.Id)
}

func (h *HTTPHandler) PutTest(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var t domain.Test
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	result, err := h.ts.Update(id, t.Name, t.Action)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	WriteAsJson(w, http.StatusOK, result.Id)
}

// utils
type JError struct {
	Error string `json:"error"`
}

func WriteAsJson(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}

func WriteError(w http.ResponseWriter, statusCode int, err error) {
	e := "error"
	if err != nil {
		e = err.Error()
	}
	WriteAsJson(w, statusCode, JError{e})
}
