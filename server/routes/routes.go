package routes

import (
	"cawall-be/responses"
	s "cawall-be/server"
	"cawall-be/server/handlers"
	"cawall-be/services/token"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func ConfigureRoutes(server *s.Server, graphQLServer *handler.Server) {
	authHandler := handlers.NewAuthHandler(server)
	registerHandler := handlers.NewRegisterHandler(server)

	server.Echo.Use(middleware.Logger())

	//bind graphql
	server.Echo.GET("/playground", echo.WrapHandler(playground.Handler("GraphQL playground", "/query")))
	server.Echo.POST("/graphql", echo.WrapHandler(graphQLServer))
	server.Echo.GET("/query", echo.WrapHandler(graphQLServer))

	server.Echo.GET("/swagger/*", echoSwagger.WrapHandler)

	server.Echo.POST("/login", authHandler.Login)
	server.Echo.POST("/register", registerHandler.Register)
	server.Echo.POST("/refresh", authHandler.RefreshToken)

	fmt.Println(server.Config.Auth.AccessSecret)
	r := server.Echo.Group("")
	r.GET("/health", dummyHealth)
	// Configure middleware with the custom claims type
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(token.JwtCustomClaims)
		},
		SigningKey: []byte(server.Config.Auth.AccessSecret),
	}
	r.Use(echojwt.WithConfig(config))

	//postHandler := handlers.NewPostHandlers(server)
	//r.GET("/posts", postHandler.GetPosts)
	//r.POST("/posts", postHandler.CreatePost)
	//r.DELETE("/posts/:id", postHandler.DeletePost)
	//r.PUT("/posts/:id", postHandler.UpdatePost)
	//r.PUT("/posts/:id", postHandler.UpdatePost)
}

func dummyHealth(c echo.Context) error {
	return responses.Response(c, http.StatusOK, map[string]interface{}{"health": true})
}
