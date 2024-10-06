// 2311102012
package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, e int
	var char1, char2, char3 rune

	// Membaca 5 buah data integer dari pengguna
	fmt.Println("Masukkan 5 angka integer :")
	fmt.Scanf("%d %d %d %d %d", &a, &b, &c, &d, &e)

	// Membaca newline setelah angka agar tidak terbaca sebagai karakter pertama
	fmt.Scanf("\n")

	// Membaca 3 buah karakter dari pengguna
	fmt.Println("Masukkan 3 karakter:")
	fmt.Scanf("%c%c%c", &char1, &char2, &char3)

	// Menampilkan keluaran untuk 5 integer dalam bentuk karakter (ASCII)
	fmt.Printf("Hasil konversi angka ke karakter: %c%c%c%c%c\n", a, b, c, d, e)

	// Menampilkan 3 karakter
	fmt.Printf("Karakter yang diinput: %c%c%c\n", char1, char2, char3)
}
