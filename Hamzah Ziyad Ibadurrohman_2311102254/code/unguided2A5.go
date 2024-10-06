package main

import (
	"fmt"
)

func main() {
	var integers [5]int
	var chars [3]byte

	// Membaca 5 data integer

	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &integers[i])

		// Validasi input
		if integers[i] < 32 || integers[i] > 127 {
			fmt.Println("Input integer harus di antara 32 dan 127.")
			return
		}
	}

	// Membaca 3 data karakter

	fmt.Scanf("\n%c%c%c", &chars[0], &chars[1], &chars[2])

	// Mencetak 5 karakter dari integer
	fmt.Print("Karakter dari Integer: ")
	for i := 0; i < 5; i++ {
		fmt.Printf("%c", integers[i])
	}
	fmt.Println()

	// Mencetak 3 karakter
	fmt.Print("Karakter: ")
	for i := 0; i < 3; i++ {
		fmt.Printf("%c", chars[i])
	}
	fmt.Println()
}
