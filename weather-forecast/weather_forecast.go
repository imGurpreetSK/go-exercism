// Package weather describes the weather forecast for the country.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string

// CurrentLocation represents the current location.
var CurrentLocation string

// Forecast returns the weather forecast for the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
