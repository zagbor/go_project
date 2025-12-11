package handlers

import (
	"encoding/json"
	"go-microservice/models"
	"go-microservice/services"
	"go-microservice/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// UserHandler holds methods to handle user requests.
type UserHandler struct {
	Service *services.UserService
}

// NewUserHandler creates a new handler with the given service.
func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

// CreateUser handles POST /api/users
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	savedUser := h.Service.Create(user)

	// Async logging
	go utils.LogUserAction("CREATE", savedUser.ID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(savedUser)
}

// GetUser handles GET /api/users/{id}
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.Service.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// GetAllUsers handles GET /api/users
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := h.Service.GetAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// UpdateUser handles PUT /api/users/{id}
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedUser, err := h.Service.Update(id, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	go utils.LogUserAction("UPDATE", id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

// DeleteUser handles DELETE /api/users/{id}
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := h.Service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	go utils.LogUserAction("DELETE", id)

	w.WriteHeader(http.StatusNoContent)
}
