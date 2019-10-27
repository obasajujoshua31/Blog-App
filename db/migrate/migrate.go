package main

import (
	"fmt"

	"blog-app/app/config"
	"blog-app/app/services"
	"io/ioutil"
	"log"
)

const (
	userFileName = "db/migrate/user.sql"
	blogFileName = "db/migrate/blog.sql"
	dropFileName = "db/migrate/drop.sql"
)


func main() {
	appConfig, err := config.GetConfig()

	if err != nil {
		log.Fatal("Unable to get application config")
	}

	db, err := services.ConnectToDB(appConfig)

	if err != nil {
		log.Fatal("Server was unable to connect to Database", err)
	}

	userSQL, err := ioutil.ReadFile(userFileName)
	blogSQL, err := ioutil.ReadFile(blogFileName)

	//dropTables(db)

	if err != nil {
		log.Fatal("Error occured while reading file", err)
	}

	_, err = db.Exec(string(userSQL))
	_, err = db.Exec(string(blogSQL))

	if err != nil {
		log.Fatal("Unable to create user migration")
	}

	fmt.Println("Rows created successfully ...")

}

func dropTables (db *services.DB) {

	dropSQL, err := ioutil.ReadFile(dropFileName)

	if err != nil {
		log.Fatal("Unable to read drop sql", err)
	}

	_, err = db.Exec(string(dropSQL))

	if err != nil {
		log.Fatal("Unable to drop tables", err)
	}

}