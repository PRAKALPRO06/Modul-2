package main

import (
	"fmt"
)

func main() {
	// Membaca 5 angka bulat yang dipisahkan oleh spasi
	var num1, num2, num3, num4, num5 int
	fmt.Println("Masukkan 5 angka bulat (32-127) dipisahkan oleh spasi:")
	fmt.Scanf("%d %d %d %d %d", &num1, &num2, &num3, &num4, &num5)

	// Memastikan bahwa angka berada dalam rentang 32 sampai 127
	if num1 < 32 || num1 > 127 || num2 < 32 || num2 > 127 || num3 < 32 || num3 > 127 || num4 < 32 || num4 > 127 || num5 < 32 || num5 > 127 {
		fmt.Println("Semua angka harus berada dalam rentang 32 sampai 127")
		return
	}

	// Konversi angka bulat menjadi karakter ASCII
	fmt.Print("Karakter hasil konversi dari angka: ")
	fmt.Printf("%c%c%c%c%c\n", num1, num2, num3, num4, num5)

	// Membaca 3 karakter yang ditulis berdekatan tanpa spasi
	var char1, char2, char3 rune
	fmt.Println("Masukkan 3 karakter berdekatan tanpa spasi:")
	fmt.Scanf("%c%c%c", &char1, &char2, &char3)

	// Mencetak karakter berikutnya dalam tabel ASCII
	fmt.Print("Karakter setelah 3 karakter input: ")
	fmt.Printf("%c%c%c\n", char1+1, char2+1, char3+1)
}
