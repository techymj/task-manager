package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/techymj/task-manager/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(cfg *config.Config) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("DB open error:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("DB ping error:", err)
	}

	log.Println("Connected to MySQL")
	return db
}
