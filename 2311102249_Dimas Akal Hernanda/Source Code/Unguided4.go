package main

import (
	"fmt"
	"strings"
)

func main() {
	var totalBunga int
	fmt.Print("Masukkan jumlah bunga yang ingin Anda masukkan: ")
	fmt.Scan(&totalBunga)

	var karangan strings.Builder
	var namaBunga string
	jumlahBunga := 0

	for i := 1; i <= totalBunga; i++ {
		fmt.Printf("Masukkan nama bunga ke-%d (ketik 'SELESAI' untuk selesai): ", i)
		fmt.Scan(&namaBunga)

		if strings.EqualFold(namaBunga, "SELESAI") {
			break
		}

		if karangan.Len() > 0 {
			karangan.WriteString(" - ")
		}
		karangan.WriteString(namaBunga)
		jumlahBunga++
	}

	fmt.Printf("Karangan bunga: %s\n", karangan.String())
	fmt.Printf("Bunga: %d\n", jumlahBunga)
}
