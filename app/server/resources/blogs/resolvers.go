package blogs


import (
	"blog-app/app/services"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
)


type BlogResolver struct {
	services.Resolver
}


func (r *BlogResolver) BlogResolverByID(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int)

	if ok {
		blog, err := r.DB.GetBlogByID(id)

		if err != nil {
			return nil, err
		}

		if blog.Content == "" {

			err := fmt.Sprintf("no Blog Found for ID :%d", id)
			return nil, errors.New(err)
		}

		return blog, nil

	}

	return nil, nil

}

func (r *BlogResolver) CreateBlogResolver(p graphql.ResolveParams) (interface{}, error) {
	blog, ok := p.Args["newBlog"].(map[string]interface{})

	if ok {
		newBlog, err := r.DB.CreateBlog(blog)

		if err != nil {
			return nil, err
		}

		return newBlog, nil
	}

	return nil, nil

}

func (r *BlogResolver) GetAllBlogResolver(p graphql.ResolveParams) (interface{}, error) {
	allBlogs, err := r.DB.GetAllBlogs()

	if err != nil {
		return nil, err
	}
	return allBlogs, nil
}