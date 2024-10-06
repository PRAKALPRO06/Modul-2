package main

import "fmt"

func celsiusKeFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func celsiusKeReamur(celsius float64) float64 {
	return celsius * 4 / 5
}

func celsiusKeKelvin(celsius float64) float64 {
	return celsius + 273.15
}

func main() {
	var celsius float64
	fmt.Print("Masukkan suhu dalam Celsius: ")
	fmt.Scanln(&celsius)

	fahrenheit := celsiusKeFahrenheit(celsius)
	reamur := celsiusKeReamur(celsius)
	kelvin := celsiusKeKelvin(celsius)

	fmt.Printf("Temperatur Celsius: %.2f\n", celsius)
	fmt.Printf("Derajat Reamur: %.2f\n", reamur)
	fmt.Printf("Derajat Fahrenheit: %.2f\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %.2f\n", kelvin)
}
