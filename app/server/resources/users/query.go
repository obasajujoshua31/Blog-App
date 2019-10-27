package users

import (
	"github.com/graphql-go/graphql"
)

func GetQuery(resolver UserResolver) (query map[string]*graphql.Field) {
		users:= &graphql.Field{
			Type: graphql.NewList(QQLUser),
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolver.UserResolverByName,
		}

			allUsers:= &graphql.Field{
				Type: graphql.NewList(QQLUser),
				Resolve: resolver.GetAllUserResolver,
		}

		user:= &graphql.Field{
			Type: QQLUser,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: resolver.UserResolverByID,
		}

	query = map[string]*graphql.Field{
		"user": user,
		"users": users,
		"allUsers": allUsers,
	}

	return query
}
