package main

import "fmt"

func main() {
	var (
		satu_2311102021, dua, tiga string
		temp            string
	)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu_2311102021)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu_2311102021 + " " + dua + " " + tiga)
	temp = satu_2311102021
	satu_2311102021 = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu_2311102021 + " " + dua + " " + tiga)
}