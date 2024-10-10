package main

import (
    "fmt"
    "math"
)

func main() {
    var radius int

    // Input jari-jari
    fmt.Print("Masukkan jari-jari bola: ")
    fmt.Scan(&radius)

    // Menghitung volume bola
    volume := (4.0 / 3.0) * math.Pi * math.Pow(float64(radius), 3)

    // Menghitung luas permukaan bola
    luas := 4 * math.Pi * math.Pow(float64(radius), 2)

    // Tampilkan hasil
    fmt.Printf("Volume bola = %.2f\n", volume)
    fmt.Printf("Luas permukaan bola = %.2f\n", luas)
}
