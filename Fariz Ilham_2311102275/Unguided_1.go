package main

import (
	"fmt"
)

func main() {
	var celsius float64

	fmt.Print("Temperatur Celsius: ")
	_, err := fmt.Scan(&celsius)
	if err != nil {
		fmt.Println("Input tidak valid. Harap masukkan angka.")
		return
	}

	fahrenheit := (celsius * 9 / 5) + 32
	reamur := celsius * 4 / 5
	kelvin := celsius + 273.15

	fmt.Printf("Derajat Fahrenheit: %.2f\n", fahrenheit)
	fmt.Printf("Derajat Reamur: %.2f\n", reamur)
	fmt.Printf("Derajat Kelvin: %.2f\n", kelvin)
	fmt.Println()
}
