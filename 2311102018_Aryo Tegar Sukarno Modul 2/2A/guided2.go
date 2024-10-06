package main

import "fmt"

func kabisat(tahun int) bool {
    if tahun % 400 == 0 {
        return true
    } else if tahun % 100 == 0 {
        return false 
    } else if tahun % 4 == 0 {
        return true
    } else {
        return false
    }
}

func main() {
    var tahun int
    fmt.Print("Ketik Tahun : ")
    fmt.Scanln(&tahun)
    if kabisat(tahun) {
        fmt.Println(tahun, "ini adalah tahun kabisat")
    } else {
        fmt.Println(tahun, "ini bukan tahun kabisat")
    }
}