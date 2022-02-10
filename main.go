package main

import (
	"fiber-mongo-api/configs"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    //run database connection
    configs.ConnectDB()

    log.Fatal(app.Listen(":3000"))
}