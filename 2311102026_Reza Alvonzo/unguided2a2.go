package main

import "fmt"


//reza alvonzo 2311102026 IF 11 06
func main() {
	var tahun int
	var kabisat bool

	fmt.Print("Tahun : ")
	fmt.Scanln(&tahun)

	if tahun%400 == 0 {
		kabisat = true
	} else if tahun%100 == 0 {
		kabisat = false
	} else if tahun%4 == 0 {
		kabisat = true
	} else {
		kabisat = false
	}

	fmt.Println("Kabisat:", kabisat)
}