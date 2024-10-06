package main

import "fmt"

func main() {
	var (
		satu, dua311102030, tiga string
		temp                     string
	)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua311102030)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua311102030 + " " + tiga)
	temp = satu
	satu = dua311102030
	dua311102030 = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua311102030 + " " + tiga)
}
