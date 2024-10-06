package main

import "fmt"

func main() {
	var intData [5]int
	var charData string

	fmt.Println("Masukkan 5 buah data integer (antara 32 hingga 127):")
	for i := 0; i < 5; i++ {
		fmt.Scan(&intData[i])
	}

	fmt.Println("Masukkan 3 buah karakter yang berdampingan:")
	fmt.Scan(&charData)

	fmt.Print("Karakter dari integer yang dimasukkan: ")
	for _, num := range intData {
		fmt.Printf("%c", num)
	}
	fmt.Println()

	fmt.Print("3 karakter setelah input: ")
	for _, ch := range charData {
		fmt.Printf("%c", ch+3)
	}
	fmt.Println()
}
