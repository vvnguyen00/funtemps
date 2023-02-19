package conv

import (
	"testing"
	"math"
)

/*
*

	Mal for testfunksjoner
	Du skal skrive alle funksjonene basert på denne malen
	For alle konverteringsfunksjonene (tilsammen 6)
	kan du bruke malen som den er (du må selvsagt endre
	funksjonsnavn og testverdier)
*/
func TestFahrenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 56.67},
	}

	for _, tc := range tests {
		got := FahrenheitToCelsius(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

// De andre testfunksjonene implementeres her
// ...

func TestCelsiusToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 20, want: 68},
	}

	for _, tc := range tests {
		got := CelsiusToFahrenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

//Test Kelvin til Fahrenheit


func TestKelvinToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 10, want: -441.67},
	}

	for _, tc := range tests {
		got := KelvinToFahrenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

//Farhenheit til Kelvin
func TestFahrenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 5, want: 258.15},
	}

	for _, tc := range tests {
		got := FahrenheitToKelvin(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

//konverterer Celcius til Kelvin

func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 3, want: 276.15},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

////konverterer Kelvin til Celcius {

func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 2, want: -271.15},
	}

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)
		if !withinTolerance(tc.want, got, 1e-2) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}





//withinTolerance

func withinTolerance(a, b, error float64) bool {
	// Først sjekk om tallene er nøyaktig like
	if a == b {
	  return true
	}
  
	difference := math.Abs(a - b)
  
	// Siden vi skal dele med b, må vi sjekke om den er 0
	// Hvis b er 0, returner avgjørelsen om d er mindre enn feilmarginen
	// som vi aksepterer
	if b == 0 {
	  return difference < error
	}
  
	// Tilslutt sjekk den relative differanse mot feilmargin
	return (difference/math.Abs(b)) < error
  }