package routes

import (
	"github.com/farinchan/golang_backend/controller/api"
	"github.com/gofiber/fiber/v2"
)

func ApiRoute(app *fiber.App) {
	apiRoute := app.Group("/api")

	v1 := apiRoute.Group("/v1")

	auth := v1.Group("/auth")
	auth.Post("/register", api.Register)
	auth.Get("/login", api.Login)
	auth.Post("/forgot-password", api.ForgotPassword)

	user := v1.Group("/user")
	user.Post("/user", api.UserCreate)
	user.Get("/user", api.UserIndex)
	user.Get("/user/:userId", api.UserShow)
	user.Put("/user/:userId", api.UserUpdate)
	user.Delete("/user/:userId", api.UserDelete)
}
