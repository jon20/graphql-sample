package main

import (
	"github.com/jon20/graphql-sample/graphql"
	"github.com/labstack/echo"
)

func main() {
	route := echo.New()

	route.POST("/graphql", echo.WrapHandler(graphql.GraphqlSetting()))
	route.GET("/graphql", echo.WrapHandler(graphql.GraphqlSetting()))

	route.Logger.Fatal(route.Start("localhost:8080"))
}
