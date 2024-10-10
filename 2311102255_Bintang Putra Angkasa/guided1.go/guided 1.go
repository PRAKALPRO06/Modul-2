package main

import (
    "fmt"
)

func cekTahunKabisat() {
    var tahun int
    fmt.Print("Masukkan sebuah tahun: ")
    fmt.Scanln(&tahun)

    if (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0) {
        fmt.Println(tahun, "adalah tahun kabisat")
    } else {
        fmt.Println(tahun, "bukan tahun kabisat")
    }
}

func main() {
    cekTahunKabisat()
}
