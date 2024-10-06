package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, e int
	var x, y, z string

	fmt.Print("Masukkan 5 data integer (32 - 127): ")
	fmt.Scanf("%d %d %d %d %d", &a, &b, &c, &d, &e)

	fmt.Printf("Output karakter: %c%c%c%c%c\n", a, b, c, d, e)

	fmt.Print("Masukkan 3 karakter: ")
	fmt.Scanf("%s %s %s", &x, &y, &z)

	fmt.Printf("Output karakter: %s%s%s\n", x, y, z)
}

