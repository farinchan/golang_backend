package main

import (
	"github.com/farinchan/golang_backend/config"
	"github.com/farinchan/golang_backend/database/migration"
	"github.com/farinchan/golang_backend/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	//init Database
	config.ConnectDB()
	migration.RunMigration()

	app := fiber.New()

	//init Route
	routes.ApiRoute(app)
	

	app.Listen(":3000")
}
