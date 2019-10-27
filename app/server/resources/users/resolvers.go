package users

import (
	"blog-app/app/services"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
)

type UserResolver struct {
	services.Resolver
}

func (r *UserResolver) UserResolverByName(p graphql.ResolveParams) (interface{}, error) {
	name, ok := p.Args["name"].(string)

	if ok {

		users, err := r.DB.GetUserByName(name)

		if err != nil {
			return nil, err
		}
		return users, nil
	}

	return nil, nil
}

func (r *UserResolver) UserResolverByID(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int)

	if ok {
		user, err := r.DB.GetUserByID(id)

		if err != nil {
			return nil, err
		}

		if user.Profession == "" {

			err := fmt.Sprintf("no User Found for ID :%d", id)
			return nil, errors.New(err)
		}

		return user, nil

	}

	return nil, nil

}

func (r *UserResolver) CreateUserResolver(p graphql.ResolveParams) (interface{}, error) {
	user, ok := p.Args["newUser"].(map[string]interface{})

	if ok {
		newUser, err := r.DB.CreateUser(user)

		if err != nil {
			return nil, err
		}

		return newUser, nil
	}

	return nil, nil

}

func (r *UserResolver) GetAllUserResolver(p graphql.ResolveParams) (interface{}, error) {
	allUsers, err := r.DB.GetAllUsers()

	if err != nil {
		return nil, err
	}

	return allUsers, nil
}