package main

// 2311102254_Hamzah Ziyad I IF-11-06
import (
	"fmt"
	"math"
)

func main() {
	var berat1, berat2, total float64

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&berat1, &berat2)

		if berat1 < 0 || berat2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		selisih := math.Abs(berat1 - berat2)

		if selisih >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}

		total = berat1 + berat2

		if total > 150 {
			fmt.Println("Proses selesai.")
			break
		}
	}
}
