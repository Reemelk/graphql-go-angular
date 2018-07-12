package queries

import "github.com/graphql-go/graphql"

// GetRootFields all resolvers
func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"users": GetUsersQuery(),
		"user": GetUserQuery(),
	}
}
