package main

import (
	"fmt"
	"math"
)

func main() {
	var beratKiri, beratKanan float64
	totalBerat := 0.0
	prosesSelesai := false

	for {
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scanln(&beratKiri, &beratKanan)

		fmt.Printf("Berat di kedua kantong adalah: %.1f kg dan %.1f kg\n", beratKiri, beratKanan)

		totalBerat += beratKiri + beratKanan

		if beratKiri >= 9 || beratKanan >= 9 {
			prosesSelesai = true
		}

		selisih := math.Abs(beratKiri - beratKanan)

		if selisih >= 9 {
			fmt.Println("Sepeda motor Pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor Pak Andi akan oleng: false")
		}
		if totalBerat > 150 {
			prosesSelesai = true
		}
		if prosesSelesai {
			fmt.Println("Proses selesai.")
			break
		}
	}
}
