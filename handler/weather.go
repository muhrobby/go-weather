package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/muhrobby/go-weather/model"
	"github.com/muhrobby/go-weather/utils"
)

func WeatherShow(c *fiber.Ctx) error {
	ApiKey, err := utils.LoadApiConfig(".apiConfig")
	city := c.Query("city")

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"data":   model.WeatherData{},
		})
	}

	resp, errG := http.Get("https://api.openweathermap.org/data/2.5/weather?appid=" + ApiKey.ApiKey + "&q=" + city)

	if errG != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed Fetch data",
		})
	}

	defer resp.Body.Close()

	var WeatherData model.WeatherData

	err = json.NewDecoder(resp.Body).Decode(&WeatherData)

	if err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "failed decode",
		})
	}
	WeatherData.Main.Metric = utils.Round(utils.KelvinToCelcius(WeatherData.Main.Metric))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   WeatherData,
	})

}
