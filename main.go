package main

import "github.com/labstack/echo"

func main() {
	route := echo.New()

	route.POST("/graphql", echo.WrapHandler())
	route.GET("/graphql", echo.WrapHandler())
}
