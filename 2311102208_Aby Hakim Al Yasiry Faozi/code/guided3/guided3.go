package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64
	const pi = 3.1415926535

	fmt.Print("Masukkan jari-jari bola: ")
	fmt.Scanln(&radius)

	volume := (4.0 / 3.0) * pi * math.Pow(radius, 3)
	surfaceArea := 4 * pi * math.Pow(radius, 2)

	fmt.Printf("Bola dengan jari-jari %.2f memiliki volume %.4f dan luas permukaan %.4f\n", radius, volume, surfaceArea)
}
