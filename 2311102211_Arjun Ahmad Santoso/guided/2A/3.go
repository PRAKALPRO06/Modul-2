package main

import (
	"fmt"
	"math"
)

func main() {
	const pi = 3.1415926535
	var r int
	var (
		v, lp float64
	)

	fmt.Print("Masukkan jari-jari bola: ")
	fmt.Scanln(&r)

	if r < 0 {
		fmt.Print("Input tidak valid!")
		return
	}
	v = pi * 4 / 3 * math.Pow(float64(r), 3)
	lp = pi * 4 * pi * math.Pow(float64(r), 2)

	fmt.Printf("Bola dengn jari-jari %d memiliki volume %f dan luas permukaan %f", r, v, lp)
}