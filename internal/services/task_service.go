package services

import (
	"errors"
	"log"

	"github.com/google/uuid"

	"github.com/techymj/task-manager/internal/models"
	"github.com/techymj/task-manager/internal/repositories"
)

type TaskService struct {
	Repo  repositories.TaskRepository
	Queue chan string // task IDs for worker
}

func NewTaskService(repo repositories.TaskRepository, queue chan string) *TaskService {
	return &TaskService{
		Repo:  repo,
		Queue: queue,
	}
}

func (s *TaskService) CreateTask(title, desc, userID string) (*models.Task, error) {
	task := &models.Task{
		ID:          uuid.NewString(),
		Title:       title,
		Description: desc,
		Status:      "pending",
		UserID:      userID,
	}

	err := s.Repo.Create(task)
	if err != nil {
		return nil, err
	}

	// send to background worker
	log.Println("Sending to worker:", task.ID)
	s.Queue <- task.ID

	return task, nil
}

func (s *TaskService) GetTasks(userID, role, status string, page, limit int) ([]models.Task, error) {
	offset := (page - 1) * limit
	if role == "admin" {
		return s.Repo.GetAllWithFilter(status, limit, offset)
	}
	return s.Repo.GetByUserWithFilter(userID, status, limit, offset)
}

func (s *TaskService) GetTaskByID(id, userID, role string) (*models.Task, error) {
	task, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if role != "admin" && task.UserID != userID {
		return nil, errors.New("forbidden")
	}
	return task, nil
}

func (s *TaskService) DeleteTask(id, userID, role string) error {
	task, err := s.Repo.GetByID(id)
	if err != nil {
		return err
	}
	if role != "admin" && task.UserID != userID {
		return errors.New("forbidden")
	}
	return s.Repo.Delete(id)
}

func (s *TaskService) CompleteTaskManually(id, userID, role string) error {
	task, err := s.Repo.GetByID(id)
	if err != nil {
		return err
	}
	if role != "admin" && task.UserID != userID {
		return errors.New("forbidden")
	}
	return s.Repo.UpdateStatus(id, "completed")
}
