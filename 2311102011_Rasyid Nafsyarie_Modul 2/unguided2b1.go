package main

import (
	"fmt"
)

func main() {
	var warna_2311102011 [5][4]string
	urutanBenar := [4]string{"merah", "kuning", "hijau", "ungu"}
	berhasil := true

	for i := 0; i < 5; i++ {
		fmt.Printf("Masukkan warna untuk percobaan %d (pisahkan dengan spasi): ", i+1)
		fmt.Scan(&warna_2311102011[i][0], &warna_2311102011[i][1], &warna_2311102011[i][2], &warna_2311102011[i][3])
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("percobaan %d: %s %s %s %s\n", i+1, warna_2311102011[i][0], warna_2311102011[i][1], warna_2311102011[i][2], warna_2311102011[i][3])
		if warna_2311102011[i] != urutanBenar {
			berhasil = false
		}
	}

	fmt.Printf("berhasil : %t\n", berhasil)
}
