package conv

import "fmt"

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...

		1. Celsius = Kelvin - 273.15
	  2. Kelvin = Celsius + 273.15
	  3. Farhrenheit = Celsius*(9/5) + 32
	  4. Celsius = (Farhrenheit - 32)*(5/9)
	  5. Farhrenheit = (Kelvin - 273.15)*(9/5) + 32
	  6. Kelvin = (Farhenheit - 32) * (5/9) + 273.15
*/


// Konverterer Farhenheit til Celsius
func FahrenheitToCelsius(value float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
	return (value - 32)*(5.0/9.0)
}

func CelsiusToFahrenheit(value float64) float64 {
	return value*(9.0/5.0) + 32
}

func CelsiusToKelvin(value float64) float64 {
	return value + 273.15
}

func KelvinToCelsius(value float64) float64 {
	return value - 273.15
}

func KelvinToFahrenheit(value float64) float64 {
	return (value - 273.15)*(9.0/5.0) + 32
}

func FahrenheitToKelvin(value float64) float64 {
	return (value - 32) * (5.0/9.0) + 273.15
}

/*
%f     default width, default precision
%9f    width 9, default precision
%.2f   default width, precision 2
%9.2f  width 9, precision 2
%9.f   width 9, precision 0
*/
func FormatBigInteger(number float64) string {

	s := fmt.Sprintf("%.f", number)

	var sc string
	var scc string
	var neg bool

	for i := len(s)-1; i >= 0; i-- {
    if string(s[i]) == "-" {
			neg = true
			continue
		}
		if len(sc) < 3 {
			sc = string(s[i]) + sc
			//fmt.Println("IF: i", i, "sc = ", sc, "len(sc) = ", len(sc))
	  } else {
			//fmt.Println("ELSE START: scc = ", scc, "sc = ", sc)
			scc = sc + " " + scc
			sc = string(s[i])
			//fmt.Println("ELSE END: scc = ", scc, "sc = ", sc)
		}
	}
  //fmt.Println("len(sc) = ", len(sc), "sc = ", sc, "scc = ", scc)
	if neg {
		sc = "-" + sc
	}
	if len(scc) > 0 {
		return sc + " " + scc[:len(scc)-1]
	} else {
		return sc
	}
}
