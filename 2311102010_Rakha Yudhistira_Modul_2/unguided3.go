package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	urutanWarna := []string{"merah", "kuning", "hijau", "ungu"}

	cek := bufio.NewScanner(os.Stdin)
	berhasil := true

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)
		cek.Scan()
		input := cek.Text()
		colors := strings.Split(input, " ")
		
		for j := 0; j < 4; j++ {
			if colors[j] != urutanWarna[j] {
				berhasil = false
			}
		}
	}

	if berhasil {
		fmt.Println("BERHASIL: true")
	} else {
		fmt.Println("BERHASIL: false")
	}
}
