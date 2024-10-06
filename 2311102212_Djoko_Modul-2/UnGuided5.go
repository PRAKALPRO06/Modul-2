package main

import (
	"fmt"
	"math"
)

func main() {
	var beratKiri, beratKanan float64

	for {
		// Meminta input berat barang di kantong kiri dan kanan
		fmt.Print("Masukkan berat belanjaan di kantong kiri: ")
		fmt.Scan(&beratKiri)
		fmt.Print("Masukkan berat belanjaan di kantong kanan: ")
		fmt.Scan(&beratKanan)

		// Menghitung selisih berat antara kedua kantong
		selisihBerat := math.Abs(beratKiri - beratKanan)

		// Tentukan apakah sepeda motor oleng (true) atau tidak oleng (false)
		oleng := selisihBerat > 9

		// Tampilkan hasil pengecekan
		if oleng {
			fmt.Println("Sepeda motor Pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor Pak Andi tidak akan oleng: false")
		}
		
		totalBarang := math.Abs(beratKiri + beratKanan)
		
		if totalBarang > 150 {
		    fmt.Println("Program Selesai")
		    break
		}
		
	}
}
