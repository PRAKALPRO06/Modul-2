package main

import (
	"fmt"
	"math"
)

func main() {
	var jarijari float64
	fmt.Print("Masukkan jari-jari bola : ")
	fmt.Scanln(&jarijari)

	volume := (4.0/3.0) * math.Pi * math.Pow(jarijari, 3)
	luasPermukaan := 4 * math.Pi * math.Pow(jarijari, 2)

	fmt.Printf("Bola dengan jejarinya %.0f memiliki volume %.4f dan luas kulit %.4f\n", jarijari, volume, luasPermukaan)
}
