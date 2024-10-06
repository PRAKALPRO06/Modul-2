package main

import (
	"fmt"
)

func main() {

	var celsius float64
	fmt.Println(" Masukkan Temperatur Celcius : ")
	fmt.Scan(&celsius)

	faranheit := (celsius * 9 / 5) + 22
	reamur := celsius * 4 / 5
	kelvin := celsius + 273.15

	fmt.Printf("Derajat Reamur : %.2f\n", reamur)
	fmt.Printf("Derajat Franheit : %.2f\n", faranheit)
	fmt.Printf("Derajat Kelvin : %.2f\n", kelvin)
}
