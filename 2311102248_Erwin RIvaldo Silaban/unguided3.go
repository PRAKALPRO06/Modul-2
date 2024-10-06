package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	win := []string{"merah", "kuning", "hijau", "ungu"}
	berhasil := true

	for i := 1; i <= 5; i++ {
		fmt.Println("Warna : Merah, Kuning, Hijau, Ungu")
		fmt.Printf("Percobaan %d: ", i)
		input, _ := reader.ReadString('\n')
		warna := strings.Fields(strings.TrimSpace(input))

		if len(warna) != 4 {
			fmt.Println("Error: Masukkan 4 warna untuk setiap percobaan.")
			return
		}

		fmt.Printf("Percobaan %d: %s %s %s %s\n", i, warna[0], warna[1], warna[2], warna[3])

		for j := 0; j < 4; j++ {
			if warna[j] != win[j] {
				berhasil = false
			}
		}
	}

	fmt.Printf("BERHASIL: %v\n", berhasil)
}
