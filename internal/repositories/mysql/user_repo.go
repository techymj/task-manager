package mysql

import (
	"database/sql"

	"github.com/techymj/task-manager/internal/models"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) Create(user *models.User) error {
	query := `INSERT INTO users (id,email,password,role) VALUES (?,?,?,?)`
	_, err := r.DB.Exec(query, user.ID, user.Email, user.Password, user.Role)
	return err
}

func (r *UserRepo) GetByEmail(email string) (*models.User, error) {
	row := r.DB.QueryRow(`SELECT id,email,password,role,created_at FROM users WHERE email=?`, email)
	var u models.User
	err := row.Scan(&u.ID, &u.Email, &u.Password, &u.Role, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepo) GetByID(id string) (*models.User, error) {
	row := r.DB.QueryRow(`SELECT id,email,password,role,created_at FROM users WHERE id=?`, id)
	var u models.User
	err := row.Scan(&u.ID, &u.Email, &u.Password, &u.Role, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
