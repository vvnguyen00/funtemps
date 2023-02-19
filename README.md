Klarte en liten grad å utføre dette på terminalen. 

funtemps -K 273.15 -out C
Output: 273.15K er 0°C
funtemps -F 32 -out C
Output: 32°F er 0°C
funtemps -C -89.4 -out F
Output: -89.4°C er -128.92°F
funtemps -C -89.4 -out K
Output: -89.4°C er -183.75K



Jeg gjorde dette først på main.go. Det fungerte og, men jeg skrev heller det som chat.openAi anbefalte istedenfor.
func main() {

    flag.Parse()

  //fahrenheit to celcius
    if out == "C" && isFlagPassed("F") {
        fmt.Println(conv.FahrenheitToCelsius(fahr))
    }


//Celsiustofarenheit
    if out == "F" && isFlagPassed("C") {
        fmt.Println(conv.CelsiusToFahrenheit(celsius))
    }

//kelvin til farenheit
if out == "F" && isFlagPassed("K") {
    fmt.Println(conv.KelvinToFahrenheit(fahr))
}

//farenheit til Kelvin  

if out == "K" && isFlagPassed("F") {
    fmt.Println(conv.FahrenheitToKelvin(fahr))
}

//Celsius til kelvin

if out == "K" && isFlagPassed("C") {
    fmt.Println(conv.CelsiusToKelvin(celsius))
}

//konverterer Kelvin til Celcius
if out == "C" && isFlagPassed("K") {
    fmt.Println(conv.CelsiusToKelvin(fahr))
}

}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
    found := false
    flag.Visit(func(f *flag.Flag) {
        if f.Name == name {
            found = true
        }
    })
    return found
}



Klarte å utføre greit med test-drevet utvikling.


‌
