package main

import (
	"fmt"
)

func main() {

	var celsius_2311102021 float64


	fmt.Print("Masukkan suhu dalam derajat Celsius: ")
	fmt.Scanf("%f", &celsius_2311102021)


	reaumur := (4.0 / 5.0) * celsius_2311102021


	fahrenheit := (celsius_2311102021 * 9.0 / 5.0) + 32.0


	kelvin := celsius_2311102021 + 273.15

	fmt.Printf("Temperatur : %.2f\n", celsius_2311102021)
	fmt.Printf("Derajat Reaumur: %.2f\n", reaumur)
	fmt.Printf("Derajat Fahrenheit: %.2f\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %.2f\n", kelvin)
}