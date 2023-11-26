package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhrobby/go-weather/handler"
)

func Router(app *fiber.App) {
	api := app.Group("api")
	api.Get("/wheater/", handler.WeatherShow)

}
