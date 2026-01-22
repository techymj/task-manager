// @title Task Manager API
// @version 1.0
// @description Golang Task Management Service
// @host localhost:8080
// @BasePath /

package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/techymj/task-manager/internal/config"
	"github.com/techymj/task-manager/internal/database"
	"github.com/techymj/task-manager/internal/handlers"
	"github.com/techymj/task-manager/internal/repositories/mysql"
	"github.com/techymj/task-manager/internal/routes"
	"github.com/techymj/task-manager/internal/services"

	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/techymj/task-manager/docs"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.Load()
	log.Println("Auto complete minutes:", cfg.AutoMin)

	db := database.Connect(cfg)
	defer db.Close()

	taskQueue := make(chan string, 100)

	taskRepo := mysql.NewTaskRepo(db)
	userRepo := mysql.NewUserRepo(db)

	authService := services.NewAuthService(userRepo, cfg.JWTKey)
	taskService := services.NewTaskService(taskRepo, taskQueue)

	worker := services.NewWorker(taskRepo, taskQueue, cfg.AutoMin)
	worker.Start()

	authHandler := handlers.NewAuthHandler(authService)
	taskHandler := handlers.NewTaskHandler(taskService)

	mux := http.NewServeMux()
	routes.Register(mux, authHandler, taskHandler, cfg.JWTKey)

	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
