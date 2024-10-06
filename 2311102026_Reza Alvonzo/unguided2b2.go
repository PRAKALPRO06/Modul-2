package main

import (
	"fmt"
	"strings"
)

//reza alvonzo 2311102026 IF 11 06

func main() {
	var N int
	fmt.Print("Masukkan jumlah bunga (N): ")
	fmt.Scan(&N)

	var pita strings.Builder
	var bunga string
	count := 0

	for i := 1; i <= N; i++ {
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