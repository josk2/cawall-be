package main

import (
	"cawall-be/graph"
	"cawall-be/graph/resolvers"
	"github.com/99designs/gqlgen/graphql/handler"
)

const defaultPort = "8080"

func initGraphqlServer(gqlPort string) *handler.Server {
	//port := os.Getenv("PORT")
	if gqlPort == "" {
		gqlPort = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{}}))

	//http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	//http.Handle("/query", srv)

	//log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	//log.Fatal(http.ListenAndServe(":"+port, nil))

	return srv
}
