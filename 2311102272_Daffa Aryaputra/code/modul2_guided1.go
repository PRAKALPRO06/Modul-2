package main

import "fmt"

func main () {
	var (
		satu, dua, tiga string
		temp            string
	)

	fmt.Println("Masukkan input string pertama: ")
	fmt.Scanln(&satu)
	fmt.Println("Masukkan input string kedua: ")
	fmt.Scanln(&dua)
	fmt.Println("Masukkan input string ketiga: ")
	fmt.Scanln(&tiga)

	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)

	// Proses penukaran
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp

	fmt.Println("Output setelah ditukar = " + satu + " " + dua + " " + tiga)
}