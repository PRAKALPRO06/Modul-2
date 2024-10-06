package main

import (
	"fmt"
)

func main() {
	var berat1, berat2 float64

	for {
		fmt.Print("Masukkan berat belanjaan di kedua kantong (berat1 berat2), atau ketik -1 untuk keluar: ")
		fmt.Scan(&berat1, &berat2)

		// Berikan opsi untuk keluar dari loop
		if berat1 == -1 || berat2 == -1 {
			fmt.Println("Program selesai.")
			break
		}

		// Cek apakah salah satu kantong lebih dari atau sama dengan 9 kg
		if berat1 >= 9 || berat2 >= 9 {
			fmt.Println("Sepeda motor Pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor Pak Andi akan oleng: false")
		}
	}
}
