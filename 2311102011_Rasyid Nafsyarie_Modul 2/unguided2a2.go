package main

import "fmt"

func main() {
	var tahun_2311102011 int
	var kabisat bool

	fmt.Print("Tahun : ")
	fmt.Scanln(&tahun_2311102011)

	if tahun_2311102011%400 == 0 {
		kabisat = true
	} else if tahun_2311102011%100 == 0 {
		kabisat = false
	} else if tahun_2311102011%4 == 0 {
		kabisat = true
	} else {
		kabisat = false
	}

	fmt.Println("Kabisat:", kabisat)
}