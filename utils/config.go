package utils

import (
	"encoding/json"
	"math"
	"os"
)

type apiConfigData struct {
	ApiKey string `json:"apikey"`
}

func LoadApiConfig(filename string) (apiConfigData, error) {

	bytes, err := os.ReadFile(filename)

	if err != nil {
		return apiConfigData{}, err
	}

	var c apiConfigData

	err = json.Unmarshal(bytes, &c)

	if err != nil {
		return c, err
	}

	return c, nil

}

func KelvinToCelcius(kelvin float64) float64 {
	return kelvin - 273.15
}

func Round(num float64) float64 {
	return math.Round(num*100) / 100
}
