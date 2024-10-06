package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, e int
	var char1, char2, char3 rune

	
	fmt.Println("Masukkan 5 angka integer :")
	fmt.Scanf("%d %d %d %d %d", &a, &b, &c, &d, &e)

	
	fmt.Scanf("\n")

	fmt.Println("Masukkan 3 karakter:")
	fmt.Scanf("%c%c%c", &char1, &char2, &char3)

	
	fmt.Printf("Hasil konversi angka ke karakter: %c%c%c%c%c\n", a, b, c, d, e)
	
	fmt.Printf("Karakter yang diinput: %c%c%c\n", char1, char2, char3)
}