package main

import "fmt"

func main() {
    var tahun int
    fmt.Print("Masukkan tahun: ")
    fmt.Scanln(&tahun)

    if (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0) {
        fmt.Printf("%d adalah tahun kabisat: true\n", tahun)
    } else {
        fmt.Printf("%d bukan tahun kabisat: false\n", tahun)
    }
}
