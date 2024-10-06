package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, e int
	var char1, char2, char3 rune

	fmt.Println("Masukkan 5 buah integer (spasi dipisahkan)(Antara 32 s.d. 127): ")
	fmt.Scanf("%d %d %d %d %d", &a, &b, &c, &d, &e)

	fmt.Println("Masukkan 3 buah karakter (tanpa spasi): ")
	fmt.Scanf("%c%c%c", &char1, &char2, &char3)

	fmt.Printf("%c%c%c%c%c\n", a, b, c, d, e)

	fmt.Printf("%c%c%c\n", char1, char2, char3)
}
