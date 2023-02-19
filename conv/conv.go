package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

//Konverterer farhenheit til Celcius
func FahrenheitToCelsius(value float64) float64{
  return (value - 32.0) * (5.0/9.0)
}

//Konverterer Celcius til Farhenheit
func CelsiusToFahrenheit(value float64) float64 {
  return (value * 9/5) + 32
}

//konverterer kelvin til farhenheit
func KelvinToFahrenheit(value float64) float64 {
  return (value-273.15)*9/5 + 32
}

//konverterer Farhenheit til Kelvin
func FahrenheitToKelvin(value float64) float64 {
  return (value-32)*5/9 + 273.15
}

//konverterer Celcius til Kelvin
func CelsiusToKelvin(value float64) float64 {
  return (value + 273.15)
}

//konverterer Kelvin til Celcius 
func KelvinToCelsius(value float64) float64 {
  return (value-273.15)
}