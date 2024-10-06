package main

import "fmt"

func main() {
        var nilai float64
        fmt.Print("Masukkan nilai akhir mata kuliah: ")
        fmt.Scan(&nilai)

        var nilaiHuruf string

        switch {
        case nilai >= 80:
                nilaiHuruf = "A"
        case nilai >= 72.5:
                nilaiHuruf = "AB"
        case nilai >= 65:
                nilaiHuruf = "B"
        case nilai >= 57.5:
                nilaiHuruf = "BC"
        case nilai >= 50:
                nilaiHuruf = "C"
        case nilai >= 40:
                nilaiHuruf = "D"
        default:
                nilaiHuruf = "E"
        }

        fmt.Printf("Nilai mata kuliah: %s\n", nilaiHuruf)
}