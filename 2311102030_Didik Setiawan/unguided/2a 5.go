package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var nums [5]int
	reader := bufio.NewReader(os.Stdin)

	// Input angka dari 32 hingga 127 dan cetak sebagai karakter
	fmt.Println("Masukkan 5 buah angka (nilai antara 32 hingga 127):")
	fmt.Scanf("%d %d %d %d %d", &nums[0], &nums[1], &nums[2], &nums[3], &nums[4])

	fmt.Print("Hasil karakter dari angka yang diinput: ")
	for _, num := range nums {
		fmt.Printf("%c", num)
	}
	fmt.Println()

	// Input string karakter dan geser setiap karakter
	fmt.Println("Masukkan karakter (contoh input: 123):")
	inputStr, _ := reader.ReadString('\n')

	fmt.Print("Hasil karakter dari inputan angka yang digeser: ")
	for _, char := range inputStr {
		if char != '\n' && char != '\r' {
			fmt.Printf("%c", char+1)
		}
	}
	fmt.Println()
}
