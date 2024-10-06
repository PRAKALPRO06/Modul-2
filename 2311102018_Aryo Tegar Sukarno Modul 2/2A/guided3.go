package main

import (
	"fmt"
	"math"
)

func main() {
	var jariJari float64
	fmt.Print("Masukan jari jari bola : ")
	fmt.Scanln(&jariJari)

	volume := (4.0/3.0) * math.Pi * math.Pow(jariJari, 3)
	luasPermukaan := 4 * math.Pi * math.Pow(jariJari, 2)

	fmt.Printf("Bola dengan jari-jari %.2f memiliki volume %.5f dan luas permukaan %.5f\n", jariJari, volume, luasPermukaan)
}