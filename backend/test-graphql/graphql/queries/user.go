package queries

import (
	"strconv"
	"web/test-graphql/graphql/types"
	. "web/test-graphql/models"

	"github.com/graphql-go/graphql"
)

var testUsers []User

func init() {
	testUsers = []User{
		User{
			ID:        1,
			Lastname:  "lesieur",
			Firstname: "julien",
		},
		User{
			ID:        2,
			Lastname:  "el hormi",
			Firstname: "nassim",
		},
	}
}

// GetUsersQuery users
func GetUsersQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.UserType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var users []User
			users = testUsers
			return users, nil
		},
	}
}

// GetUserQuery user
func GetUserQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.UserType),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var users []User
			userID, isOK := p.Args["id"].(string)
			if isOK {
				for _, u := range testUsers {
					if userID == strconv.Itoa(u.ID) {
						users = append(users, u)
						return users, nil
					}
				}
			}
			return nil, nil
		},
	}
}
