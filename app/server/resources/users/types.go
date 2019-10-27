package users

import "github.com/graphql-go/graphql"

var QQLUser = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",

		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"age": &graphql.Field{
				Type: graphql.Int,
			},
			"profession": &graphql.Field{
				Type: graphql.String,
			},
			"friendly": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	})

var UserInput = graphql.InputObjectConfig{
	Name:        "newUser",
	Fields:      QQLUserInput,
	Description: "This is to create new Object",
}

var QQLUserInput = graphql.InputObjectConfigFieldMap{
	"name": &graphql.InputObjectFieldConfig{
		Type: graphql.String,
	},
	"age": &graphql.InputObjectFieldConfig{
		Type: graphql.Int,
	},
	"profession": &graphql.InputObjectFieldConfig{
		Type: graphql.String,
	},
	"friendly": &graphql.InputObjectFieldConfig{
		Type: graphql.Boolean,
	},
}
