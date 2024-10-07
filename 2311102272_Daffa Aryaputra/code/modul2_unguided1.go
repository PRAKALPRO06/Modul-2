package main

import (
	"fmt"
	"strconv"
)

func main() {
	var celsiusStr string
	fmt.Print("Masukkan suhu dalam Celcius: ")
	fmt.Scanln(&celsiusStr)

	celsius, err := strconv.ParseFloat(celsiusStr, 64)
	if err != nil {
		fmt.Println("Input tidak valid. Mohon masukkan angka.")
		return
	}

	fahrenheit := (celsius * 9 / 5) + 32
	reamur := celsius * 4 / 5
	kelvin := celsius + 273.15

	fmt.Printf("Suhu dalam Fahrenheit: %.2f\n", fahrenheit)
	fmt.Printf("Suhu dalam Reaumur: %.2f\n", reamur)
	fmt.Printf("Suhu dalam Kelvin: %.2f\n", kelvin)
}