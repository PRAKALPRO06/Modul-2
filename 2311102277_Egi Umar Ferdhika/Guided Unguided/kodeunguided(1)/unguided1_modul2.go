package main

import (
	"fmt"
)

func main() {
	// Input: Celsius temperature
	var celsius float64
	fmt.Print("Masukkan suhu dalam Celsius: ")
	fmt.Scanln(&celsius)

	// Convert to Reamur
	reamur := celsius * 4 / 5

	// Convert to Fahrenheit
	fahrenheit := (celsius * 9 / 5) + 32

	// Convert to Kelvin
	kelvin := celsius + 273.15

	// Output the results
	fmt.Printf("Temperatur Celsius: %.2f\n", celsius)
	fmt.Printf("Derajat Reamur: %.2f\n", reamur)
	fmt.Printf("Derajat Fahrenheit: %.2f\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %.2f\n", kelvin)
}
