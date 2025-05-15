package main

import (
	"github.com/farinchan/golang_backend/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//init Route
	routes.ApiRoute(app)

	app.Listen(":3000")
}
