package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	cek := bufio.NewScanner(os.Stdin)

	pita := ""
	jumlah := 0

	for {
		fmt.Printf("Bunga %d (ketik SELESAI untuk berhenti): ", jumlah+1)
		cek.Scan()
		bunga := cek.Text()

		if strings.ToUpper(bunga) == "SELESAI" {
			break
		}

		if pita == "" {
			pita = bunga
		} else {
			pita += " - " + bunga
		}

		jumlah++
	}

	fmt.Printf("\nPita: %s\n", pita)
	fmt.Printf("Jumlah bunga: %d\n", jumlah)
}
