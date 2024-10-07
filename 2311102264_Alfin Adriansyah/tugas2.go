package main

import "fmt"

func main() {
	var a, b, c, d, e int
	var x, y, z rune

	fmt.Print("Masukkan 5 angka integer: ")
	fmt.Scanf("%d %d %d %d %d", &a, &b, &c, &d, &e)

	fmt.Print("Masukkan 3 karakter tanpa spasi: ")
	fmt.Scanf("%c%c%c", &x, &y, &z)

	fmt.Printf("OUTPUT AWAL = %c%c%c%c%c\n", a, b, c, d, e)
	fmt.Printf("OUTPUT AKHIR = %c%c%c\n", x, y, z)
}
