package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	correctSequence := []string{"merah", "kuning", "hijau", "ungu"}

	scanner := bufio.NewScanner(os.Stdin)

	experiments := make([][]string, 5)
	for i := 0; i < 5; i++ {
		fmt.Printf("Percobaan %d: ", i+1)
		if scanner.Scan() {
			experiments[i] = strings.Fields(scanner.Text())
		}
	}

	success := true
	for _, experiment := range experiments {
		if !equal(experiment, correctSequence) {
			success = false
			break
		}
	}

	if success {
		fmt.Println("BERHASIL: true")
	} else {
		fmt.Println("BERHASIL: false")
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
