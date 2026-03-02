/*
Package weather provides tools to tell the weather condition in a particular city.
*/
package weather

var (
    // CurrentCondition is a string representing the weather condition.
	CurrentCondition string
    // CurrentLocation is a string representing a location (city).
	CurrentLocation  string
)

// Forecast takes two strings, city and condition, and returns a string containing a statement about the current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
