package main

import (
	"fmt"
	"strings"
)

func main() {
	var warna1, warna2, warna3, warna4 string
	var berhasil bool = true

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)
		fmt.Scanln(&warna1, &warna2, &warna3, &warna4)

		if !strings.EqualFold(warna1, "merah") || !strings.EqualFold(warna2, "kuning") || !strings.EqualFold(warna3, "hijau") || !strings.EqualFold(warna4, "ungu") {
			berhasil = false
		}
	}

	fmt.Println("BERHASIL:", berhasil)
}
