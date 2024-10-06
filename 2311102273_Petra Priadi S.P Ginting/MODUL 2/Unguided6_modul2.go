package main

import (
	"fmt"
	"math"
)

func main() {
	var k int

	fmt.Print("Masukkan nilai K: ")
	_, err := fmt.Scanf("%d", &k)
	if err != nil {
		fmt.Println("Error saat membaca input:", err)
		return
	}

	result := calculateSqrt2(k)
	fmt.Printf("Nilai K = %d\n", k)
	fmt.Printf("Nilai akar 2 = %.10f\n", result)
}

func calculateSqrt2(k int) float64 {
	result := 1.0
	for i := 0; i < k; i++ {
		result *= calculateF(float64(i))
	}
	return result
}

func calculateF(k float64) float64 {
	numerator := math.Pow(4*k+2, 2)
	denominator := (4*k + 1) * (4*k + 3)
	return numerator / denominator
}
