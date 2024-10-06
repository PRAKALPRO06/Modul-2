package main

import "fmt"

func main() {
        var satu, dua, tiga string
        var temp string

        fmt.Print("Masukkan input string pertama: ")
        fmt.Scanln(&satu)

        fmt.Print("Masukkan input string kedua: ")
        fmt.Scanln(&dua)

        fmt.Print("Masukkan input string ketiga: ")
        fmt.Scanln(&tiga)

        fmt.Println("Output awal =", satu, dua, tiga)

        // Menukar nilai satu dan dua
        temp = satu
        satu = dua
        dua = temp

        fmt.Println("Output akhir =", satu, dua, tiga)
}