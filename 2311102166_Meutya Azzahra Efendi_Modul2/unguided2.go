// Meutya Azzahra Efendi
// IF-11-06
// 2311102166
package main

import "fmt"

func main() {
	var numbers [5]int
	var chars [3]byte

	fmt.Println("Masukkan 5 buah data integer (32-127):")
	for i := 0; i < 5; i++ {
		fmt.Scan(&numbers[i])
	}

	fmt.Println("Masukkan 3 buah karakter:")
	for i := 0; i < 3; i++ {
		fmt.Scanf("%c", &chars[i])
	}

	fmt.Print("Karakter dari nilai integer: ")
	for _, num := range numbers {
		fmt.Printf("%c", num)
	}
	fmt.Println()

	fmt.Print("Karakter setelahnya: ")
	for _, char := range chars {
		fmt.Printf("%c", char+1)
	}
	fmt.Println()
}