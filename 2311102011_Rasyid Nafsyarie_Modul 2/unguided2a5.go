package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b, c, d, e int
	reader := bufio.NewReader(os.Stdin) // Membuat reader untuk menangani input karakter

	// Meminta input 5 buah angka integer dari pengguna
	fmt.Println("Masukkan 5 buah angka (nilai antara 32 hingga 127):")
	fmt.Scanf("%d %d %d %d %d", &a, &b, &c, &d, &e)

	// Mencetak hasil karakter dari input angka (berdasarkan tabel ASCII)
	fmt.Printf("Hasil karakter dari angka yang diinput: %c%c%c%c%c\n", a, b, c, d, e)

	// Meminta input karakter berupa angka "123" dari pengguna
	fmt.Println("Masukkan karakter (contoh input: 123):")
	inputStr, _ := reader.ReadString('\n') // Membaca input berupa string sampai newline

	// Menggeser karakter sesuai dengan keinginan (contoh: 1 -> 2, 2 -> 3, 3 -> 4)
	shiftedStr := ""
	for _, char := range inputStr {
		// Skip newline or other non-visible characters
		if char != '\n' && char != '\r' {
			shiftedStr += string(char + 1) // Menggeser setiap karakter
		}
	}

	// Mencetak hasil karakter yang sudah digeser
	fmt.Printf("Hasil karakter dari inputan angka yang digeser: %s\n", shiftedStr)
}