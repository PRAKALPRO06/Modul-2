package main

import "fmt"

func main() {
        var numbers [5]int
        var chars [3]byte

        // Membaca 5 buah data integer
        fmt.Println("Masukkan 5 buah data integer (32-127):")
        for i := 0; i < 5; i++ {
                fmt.Scanf("%d", &numbers[i])
        }

        // Membaca 3 buah karakter
        fmt.Println("Masukkan 3 buah karakter:")
        for i := 0; i < 3; i++ {
                fmt.Scanf("%c", &chars[i])
        }

        // Mencetak representasi karakter dari data integer
        fmt.Print("Representasi karakter: ")
        for _, num := range numbers {
                fmt.Printf("%c", num)
        }
        fmt.Println()

        // Mencetak 3 karakter setelah karakter yang diberikan
        fmt.Print("3 karakter setelahnya: ")
        for _, char := range chars {
                fmt.Printf("%c", char+1)
        }
        fmt.Println()
}