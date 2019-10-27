package users

import (
	"github.com/graphql-go/graphql"
)

func getQuery(resolver Resolver) graphql.Fields {

	fields := graphql.Fields{
		"users": &graphql.Field{
			Type: graphql.NewList(QQLUser),
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolver.UserResolver,
		},
		"user": &graphql.Field{
			Type: QQLUser,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: resolver.UserResolverByID,
		},
	}
	return fields

}
