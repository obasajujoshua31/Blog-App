package users

import (
	"blog-app/app/services"
	"github.com/graphql-go/graphql"
)

type Root struct {
	Mutation *graphql.Object
	Query    *graphql.Object
}

func NewRoot(db *services.DB) *Root {
	resolver := Resolver{db}

	root := Root{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "Query",
			Fields: getQuery(resolver),
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "Mutation",
			Fields: getMutation(resolver),
		}),
	}
	return &root
}
