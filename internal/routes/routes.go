package routes

import (
	"net/http"

	"github.com/techymj/task-manager/internal/handlers"
	"github.com/techymj/task-manager/internal/middleware"
)

func Register(
	mux *http.ServeMux,
	auth *handlers.AuthHandler,
	task *handlers.TaskHandler,
	jwtKey string,
) {

	mux.HandleFunc("/register", auth.Register)
	mux.HandleFunc("/login", auth.Login)

	protected := middleware.JWTAuth(jwtKey)

	// /tasks -> GET or POST
	mux.Handle("/tasks", protected(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			task.CreateTask(w, r)
		} else if r.Method == http.MethodGet {
			task.GetTasks(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})))

	// /tasks/{id}
	mux.Handle("/tasks/", protected(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			task.GetTaskByID(w, r)
		} else if r.Method == http.MethodDelete {
			task.DeleteTask(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})))
}
