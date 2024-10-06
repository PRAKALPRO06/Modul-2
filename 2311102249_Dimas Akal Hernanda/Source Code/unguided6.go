package main

import (
	"fmt"
)

// Fungsi untuk menghitung ekspresi matematis berdasarkan nilai k
func f(k int) float64 {
	// Menghitung pembilang (atas)
	atas := (4*float64(k) + 2) * (4*float64(k) + 2)
	// Menghitung penyebut (bawah)
	bawah := (4*float64(k) + 1) * (4*float64(k) + 3)
	// Mengembalikan hasil dari ekspresi
	return atas / bawah
}

func main() {
	var k int
	// Meminta input dari pengguna
	fmt.Print("Nilai K = ")
	fmt.Scan(&k)
	// Menghitung hasil menggunakan fungsi f
	hasil := f(k)
	// Menampilkan hasil dengan format 10 angka desimal
	fmt.Printf("Nilai f(k) = %.10f\n", hasil)
}
