package config

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strconv"
)

type AppConfig struct {
	Host     string
	User     string
	Port     int
	Password string
	DBName   string
	AppPort  string
}

func GetConfig() (AppConfig, error) {

	//err := godotenv.Load()
	//
	//if err != nil {
	//	return AppConfig{}, err
	//}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		return AppConfig{}, err
	}

	config := AppConfig{
		Host:     os.Getenv("HOST"),
		User:     os.Getenv("DBUSER"),
		Port:     port,
		Password: os.Getenv("PASSWORD"),
		DBName:   os.Getenv("DBNAME"),
		AppPort:  fmt.Sprintf(":%s", os.Getenv("PORT")),
	}
	return config, nil

}
