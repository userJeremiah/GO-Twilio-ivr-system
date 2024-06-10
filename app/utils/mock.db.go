package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// Weather reports
var WeatherConditions = []string{
	"Sunny, high of 75 degrees.",
	"Cloudy, high of 68 degrees.",
	"Partly cloudy, high of 72 degrees.",
	"Foggy, high of 64 degrees.",
	"Sunny, high of 80 degrees.",
	"Clear skies, high of 78 degrees.",
	"Breezy, high of 74 degrees.",
	"Mild, high of 70 degrees.",
	"Hot, high of 85 degrees.",
	"Pleasant, high of 76 degrees.",
}

var Cities = []string{
	"Los Angeles",
	"San Diego",
	"San Jose",
	"San Francisco",
	"Fresno",
	"Sacramento",
	"Long Beach",
	"Oakland",
	"Bakersfield",
	"Anaheim",
}

func GetCityWeather(index int) string {
	if index < 0 || index >= len(Cities) {
		return "Invalid city index."
	}

	city := Cities[index]
	weatherReport := GetRandomWeatherReport()
	return fmt.Sprintf("The weather in %s: %s", city, weatherReport)
}

// getRandomWeatherReport returns a random weather report
func GetRandomWeatherReport() string {
	rand.Seed(time.Now().UnixNano())
	return WeatherConditions[rand.Intn(len(WeatherConditions))]
}
