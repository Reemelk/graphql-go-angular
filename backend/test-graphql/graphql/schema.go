package graphql

import (
	"web/test-graphql/graphql/queries"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var schemaConfig = graphql.SchemaConfig{
	Query: graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: queries.GetRootFields(),
	}),
	// Mutation: graphql.NewObject(graphql.ObjectConfig{
	// 	Name:   "RootMutation",
	// 	Fields: mutations.GetRootFields(),
	// }),
}

var graphqlSchema, _ = graphql.NewSchema(schemaConfig)

//Schema _
var Schema = handler.New(&handler.Config{
	Schema:     &graphqlSchema,
	Pretty:     true,
	GraphiQL:   true,
	Playground: true,
})
