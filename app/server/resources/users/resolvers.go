package users

import (
	"blog-app/app/services"
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
