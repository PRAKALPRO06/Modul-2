package main

import (
	"fmt"
)

func main() {
	var integers [5]int
	var characters [3]byte

	fmt.Println("Masukkan 5 angka integer (32-127) dipisahkan spasi:")
	for i := 0; i < 5; i++ {
		fmt.Scan(&integers[i])
	}

	fmt.Println("Masukkan 3 karakter tanpa spasi:")
	var input string
	fmt.Scan(&input)
	for i := 0; i < 3 && i < len(input); i++ {
		characters[i] = input[i]
	}

	fmt.Print("Output baris pertama: ")
	for _, i := range integers {
		fmt.Printf("%c", rune(i))
	}
	fmt.Println()

	fmt.Print("Output baris kedua: ")
	for _, c := range characters {
		fmt.Printf("%c", c+1)
	}
	fmt.Println()
}
