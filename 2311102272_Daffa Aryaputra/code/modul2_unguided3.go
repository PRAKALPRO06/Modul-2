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
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scanln(&kantongKiri, &kantongKanan)

		fmt.Printf("Masukan berat belanjaan di kedua kantong: %.1f %.1f\n", kantongKiri, kantongKanan)

		totalBerat += kantongKiri + kantongKanan

		if kantongKiri >= 9 || kantongKanan >= 9 {
			prosesSelesai = true
		}

		selisih := math.Abs(kantongKiri - kantongKanan)

		if selisih >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
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