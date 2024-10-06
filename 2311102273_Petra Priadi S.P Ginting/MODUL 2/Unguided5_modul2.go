package main

import (
	"fmt"
)

func main() {
	var berat1, berat2, totalBerat float64

	for {
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scan(&berat1, &berat2)

		if berat1 < 0 || berat2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		totalBerat = berat1 + berat2
		if totalBerat > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		selisih := berat1 - berat2
		if selisih < 0 {
			selisih = -selisih
		}

		if selisih >= 9 {
			fmt.Println("Sepeda motor Pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor Pak Andi akan oleng: false")
		}
	}
}
