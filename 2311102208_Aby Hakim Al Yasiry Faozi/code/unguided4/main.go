package main

import (
	"fmt"
	"strings"
)

func gantiNamaBunga(bunga string) string {
	pengganti := map[string]string{
		"Kertas": "Tulip",
		"Mawar":  "Melati",
		"Tulip":  "Anggrek",
	}

	if namaBaru, ok := pengganti[bunga]; ok {
		return namaBaru
	}
	return bunga
}

func main() {
	for {
		var input string
		fmt.Print("Masukkan nama-nama bunga (pisahkan dengan '-'), atau ketik 'SELESAI' untuk berhenti: ")
		fmt.Scanln(&input)

		if input == "SELESAI" {
			break
		}

		bungaList := strings.Split(input, "-")

		for _, bunga := range bungaList {
			namaBaru := gantiNamaBunga(bunga)
			fmt.Printf("Nama asli: %s, Nama baru: %s\n", bunga, namaBaru)
		}
	}
}
