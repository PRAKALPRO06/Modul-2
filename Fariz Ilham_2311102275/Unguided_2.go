package main

import (
	"fmt"
	"unicode"
)

func main() {
	var intArr [5]int
	var charArr [3]rune

	fmt.Print("Masukkan 5 buah data integer (32-127): ")
	for i := 0; i < 5; i++ {
		fmt.Scan(&intArr[i])

		if intArr[i] < 32 || intArr[i] > 127 {
			fmt.Println("Input harus berupa bilangan antara 32 sampai 127.")
			return
		}
	}

	fmt.Print("Masukkan 3 buah karakter berdampingan (tanpa spasi): ")
	var input string
	fmt.Scan(&input)

	if len(input) != 3 {
		fmt.Println("Input harus terdiri dari 3 karakter.")
		return
	}

	for _, char := range input {
		if unicode.IsDigit(char) {
			fmt.Println("Input harus berupa karakter, bukan angka.")
			return
		}
	}

	for i := 0; i < 3; i++ {
		charArr[i] = rune(input[i])
	}

	fmt.Println()
	fmt.Print("Keluaran:\n")
	for i := 0; i < 5; i++ {
		fmt.Print(string(intArr[i]))
	}
	fmt.Println()

	for i := 0; i < 3; i++ {
		fmt.Print(string(charArr[i] + 1))
	}
	fmt.Println()
}
