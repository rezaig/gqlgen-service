package main

import (
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	"github.com/rezaindra/gqlgen-todos/internal/delivery/gqlsvc"
	"github.com/rezaindra/gqlgen-todos/internal/delivery/gqlsvc/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &gqlsvc.Resolver{}}))

	e := echo.New()
	e.Any("/", func(ctx echo.Context) error {
		playground.Handler("GraphQL playground", "/query").ServeHTTP(ctx.Response(), ctx.Request())
		return nil
	})
	e.Any("/query", func(ctx echo.Context) error {
		srv.ServeHTTP(ctx.Response(), ctx.Request())
		return nil
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	e.Logger.Fatal(e.Start(":" + port))
}
