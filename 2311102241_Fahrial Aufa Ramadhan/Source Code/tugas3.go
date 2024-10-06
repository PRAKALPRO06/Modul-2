package main

import "fmt"

func main() {
	var jariJari float64
	const pi = 3.14159

	fmt.Print("Masukkan jari-jari bola: ")
	fmt.Scan(&jariJari)

	luasKulit := 4 * pi * (jariJari * jariJari)
	volume := (4.0 / 3.0) * pi * (jariJari * jariJari * jariJari)

	fmt.Println("bola dengan jari jari", jariJari, "memiliki volume ", volume, "dengan luas kulit", luasKulit)
}
