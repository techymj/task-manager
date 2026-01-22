package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/techymj/task-manager/internal/middleware"
	"github.com/techymj/task-manager/internal/services"
)

type TaskHandler struct {
	Service *services.TaskService
}

func NewTaskHandler(s *services.TaskService) *TaskHandler {
	return &TaskHandler{Service: s}
}

// CreateTask godoc
// @Summary Create task
// @Description Create a new task
// @Tags Tasks
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param task body map[string]string true "Task payload"
// @Success 200 {object} models.Task
// @Failure 401 {string} string
// @Router /tasks [post]
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	claimsVal := r.Context().Value(middleware.UserContextKey)
	if claimsVal == nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	claims := claimsVal.(jwt.MapClaims)

	userID := claims["user_id"].(string)

	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	task, err := h.Service.CreateTask(req.Title, req.Description, userID)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(task)
}

// GetTasks godoc
// @Summary List tasks
// @Description Get tasks with pagination and filtering
// @Tags Tasks
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Param status query string false "pending | in_progress | completed"
// @Success 200 {array} models.Task
// @Failure 401 {string} string
// @Router /tasks [get]
func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	claimsVal := r.Context().Value(middleware.UserContextKey)
	if claimsVal == nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	claims := claimsVal.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	// Query params
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	status := r.URL.Query().Get("status")

	// Defaults
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	tasks, err := h.Service.GetTasks(userID, role, status, page, limit)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

// GetTaskByID godoc
// @Summary Get task by ID
// @Description Get single task
// @Tags Tasks
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Task ID"
// @Success 200 {object} models.Task
// @Failure 403 {string} string
// @Router /tasks/{id} [get]
func (h *TaskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	claimsVal := r.Context().Value(middleware.UserContextKey)
	if claimsVal == nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	claims := claimsVal.(jwt.MapClaims)

	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	task, err := h.Service.GetTaskByID(id, userID, role)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}
	json.NewEncoder(w).Encode(task)
}

// DeleteTask godoc
// @Summary Delete task
// @Description Delete a task
// @Tags Tasks
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Task ID"
// @Success 200 {string} string
// @Failure 403 {string} string
// @Router /tasks/{id} [delete]
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	claimsVal := r.Context().Value(middleware.UserContextKey)
	if claimsVal == nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	claims := claimsVal.(jwt.MapClaims)

	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	err := h.Service.DeleteTask(id, userID, role)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}
	w.Write([]byte("deleted"))
}
