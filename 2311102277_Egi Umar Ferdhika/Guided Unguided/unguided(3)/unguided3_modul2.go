package main

import (
	"fmt"
)

func main() {
	var warna [5][4]string
	urutanBenar := [4]string{"merah", "kuning", "hijau", "ungu"}
	var berhasil bool = true

	for i := range warna {
		fmt.Printf("Masukkan warna untuk percobaan %d (pisahkan dengan spasi): ", i+1)
		fmt.Scan(&warna[i][0], &warna[i][1], &warna[i][2], &warna[i][3])
	}

	for i, percobaan := range warna {
		fmt.Printf("Percobaan %d: %s %s %s %s\n", i+1, percobaan[0], percobaan[1], percobaan[2], percobaan[3])
		if percobaan != urutanBenar {
			berhasil = false
		}
	}

	fmt.Printf("Berhasil: %t\n", berhasil)
}
