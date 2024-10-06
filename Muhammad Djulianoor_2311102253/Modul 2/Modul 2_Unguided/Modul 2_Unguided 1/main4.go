package main

import "fmt"

func main() {
	var celcius float32

	fmt.Print("Temperatur Celcius: ")
	fmt.Scanln(&celcius)

	fahrenheit := (celcius * 9 / 5) + 32
	reamur := celcius * 4 / 5
	kelvin := celcius + 273.15

	fmt.Println("Derajat Reamur: ", reamur)
	fmt.Println("Derajat Fahrenheit: ", fahrenheit)
	fmt.Println("Derajat Kelvin: ", kelvin)
}
