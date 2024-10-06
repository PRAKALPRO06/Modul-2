package main

import (
	"fmt"
)

func main() {
	var int1, int2, int3, int4, int5 int
	var var1, var2, var3 byte

	// Input 5 angka integer
	fmt.Print("Masukkan 5 buah data integer: ")
	fmt.Scanf("%d %d %d %d %d", &int1, &int2, &int3, &int4, &int5)

	// Membersihkan newline yang tersisa dari input sebelumnya
	fmt.Scanf("%c", new(byte))

	// Input 3 karakter tanpa spasi
	fmt.Print("Masukkan 3 karakter : ")
	fmt.Scanf("%c", &var1) // Input karakter pertama
	fmt.Scanf("%c", &var2) // Input karakter kedua
	fmt.Scanf("%c", &var3) // Input karakter ketiga

	// Mencetak hasil konversi integer ke karakter (ASCII)
	fmt.Printf("Hasil konversi integer ke karakter : %c%c%c%c%c\n", int1, int2, int3, int4, int5)

	// Mencetak karakter setelah karakter yang diinput
	fmt.Printf("Konversi karakter : %c%c%c\n", var1+1, var2+1, var3+1)
}
