package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhrobby/go-weather/router"
)

func main() {

	app := fiber.New()

	router.Router(app)

	app.Listen(":3030")

}
