package main

import (
        "fmt"
        "math"
)

func hitungF(k int) float64 {
        pembilang := math.Pow(float64(4*k+2) , 2)
        penyebut := float64((4*k+1) * (4*k+3))
        return pembilang / penyebut
}

func hitungAkar2(k int) float64 {
        hasil := 1.0
        for i := 0; i <= k; i++ {
                hasil *= hitungF(i)
        }
        return hasil
}

func main() {
        var k_2311102013 int
        fmt.Print("Masukkan nilai K: ")
        fmt.Scan(&k_2311102013)

        nilaiF := hitungF(k_2311102013)
        fmt.Printf("Nilai f(%d) = %.10f\n", k_2311102013, nilaiF)

        akar2 := hitungAkar2(k_2311102013)
        fmt.Printf("Nilai akar 2 ≈ %.10f\n", akar2)
}