package main
import (
	"fmt"
)

func hitungAkar2(k int) float64 {

	//Rumusnya ketika di code kan
	hasil := 1.0
	for i := 0; i < k; i++ {
		NilaiAtas := (4*float64(i) + 2) * (4*float64(i) + 2)
		NilaiBawah := (4*float64(i) + 1) * (4*float64(i) + 3)
		hasil *= NilaiAtas / NilaiBawah
	}
	return hasil
}

func main() {
	var k int
	fmt.Print("Nilai K = ")
	fmt.Scan(&k)

	akar2 := hitungAkar2(k)
	fmt.Printf("Nilai akar 2 = %.10f\n", akar2)
}