package main

import (
	"fmt"
)

//reza alvonzo 2311102026 IF 11 06
func main() {
	var warna [5][4]string
	urutanBenar := [4]string{"merah", "kuning", "hijau", "ungu"}
	berhasil := true

	for i := 0; i < 5; i++ {
		fmt.Printf("Masukkan warna untuk percobaan %d (pisahkan dengan spasi): ", i+1)
		fmt.Scan(&warna[i][0], &warna[i][1], &warna[i][2], &warna[i][3])
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("percobaan %d: %s %s %s %s\n", i+1, warna[i][0], warna[i][1], warna[i][2], warna[i][3])
		if warna[i] != urutanBenar {
			berhasil = false
		}
	}

	fmt.Printf("berhasil : %t\n", berhasil)
}
