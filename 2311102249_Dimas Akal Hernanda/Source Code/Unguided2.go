package main

import "fmt"

func main() {
	var a, b, c, d, e int
	var char1, char2, char3 rune

	fmt.Print("Masukkan 5 buah data integer (nilai antara 32 s.d. 127):")
	fmt.Scanf("%d %d %d %d %d", &a, &b, &c, &d, &e)

	fmt.Scanf("\n")

	fmt.Print("Masukkan 3 buah karakter (tanpa dipisah spasi):")
	fmt.Scanf("%c%c%c\n", &char1, &char2, &char3)

	fmt.Println(" ")

	fmt.Printf("Konversi ASCII: %c %c %c %c %c\n", a, b, c, d, e)
	fmt.Printf("Output : %c%c%c\n", char1+1, char2+1, char3+1)
}
