package main

import (
	"fmt"
	"strings"
)

func main() {
	var pita []string // Slice untuk menyimpan nama bunga
	var n int         // Variabel untuk menyimpan jumlah input bunga yang diinginkan

	// Meminta input jumlah bunga yang akan dimasukkan
	fmt.Print("N: ")
	fmt.Scanln(&n)

	// Jika input N adalah 0, langsung tampilkan hasil kosong
	if n == 0 {
		fmt.Println("Pita :")
		fmt.Println("Bunga: 0")
		return
	}

	count := 0 // Variabel untuk menghitung jumlah bunga yang dimasukkan

	// Meminta input bunga sebanyak n kali
	for count < n {
		var bunga string
		fmt.Printf("Bunga %d: ", count+1)
		fmt.Scanln(&bunga)

		if strings.ToUpper(bunga) == "SELESAI" { // Jika pengguna mengetik 'SELESAI', keluar dari loop
			break
		}

		pita = append(pita, bunga) // Menambahkan bunga ke pita
		count++
	}

	// Jika tidak ada bunga yang dimasukkan atau pengguna langsung mengetik "SELESAI"
	if count == 0 {
		fmt.Println("Pita :")
		fmt.Println("Bunga: 0")
		return
	}

	// Menampilkan pita dan jumlah bunga
	fmt.Printf("Pita: %s - \n", strings.Join(pita, " - ")) // Menampilkan pita dengan tanda " - " di antara bunga dan " - " di akhir
	fmt.Printf("Bunga: %d\n", count)                       // Menampilkan jumlah bunga yang sudah dimasukkan
}
