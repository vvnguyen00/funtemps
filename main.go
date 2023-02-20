package main

import (
	"flag"
	"fmt"
	"github.com/vvnguyen00/funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var kelvin float64
var celsius float64
var out string
var funfacts string
var tempScale string



// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	// Du må selv definere flag-variablene for "C" og "K"
	flag.Float64Var(&kelvin, "K", 0.0, "temperatur i grader Kelvin")
	flag.Float64Var(&celsius, "C", 0.0, "temperatur i grader Celsius")


	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises
	flag.StringVar(&tempScale, "t", "C", "temperaturskala for funfacts: C - Celsius, F - Fahrenheit, K - Kelvin")


}

func main() {

	flag.Parse()


	/**
	    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
	    pakkene implementeres.

	    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
	    av flagg. flag-pakken har funksjoner som man kan bruke for å teste
	    hvor mange flagg og argumenter er spesifisert på kommandolinje.

	        fmt.Println("len(flag.Args())", len(flag.Args()))
			    fmt.Println("flag.NFlag()", flag.NFlag())

	    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
	    brukes for å utelukke ugyldige kombinasjoner:
	    -F, -C, -K kan ikke brukes samtidig
	    disse tre kan brukes med -out, men ikke med -funfacts
	    -funfacts kan brukes kun med -t
	    ...
	    Jobb deg gjennom alle tilfellene. Vær obs på at det er en del sjekk
	    implementert i flag-pakken og at den vil skrive ut "Usage" med
	    beskrivelsene av flagg-variablene, som angitt i parameter fire til
	    funksjonene Float64Var og StringVar
	*/

	// Her er noen eksempler du kan bruke i den manuelle testingen
	fmt.Println(fahr, out, funfacts)

	//fmt.Println("len(flag.Args())", len(flag.Args()))
	//fmt.Println("flag.NFlag()", flag.NFlag())

	// Check for invalid flag combinations
	if (isFlagPassed("F") && isFlagPassed("C")) ||
	   (isFlagPassed("F") && isFlagPassed("K")) ||
	   (isFlagPassed("C") && isFlagPassed("K")) {
		fmt.Println("Invalid flag combination")
		return
	}

	// Check for unsupported flag combinations
	if isFlagPassed("funfacts") && !isFlagPassed("t") {
		fmt.Println("-funfacts can only be used with -t")
		return
	}
	if isFlagPassed("t") && (isFlagPassed("F") || isFlagPassed("C") || isFlagPassed("K")) {
		fmt.Println("-t cannot be used with -F, -C, or -K")
		return
	}

	
	
	
	
	
	/** 
	Kode hentet fra ChatopenAi. Convert the temperature.
	Hvis "F" flagget er sendt, vil programmet konvertere variabelen 
	f.eks. med navnet "fahr" ved funksjonen "FahrenheittoCelsius"
	som hentes fra conv-pakken, og skrive ut resultatet ved hjelp av fmt.Println".
	Hvis "K" flagget er sendt, vil programmet konvertere en variabel med navnet "kelvin" 
	fra Kelvin til Celsius ved hjelp av en funksjon som heter "KelvinToCelsius",
	og skrive ut resultatet på samme måte som for Fahrenheit-til-Celsius konverteringen.
	Hvis ingen av flaggene er sendt, vil ingen av konverteringene skje, og det vil ikke bli skrevet ut noe.

*/


	if out == "C" {
		if isFlagPassed("F") {
			fmt.Println(conv.FahrenheitToCelsius(fahr))
		} else if isFlagPassed("K") {
			fmt.Println(conv.KelvinToCelsius(kelvin))
		}
	} else if out == "F" {
		if isFlagPassed("C") {
			fmt.Println(conv.CelsiusToFahrenheit(celsius))
		} else if isFlagPassed("K") {
			fmt.Println(conv.KelvinToFahrenheit(kelvin))
		}
	} else if out == "K" {
		if isFlagPassed("C") {
			fmt.Println(conv.CelsiusToKelvin(celsius))
		} else if isFlagPassed("F") {
			fmt.Println(conv.FahrenheitToKelvin(fahr))
		}
	}
}

// Helper function to check if a flag is passed
func isFlagPassed(flagName string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == flagName {
			found = true
		}
	})
	return found
}


