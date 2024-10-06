package main

import "fmt"

func celsiuskefahrenheit(suhu float64) float64 {
	return (9.0 / 5.0 * suhu) + 32
}

func celsiuskekelvin(suhu float64) float64 {
	return suhu + 273.15
}
func celsiuskereamur(suhu float64) float64 {
	return 4.0 / 5.0 * suhu
}

func main() {

	var suhu float64
	fmt.Print("Temperatur celsius: ")
	fmt.Scanln(&suhu)

	fahrenheit := celsiuskefahrenheit(suhu)
	kelvin := celsiuskekelvin(suhu)
	reamur := celsiuskereamur(suhu)

	fmt.Println("Temperatur Celsius: %.0f\n", suhu)
	fmt.Println("Derajat Reamur: %.0f\n", reamur)
	fmt.Println("Derajat Fahrenheit: %.0f\n", fahrenheit)
	fmt.Println("Derajat Kelvin: %.0f\n", kelvin)

}
