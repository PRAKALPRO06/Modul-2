package main

import "fmt"

func main() {

	var n int
	var pita string

	fmt.Print("N: ")
	fmt.Scanln(&n)

	for i := 1; i<=n; i++ {
		fmt.Printf("Bunga %d: ", i)
		var bunga string
		fmt.Scanln(&bunga)
		pita += bunga + " - "
	}
	fmt.Println("Pita: ", pita)

}