package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung nilai f(k)
func hitungF(k int) float64 {
	pembilang := math.Pow(float64(4*k+2), 2)
	penyebut := float64((4*k + 1) * (4*k + 3))
	return pembilang / penyebut
}

// Fungsi untuk menghitung nilai perkiraan akar 2 dengan iterasi hingga k
func hitungAkarDua(k int) float64 {
	hasil := 1.0
	for i := 0; i <= k; i++ {
		hasil *= hitungF(i)
	}
	return hasil
}

func main() {
	var k int
	fmt.Print("Masukkan nilai K (>= 0): ")
	fmt.Scan(&k)

	// Validasi nilai k agar tidak negatif
	if k < 0 {
		fmt.Println("Nilai K harus lebih besar atau sama dengan 0.")
		return
	}

	nilaiF := hitungF(k)
	fmt.Printf("Nilai f(%d) = %.10f\n", k, nilaiF)

	akarDua := hitungAkarDua(k)
	fmt.Printf("Nilai perkiraan akar 2 untuk K = %d: %.10f\n", k, akarDua)
}
