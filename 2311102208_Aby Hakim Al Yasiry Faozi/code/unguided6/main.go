package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung nilai f(K)
func f(k float64) float64 {
	return (4*math.Pow(k, 2) + math.Pow(2, 2)) / ((k + 1) * (4*k + 3))
}

func main() {
	var K float64
	fmt.Print("Masukkan nilai K: ")
	fmt.Scan(&K)

	// Menghitung nilai f(K)
	hasilF := f(K)
	// Menghitung akar dari hasil f(K)
	hasilAkar := math.Sqrt(hasilF)

	// Menampilkan hasil
	fmt.Printf("Nilai f(K) = %.10f\n", hasilF)
	fmt.Printf("Nilai akar dari f(K) = %.10f\n", hasilAkar)
}
