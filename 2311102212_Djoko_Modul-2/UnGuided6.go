package main

import (
	"fmt"
)

// Fungsi untuk menghitung aproksimasi akar 2 menggunakan rumus produk
func hitungAkar2(n int) float64 {
	hasil := 1.0

	// Loop untuk menghitung produk dari suku-suku
	for k := 0; k < n; k++ {
		atas := (4*float64(k) + 2) * (4*float64(k) + 2)
		bawah := (4*float64(k) + 1) * (4*float64(k) + 3)
		hasil *= atas / bawah
	}

	return hasil
}

func main() {
	var n int

	// Meminta input jumlah suku yang akan dihitung
	fmt.Print("Nilai K: ")
	fmt.Scan(&n)

	// Menghitung aproksimasi akar 2
	akar2 := hitungAkar2(n)

	// Menampilkan hasil aproksimasi
	fmt.Printf("Nilai akar 2 = %.10f\n", akar2)
}
