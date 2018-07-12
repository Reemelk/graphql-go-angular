package types

import "github.com/graphql-go/graphql"

// UserType Schema
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.String},
		"lastname":  &graphql.Field{Type: graphql.String},
		"firstname": &graphql.Field{Type: graphql.String},
	},
})
