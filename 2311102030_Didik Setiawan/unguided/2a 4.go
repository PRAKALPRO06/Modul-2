package main

import (
	"fmt"
)

func main() {
	var celsius float64

	// Input suhu dalam Celsius
	fmt.Print("Masukkan suhu dalam derajat Celsius: ")
	fmt.Scanf("%f", &celsius)

	// Hitung suhu dalam berbagai satuan
	fmt.Printf("Temperatur Celsius: %.2f\n", celsius)
	fmt.Printf("Derajat Reaumur: %.2f\n", (4.0/5.0)*celsius)
	fmt.Printf("Derajat Fahrenheit: %.2f\n", (celsius*9.0/5.0)+32.0)
	fmt.Printf("Derajat Kelvin: %.2f\n", celsius+273.15)
}
