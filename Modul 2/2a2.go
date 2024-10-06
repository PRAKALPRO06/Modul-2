package main

import "fmt"

func main() {
        var tahun int

        fmt.Print("Masukkan tahun: ")
        fmt.Scanln(&tahun)

        kabisat := (tahun%4 == 0 && tahun%100 != 0) || tahun%400 == 0

        fmt.Printf("Kabisat: %t\n", kabisat)
}