package users

import (
	"github.com/graphql-go/graphql"
)

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
			"Blogs": &graphql.Field{
				Type:             graphql.NewList(blog),
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

var blog = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Blogs",

		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"number_of_comments": &graphql.Field{
				Type: graphql.Int,
			},
			"content": &graphql.Field{
				Type: graphql.String,
			},
			"AuthorID": &graphql.Field{
				Type: graphql.Int,
				Name:   "author_id",
			},
		},
	})
