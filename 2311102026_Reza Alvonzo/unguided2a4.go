package main

import (
	"fmt"
)


//reza alvonzo 2311102026 IF 11 06
func main() {

	var celsius float64


	fmt.Print("Masukkan suhu dalam derajat Celsius: ")
	fmt.Scanf("%f", &celsius)


	reaumur := (4.0 / 5.0) * celsius


	fahrenheit := (celsius * 9.0 / 5.0) + 32.0


	kelvin := celsius + 273.15

	fmt.Printf("Temperatur : %.2f\n", celsius)
	fmt.Printf("Derajat Reaumur: %.2f\n", reaumur)
	fmt.Printf("Derajat Fahrenheit: %.2f\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %.2f\n", kelvin)
}