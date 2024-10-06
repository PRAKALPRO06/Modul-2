	package main

	import (
		"fmt"
		"math"
	)

	func fK(K float64) float64 {
		numerator := (4*K + 2) * (4*K + 2)
		denominator := (4*K + 1) * (4*K + 3)
		return numerator / denominator
	}

	func sqrt2_approx(K int) float64 {
		result := 1.0
		for i := 0; i <= K; i++ {
			numerator := (4*float64(i) + 2) * (4*float64(i) + 2)
			denominator := (4*float64(i) + 1) * (4*float64(i) + 3)
			result *= numerator / denominator
		}
		return result
	}

	func main() {
		var K float64
		fmt.Print("Masukan nilai K: ")
		fmt.Scan(&K)

		resultF := fK(K)
		fmt.Printf("Nilai f(K) = %.10f\n", resultF)

		Kint := int(math.Round(K))
		resultSqrt2 := sqrt2_approx(Kint)
		fmt.Printf("Nilai akar 2 = %.10f\n", resultSqrt2)
	}