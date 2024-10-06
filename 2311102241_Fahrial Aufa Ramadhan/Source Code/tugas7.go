package main

import (
	"fmt"
	"strings"
)

func main() {
	var totalBunga int
	fmt.Print("Masukkan jumlah bunga yang diinginkan: ")
	fmt.Scan(&totalBunga)

	var karangan strings.Builder
	var namaBunga string
	jumlahBunga := 0

	for urutan := 1; urutan <= totalBunga; urutan++ {
		fmt.Printf("Nama bunga ke-%d: ", urutan)
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
	fmt.Printf("Total bunga yang terhitung: %d\n", jumlahBunga)
}
