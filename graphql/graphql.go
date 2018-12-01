package graphql

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/jon20/graphql-sample/graphql/fields"
)

func GraphqlSetting() http.Handler {
	query := fields.SetQuery()
	mutateQuery := fields.SetMutation()
	rootMutation := graphql.NewObject(graphql.ObjectConfig{Name: "RootMutation", Fields: mutateQuery})
	rootQuery := graphql.NewObject(graphql.ObjectConfig{Name: "RootQuery", Fields: query})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
	if err != nil {
		panic(err)
	}
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
	return h
}
