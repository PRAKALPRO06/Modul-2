package main

import "fmt"

var a, b int

// Fungsi untuk menghitung faktorial dari bilangan n
func faktorial(n int) int {
	hasil := 1
	// Loop untuk menghitung faktorial, dimulai dari 1 hingga n
	for i := 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

// Fungsi untuk menghitung permutasi P(n, r) = n! / (n-r)!
func permutasi(n, r int) int {
	// P(n, r) = faktorial(n) / faktorial(n-r)
	return faktorial(n) / faktorial(n-r)
}

func main() {
	// Input dari pengguna untuk nilai a dan b
	fmt.Print("Inputkan a = ")
	fmt.Scan(&a)
	fmt.Print("Inputkan b = ")
	fmt.Scan(&b)

	// Memastikan a >= b agar menghitung permutasi dengan benar
	if a >= b {
		// Cetak hasil permutasi P(a, b)
		fmt.Println("Hasil permutasi:", permutasi(a, b))
	} else {
		// Jika b lebih besar dari a, hitung permutasi P(b, a)
		fmt.Println("Hasil permutasi:", permutasi(b, a))
	}
}
