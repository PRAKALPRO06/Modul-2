// 2311102012
package main

import (
	"fmt"
	"math"
)

func f(k int) float64 {
	// Menghitung f(k) berdasarkan rumus yang diberikan
	numerator := math.Pow(float64(4*k+2), 2)
	denominator := float64((4*k + 1) * (4*k + 3))
	return numerator / denominator
}

func main() {
	var k int

	// Meminta input dari pengguna
	fmt.Print("Nilai K = ")
	fmt.Scan(&k)

	// Menghitung dan menampilkan f(K)
	fK := f(k)
	fmt.Printf("Nilai f(k) = %.10f\n", fK)

	// Menghitung sqrt(2) menggunakan rumus hingga K
	sqrt2Approx := 1.0 // Inisialisasi dengan nilai awal
	for i := 0; i <= k; i++ {
		sqrt2Approx *= f(i)
	}

	// Menampilkan hasil perhitungan sqrt(2)
	fmt.Printf("Nilai akar 2 = %.10f\n", sqrt2Approx)
}
