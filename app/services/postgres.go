package services

import (
	"blog-app/app/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	driverName = "postgres"
)

type DB struct {
	*sql.DB
}


func New(connString string) (*DB, error) {
	db, err := sql.Open(driverName, connString)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func ConnString(connString config.AppConfig) string {
	conn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", connString.Host, connString.Port, connString.User, connString.DBName, connString.Password)

	return conn
}

