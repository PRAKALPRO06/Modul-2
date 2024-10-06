package main

import (
	"fmt"
)

func main() {
	var celsius float64
	fmt.Print("Masukkan suhu dalam derajat Celsius: ")
	fmt.Scanln(&celsius)

	fahrenheit := (celsius * 9 / 5) + 32
	
	reamur := celsius * 4 / 5
	
	kelvin := celsius + 273.15

	fmt.Printf("Temperatur Celsius: %.0f\n", celsius)
	fmt.Printf("Derajat Reamur: %.0f\n", reamur)
	fmt.Printf("Derajat Fahrenheit: %.0f\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %.0f\n", kelvin)
}
