package users

import (
	"blog-app/app/services"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
)

type Resolver struct {
	db *services.DB
}

func (r *Resolver) UserResolver(p graphql.ResolveParams) (interface{}, error) {
	name, ok := p.Args["name"].(string)

	if ok {

		users, err := r.db.GetUserByName(name)

		if err != nil {
			return nil, err
		}
		return users, nil
	}

	return nil, nil
}

func (r *Resolver) UserResolverByID(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int)

	if ok {
		user, err := r.db.GetUserByID(id)

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

func (r *Resolver) CreateUserResolver(p graphql.ResolveParams) (interface{}, error) {
	user, ok := p.Args["newUser"].(map[string]interface{})

	if ok {
		newUser, err := r.db.CreateUser(user)

		if err != nil {
			return nil, err
		}

		return newUser, nil
	}

	return nil, nil

}

func (r *Resolver) GetAllUserResolver(p graphql.ResolveParams) (interface{}, error) {
	allUsers, err := r.db.GetAllUsers()

	if err != nil {
		return nil, err
	}

	return allUsers, nil
}