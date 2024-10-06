package main

import (
	"fmt"
)

// Fungsi untuk menghitung f(K) menggunakan rumus yang diberikan
func calculateF(k int) float64 {
	numerator := (4*float64(k) + 2) * (4*float64(k) + 2)
	denominator := (4*float64(k) + 1) * (4*float64(k) + 3)
	return numerator / denominator
}

// Fungsi untuk menghitung hampiran sqrt(2) menggunakan rumus yang diberikan
func calculateSqrt2(k int) float64 {
	sqrt2 := 1.0
	for i := 0; i <= k; i++ {
		numerator := (4*float64(i) + 2) * (4*float64(i) + 2)
		denominator := (4*float64(i) + 1) * (4*float64(i) + 3)
		sqrt2 *= numerator / denominator
	}
	return sqrt2
}

func main() {
	var k int
	fmt.Print("Masukkan nilai K: ")
	fmt.Scan(&k)

	// Hitung f(K)
	fK := calculateF(k)
	fmt.Printf("Nilai f(K) = %.10f\n", fK)

	// Hitung nilai hampiran sqrt(2)
	sqrt2 := calculateSqrt2(k)
	fmt.Printf("Nilai akar 2 = %.10f\n", sqrt2)
}
