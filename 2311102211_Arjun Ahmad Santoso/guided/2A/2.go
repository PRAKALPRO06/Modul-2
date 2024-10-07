package main

import "fmt"

func main() {
	var tahun int

	fmt.Print("Masukkan tahun: ")
	fmt.Scanln(&tahun)

	if tahun%4 == 0 {
		fmt.Printf("%d adalah tahun kabisat", tahun)
	}

}