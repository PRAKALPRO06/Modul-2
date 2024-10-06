package main

import (
	"fmt"
	"math"
)


//reza alvonzo 2311102026 IF 11 06
func main() {
	var jejari float64
	const pi = 3.1415926535

	fmt.Print("Jejari = ")
	fmt.Scanln(&jejari)

	// Menghitung volume bola
	volume := (4.0 / 3.0) * pi * math.Pow(jejari, 3)

	// Menghitung luas permukaan bola
	luas := 4 * pi * math.Pow(jejari, 2)

	// Menampilkan hasil
	fmt.Printf("Bola dengan jejari %.0f memiliki volume %.4f dan luas kulit %.4f\n", jejari, volume, luas)
}