package users

import "github.com/graphql-go/graphql"

func GetMutation(resolver UserResolver) map[string]*graphql.Field {
		fields:=  &graphql.Field{
			Type: QQLUser,
			Args: graphql.FieldConfigArgument{
				"newUser": &graphql.ArgumentConfig{
					Type:        graphql.NewInputObject(UserInput),
					Description: "This is create User mutation",
				},
			},
			Resolve: resolver.CreateUserResolver,
		}

		return map[string]*graphql.Field{
			"create": fields,
		}
}
