package services

import (
	"blog-app/app/config"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

const (
	driverName = "postgres"
)

type DB struct {
	*gorm.DB
}

func New(connString string) (*DB, error) {
	//db, err := sql.Open(driverName, connString)

	db, err := gorm.Open(driverName, connString)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func connString(connString config.AppConfig) string {
	conn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=require", connString.Host, connString.Port, connString.User, connString.DBName, connString.Password)

	return conn
}


func ConnectToDB(appConfig config.AppConfig) (*DB, error) {

	db, err := New(
		connString(appConfig))

	if err != nil {
		return nil, err
	}

	return db, nil
}
