package main

import "fmt"

func main() {
	var N int
	var pita string
	var count int

	fmt.Print("Masukkan jumlah bunga (input berhenti saat mengetik 'SELESAI'): ")
	fmt.Scanln(&N)

	for i := 0; i < N; i++ {
		var namaBunga string
		fmt.Printf("Masukkan nama bunga ke-%d: ", i+1)
		fmt.Scanln(&namaBunga)

		if namaBunga == "SELESAI" {
			break
		}

		// Menambahkan nama bunga ke pita jika bukan "SELESAI"
		if count > 0 {
			pita += " - "
		}
		pita += namaBunga
		count++
	}

	// Menampilkan pita dan jumlah bunga
	fmt.Println("Pita:", pita)
	fmt.Println("Bunga:", count)
}
