package main

import (
	"fmt"
	"math"
)

func main() {
	var kantong1, kantong2, totalBerat float64
	const batasBerat = 150.0

	for {
		fmt.Print("Masukkan berat belanjaan di kedua kantong (kg): ")
		fmt.Scan(&kantong1, &kantong2)

		if kantong1 < 0 || kantong2 < 0 {
			fmt.Println("Proses selesai karena ada berat negatif.")
			break
		}

		totalBerat += kantong1 + kantong2

		if totalBerat > batasBerat {
			fmt.Printf("Total berat sudah melebihi batas %.2f kg.\nProses selesai.\n", batasBerat)
			break
		}

		if math.Abs(kantong1-kantong2) >= 9 {
			fmt.Println("Sepeda motor Pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor Pak Andi akan oleng: false")
		}

		if kantong1 >= 9 || kantong2 >= 9 {
			fmt.Println("Proses selesai.")
			break
		}
	}
}
