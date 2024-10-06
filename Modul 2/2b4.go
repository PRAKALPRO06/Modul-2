package main

import (
        "fmt"
        "math"
)

func hitungF(k int) float64 {
        pembilang := math.Pow(float64(4*k+2), 2)
        penyebut := float64((4*k+1)*(4*k+3))
        return pembilang / penyebut
}

func hitungAkarDua(k int) float64 {
        hasil := 1.0
        for i := 0; i <= k; i++ {
                hasil *= hitungF(i)
        }
        return hasil
}

func main() {
        var k int
        fmt.Print("Masukkan nilai K: ")
        fmt.Scan(&k)

        nilaiF := hitungF(k)
        fmt.Printf("Nilai f(%d) = %.10f\n", k, nilaiF)

        akarDua := hitungAkarDua(k)
        fmt.Printf("Nilai akar 2 untuk K = %d: %.10f\n", k, akarDua)
}