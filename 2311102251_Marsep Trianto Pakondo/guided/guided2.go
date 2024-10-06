package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64

	fmt.Print("Masukkan jari-jari Bola : ")
	fmt.Scanln(&r)

	volumebola := (4.0/3.0) * math.Pi * math.Pow(r, 3)
	luasbola := 4.0 * math.Pi * math.Pow(r, 2)

	fmt.Printf("Bola dengan jejari %v memiliki volume %.4f dan luas kulit %.4f", r, volumebola, luasbola)
}
