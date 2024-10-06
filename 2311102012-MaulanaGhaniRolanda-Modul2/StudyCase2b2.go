//2311102012
package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	var pita []string

	// Meminta input bilangan bulat positif untuk N
	fmt.Print("N: ")
	fmt.Scan(&N)

	// Memastikan N positif dan tidak nol
	if N <= 0 {
		fmt.Println("N harus merupakan bilangan bulat positif dan tidak nol.")
		return
	}

	// Meminta input nama bunga sebanyak N kali
	for i := 1; i <= N; i++ {
		var bunga string
		fmt.Printf("Bunga %d: ", i)
		fmt.Scan(&bunga)

		// Memeriksa apakah input adalah "SELESAI"
		if strings.ToUpper(bunga) == "SELESAI" {
			break
		}

		// Menambahkan nama bunga ke dalam pita
		pita = append(pita, bunga)
	}

	// Menampilkan isi pita dan jumlah bunga
	fmt.Printf("Pita: %s\n", strings.Join(pita, " - "))
	fmt.Printf("Jumlah bunga: %d\n", len(pita))
}
