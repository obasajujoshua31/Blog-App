package main


import (
	"blog-app/app/config"
	"blog-app/app/services"
	"fmt"
	"io/ioutil"
	"log"
)

const (
	userFileName = "db/seeders/users.sql"
	blogFileName = "db/seeders/blog.sql"
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

	if err != nil {
		log.Fatal("Error occured while reading file", err)
	}

	_, err = db.Exec(string(userSQL))


	if err != nil {
		log.Fatal("Unable to create user migration", err)
	}

	_, err = db.Exec(string(blogSQL))

	if err != nil {
		log.Fatal("Unable to create user migration", err)
	}

	fmt.Println("Seed Created successfully ...")

}
