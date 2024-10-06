package main

// 2311102254_Hamzah Ziyad I IF-11-06
import (
	"fmt"
	"math"
)

func f(k int) float64 {
	numerator := math.Pow(float64(4*k+2), 2)
	denominator := float64((4*k + 1) * (4*k + 3))
	return numerator / denominator
}

func main() {

	var k int
	fmt.Print("Nilai K: ")
	fmt.Scan(&k)

	result := f(k)
	fmt.Printf("Nilai f(K) = %.10f\n", result)
}
