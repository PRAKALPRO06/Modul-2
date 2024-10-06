package main

import "fmt"

func main() {

	var tahun int

	fmt.Print("masukan Tahun = ")
	fmt.Scanln(&tahun)

	if tahun%400 == 0 {
		fmt.Println("ini tahun kabisat")
	} else if tahun%100 == 0 {
		fmt.Println("ini bukan tahun kabisat")
	} else if tahun%4 == 0 {
		fmt.Println("ini tahun tahun kabisat")
	} else {
		fmt.Print("ini bukan tahun kabisat")
	}
}
