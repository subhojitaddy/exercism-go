/* Package weather provides forecast of
   any city based on current condition.
*/
package weather

var (
    //CurrentCondition: Represents the current weather condition in Celsius as string.
    CurrentCondition string
    
    //CurrentLocation: Represents the current city as string.
	CurrentLocation  string
)

/* Forecast returns a string which is the concatenation of
   CurrentLocation and CurrentCondition.
*/
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
