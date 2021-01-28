package helpers

import (
	"echo-demo-project/server"
	"echo-demo-project/server/validation"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

func NewServer() *server.Server {
	s := &server.Server{
		Echo: echo.New(),
		DB:   Init(),
	}
	s.Echo.Validator = validation.NewCustomValidator(validator.New())

	return s
}
