package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	fmt.Print("Masukkan jumlah bunga (N): ")
	fmt.Scan(&N)

	var pita_2311102013 strings.Builder
	var bunga_2311102013 string
	count := 0

	for i := 1; i <= N; i++ {
		fmt.Printf("Bunga %d: ", i)
		fmt.Scan(&bunga_2311102013)

		if strings.ToUpper(bunga_2311102013) == "SELESAI" {
			break
		}

		if pita_2311102013.Len() > 0 {
			pita_2311102013.WriteString(" - ")
		}
		pita_2311102013.WriteString(bunga_2311102013)
		count++
	}

	fmt.Printf("Pita: %s\n", pita_2311102013.String())
	fmt.Printf("Bunga: %d\n", count)
}