package main

import (
	"fmt"
)

func main() {

	var celsius_2311102011 float64


	fmt.Print("Masukkan suhu dalam derajat Celsius: ")
	fmt.Scanf("%f", &celsius_2311102011)


	reaumur := (4.0 / 5.0) * celsius_2311102011


	fahrenheit := (celsius_2311102011 * 9.0 / 5.0) + 32.0


	kelvin := celsius_2311102011 + 273.15

	fmt.Printf("Temperatur : %.2f\n", celsius_2311102011)
	fmt.Printf("Derajat Reaumur: %.2f\n", reaumur)
	fmt.Printf("Derajat Fahrenheit: %.2f\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %.2f\n", kelvin)
}