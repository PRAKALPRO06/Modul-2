package main

import (
	"fmt"
)

func main() {
	var integer [5]int
	var karakter [3]byte

	fmt.Println("Silahkan masukkan 5 angka integer (32-127) dipisahkan dengan spasi:")
	for i := 0; i < 5; i++ {
		fmt.Scan(&integer[i])
	}

	fmt.Println("Silahkan masukkan 3 karakter tanpa spasi:")
	var input string
	fmt.Scan(&input)
	for i := 0; i < 3 && i < len(input); i++ {
		karakter[i] = input[i]
	}

	fmt.Print("Output baris pertama: ")
	for _, i := range integer {
		fmt.Printf("%c", rune(i))
	}
	fmt.Println()

	fmt.Print("Output baris kedua: ")
	for _, c := range karakter {
		fmt.Printf("%c", c+1)
	}
	fmt.Println()
}
