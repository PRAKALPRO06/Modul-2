package main

import (
	"fmt"
)

func main() {
	var k int

	fmt.Print("Nilai K = ")
	fmt.Scanln(&k)

	f_k := (float64(4*k+2) * float64(4*k+2)) / (float64(4*k+1) * float64(4*k+3))

	fmt.Printf("Nilai f(K) = %.10f\n", f_k)

	var hitungfk float64

	for i := 0; i <= k; i++ {
		hitungfk += (float64(4*i+2) * float64(4*i+2)) / (float64(4*i+1) * float64(4*i+3))
	}

	fmt.Printf("Nilai hampiran √2 = %.10f\n", hitungfk)
}
