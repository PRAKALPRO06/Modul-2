package main

import (
	"fmt"
)

func hitungFungsiFK(K int) float64 {
	return float64((4*K+2)*(4*K+2)) / float64((4*K+1)*(4*K+3))
}
func hitungAkar2(K int) float64 {
	akar2 := 1.0
	for k := 0; k <= K; k++ {
		akar2 *= float64((4*k+2)*(4*k+2)) / float64((4*k+1)*(4*k+3))
	}
	return akar2
}

func main() {
	var pilihan int
	var K int

	fmt.Println("Pilih opsi:")
	fmt.Println("1. Hitung f(K)")
	fmt.Println("2. Hitung akar 2")
	fmt.Print("Masukkan pilihan (1/2): ")
	fmt.Scan(&pilihan)

	fmt.Print("Nilai K: ")
	fmt.Scan(&K)

	if pilihan == 1 {
		fK := hitungFungsiFK(K)
		fmt.Printf("Nilai f(K) = %.10f\n", fK)
	} else if pilihan == 2 {
		akar2 := hitungAkar2(K)
		fmt.Printf("Nilai akar 2 = %.10f\n", akar2)
	} else {
		fmt.Println("Pilihan tidak valid!")
	}
}
