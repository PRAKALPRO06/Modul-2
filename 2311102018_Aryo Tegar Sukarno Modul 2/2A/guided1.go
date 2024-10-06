package main

import "fmt"

func main() {
    var satu, dua, tiga string
    fmt.Print("Masukan input string 1: ")
    fmt.Scanln(&satu)
    fmt.Print("Masukan input string 2: ")
    fmt.Scanln(&dua)
    fmt.Print("Masukan input string 3: ")
    fmt.Scanln(&tiga)

    fmt.Println("Output awal = ", satu, dua, tiga)
    temp := satu
    satu = dua
    dua = tiga
    tiga = temp

    fmt.Println("Output akhir = ", satu, dua, tiga)
}