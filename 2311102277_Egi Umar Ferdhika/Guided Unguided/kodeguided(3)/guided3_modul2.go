package main

import (
	"fmt"
	"math"
)

func main() {
	var jariJari float64
	fmt.Print("Masukan jari-jari bola: ")
	fmt.Scanln(&jariJari)

	volume := (4.0 / 3.0) * math.Pi * math.Pow(jariJari, 3)
	luasPermukaan := 4 * math.Pi * math.Pow(jariJari, 2)

	fmt.Printf("Volume bola dengan jari-jari %.2f adalah %.2f\n", jariJari, volume)
	fmt.Printf("Luas permukaan bola dengan jari-jari %2.f adalah %.2f\n", jariJari, luasPermukaan)

}
