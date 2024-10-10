package main

import (
    "fmt"
    "math"
)

func main() {
    for {
        var kantong1, kantong2 float64

        // Meminta input berat belanjaan di dua kantong dari pengguna
        fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
        fmt.Scan(&kantong1, &kantong2)

        // Menghentikan proses jika salah satu berat kantong negatif
        if kantong1 < 0 || kantong2 < 0 {
            fmt.Println("Proses selesai.")
            break
        }

        // Menghentikan proses jika total berat di kedua kantong lebih dari 150 kg
        if kantong1+kantong2 > 150 {
            fmt.Println("Proses selesai.")
            break
        }

        // Memeriksa apakah sepeda motor oleng (selisih berat kedua kantong >= 9 kg)
        oleng := math.Abs(kantong1-kantong2) >= 9
        fmt.Printf("Sepeda motor Pak Andi akan oleng: %v\n", oleng)

        // Menghentikan proses jika salah satu kantong memiliki berat 9 kg atau lebih
        if kantong1 >= 9 || kantong2 >= 9 {
            fmt.Println("Proses selesai.")
            break
        }
    }
}
