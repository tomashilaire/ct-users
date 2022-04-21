package entityhttphdl

import (
	"encoding/json"
	"entity/internal/core/domain"
	"entity/internal/core/ports"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPHandler struct {
	ts ports.EntityService
}

func NewHTTPHandler(ts ports.EntityService) *HTTPHandler {
	return &HTTPHandler{ts: ts}
}

// swagger:route GET /entities/{id} entities findEntity
// Returns a entity
// responses:
//	200: entityResponse
//	422: errorResponse

// GetEntity returns a entity document from the database
func (h *HTTPHandler) GetEntity(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	result, err := h.ts.ShowById(id)
	if err != nil {
		WriteAsJson(w, http.StatusUnprocessableEntity, err)
		return
	}
	WriteAsJson(w, http.StatusOK, result)
}

// swagger:route GET /entities entities listEntities
// Returns a list of entities
// responses:
//	200: entitiesResponse
//	422: errorResponse

// GetAllEntity returns the entity documents from the database
func (h *HTTPHandler) GetAllEntities(w http.ResponseWriter, r *http.Request) {
	result, err := h.ts.ShowAll()
	if err != nil {
		WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}

	WriteAsJson(w, http.StatusOK, result)
}

// swagger:route DELETE /entities/{id} entities deleteEntity
// Deletes a entity
// responses:
//	200: errorResponse
//  500: errorResponse

// DeleteEntity deletes an entity document from the database
func (h *HTTPHandler) DeleteEntity(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := h.ts.Delete(id)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	WriteAsJson(w, http.StatusOK, fmt.Sprintf("Succesfully deleted document of id: %s", id))
}

// swagger:route POST /entities entities createEntity
// Creates a entity
// responses:
//  200: idResponse
//  400: errorResponse
//  500: errorResponse

// PostEntity creates an entity document in the database
func (h *HTTPHandler) PostEntity(w http.ResponseWriter, r *http.Request) {
	var t domain.Entity
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

// swagger:route PUT /entities/{id} entities updateEntity
// Updates a entity
// responses:
//	200: idResponse
//	400: errorResponse
//	500: errorResponse

// PutEntity updates an entity document in the database
func (h *HTTPHandler) PutEntity(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var t domain.Entity
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
