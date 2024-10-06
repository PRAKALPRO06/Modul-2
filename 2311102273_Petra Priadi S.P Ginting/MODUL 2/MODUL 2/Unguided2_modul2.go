package main

import (
	"fmt"
)

func main() {
	var intData [5]int
	fmt.Println("Masukkan 5 buah data integer (nilai antara 32 s.d. 127):")
	for i := 0; i < 5; i++ {
		fmt.Scan(&intData[i])
	}

	for _, val := range intData {
		fmt.Printf("%c", val)
	}
	fmt.Println()

	var charData string
	fmt.Println("Masukkan 3 buah karakter berdampingan:")
	fmt.Scan(&charData)

	if len(charData) != 3 {
		fmt.Println("Masukkan harus terdiri dari 3 karakter.")
		return
	}

	for _, char := range charData {
		fmt.Printf("%c", char+1)
	}
	fmt.Println()
}
