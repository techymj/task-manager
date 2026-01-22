package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"github.com/techymj/task-manager/internal/models"
	"github.com/techymj/task-manager/internal/services"
)

type AuthHandler struct {
	Service *services.AuthService
}

func NewAuthHandler(s *services.AuthService) *AuthHandler {
	return &AuthHandler{Service: s}
}

// Register godoc
// @Summary Register new user
// @Description Create a new user account
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body map[string]string true "User credentials"
// @Success 201 {string} string
// @Failure 400 {string} string
// @Router /register [post]
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	user := &models.User{
		ID:       uuid.NewString(),
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}

	err := h.Service.Register(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user registered"))
}

// Login godoc
// @Summary Login user
// @Description Login and get JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body map[string]string true "Login credentials"
// @Success 200 {object} map[string]string
// @Failure 401 {string} string
// @Router /login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	token, err := h.Service.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
