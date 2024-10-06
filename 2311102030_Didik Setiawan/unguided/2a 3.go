package main

import (
	"fmt"
	"math"
)

func main() {
	var jari2 float64
	const pi = 3.1415926535

	fmt.Print("Jejari = ")
	fmt.Scanln(&jari2)

	volume := (4.0 / 3.0) * pi * math.Pow(jari2, 3)

	luas := 4 * pi * math.Pow(jari2, 2)

	fmt.Printf("Bola dengan jari2 %.0f mempunyai volume %.4f dan luas kulit %.4f\n", jari2, volume, luas)
}
