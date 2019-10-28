package main

import (
	"blog-app/app/config"
	"blog-app/app/services"
	"log"
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



	err = db.Debug().DropTableIfExists(&services.Blog{}, &services.User{}).Error

	if err != nil {
		log.Fatal("Unable to drop table tables", err)
	}

	err = db.Debug().AutoMigrate(&services.Blog{}, &services.User{}).Error

	if err != nil {
		log.Fatal("Migration was not successful", err)
	}

	err = db.Debug().Model(&services.Blog{}).AddForeignKey("author_id",
		"users(id)", "RESTRICT", "RESTRICT").Error

	if err != nil {
		log.Fatal("Unable to add Foreign Key",err)
	}

	log.Println("Migration successful")

}
