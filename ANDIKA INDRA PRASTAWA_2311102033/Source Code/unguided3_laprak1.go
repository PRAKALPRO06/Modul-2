package main

import (
	"fmt"
)

func main() {
	var kombinasiWarna [5][4]string
	polaBenar := [4]string{"merah", "kuning", "hijau", "ungu"}
	tepat := true

	for percobaan := 0; percobaan < 5; percobaan++ {
		fmt.Printf("Masukkan kombinasi warna untuk percobaan ke-%d : ", percobaan+1)
		fmt.Scan(&kombinasiWarna[percobaan][0], &kombinasiWarna[percobaan][1], &kombinasiWarna[percobaan][2], &kombinasiWarna[percobaan][3])
	}

	for percobaan := 0; percobaan < 5; percobaan++ {
		fmt.Printf("Hasil percobaan ke-%d: %s %s %s %s\n", percobaan+1, kombinasiWarna[percobaan][0], kombinasiWarna[percobaan][1], kombinasiWarna[percobaan][2], kombinasiWarna[percobaan][3])
		if kombinasiWarna[percobaan] != polaBenar {
			tepat = false
		}
	}

	fmt.Printf("Apakah kombinasi benar? : %t\n", tepat)
}
