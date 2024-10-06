package main
import (
	"fmt"
	"math"
)

func main() {
	var jari2 float64
	const pi = 3.1415926535

	fmt.Print("Jari2 = ")
	fmt.Scanln(&jari2)

	// Menghitung volume bola
	volume := (4.0 / 3.0) * pi * math.Pow(jari2, 3)

	// Menghitung luas permukaan bola
	luas := 4 * pi * math.Pow(jari2, 2)

	// Menampilkan hasil
	fmt.Printf("Bola dengan jari2 %.0f memiliki volume %.4f dan luas permukaan %.4f\n", jari2, volume, luas)
}