package main

import "fmt"

func main() {
	var tahun_2311102021 int
	var kabisat bool

	fmt.Print("Tahun : ")
	fmt.Scanln(&tahun_2311102021)

	if tahun_2311102021%400 == 0 {
		kabisat = true
	} else if tahun_2311102021%100 == 0 {
		kabisat = false
	} else if tahun_2311102021%4 == 0 {
		kabisat = true
	} else {
		kabisat = false
	}

	fmt.Println("Kabisat:", kabisat)
}