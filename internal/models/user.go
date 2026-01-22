package models

import "time"

type User struct {
	ID        string    `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"` // never return password
	Role      string    `json:"role" db:"role"`  // user or admin
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
