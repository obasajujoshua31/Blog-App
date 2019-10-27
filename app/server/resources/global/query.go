
package global

import (
	"blog-app/app/server/resources/blogs"
	"blog-app/app/server/resources/users"
	"blog-app/app/services"
	"github.com/graphql-go/graphql"
)

func getQuery(resolver services.Resolver) graphql.Fields {

	blogResolver := blogs.BlogResolver{Resolver:resolver}
	userResolver := users.UserResolver{Resolver:resolver}

	blogFields := blogs.GetQuery(blogResolver)
	userFields := users.GetQuery(userResolver)


	for key, value := range blogFields {
		userFields[key] = value
	}

	return userFields
}
