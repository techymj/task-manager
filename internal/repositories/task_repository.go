package repositories

import "github.com/techymj/task-manager/internal/models"

type TaskRepository interface {
	Create(task *models.Task) error
	GetByID(id string) (*models.Task, error)
	GetAllByUser(userID string) ([]models.Task, error)
	GetAll() ([]models.Task, error)
	GetByUserWithFilter(userID, status string, limit, offset int) ([]models.Task, error)
	GetAllWithFilter(status string, limit, offset int) ([]models.Task, error)
	UpdateStatus(id string, status string) error
	Delete(id string) error
}
