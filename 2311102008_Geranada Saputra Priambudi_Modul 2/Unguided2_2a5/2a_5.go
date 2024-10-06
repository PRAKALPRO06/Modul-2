package main

import (
	"fmt"
)

func main() {
	var (
		integers   [5]int
		characters [3]rune
	)

	fmt.Print("Masukkan 5 data integer (antara 32 s.d. 127): ")
	for i := 0; i < 5; i++ {
		fmt.Scan(&integers[i])
	}

	fmt.Print("Output karakter: ")
	for _, value := range integers {
		fmt.Printf("%c", value)
	}
	fmt.Println()

	fmt.Print("Masukkan 3 huruf: ")
	var input string
	fmt.Scan(&input)

	for i, char := range input {
		if i < 3 {
			characters[i] = char
		}
	}

	fmt.Print("Output huruf : ")
	for _, char := range characters {
		nextChar := char + 1
		fmt.Printf("%c", nextChar)
	}
	fmt.Println()
}
