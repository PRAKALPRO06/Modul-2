package main

import (
	"fmt"
)

func main() {
	var celsius float64

	// Membaca input temperatur dalam Celsius
	fmt.Print("Masukkan suhu dalam derajat Celsius: ")
	fmt.Scan(&celsius)

	// Konversi ke Fahrenheit
	fahrenheit := (celsius * 9 / 5) + 32

	// Konversi ke Reamur
	reamur := celsius * 4 / 5

	// Konversi ke Kelvin
	kelvin := celsius + 273.15

	// Menampilkan hasil konversi
	fmt.Printf("\nSuhu dalam Fahrenheit: %.2f°F\n", fahrenheit)
	fmt.Printf("Suhu dalam Reamur: %.2f°R\n", reamur)
	fmt.Printf("Suhu dalam Kelvin: %.2f°K\n", kelvin)
}
