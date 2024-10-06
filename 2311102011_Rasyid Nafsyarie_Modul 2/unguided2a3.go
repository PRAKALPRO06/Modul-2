package main

import (
	"fmt"
	"math"
)

func main() {
	var jejari_2311102011 float64
	const pi = 3.1415926535

	fmt.Print("Jejari = ")
	fmt.Scanln(&jejari_2311102011)

	// Menghitung volume bola
	volume := (4.0 / 3.0) * pi * math.Pow(jejari_2311102011, 3)

	// Menghitung luas permukaan bola
	luas := 4 * pi * math.Pow(jejari_2311102011, 2)

	// Menampilkan hasil
	fmt.Printf("Bola dengan jejari %.0f memiliki volume %.4f dan luas kulit %.4f\n", jejari_2311102011, volume, luas)
}