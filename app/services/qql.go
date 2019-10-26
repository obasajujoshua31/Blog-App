package services



import (
	"fmt"
	"github.com/graphql-go/graphql"
	)


type Root struct {
	Query *graphql.Object
}



func ExecuteQuery(query string, schema graphql.Schema) * graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("Unexpected errors occured inside Execute Query: %v", result.Errors)
	}

	return result
}