package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func cekKeseimbangan(berat1, berat2 float64) bool {
	selisih := math.Abs(berat1 - berat2)
	return selisih <= 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Masukkan berat belanjaan di kedua kantong (pisahkan dengan spasi), atau ketik 'SELESAI' untuk berhenti: ")
		scanner.Scan()
		input := scanner.Text()

		if strings.ToUpper(input) == "SELESAI" {
			fmt.Println("Proses selesai.")
			break
		}

		var berat1, berat2 float64
		_, err := fmt.Sscanf(input, "%f %f", &berat1, &berat2)
		if err != nil {
			fmt.Println("Input tidak valid. Silakan coba lagi.")
			continue
		}

		seimbang := cekKeseimbangan(berat1, berat2)
		fmt.Printf("Sepeda motor Pak Andi akan oleng: %t\n", !seimbang)
	}
}
