package blogs

import (
	"github.com/graphql-go/graphql"
)


func GetMutation(resolver BlogResolver) map[string]*graphql.Field {
	fields:=  &graphql.Field{
		Type: QQLBlog,
		Args: graphql.FieldConfigArgument{
			"newBlog": &graphql.ArgumentConfig{
				Type:        graphql.NewInputObject(BlogInput),
				Description: "This is create Blog mutation",
			},
		},
		Resolve: resolver.CreateBlogResolver,
	}

	return map[string]*graphql.Field{
		"createBlog": fields,
	}

}
