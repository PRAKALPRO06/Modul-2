package main

import (
	"fmt"
	"strings"
)

func main() {
	var N_2311102021 int
	fmt.Print("Masukkan jumlah bunga (N): ")
	fmt.Scan(&N_2311102021)

	var pita strings.Builder
	var bunga string
	count := 0

	for i := 1; i <= N_2311102021; i++ {
		fmt.Printf("Bunga %d: ", i)
		fmt.Scan(&bunga)

		if strings.ToUpper(bunga) == "SELESAI" {
			break
		}

		if pita.Len() > 0 {
			pita.WriteString(" - ")
		}
		pita.WriteString(bunga)
		count++
	}

	fmt.Printf("Pita: %s\n", pita.String())
	fmt.Printf("Bunga: %d\n", count)
}