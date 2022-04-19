package testhttphdl

import (
	"encoding/json"
	"net/http"
	"test/internal/core/ports"

	"github.com/gorilla/mux"
)

type HTTPHandler struct {
	ts ports.TestService
}

func NewHTTPHandler(ts ports.TestService) *HTTPHandler {
	return &HTTPHandler{
		ts: ts,
	}
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
	panic("unimplemented")
}

func (h *HTTPHandler) PostTest(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func (h *HTTPHandler) PutTest(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
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
