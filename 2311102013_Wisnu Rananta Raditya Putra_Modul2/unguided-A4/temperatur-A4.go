package main

import (
	"fmt"
)

func main() {
	var celcius_2311102013, fahrenheit, kelvin, reamur float64

	fmt.Print("Temperatur Celsius: ")
	fmt.Scanln(&celcius_2311102013)

	reamur = celcius_2311102013 * 4 / 5
	fahrenheit = (celcius_2311102013 * 9 / 5) + 32
	kelvin = celcius_2311102013 + 273.15
	
	fmt.Println("Derajat Reamur: ", reamur)
	fmt.Println("Derajat Fahrenheit: ", fahrenheit)
	fmt.Println("Derajat Kelvin: ", kelvin)
}