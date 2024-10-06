package main

import (
    "fmt"
    "math"
)

func main() {
    var jarijari float64
    fmt.Print("Masukkan jari-jari bola : ")
    fmt.Scanln(&jarijari)

    volume := (4.0 / 3.0) * math.Pi * math.Pow(jarijari, 3)
    luasPermukaan := 4 * math.Pi * math.Pow(jarijari, 2)

    fmt.Printf("Volume bola dengan jari-jari %.2f adalah %.2f\n", jarijari, volume)
    fmt.Printf("Luas permukaan bola dengan jari-jari %.2f adalah %.2f\n", jarijari, luasPermukaan)
}