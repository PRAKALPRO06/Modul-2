package main

import "fmt"

func main() {
	var tahun int
	fmt.Print("Tahun: ")
	fmt.Scanf("%d", &tahun)

	kabisat := cekTahunKabisat(tahun)
	fmt.Println("Kabisat: ", kabisat)
}

func cekTahunKabisat(tahun int) bool {
	if tahun%400 == 0 {
		return true
	} else if tahun%4 == 0 {
		return true
	} else if tahun%100 == 0 {
		return false
	}
	return false
}
