// 2311102012
package main

import (
	"fmt"
)

func main() {
	var kantong1, kantong2 float64

	for {
		// Meminta input berat dari pengguna
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scanf("%f %f", &kantong1, &kantong2)

		// Mengecek apakah berat negatif
		if kantong1 < 0 || kantong2 < 0 {
			break // Hentikan proses jika ada berat negatif
		}

		// Mengecek total berat apakah melebihi 150 kg
		totalBerat := kantong1 + kantong2
		if totalBerat > 150 {
			break // Hentikan proses jika total berat lebih dari 150 kg
		}

		// Menghitung selisih berat
		selisih := kantong1 - kantong2
		if selisih < 0 {
			selisih = -selisih // Mengambil nilai absolut
		}

		// Menampilkan hasil apakah sepeda motor akan oleng
		if selisih >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}

	fmt.Println("Proses selesai")
}
