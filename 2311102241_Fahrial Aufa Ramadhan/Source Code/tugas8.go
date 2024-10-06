package main

import (
	"fmt"
	"math"
)

func main() {
	var beratBelajaan, berat float64

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&beratBelajaan, &berat)

		totalBerat := beratBelajaan + berat
		selisih := math.Abs(beratBelajaan - berat)

		if totalBerat > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		if selisih >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")

		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}
}
