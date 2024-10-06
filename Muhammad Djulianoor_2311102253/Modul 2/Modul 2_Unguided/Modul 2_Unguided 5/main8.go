package main

import (
	"fmt"
	"math"
)

func main() {
	var kantong1, kantong2 float64
	var total1, total2 float64
	total1, total2 = 0, 0

	for {
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kantong1, &kantong2)

		if kantong1 < 0 || kantong2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		total1 += kantong1
		total2 += kantong2

		if math.Abs(total1-total2) > 9 {
			fmt.Println("Sepeda motor Pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor Pak Andi akan oleng: false")
		}

		if total1+total2 >= 150 {
			fmt.Println("Proses selesai.")
			break
		}
	}
}
