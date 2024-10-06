package main

import "fmt"

func main() {
        var angka [5]int
        var karakter [3]byte

        fmt.Println("Masukkan 5 angka integer (32-127):")
        for i := 0; i < 5; i++ {
                fmt.Scan(&angka[i])
        }

        fmt.Println("Masukkan 3 karakter:")
        fmt.Scanf("%c%c%c", &karakter[0], &karakter[1], &karakter[2])

        fmt.Print("Karakter dari angka: ")
        for i := 0; i < 5; i++ {
                fmt.Printf("%c", angka[i])
        }
        fmt.Println()

        fmt.Print("Karakter setelahnya: ")
        for i := 0; i < 3; i++ {
                fmt.Printf("%c", karakter[i]+1)
        }
        fmt.Println()
}