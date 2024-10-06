package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Urutan warna yang benar
	correctOrder := []string{"merah", "kuning", "hijau", "ungu"}

	// Membaca input dari pengguna sebanyak 5 kali percobaan
	scanner := bufio.NewScanner(os.Stdin)
	var allCorrect bool = true

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)
		scanner.Scan()
		input := scanner.Text()

		// Memecah input menjadi array warna
		colors := strings.Split(input, " ")

		// Cek apakah input sesuai dengan urutan yang benar
		if len(colors) != 4 {
			fmt.Println("Input tidak valid, harap masukkan 4 warna.")
			allCorrect = false
			continue
		}

		for j, color := range colors {
			if color != correctOrder[j] {
				allCorrect = false
				break
			}
		}
	}

	// Output hasil setelah 5 percobaan
	if allCorrect {
		fmt.Println("HASIL: true")
	} else {
		fmt.Println("HASIL: false")
	}
}
