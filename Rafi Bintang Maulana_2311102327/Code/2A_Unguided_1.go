package main

import (
	"fmt"
)

func main() {
	var celsius float64

	fmt.Print("Masukkan Temperatur Celsius: ")
	fmt.Scan(&celsius)

	untukReamur := int(celsius * 4 / 5)
	untukFahrenheit := int((celsius * 9 / 5) + 32)
	untukKelvin := int(celsius + 273.15)

	fmt.Printf("Derajat Reamur: %d\n", untukReamur)
	fmt.Printf("Derajat Fahrenheit: %d\n", untukFahrenheit)
	fmt.Printf("Derajat Kelvin: %d\n", untukKelvin)
}