package main

import (
	"fmt"
	"strings"
)

func main() {
	var TOTALPADABUNGA int
	fmt.Print("Silahkan Masukkan jumlah bunga yang anda diinginkan: ")
	fmt.Scan(&TOTALPADABUNGA)

	var karangan strings.Builder
	var namaBunga string
	jumlahBunga := 0

	for urutan := 1; urutan <= TOTALPADABUNGA; urutan++ {
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
