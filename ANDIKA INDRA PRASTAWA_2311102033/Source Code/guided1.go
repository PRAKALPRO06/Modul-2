package main

import "fmt"

func main() {

	var satu, dua, tiga, temp string

	fmt.Print("masukian string pertama : ")
	fmt.Scanln(&satu)
	fmt.Print("masukian string kedua :")
	fmt.Scanln(&dua)
	fmt.Print("masukian string Tiga : ")
	fmt.Scanln(&tiga)
	fmt.Println("OUTPUT AWAL = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("OUTPUT AKHIR = " + satu + " " + dua + " " + tiga)

}
