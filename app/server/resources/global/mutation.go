package global

import (
	"blog-app/app/server/resources/blogs"
	"blog-app/app/server/resources/users"
	"blog-app/app/services"
	"github.com/graphql-go/graphql"
)

func getMutation(resolver services.Resolver) graphql.Fields {
	userResolver := users.UserResolver{Resolver:resolver}
	blogResolver := blogs.BlogResolver{Resolver:resolver}

	blogFields := blogs.GetMutation(blogResolver)
	userFields := users.GetMutation(userResolver)

	for key, value := range blogFields {
		userFields[key] = value
	}

	return userFields
}

