package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	var bunga, pita string
	var count int

	// Meminta input jumlah bunga
	fmt.Print("Masukkan jumlah bunga (N): ")
	fmt.Scan(&N)

	// Proses input nama bunga sebanyak N kali atau hingga 'SELESAI'
	for i := 1; i <= N; i++ {
		fmt.Printf("Bunga %d: ", i)
		fmt.Scan(&bunga)

		// Jika input 'SELESAI', hentikan input
		if strings.ToUpper(bunga) == "SELESAI" {
			break
		}

		// Gabungkan nama bunga ke pita
		if pita == "" {
			pita = bunga
		} else {
			pita = pita + " – " + bunga
		}
		count++
	}

	// Menampilkan hasil pita dan jumlah bunga yang dimasukkan
	fmt.Printf("Pita: %s\n", pita)
	fmt.Printf("Banyak bunga: %d\n", count)
}
