package main

import (
	"fmt"
	"math"
)

func hitungF(k int) float64 {
	pembilang := math.Pow(float64(4*k+2), 2)
	penyebut := float64((4*k+1)*(4*k+3))
	return pembilang / penyebut
}

func hitungAkar2(K int) float64 {
	hasil := 1.0
	for k := 0; k <= K; k++ {
		pembilang := math.Pow(float64(4*k+2), 2)
		penyebut := float64((4*k+1)*(4*k+3))
		hasil *= pembilang / penyebut
	}
	return hasil
}

func main() {
	var k int
	fmt.Print("Masukkan nilai k: ")
	fmt.Scan(&k)

	fk := hitungF(k)
	fmt.Printf("Nilai f(%d) = %.10f\n", k, fk)

	var K int
	fmt.Print("Masukkan nilai K untuk menghitung akar 2: ")
	fmt.Scan(&K)

	hampiran := hitungAkar2(K)
	fmt.Printf("Akar dua untuk K = %d adalah %.10f\n", K, hampiran)
}
