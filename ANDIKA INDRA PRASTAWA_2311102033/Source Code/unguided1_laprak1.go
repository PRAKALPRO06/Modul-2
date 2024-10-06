package main

import (
	"fmt"
)

func main() {
	var celsius float64

	fmt.Print(" Masukkan suhu ke dalam Celsius: ")
	fmt.Scan(&celsius)

	fahrenheit := (celsius * 9 / 5) + 32
	kelvin := celsius + 273.15

	fmt.Printf("%.2f°C Sama Dengan:\n", celsius)
	fmt.Printf("%.2f°F\n", fahrenheit)
	fmt.Printf("%.2f K\n", kelvin)
}
