package main

import "fmt"

func main() {
	var celcius float64

	fmt.Print("Masukan Temperatur Celcius : ")
	fmt.Scan(&celcius)

	fahrenheit := (celcius * 9 / 5) + 32
	kelvin := celcius + 273
	reamur := (celcius * 4 / 5)

	fmt.Printf("Derajat Reamur : %.2f\n", reamur)
	fmt.Printf("Derajat Fahrenheit : %.2f\n", fahrenheit)
	fmt.Printf("Derajat Kelvin : %.2f\n", kelvin)

}
