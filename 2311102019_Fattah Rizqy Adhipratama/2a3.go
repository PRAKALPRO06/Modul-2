package main

import (
        "fmt"
        "math"
)

func main() {
        var jariJari float64
        const pi = 3.1415926535

        fmt.Print("Masukkan jari-jari bola: ")
        fmt.Scanln(&jariJari)

        volume := (4 / 3) * pi * math.Pow(jariJari, 3)
        luas := 4 * pi * math.Pow(jariJari, 2)

        fmt.Printf("Bola dengan jari-jari %.0f memiliki volume %.6f dan luas %.6f\n", jariJari, volume, luas)
}