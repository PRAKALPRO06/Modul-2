package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Membuat scanner untuk membaca input dari pengguna
	scanner := bufio.NewScanner(os.Stdin)

	// Inisialisasi pita untuk menyimpan nama-nama bunga
	var pita string
	jumlahBunga := 0

	// Meminta input bunga secara berulang hingga pengguna mengetik "SELESAI"
	for {
		fmt.Printf("Bunga %d: ", jumlahBunga+1)
		scanner.Scan()
		namaBunga := scanner.Text()

		// Jika pengguna mengetik "SELESAI", hentikan input
		if namaBunga == "SELESAI" || namaBunga == "selesai"{
			break
		}

		// Menggabungkan nama bunga dengan pita, dipisahkan dengan " – "
		if jumlahBunga == 0 {
			pita = namaBunga
		} else {
			pita += " – " + namaBunga
		}

		// Menambah jumlah bunga
		jumlahBunga++
	}

	// Menampilkan pita yang sudah dibuat dan jumlah bunga yang dimasukkan
	fmt.Printf("Pita: %s\n", pita)
	fmt.Printf("Bunga: %d\n", jumlahBunga)
}
