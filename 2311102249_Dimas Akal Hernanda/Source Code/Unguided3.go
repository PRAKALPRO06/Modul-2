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
	numTrials := 5 // jumlah percobaan

	// Membaca input untuk percobaan
	reader := bufio.NewReader(os.Stdin)
	success := true

	for i := 1; i <= numTrials; i++ {
		fmt.Printf("Percobaan %d: ", i)

		// Membaca input dari pengguna
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Terjadi kesalahan saat membaca input:", err)
			return
		}

		input = strings.TrimSpace(input)
		// Memisahkan input berdasarkan spasi
		colors := strings.Split(input, " ")

		// Mengecek apakah urutan warna sesuai
		if len(colors) != len(correctOrder) {
			success = false
			break
		}

		for j := 0; j < len(correctOrder); j++ {
			if colors[j] != correctOrder[j] {
				success = false
				break
			}
		}

		// Jika ada percobaan yang tidak sesuai, keluar dari loop
		if !success {
			break
		}
	}

	// Menampilkan hasil
	if success {
		fmt.Println("BERHASIL: true")
	} else {
		fmt.Println("BERHASIL: false")
	}
}
