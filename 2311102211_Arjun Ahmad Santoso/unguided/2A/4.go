package main

import (
	"fmt"
)

func c_to_f(c int) int {
	return c*9/5 + 32
}

func main() {
	var c int = 5

	fmt.Print("Temperatur Celcius: ")
	fmt.Scanln(&c)
	f := c_to_f(c)
	fmt.Printf("Derajat Fahrenheit: %d", f)
}
