package main

import (
	"fmt"
)

func akar2(num int) float64 {
	result := 1.0

	for k := 0; k < num; k++ {
		term := float64((4*k+2)*(4*k+2)) / float64((4*k+1)*(4*k+3))
		result *= term
	}

	return result
}

func main() {
	var num int

	fmt.Print("Nilai K = ")
	fmt.Scan(&num)

	hasil := akar2(num)
	//fmt.Printf("Nilai akar 2 dengan nilai k %d = %.10f\n", num, hasil)
	fmt.Println("Nilai akar 2: ", hasil)
}
