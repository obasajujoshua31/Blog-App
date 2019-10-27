package server

import (
	"blog-app/app/config"
	"blog-app/app/services"
	"fmt"
	"github.com/graphql-go/graphql"
	"net/http"

	"blog-app/app/server/resources/users"
)

func Start() error {
	appConfiguration, err := config.GetConfig()

	if err != nil {
		return err
	}

	db, err := services.ConnectToDB(appConfiguration)

	if err != nil {
		return err
	}

	rootQuery := users.NewRoot(db)

	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    rootQuery.Query,
			Mutation: rootQuery.Mutation,
		})

	if err != nil {
		fmt.Println("Error creating Schema", err)
	}

	server := NewServer(&appConfiguration, &sc)

	server.InitMiddlewares()

	err = server.setUpGraphiQL()

	if err != nil {
		return  err
	}

	defer db.Close()

	fmt.Printf("Starting GraphQL Server on port %s .....", appConfiguration.AppPort)

	return http.ListenAndServe(appConfiguration.AppPort, nil)
}
