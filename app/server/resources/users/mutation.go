package users

import "github.com/graphql-go/graphql"

func getMutation(resolver Resolver) graphql.Fields {
	return graphql.Fields{
		"create": &graphql.Field{
			Type: QQLUser,
			Args: graphql.FieldConfigArgument{
				"newUser": &graphql.ArgumentConfig{
					Type:        graphql.NewInputObject(UserInput),
					Description: "This is create User mutation",
				},
			},
			Resolve: resolver.CreateUserResolver,
		},
	}
}
