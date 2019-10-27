package blogs

import (
	"blog-app/app/server/resources/users"
	"github.com/graphql-go/graphql"
)


var QQLBlog = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Blog",

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
			"User": &graphql.Field{
				Type:       users.QQLUser,
				Name:		"author",
			},
		},
	})

var BlogInput = graphql.InputObjectConfig{
	Name:        "newBlog",
	Fields:      QQLBlogInput,
	Description: "This is to create new Object",
}

var QQLBlogInput = graphql.InputObjectConfigFieldMap{
	"title": &graphql.InputObjectFieldConfig{
		Type: graphql.String,
	},
	"content": &graphql.InputObjectFieldConfig{
		Type: graphql.String,
	},
	"author_id": &graphql.InputObjectFieldConfig{
		Type: graphql.Int,
	},
}

