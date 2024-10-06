package main

import (
	"fmt"
)

func main() {
	var warna1, warna2, warna3, warna4 string
	berhasil := true
	urutanBenar := [4]string{"merah", "kuning", "hijau", "ungu"}

	for percobaan := 1; percobaan <= 5; percobaan++ {
		fmt.Printf("Percobaan %d: ", percobaan)
		fmt.Scan(&warna1, &warna2, &warna3, &warna4)

		if warna1 != urutanBenar[0] || warna2 != urutanBenar[1] ||
			warna3 != urutanBenar[2] || warna4 != urutanBenar[3] {
			berhasil = false
		}
	}

	if berhasil {
		fmt.Println("BERHASIL: TRUE")
	} else {
		fmt.Println("BERHASIL: FALSE")
	}
}
