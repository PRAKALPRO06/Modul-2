package main

import "fmt"

func main() {
	var a, pita string
	i := 0
	for a != "SELESAI" {
		fmt.Printf("\nBunga %v : ", i+1)
		fmt.Scanln(&a)

		if a != "SELESAI" {
			a += " - "
			pita += a
			i++
		}

	}
	fmt.Printf("\nPita  : %v\n", pita)
	fmt.Printf("Bunga : %v\n", i)
}
