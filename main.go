package main

import (
	"fiber-mongo-api/configs"
	"log"
    "fiber-mongo-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    //run database connection
    configs.ConnectDB()

    //run routes
    routes.UserRoute(app)

    log.Fatal(app.Listen(":3000"))
}