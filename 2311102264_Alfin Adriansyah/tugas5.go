package main

import (
	"fmt"
	"math"
)

func main() {
	var kantongKiri, kantongKanan float64
	totalBerat := 0.0
	prosesSelesai := false

	for {
		// Input berat dari kedua kantong
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scanln(&kantongKiri, &kantongKanan)

		// Tampilkan berat yang diinputkan
		fmt.Printf("Masukan berat belanjaan di kedua kantong: %.1f %.1f\n", kantongKiri, kantongKanan)

		// Tambahkan berat ke total berat
		totalBerat += kantongKiri + kantongKanan

		// Cek jika salah satu kantong >= 9 kg
		if kantongKiri >= 9 || kantongKanan >= 9 {
			prosesSelesai = true
		}

		// Cek selisih antara kantong kiri dan kanan
		selisih := math.Abs(kantongKiri - kantongKanan)

		// Tampilkan hasil apakah sepeda motor akan oleng
		if selisih >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}

		// Cek jika total berat > 150 kg
		if totalBerat > 150 {
			prosesSelesai = true
		}

		// Jika proses selesai, berhenti
		if prosesSelesai {
			fmt.Println("Proses selesai.")
			break
		}
	}
}
