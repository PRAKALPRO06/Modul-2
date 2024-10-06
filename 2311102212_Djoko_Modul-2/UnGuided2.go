package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, e int
	var chars string

	fmt.Print("Masukkan 5 data integer (32 - 127): ")
	fmt.Scanf("%d %d %d %d %d", &a, &b, &c, &d, &e)

	fmt.Printf("Output karakter: %c%c%c%c%c\n", a, b, c, d, e)

	fmt.Print("Masukkan 3 karakter berdampingan: ")
	fmt.Scanf("%s", &chars)

	nextChars := string(chars[0]+1) + string(chars[1]+1) + string(chars[2]+1)
	fmt.Printf("Karakter setelah: %s\n", nextChars)
}
