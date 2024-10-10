package main

import (
    "fmt"
)

// Fungsi untuk menghitung nilai f(K) sesuai dengan rumus yang diberikan
func calculateF(K float64) float64 {
    return (4*K + 2) * (4*K + 2) / ((4*K + 1) * (4*K + 3))
}

// Fungsi untuk menghampiri akar kuadrat dari 2 dengan iterasi tertentu
func approximateSqrt2(iterations int) float64 {
    sqrt2Approx := 1.0
    for i := 0; i < iterations; i++ {
        sqrt2Approx = (sqrt2Approx + 2/sqrt2Approx) / 2
    }
    return sqrt2Approx
}

func main() {
    // Meminta pengguna untuk memasukkan nilai K
    var K float64
    fmt.Print("Masukkan nilai K: ")
    fmt.Scanln(&K)

    // Menghitung nilai f(K) menggunakan fungsi calculateF
    fK := calculateF(K)

    // Menampilkan hasil perhitungan nilai f(K)
    fmt.Println("Nilai f(K) =", fK)

    // Menghitung hampiran nilai akar kuadrat dari 2 dengan iterasi sebanyak 10 kali
    sqrt2Approx := approximateSqrt2(10)

    // Menampilkan informasi mengenai akar kuadrat dari 2 dan hasil hampiran
    fmt.Println("\n√2 merupakan bilangan irasional.")
    fmt.Println("Meskipun demikian, nilai tersebut dapat dihampiri dengan rumus berikut:")
    fmt.Println("√2 ≈", sqrt2Approx)
}
