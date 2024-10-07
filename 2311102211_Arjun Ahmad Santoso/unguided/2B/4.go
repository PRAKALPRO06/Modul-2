package main

import (
	"fmt"
	"math"
)

func f(k int) float64 {
	return math.Pow(4 * float64(k) + 2, 2)/((4 * float64(k) + 1)*(4 * float64(k) + 3))
}
func f2(k int) float64 {
	var result float64 = 1

	for i:=0; i<=k; i++ {
		result *= math.Pow(4 * float64(i) + 2, 2)/((4 * float64(i) + 1)*(4 * float64(i) + 3))
	}
	return result
}

func main() {

	var k int
	fmt.Print("Nilai k: ")
	fmt.Scanln(&k)
	fmt.Printf("Nilai f(k): %.10f \n", f(k))
	fmt.Printf("Nilai akar 2: %.10f", f2(k))
}