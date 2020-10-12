
package main

import (
	"fmt"
	"main/app/dal"
	"main/app/routes"
	"main/config"
	"main/config/database"
	"main/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.Connect()
	database.Migrate(&dal.User{})

	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	app.Use(logger.New())

	routes.AuthRoutes(app)

	app.Listen(fmt.Sprintf(":%v", config.PORT))
}