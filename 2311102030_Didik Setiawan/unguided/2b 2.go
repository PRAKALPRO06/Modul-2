package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	fmt.Print("Masukkan jumlah bunga (N): ")
	fmt.Scan(&N)

	var pita strings.Builder
	var bunga string

	for i := 1; i <= N; i++ {
		fmt.Printf("Bunga %d: ", i)
		fmt.Scan(&bunga)

		// Jika input adalah "SELESAI", hentikan input
		if strings.ToUpper(bunga) == "SELESAI" {
			break
		}

		// Tambahkan tanda "-" sebelum bunga, kecuali bunga pertama
		if pita.Len() > 0 {
			pita.WriteString(" - ")
		}
		pita.WriteString(bunga)
	}

	// Output hasil jika ada bunga yang dimasukkan
	if pita.Len() > 0 {
		fmt.Printf("Pita: %s\n", pita.String())
		fmt.Printf("Bunga: %d\n", strings.Count(pita.String(), " - ")+1)
	} else {
		fmt.Println("Tidak ada bunga yang dimasukkan.")
	}
}
