package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	experiment := []string{"merah", "kuning", "hijau", "ungu"}
	scanner := bufio.NewScanner(os.Stdin)

	allCorrect := true

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)

		scanner.Scan()
		input := scanner.Text()

		warna := strings.Fields(input)

		if len(warna) != 4 || !isCorrectOrder(warna, experiment) {
			allCorrect = false
		}
	}

	if allCorrect {
		fmt.Println("BERHASIL: true")
	} else {
		fmt.Println("BERHASIL: false")
	}
}

func isCorrectOrder(warna, experiment []string) bool {
	for i := range warna {
		if warna[i] != experiment[i] {
			return false
		}
	}
	return true
}
