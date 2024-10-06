package main

import (
	"fmt"
	"math"
)

const (
	maxWeight    = 150
	unstableDiff = 9
)

func main() {
	for {
		var kantong1, kantong2 float64
		fmt.Print("Masukkan berat belanjaan di kedua kantong (format: berat1 berat2): ")

		if _, err := fmt.Scan(&kantong1, &kantong2); err != nil {
			fmt.Println("Input tidak valid. Silakan coba lagi.")
			continue
		}

		if kantong1 < 0 || kantong2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		totalWeight := kantong1 + kantong2
		if totalWeight > maxWeight {
			fmt.Println("Proses selesai. ")
			break
		}

		weightDifference := math.Abs(kantong1 - kantong2)
		isUnstable := weightDifference >= unstableDiff
		fmt.Printf("Sepeda motor Pak Andi akan oleng: %t\n", isUnstable)
	}
}
