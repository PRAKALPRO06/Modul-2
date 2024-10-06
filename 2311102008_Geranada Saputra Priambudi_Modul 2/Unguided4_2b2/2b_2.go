package main

import (
	"fmt"
)

func main() {
	var N int
	var pita string
	fmt.Print("N: ")
	fmt.Scanln(&N)
	for i := 1; i <= N; i++ {
		var bunga string
		fmt.Printf("Bunga %d: ", i)
		fmt.Scanln(&bunga)
		pita += bunga + " - "
	}
	fmt.Println("Pita:", pita)
}
