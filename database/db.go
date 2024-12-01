package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

// InitDB инициализирует подключение к БД.
func InitDB() (*sqlx.DB, error) {
	connStr := "user=username password=password dbname=yourdbname sslmode=disable"

	var err error
	db, err = sqlx.Connect("postgress", connStr)
	return db, err
}
