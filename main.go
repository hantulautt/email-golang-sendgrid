package main

import (
	"email/config"
	"email/controller"
	"email/exception"
	"email/repository"
	"email/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Setup Configuration
	configuration := config.New()
	//database := config.NewMongoDatabase(configuration)
	database := config.NewMysqlDatabase(configuration)

	// Setup Repository
	emailRepository := repository.NewEmailRepository(database)

	// Setup Service
	emailService := service.NewEmailService(&emailRepository)

	// Setup Controller
	emailController := controller.NewEmailController(&emailService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Setup Routing
	emailController.Route(app)

	// Start App
	err := app.Listen(":3000")
	exception.PanicIfNeeded(err)
}
