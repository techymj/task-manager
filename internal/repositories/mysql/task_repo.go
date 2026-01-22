package mysql

import (
	"database/sql"

	"github.com/techymj/task-manager/internal/models"
)

type TaskRepo struct {
	DB *sql.DB
}

func NewTaskRepo(db *sql.DB) *TaskRepo {
	return &TaskRepo{DB: db}
}

func (r *TaskRepo) Create(task *models.Task) error {
	query := `INSERT INTO tasks (id,title,description,status,user_id) VALUES (?,?,?,?,?)`
	_, err := r.DB.Exec(query, task.ID, task.Title, task.Description, task.Status, task.UserID)
	return err
}

func (r *TaskRepo) GetByID(id string) (*models.Task, error) {
	row := r.DB.QueryRow(`SELECT id,title,description,status,user_id,created_at,updated_at FROM tasks WHERE id=?`, id)
	var t models.Task
	err := row.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.UserID, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *TaskRepo) GetAllByUser(userID string) ([]models.Task, error) {
	rows, err := r.DB.Query(`SELECT id,title,description,status,user_id,created_at,updated_at FROM tasks WHERE user_id=?`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.UserID, &t.CreatedAt, &t.UpdatedAt)
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *TaskRepo) GetAll() ([]models.Task, error) {
	rows, err := r.DB.Query(`SELECT id,title,description,status,user_id,created_at,updated_at FROM tasks`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.UserID, &t.CreatedAt, &t.UpdatedAt)
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *TaskRepo) GetByUserWithFilter(userID, status string, limit, offset int) ([]models.Task, error) {
	query := `SELECT id,title,description,status,user_id,created_at,updated_at 
	          FROM tasks WHERE user_id=?`
	args := []interface{}{userID}

	if status != "" {
		query += " AND status=?"
		args = append(args, status)
	}
	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.UserID, &t.CreatedAt, &t.UpdatedAt)
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *TaskRepo) GetAllWithFilter(status string, limit, offset int) ([]models.Task, error) {
	query := `SELECT id,title,description,status,user_id,created_at,updated_at FROM tasks`
	args := []interface{}{}

	if status != "" {
		query += " WHERE status=?"
		args = append(args, status)
	}
	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.UserID, &t.CreatedAt, &t.UpdatedAt)
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *TaskRepo) UpdateStatus(id string, status string) error {
	_, err := r.DB.Exec(`UPDATE tasks SET status=? WHERE id=?`, status, id)
	return err
}

func (r *TaskRepo) Delete(id string) error {
	_, err := r.DB.Exec(`DELETE FROM tasks WHERE id=?`, id)
	return err
}
