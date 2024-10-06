// 2311102012
package main

import (
	"fmt"
	"math"
)

func main() {
	var jejari float64
	const pi = 3.1415926535

	fmt.Print("Jejari = ")
	fmt.Scanln(&jejari)

	volume := (4.0 / 3.0) * pi * math.Pow(jejari, 3)

	luas := 4 * pi * math.Pow(jejari, 2)

	fmt.Printf("Bola dengan jejari %.0f memiliki volume %.4f dan luas kulit %.4f\n", jejari, volume, luas)
}
