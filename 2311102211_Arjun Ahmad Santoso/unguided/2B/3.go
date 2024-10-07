package main

import (
	"fmt"
	"math"
)

func main() {
	berhenti := false
	for berhenti == false {
		var b1, b2 float64
		oleng := false
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scanln(&b1, &b2)
		if b1 + b2 > 150 || b1 < 0 || b2 < 0 {
			break
		}
		if math.Abs(b1 - b2) >= 9 {
			oleng = true
		}
		fmt.Println("Sepeda motor Pak Andi akan oleng: ", oleng)
	}
}