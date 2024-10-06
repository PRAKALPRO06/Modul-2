package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Fungsi untuk mencatat hasil percobaan warna
func percobaan() {
	var percobaan [5][4]string
	expectedColors := []string{"merah", "kuning", "hijau", "ungu"}

	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 5; i++ {
		fmt.Printf("Percobaan %d : ", i+1)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		colors := strings.Fields(input)

		if len(colors) != 4 {
			fmt.Println("Input harus terdiri dari 4 warna. Silakan coba lagi.")
			i--
			continue
		}
		for j := 0; j < 4; j++ {
			percobaan[i][j] = strings.ToLower(colors[j])
		}
	}
	warna := true
	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			if percobaan[i][j] != expectedColors[j] {
				warna = false
				break
			}
		}
		if !warna {
			break
		}
	}
	fmt.Printf("BERHASIL: = %t\n", warna)
	fmt.Println()
}

func bunga() {
	fmt.Print("Masukkan jumlah bunga: ")
	var n int
	fmt.Scanln(&n)

	var flowers []string

	for i := 0; i < n; i++ {
		var flower string
		fmt.Printf("Masukkan nama bunga %d: ", i+1)
		fmt.Scanln(&flower)

		if strings.ToUpper(flower) == "SELESAI" {
			fmt.Println("Program selesai.")
			break
		}

		flowers = append(flowers, flower)
	}

	fmt.Println()

	pita := strings.Join(flowers, " - ")
	fmt.Println("Pita:", pita)
	fmt.Printf("Jumlah bunga: %d\n", len(flowers))
}

func main() {
	percobaan()
	bunga()
}
