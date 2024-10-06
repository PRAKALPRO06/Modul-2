package main

import "fmt"

func main() {

	var satu, dua, tiga, temp string

	fmt.Print("Masukkan input string : ")
	fmt.Scanln(&satu)
	fmt.Print("Masukkan input string : ")
	fmt.Scanln(&dua)
	fmt.Print("Masukkan input string : ")
	fmt.Scanln(&tiga)
	fmt.Println("OUTPUT AWAL = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("OUTPUT AKHIR = " + satu + " " + dua + " " + tiga)

}