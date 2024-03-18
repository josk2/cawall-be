package main

import (
	"cawall-be/config"
	"cawall-be/graph"
	"cawall-be/server"
	"cawall-be/server/routes"
	"github.com/99designs/gqlgen/graphql/handler"
	"log"
)

func Start(cfg *config.Config) {
	app := server.NewServer(cfg)

	//graphQLServer := initGraphqlServer("8080")
	graphQLServer := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	routes.ConfigureRoutes(app, graphQLServer)

	err := app.Start(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
