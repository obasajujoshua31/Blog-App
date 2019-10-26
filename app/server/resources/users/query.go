package users

import (
	"blog-app/app/services"
	"github.com/graphql-go/graphql"
	)



type Root struct {
	Query *graphql.Object
}


func NewRoot (db *services.DB) *Root{
	resolver := Resolver{db}

	root := Root{Query: graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"users": &graphql.Field{
				Type: graphql.NewList(QQLUser),
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.UserResolver,
			},
		},
	},
	),
	}
	return &root
}