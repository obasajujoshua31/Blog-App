package blogs

import (
	"github.com/graphql-go/graphql"
)

func GetQuery(resolver BlogResolver) map[string]*graphql.Field {
		allBlogs  := &graphql.Field{
			Type: graphql.NewList(QQLBlog),
			Resolve: resolver.GetAllBlogResolver,
		}
		blog:= &graphql.Field{
			Type: QQLBlog,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: resolver.BlogResolverByID,
		}

	return map[string]*graphql.Field{
		"allBlogs": allBlogs,
		"blog": blog,
	}

}
