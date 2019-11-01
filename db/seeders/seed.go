package seeders

import (
	"blog-app/app/config"
	"blog-app/app/services"
	"log"
)

var Users = []services.User{
	services.User{
		Name:       "kelvin",
		Age:        35,
		Profession: "waiter",
		Friendly:   false,
	},
	services.User{
		Name:       "alex",
		Age:        26,
		Profession: "zoo keeper",
		Friendly:   false,
	},
	services.User{
		Name:       "becky",
		Age:        37,
		Profession: "retired",
		Friendly:   true,
	},

}

var Blogs = []services.Blog{
	services.Blog{
		Title:            "purpose",
		NumberOfComments: 1,
		Content:          "The purpose of lorem ipsum is to create a natural looking block of text (sentence, paragraph, page, etc.)",
		AuthorID:         2,

	},
	services.Blog{
		Title:            "purpose",
		NumberOfComments: 1,
		Content:          "The purpose of lorem ipsum is to create a natural looking block of text (sentence, paragraph, page, etc.)",
		AuthorID:         2,

	},
	services.Blog{
		Title:            "purpose",
		NumberOfComments: 1,
		Content:          "The purpose of lorem ipsum is to create a natural looking block of text (sentence, paragraph, page, etc.)",
		AuthorID:         1,

	},
	services.Blog{
		Title:            "purpose",
		NumberOfComments: 1,
		Content:          "The purpose of lorem ipsum is to create a natural looking block of text (sentence, paragraph, page, etc.)",
		AuthorID:         3,

	},
}

func Run() {
	appConfig, err := config.GetConfig()

	if err != nil {
		log.Fatal("Could not read app configuration")
	}

	db, err := services.ConnectToDB(appConfig)


	for index := range Users {
		err = db.Debug().Model(&services.User{}).Create(&Users[index]).Error

		if err != nil {
			log.Fatal("Error occured in seeding data", err)
		}

		//log.Println("Data seeded successfully")
	}

	for index := range Blogs {
		err = db.Debug().Model(&services.Blog{}).Create(&Blogs[index]).Error

		if err != nil {
			log.Fatal("Error occured in seeding data", err)
		}

		log.Println("Data seeded successfully")
	}


}