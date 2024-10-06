package main

import "fmt"

func main() {
	var tahun int
	var kabisat_2311102030 bool

	fmt.Print("Tahun: ")
	fmt.Scanln(&tahun)

	if tahun%400 == 0 {
		kabisat_2311102030 = true
	} else if tahun%100 == 0 {
		kabisat_2311102030 = false
	} else if tahun%4 == 0 {
		kabisat_2311102030 = true
	} else {
		kabisat_2311102030 = false
	}

	fmt.Println("Kabisat_2311102030:", kabisat_2311102030)
}
