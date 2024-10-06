package main

import (
	"fmt"
	"math"
)

func main() {
	var Jarijari float64
	fmt.Print("Masukkan jari-jari bola: ")
	fmt.Scanln(&Jarijari)
	volume := (4.0 / 3.0) * math.Pi * math.Pow(Jarijari, 3)
	Luaspermukaan := 4 * math.Pi * math.Pow(Jarijari, 2)
	fmt.Printf("Volume bola dengan jari-jari %.2f adalah %.2f\n", Jarijari, volume)
	fmt.Printf("Luas permukaan bola dengan jari-jari %.2f adalah %.2f\n", Jarijari, Luaspermukaan)
}
