package main

import (
	"fmt"
	"math"
)

func main() {
	var radius int
	var volume, surfaceArea float64

	fmt.Print("Masukkan jari-jari bola (bilangan bulat): ")
	fmt.Scanf("%d", &radius)

	volume = (4.0 / 3.0) * math.Pi * math.Pow(float64(radius), 3)

	surfaceArea = 4.0 * math.Pi * math.Pow(float64(radius), 2)

	fmt.Printf("Volume bola dengan jari-jari %d adalah: %.2f\n", radius, volume)
	fmt.Printf("Luas permukaan bola dengan jari-jari %d adalah: %.2f\n", radius, surfaceArea)
}
